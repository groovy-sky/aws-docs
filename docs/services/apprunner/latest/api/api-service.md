# Service

Describes an AWS App Runner service. It can describe a service in any state, including deleted services.

This type contains the full information about a service, including configuration details. It's returned by the [CreateService](api-createservice.md), [DescribeService](api-describeservice.md), and [DeleteService](api-deleteservice.md) actions. A subset of this
information is returned by the [ListServices](api-listservices.md) action using the [ServiceSummary](api-servicesummary.md) type.

## Contents

**AutoScalingConfigurationSummary**

Summary information for the App Runner automatic scaling configuration resource that's associated with this service.

Type: [AutoScalingConfigurationSummary](api-autoscalingconfigurationsummary.md) object

Required: Yes

**CreatedAt**

The time when the App Runner service was created. It's in the Unix time stamp format.

Type: Timestamp

Required: Yes

**InstanceConfiguration**

The runtime configuration of instances (scaling units) of this service.

Type: [InstanceConfiguration](api-instanceconfiguration.md) object

Required: Yes

**NetworkConfiguration**

Configuration settings related to network traffic of the web application that this service runs.

Type: [NetworkConfiguration](api-networkconfiguration.md) object

Required: Yes

**ServiceArn**

The Amazon Resource Name (ARN) of this service.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

**ServiceId**

An ID that App Runner generated for this service. It's unique within the AWS Region.

Type: String

Length Constraints: Fixed length of 32.

Pattern: `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}`

Required: Yes

**ServiceName**

The customer-provided service name.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 40.

Pattern: `[A-Za-z0-9][A-Za-z0-9-_]{3,39}`

Required: Yes

**SourceConfiguration**

The source deployed to the App Runner service. It can be a code or an image repository.

Type: [SourceConfiguration](api-sourceconfiguration.md) object

Required: Yes

**Status**

The current state of the App Runner service. These particular values mean the following.

- `CREATE_FAILED` – The service failed to create. The failed service isn't usable, and still counts towards your service quota. To
troubleshoot this failure, read the failure events and logs, change any parameters that need to be fixed, and rebuild your service using
`UpdateService`.

- `DELETE_FAILED` – The service failed to delete and can't be successfully recovered. Retry the service deletion call to ensure
that all related resources are removed.

Type: String

Valid Values: `CREATE_FAILED | RUNNING | DELETED | DELETE_FAILED | PAUSED | OPERATION_IN_PROGRESS`

Required: Yes

**UpdatedAt**

The time when the App Runner service was last updated at. It's in the Unix time stamp format.

Type: Timestamp

Required: Yes

**DeletedAt**

The time when the App Runner service was deleted. It's in the Unix time stamp format.

Type: Timestamp

Required: No

**EncryptionConfiguration**

The encryption key that App Runner uses to encrypt the service logs and the copy of the source repository that App Runner maintains for the service. It can be
either a customer-provided encryption key or an AWS managed key.

Type: [EncryptionConfiguration](api-encryptionconfiguration.md) object

Required: No

**HealthCheckConfiguration**

The settings for the health check that App Runner performs to monitor the health of this service.

Type: [HealthCheckConfiguration](api-healthcheckconfiguration.md) object

Required: No

**ObservabilityConfiguration**

The observability configuration of this service.

Type: [ServiceObservabilityConfiguration](api-serviceobservabilityconfiguration.md) object

Required: No

**ServiceUrl**

A subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/service.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/service.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/service.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OperationSummary

ServiceObservabilityConfiguration

All content copied from https://docs.aws.amazon.com/.
