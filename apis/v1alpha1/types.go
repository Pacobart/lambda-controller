// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Limits that are related to concurrency and storage. All file and storage
// sizes are in bytes.
type AccountLimit struct {
	CodeSizeUnzipped *int64 `json:"codeSizeUnzipped,omitempty"`
	CodeSizeZipped   *int64 `json:"codeSizeZipped,omitempty"`
	TotalCodeSize    *int64 `json:"totalCodeSize,omitempty"`
}

// The number of functions and amount of storage in use.
type AccountUsage struct {
	FunctionCount *int64 `json:"functionCount,omitempty"`
	TotalCodeSize *int64 `json:"totalCodeSize,omitempty"`
}

// Provides configuration information about a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
type AliasConfiguration struct {
	AliasARN        *string `json:"aliasARN,omitempty"`
	Description     *string `json:"description,omitempty"`
	FunctionVersion *string `json:"functionVersion,omitempty"`
	Name            *string `json:"name,omitempty"`
	RevisionID      *string `json:"revisionID,omitempty"`
	// The traffic-shifting (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
	// configuration of a Lambda function alias.
	RoutingConfig *AliasRoutingConfiguration `json:"routingConfig,omitempty"`
}

// The traffic-shifting (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
// configuration of a Lambda function alias.
type AliasRoutingConfiguration struct {
	AdditionalVersionWeights map[string]*float64 `json:"additionalVersionWeights,omitempty"`
}

// List of signing profiles that can sign a code package.
type AllowedPublishers struct {
	SigningProfileVersionARNs []*string `json:"signingProfileVersionARNs,omitempty"`
}

// Specific configuration settings for an Amazon Managed Streaming for Apache
// Kafka (Amazon MSK) event source.
type AmazonManagedKafkaEventSourceConfig struct {
	ConsumerGroupID *string `json:"consumerGroupID,omitempty"`
}

// The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
// settings for your Lambda function URL. Use CORS to grant access to your function
// URL from any origin. You can also use CORS to control access for specific
// HTTP headers and methods in requests to your function URL.
type CORS struct {
	AllowCredentials *bool     `json:"allowCredentials,omitempty"`
	AllowHeaders     []*string `json:"allowHeaders,omitempty"`
	AllowMethods     []*string `json:"allowMethods,omitempty"`
	AllowOrigins     []*string `json:"allowOrigins,omitempty"`
	ExposeHeaders    []*string `json:"exposeHeaders,omitempty"`
	MaxAge           *int64    `json:"maxAge,omitempty"`
}

// Details about a Code signing configuration (https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html).
type CodeSigningConfig_SDK struct {
	// List of signing profiles that can sign a code package.
	AllowedPublishers    *AllowedPublishers `json:"allowedPublishers,omitempty"`
	CodeSigningConfigARN *string            `json:"codeSigningConfigARN,omitempty"`
	CodeSigningConfigID  *string            `json:"codeSigningConfigID,omitempty"`
	// Code signing configuration policies (https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html#config-codesigning-policies)
	// specify the validation failure action for signature mismatch or expiry.
	CodeSigningPolicies *CodeSigningPolicies `json:"codeSigningPolicies,omitempty"`
	Description         *string              `json:"description,omitempty"`
	LastModified        *string              `json:"lastModified,omitempty"`
}

// Code signing configuration policies (https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html#config-codesigning-policies)
// specify the validation failure action for signature mismatch or expiry.
type CodeSigningPolicies struct {
	UntrustedArtifactOnDeployment *string `json:"untrustedArtifactOnDeployment,omitempty"`
}

// The dead-letter queue (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq)
// for failed asynchronous invocations.
type DeadLetterConfig struct {
	TargetARN *string `json:"targetARN,omitempty"`
}

// A configuration object that specifies the destination of an event after Lambda
// processes it.
type DestinationConfig struct {
	// A destination for events that failed processing.
	OnFailure *OnFailure `json:"onFailure,omitempty"`
	// A destination for events that were processed successfully.
	OnSuccess *OnSuccess `json:"onSuccess,omitempty"`
}

// A function's environment variable settings. You can use environment variables
// to adjust your function's behavior without updating code. An environment
// variable is a pair of strings that are stored in a function's version-specific
// configuration.
type Environment struct {
	Variables map[string]*string `json:"variables,omitempty"`
}

