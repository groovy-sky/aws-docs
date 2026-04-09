# ServiceSummary

Provides summary information for an AWS App Runner service.

This type contains limited information about a service. It doesn't include configuration details. It's returned by the [ListServices](api-listservices.md) action. Complete service information is returned by the [CreateService](api-createservice.md), [DescribeService](api-describeservice.md), and [DeleteService](api-deleteservice.md) actions using the [Service](api-service.md) type.

## Contents

**CreatedAt**

The time when the App Runner service was created. It's in the Unix time stamp format.

Type: Timestamp

Required: No

**ServiceArn**

The Amazon Resource Name (ARN) of this service.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**ServiceId**

An ID that App Runner generated for this service. It's unique within the AWS Region.

Type: String

Length Constraints: Fixed length of 32.

Pattern: `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}`

Required: No

**ServiceName**

The customer-provided service name.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 40.

Pattern: `[A-Za-z0-9][A-Za-z0-9-_]{3,39}`

Required: No

**ServiceUrl**

A subdomain URL that App Runner generated for this service. You can use this URL to access your service web application.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**Status**

The current state of the App Runner service. These particular values mean the following.

- `CREATE_FAILED` – The service failed to create. The failed service isn't usable, and still counts towards your service quota. To
troubleshoot this failure, read the failure events and logs, change any parameters that need to be fixed, and rebuild your service using
`UpdateService`.

- `DELETE_FAILED` – The service failed to delete and can't be successfully recovered. Retry the service deletion call to ensure
that all related resources are removed.

Type: String

Valid Values: `CREATE_FAILED | RUNNING | DELETED | DELETE_FAILED | PAUSED | OPERATION_IN_PROGRESS`

Required: No

**UpdatedAt**

The time when the App Runner service was last updated. It's in theUnix time stamp format.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/servicesummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/servicesummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/servicesummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceObservabilityConfiguration

SourceCodeVersion

All content copied from https://docs.aws.amazon.com/.
