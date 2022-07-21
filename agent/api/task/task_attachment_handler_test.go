// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package task

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/aws/amazon-ecs-agent/agent/acs/model/ecsacs"
	"github.com/aws/amazon-ecs-agent/agent/api/serviceconnect"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

var (
	testSCContainerName             = "ecs-service-connect"
	testInboundListener             = "testInboundListener"
	testOutboundListener            = "testOutboundListenerName"
	testHost                        = "testHostName"
	testIngressPort                 = "9090"
	testIPv4                        = "172.31.21.40"
	testIPv4CIDR                    = "127.255.0.0/16"
	testIPv6                        = "abcd:dcba:1234:4321::"
	testIPv6CIDR                    = "2002::1234:abcd:ffff:c0a8:101/64"
	testIpv4ElasticNetworkInterface = &ecsacs.ElasticNetworkInterface{
		Ipv4Addresses: []*ecsacs.IPv4AddressAssignment{
			{
				Primary:        aws.Bool(true),
				PrivateAddress: aws.String(testIPv4),
			},
		},
	}
	testIpv6ElasticNetworkInterface = &ecsacs.ElasticNetworkInterface{
		Ipv6Addresses: []*ecsacs.IPv6AddressAssignment{
			{
				Address: aws.String(testIPv6),
			},
		},
	}
)

func stringToPointer(s string) *string { return &s }

func getTestcontainerFromACS(containerName, networkMode string) *ecsacs.Container {
	return &ecsacs.Container{
		Name: aws.String(containerName),
		DockerConfig: &ecsacs.DockerConfig{
			HostConfig: aws.String(fmt.Sprintf(
				`{"NetworkMode":"%s"}`, networkMode)),
		},
	}
}

func constructTestServiceConnectConfig(
	ingressPort,
	ingressListenerName,
	egressListenerName,
	egressIPv4Cidr,
	egressIPv6Cidr,
	dnsHostName,
	dnsAddress string) string {
	testIngressConfig := fmt.Sprintf(`\"ingressConfig\":[{\"interceptPort\":%s,\"listenerName\":\"%s\"}]`, ingressPort, ingressListenerName)
	testEgressConfig := fmt.Sprintf(`\"egressConfig\":{\"listenerName\":\"%s\",\"vip\":{\"ipv4Cidr\":\"%s\",\"ipv6Cidr\":\"%s\"}}`, egressListenerName, egressIPv4Cidr, egressIPv6Cidr)
	testDnsConfig := fmt.Sprintf(`\"dnsConfig\":[{\"hostname\":\"%s\",\"address\":\"%s\"}]`, dnsHostName, dnsAddress)
	testServiceConnectConfig := strings.Join([]string{`"{`,
		testEgressConfig + `,`,
		testDnsConfig + `,`,
		testIngressConfig,
		`}"`,
	}, "")
	unquotedSCConfig, _ := strconv.Unquote(testServiceConnectConfig)
	return unquotedSCConfig
}

func TestNewAttachmentHandlers(t *testing.T) {
	handlers := NewAttachmentHandlers()
	scHandler, err := getHandlerByType(serviceConnectAttachmentType, handlers)
	assert.Nil(t, err, "Should not return error")
	assert.NotNil(t, scHandler, "Should find service connect attachment type handler")
}

func TestHandleTaskAttachmentsWithServiceConnectAttachment(t *testing.T) {
	tt := []struct {
		testName                 string
		testServiceConnectConfig string
		shouldReturnError        bool
	}{
		{
			testName: "AWSVPC IPv6 enabled without error",
			testServiceConnectConfig: constructTestServiceConnectConfig(
				testIngressPort,
				testInboundListener,
				testOutboundListener,
				testIPv4CIDR,
				testIPv6CIDR,
				testHost,
				testIPv6,
			),
			shouldReturnError: false,
		},
		{
			testName: "AWSVPC IPv6 enabled with error",
			testServiceConnectConfig: constructTestServiceConnectConfig(
				testIngressPort,
				testInboundListener,
				testOutboundListener,
				"",
				testIPv6CIDR,
				testHost,
				testIPv6,
			),
			shouldReturnError: true,
		},
	}

	testExpectedSCConfig := &serviceconnect.Config{
		ContainerName: testSCContainerName,
		IngressConfig: []serviceconnect.IngressConfigEntry{
			{
				InterceptPort: aws.Uint16(9090),
				ListenerName:  testInboundListener,
			},
		},
		EgressConfig: &serviceconnect.EgressConfig{
			ListenerName: testOutboundListener,
			VIP: serviceconnect.VIP{
				IPV4CIDR: testIPv4CIDR,
				IPV6CIDR: testIPv6CIDR,
			},
		},
		DNSConfig: []serviceconnect.DNSConfigEntry{
			{
				HostName: testHost,
				Address:  testIPv6,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			testAcsTask := &ecsacs.Task{
				ElasticNetworkInterfaces: []*ecsacs.ElasticNetworkInterface{testIpv6ElasticNetworkInterface},
				Containers: []*ecsacs.Container{
					getTestcontainerFromACS(testSCContainerName, AWSVPCNetworkMode),
				},
				Attachments: []*ecsacs.Attachment{
					{
						AttachmentArn: stringToPointer("attachmentArn"),
						AttachmentProperties: []*ecsacs.AttachmentProperty{
							{
								Name:  stringToPointer(serviceconnect.GetServiceConnectConfigKey()),
								Value: stringToPointer(tc.testServiceConnectConfig),
							},
							{
								Name:  stringToPointer(serviceconnect.GetServiceConnectContainerNameKey()),
								Value: stringToPointer(testSCContainerName),
							},
						},
						AttachmentType: stringToPointer(serviceConnectAttachmentType),
					},
				},
				NetworkMode: stringToPointer(AWSVPCNetworkMode),
			}
			testTask := &Task{}
			testTask.NetworkMode = AWSVPCNetworkMode
			err := handleTaskAttachments(testAcsTask, testTask)
			if tc.shouldReturnError {
				assert.NotNil(t, err, "Should return error")
			} else {
				assert.Nil(t, err, "Should not return error")
				assert.NotNil(t, testTask.ServiceConnectConfig, "Should get valid service connect config from attachments")
				assert.Equal(t, testExpectedSCConfig, testTask.ServiceConnectConfig)
			}
		})
	}
}

func TestHandleTaskAttachmentsWithoutAttachment(t *testing.T) {
	testAcsTask := &ecsacs.Task{
		ElasticNetworkInterfaces: []*ecsacs.ElasticNetworkInterface{testIpv4ElasticNetworkInterface},
		Containers: []*ecsacs.Container{
			getTestcontainerFromACS("C1", BridgeNetworkMode),
		},
		NetworkMode: stringToPointer(BridgeNetworkMode),
	}
	testTask := &Task{}
	err := handleTaskAttachments(testAcsTask, testTask)
	assert.Nil(t, err, "Should not return error")
	assert.Nil(t, testTask.ServiceConnectConfig, "Should not return service connect config from attachments")
}