// Error messages for environment variables that couldn't be applied.
type EnvironmentError struct {
	ErrorCode *string `json:"errorCode,omitempty"`
	Message   *string `json:"message,omitempty"`
}

// The results of an operation to update or read environment variables. If the
// operation is successful, the response contains the environment variables.
// If it failed, the response contains details about the error.
type EnvironmentResponse struct {
	// Error messages for environment variables that couldn't be applied.
	Error     *EnvironmentError  `json:"error,omitempty"`
	Variables map[string]*string `json:"variables,omitempty"`
}

// The size of the function’s /tmp directory in MB. The default value is 512,
// but can be any whole number between 512 and 10240 MB.
type EphemeralStorage struct {
	Size *int64 `json:"size,omitempty"`
}

// A mapping between an Amazon Web Services resource and a Lambda function.
// For details, see CreateEventSourceMapping.
type EventSourceMappingConfiguration struct {
	// Specific configuration settings for an Amazon Managed Streaming for Apache
	// Kafka (Amazon MSK) event source.
	AmazonManagedKafkaEventSourceConfig *AmazonManagedKafkaEventSourceConfig `json:"amazonManagedKafkaEventSourceConfig,omitempty"`
	BatchSize                           *int64                               `json:"batchSize,omitempty"`
	BisectBatchOnFunctionError          *bool                                `json:"bisectBatchOnFunctionError,omitempty"`
	// A configuration object that specifies the destination of an event after Lambda
	// processes it.
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty"`
	EventSourceARN    *string            `json:"eventSourceARN,omitempty"`
	// An object that contains the filters for an event source.
	FilterCriteria                 *FilterCriteria `json:"filterCriteria,omitempty"`
	FunctionARN                    *string         `json:"functionARN,omitempty"`
	FunctionResponseTypes          []*string       `json:"functionResponseTypes,omitempty"`
	LastModified                   *metav1.Time    `json:"lastModified,omitempty"`
	LastProcessingResult           *string         `json:"lastProcessingResult,omitempty"`
	MaximumBatchingWindowInSeconds *int64          `json:"maximumBatchingWindowInSeconds,omitempty"`
	MaximumRecordAgeInSeconds      *int64          `json:"maximumRecordAgeInSeconds,omitempty"`
	MaximumRetryAttempts           *int64          `json:"maximumRetryAttempts,omitempty"`
	ParallelizationFactor          *int64          `json:"parallelizationFactor,omitempty"`
	Queues                         []*string       `json:"queues,omitempty"`
	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource *SelfManagedEventSource `json:"selfManagedEventSource,omitempty"`
	// Specific configuration settings for a self-managed Apache Kafka event source.
	SelfManagedKafkaEventSourceConfig *SelfManagedKafkaEventSourceConfig `json:"selfManagedKafkaEventSourceConfig,omitempty"`
	SourceAccessConfigurations        []*SourceAccessConfiguration       `json:"sourceAccessConfigurations,omitempty"`
	StartingPosition                  *string                            `json:"startingPosition,omitempty"`
	StartingPositionTimestamp         *metav1.Time                       `json:"startingPositionTimestamp,omitempty"`
	State                             *string                            `json:"state,omitempty"`
	StateTransitionReason             *string                            `json:"stateTransitionReason,omitempty"`
	Topics                            []*string                          `json:"topics,omitempty"`
	TumblingWindowInSeconds           *int64                             `json:"tumblingWindowInSeconds,omitempty"`
	UUID                              *string                            `json:"uuid,omitempty"`
}

// Details about the connection between a Lambda function and an Amazon EFS
// file system (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html).
type FileSystemConfig struct {
	ARN            *string `json:"arn,omitempty"`
	LocalMountPath *string `json:"localMountPath,omitempty"`
}

// A structure within a FilterCriteria object that defines an event filtering
// pattern.
type Filter struct {
	Pattern *string `json:"pattern,omitempty"`
}

// An object that contains the filters for an event source.
type FilterCriteria struct {
	Filters []*Filter `json:"filters,omitempty"`
}

// The code for the Lambda function. You can specify either an object in Amazon
// S3, upload a .zip file archive deployment package directly, or specify the
// URI of a container image.
type FunctionCode struct {
	ImageURI        *string `json:"imageURI,omitempty"`
	S3Bucket        *string `json:"s3Bucket,omitempty"`
	S3Key           *string `json:"s3Key,omitempty"`
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty"`
	ZipFile         []byte  `json:"zipFile,omitempty"`
}

