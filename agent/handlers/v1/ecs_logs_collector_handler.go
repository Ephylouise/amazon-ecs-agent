package v1

import (
	"archive/tar"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/amazon-ecs-agent/agent/httpclient"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/awserr"
	awscreds "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/amazon-ecs-agent/agent/credentials/instancecreds"
	"github.com/cihub/seelog"
)

// AgentMetadataPath is the Agent metadata path for v1 handler.
const (
	// ECSLogsCollectorPath is the Agent metadata path for v1 logs collector.
	ECSLogsCollectorPath = "/v1/logsbundle"

	logsFilePathDir = "/var/lib/ecs/data"

	s3UploadTimeout = 5 * time.Minute
)

type logCollectorResponse struct {
	LogBundleURL string
}

// ECSLogsCollectorHandler creates response for 'v1/logsbundle' API.
func ECSLogsCollectorHandler(containerInstanceArn, region string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// create logscollect state file
		createLogscollectFile()

		seelog.Infof("Finding the logsbundle...")
		logsFound, key := isLogsCollectionSuccessful()
		if !logsFound {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// upload the ecs logs
		iamcredentials, err := instancecreds.GetCredentials(false).Get()
		if err != nil {
			seelog.Debugf("Error getting instance credentials %v", err)
		}
		arn, err := arn.Parse(containerInstanceArn)
		if err != nil {
			seelog.Debugf("Error parsing containerInstanceArn %s, err: %v", containerInstanceArn, err)
		}
		bucket := "ecs-logs-" + arn.AccountID
		err = uploadECSLogsToS3(iamcredentials, bucket, key, region)
		if err != nil {
			seelog.Debugf("Error uploading the ecs logs %v", err)
		}

		// return the presigned url
		presignedUrl := getPreSignedUrl(iamcredentials, bucket, key, region)
		seelog.Infof("Presigned URL for ECS logs: %s", presignedUrl)
		resp := logCollectorResponse{LogBundleURL: presignedUrl}
		respBuf, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respBuf)
	}
}

func createLogscollectFile() {
	logscollectFile, err := os.Create(filepath.Join(logsFilePathDir, "logscollect"))
	defer logscollectFile.Close()
	if err != nil {
		seelog.Errorf("Error creating logscollect file, err: %v", err)
	}
	seelog.Debugf("Successfully created %s file", logscollectFile.Name())
	logscollectFile.Close()
}

func readLogsbundleTar(logsbundlepath string) io.Reader {
	seelog.Debugf("Reading logsbundle from path %s", logsbundlepath)
	tarFile, err := os.Open(logsbundlepath)
	if err != nil {
		seelog.Errorf("Error: %v", err)
	}
	defer tarFile.Close()
	return tar.NewReader(tarFile)
}

func isLogsCollectionSuccessful() (bool, string) {
	var err error
	for i := 0; i < 60; i++ {
		matches, err := filepath.Glob(filepath.Join(logsFilePathDir, "collect-i*"))
		if err == nil {
			logCollectFilePath := strings.Split(matches[0], "/")
			seelog.Infof("Found the logsbundle in %s", matches[0])
			return true, logCollectFilePath[len(logCollectFilePath)-1]
		}
		time.Sleep(5 * time.Second)
	}
	seelog.Errorf("Error while trying to find matches for %s, err: %v", filepath.Join(logsFilePathDir, "collect-i*"), err)
	return false, ""
}

func uploadECSLogsToS3(iamcredentials awscreds.Value, bucket, key, region string) error {
	//s3ClientCreator := factory.NewS3ClientCreator()
	cfg := aws.NewConfig().
		WithHTTPClient(httpclient.New(s3UploadTimeout, false)).
		WithCredentials(
			awscreds.NewStaticCredentials(iamcredentials.AccessKeyID, iamcredentials.SecretAccessKey,
				iamcredentials.SessionToken)).WithRegion(region)
	sess := session.Must(session.NewSession(cfg))

	// create the bucket if it does not exist
	svc := s3.New(sess)
	_, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err.(awserr.Error).Code() == s3.ErrCodeNoSuchBucket {
		seelog.Infof("%s bucket not found, creating it", bucket)
		_, err = svc.CreateBucket(&s3.CreateBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			return seelog.Errorf("Error creating the bucket %s, err: %v", bucket, err)
		}
	}
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	seelog.Debugf("uploading the logsbundle")
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   readLogsbundleTar(filepath.Join(logsFilePathDir, key)),
	})
	if err != nil {
		return seelog.Errorf("failed to upload file, %v", err)
	}
	seelog.Debugf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	/*s3Client, err := s3ClientCreator.NewS3ClientForECSLogsUpload(region, iamcredentials)
	err = agentS3.UploadFile("ecs-logs", "instanceId", []byte(`Hello`), s3UploadTimeout, s3Client)
	if err != nil {
		return err
	}*/
	return nil
}

func getPreSignedUrl(iamcredentials awscreds.Value, bucket, key, region string) string {
	cfg := aws.NewConfig().
		WithHTTPClient(httpclient.New(s3UploadTimeout, false)).
		WithCredentials(
			awscreds.NewStaticCredentials(iamcredentials.AccessKeyID, iamcredentials.SecretAccessKey,
				iamcredentials.SessionToken)).WithRegion(region)
	sess := session.Must(session.NewSession(cfg))

	// Create S3 service client
	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(1 * time.Minute)

	if err != nil {
		seelog.Errorf("Error to signing the request: %v", err)
	}

	return urlStr
}
