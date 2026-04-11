# S3TableIntegrationSource

Represents a data source association with an S3 Table Integration, including its status
and metadata.

## Contents

**createdTimeStamp**

The timestamp when the data source association was created.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**dataSource**

The data source associated with the S3 Table Integration.

Type: [DataSource](api-datasource.md) object

Required: No

**identifier**

The unique identifier for this data source association.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**status**

The current status of the data source association.

Type: String

Valid Values: `ACTIVE | UNHEALTHY | FAILED | DATA_SOURCE_DELETE_IN_PROGRESS`

Required: No

**statusReason**

Additional information about the status of the data source association.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/s3tableintegrationsource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/s3tableintegrationsource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/s3tableintegrationsource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3DeliveryConfiguration

ScheduledQueryDestination

All content copied from https://docs.aws.amazon.com/.