// Details about a function's deployment package.
type FunctionCodeLocation struct {
	ImageURI         *string `json:"imageURI,omitempty"`
	Location         *string `json:"location,omitempty"`
	RepositoryType   *string `json:"repositoryType,omitempty"`
	ResolvedImageURI *string `json:"resolvedImageURI,omitempty"`
}

// Details about a function's configuration.
type FunctionConfiguration struct {
	Architectures []*string `json:"architectures,omitempty"`
	CodeSHA256    *string   `json:"codeSHA256,omitempty"`
	CodeSize      *int64    `json:"codeSize,omitempty"`
	// The dead-letter queue (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq)
	// for failed asynchronous invocations.
	DeadLetterConfig *DeadLetterConfig `json:"deadLetterConfig,omitempty"`
	Description      *string           `json:"description,omitempty"`
	// The results of an operation to update or read environment variables. If the
	// operation is successful, the response contains the environment variables.
	// If it failed, the response contains details about the error.
	Environment *EnvironmentResponse `json:"environment,omitempty"`
	// The size of the function’s /tmp directory in MB. The default value is 512,
	// but can be any whole number between 512 and 10240 MB.
	EphemeralStorage  *EphemeralStorage   `json:"ephemeralStorage,omitempty"`
	FileSystemConfigs []*FileSystemConfig `json:"fileSystemConfigs,omitempty"`
	FunctionARN       *string             `json:"functionARN,omitempty"`
	FunctionName      *string             `json:"functionName,omitempty"`
	Handler           *string             `json:"handler,omitempty"`
	// Response to GetFunctionConfiguration request.
	ImageConfigResponse        *ImageConfigResponse `json:"imageConfigResponse,omitempty"`
	KMSKeyARN                  *string              `json:"kmsKeyARN,omitempty"`
	LastModified               *string              `json:"lastModified,omitempty"`
	LastUpdateStatus           *string              `json:"lastUpdateStatus,omitempty"`
	LastUpdateStatusReason     *string              `json:"lastUpdateStatusReason,omitempty"`
	LastUpdateStatusReasonCode *string              `json:"lastUpdateStatusReasonCode,omitempty"`
	Layers                     []*Layer             `json:"layers,omitempty"`
	MasterARN                  *string              `json:"masterARN,omitempty"`
	MemorySize                 *int64               `json:"memorySize,omitempty"`
	PackageType                *string              `json:"packageType,omitempty"`
	RevisionID                 *string              `json:"revisionID,omitempty"`
	Role                       *string              `json:"role,omitempty"`
	Runtime                    *string              `json:"runtime,omitempty"`
	SigningJobARN              *string              `json:"signingJobARN,omitempty"`
	SigningProfileVersionARN   *string              `json:"signingProfileVersionARN,omitempty"`
	State                      *string              `json:"state,omitempty"`
	StateReason                *string              `json:"stateReason,omitempty"`
	StateReasonCode            *string              `json:"stateReasonCode,omitempty"`
	Timeout                    *int64               `json:"timeout,omitempty"`
	// The function's X-Ray tracing configuration.
	TracingConfig *TracingConfigResponse `json:"tracingConfig,omitempty"`
	Version       *string                `json:"version,omitempty"`
	// The VPC security groups and subnets that are attached to a Lambda function.
	VPCConfig *VPCConfigResponse `json:"vpcConfig,omitempty"`
}

type FunctionEventInvokeConfig struct {
	// A configuration object that specifies the destination of an event after Lambda
	// processes it.
	DestinationConfig *DestinationConfig `json:"destinationConfig,omitempty"`
	FunctionARN       *string            `json:"functionARN,omitempty"`
	LastModified      *metav1.Time       `json:"lastModified,omitempty"`
}

// Details about a Lambda function URL.
type FunctionURLConfig_SDK struct {
	AuthType *string `json:"authType,omitempty"`
	// The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
	// settings for your Lambda function URL. Use CORS to grant access to your function
	// URL from any origin. You can also use CORS to control access for specific
	// HTTP headers and methods in requests to your function URL.
	CORS             *CORS   `json:"cors,omitempty"`
	CreationTime     *string `json:"creationTime,omitempty"`
	FunctionARN      *string `json:"functionARN,omitempty"`
	FunctionURL      *string `json:"functionURL,omitempty"`
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`
}

// Configuration values that override the container image Dockerfile settings.
// See Container settings (https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms).
type ImageConfig struct {
	Command          []*string `json:"command,omitempty"`
	EntryPoint       []*string `json:"entryPoint,omitempty"`
	WorkingDirectory *string   `json:"workingDirectory,omitempty"`
}

// Error response to GetFunctionConfiguration.
type ImageConfigError struct {
	ErrorCode *string `json:"errorCode,omitempty"`
	Message   *string `json:"message,omitempty"`
}

// Response to GetFunctionConfiguration request.
type ImageConfigResponse struct {
	// Error response to GetFunctionConfiguration.
	Error *ImageConfigError `json:"error,omitempty"`
	// Configuration values that override the container image Dockerfile settings.
	// See Container settings (https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms).
	ImageConfig *ImageConfig `json:"imageConfig,omitempty"`
}

// An Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
type Layer struct {
	ARN                      *string `json:"arn,omitempty"`
	CodeSize                 *int64  `json:"codeSize,omitempty"`
	SigningJobARN            *string `json:"signingJobARN,omitempty"`
	SigningProfileVersionARN *string `json:"signingProfileVersionARN,omitempty"`
}

// A ZIP archive that contains the contents of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// You can specify either an Amazon S3 location, or upload a layer archive directly.
type LayerVersionContentInput struct {
	S3Bucket        *string `json:"s3Bucket,omitempty"`
	S3Key           *string `json:"s3Key,omitempty"`
	S3ObjectVersion *string `json:"s3ObjectVersion,omitempty"`
	ZipFile         []byte  `json:"zipFile,omitempty"`
}

// Details about a version of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
type LayerVersionContentOutput struct {
	CodeSHA256               *string `json:"codeSHA256,omitempty"`
	CodeSize                 *int64  `json:"codeSize,omitempty"`
	Location                 *string `json:"location,omitempty"`
	SigningJobARN            *string `json:"signingJobARN,omitempty"`
	SigningProfileVersionARN *string `json:"signingProfileVersionARN,omitempty"`
}

// Details about a version of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
type LayerVersionsListItem struct {
	CreatedDate     *string `json:"createdDate,omitempty"`
	Description     *string `json:"description,omitempty"`
	LayerVersionARN *string `json:"layerVersionARN,omitempty"`
}

// A destination for events that failed processing.
type OnFailure struct {
	Destination *string `json:"destination,omitempty"`
}

// A destination for events that were processed successfully.
type OnSuccess struct {
	Destination *string `json:"destination,omitempty"`
}

// Details about the provisioned concurrency configuration for a function alias
// or version.
type ProvisionedConcurrencyConfigListItem struct {
	FunctionARN  *string `json:"functionARN,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	StatusReason *string `json:"statusReason,omitempty"`
}

type PutFunctionConcurrencyOutput struct {
	ReservedConcurrentExecutions *int64 `json:"reservedConcurrentExecutions,omitempty"`
}

// The self-managed Apache Kafka cluster for your event source.
type SelfManagedEventSource struct {
	Endpoints map[string][]*string `json:"endpoints,omitempty"`
}

// Specific configuration settings for a self-managed Apache Kafka event source.
type SelfManagedKafkaEventSourceConfig struct {
	ConsumerGroupID *string `json:"consumerGroupID,omitempty"`
}

// To secure and define access to your event source, you can specify the authentication
// protocol, VPC components, or virtual host.
type SourceAccessConfiguration struct {
	Type *string `json:"type_,omitempty"`
	URI  *string `json:"uRI,omitempty"`
}

// The function's X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html)
// tracing configuration. To sample and record incoming requests, set Mode to
// Active.
type TracingConfig struct {
	Mode *string `json:"mode,omitempty"`
}

// The function's X-Ray tracing configuration.
type TracingConfigResponse struct {
	Mode *string `json:"mode,omitempty"`
}

// The VPC security groups and subnets that are attached to a Lambda function.
// For more information, see VPC Settings (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
type VPCConfig struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	SubnetIDs        []*string `json:"subnetIDs,omitempty"`
}

// The VPC security groups and subnets that are attached to a Lambda function.
type VPCConfigResponse struct {
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	SubnetIDs        []*string `json:"subnetIDs,omitempty"`
	VPCID            *string   `json:"vpcID,omitempty"`
}
