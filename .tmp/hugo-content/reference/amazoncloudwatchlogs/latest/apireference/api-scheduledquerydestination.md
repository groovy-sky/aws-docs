# ScheduledQueryDestination

Information about a destination where scheduled query results are processed, including
processing status and any error messages.

## Contents

**destinationIdentifier**

The identifier for the destination where results are delivered.

Type: String

Required: No

**destinationType**

The type of destination for query results.

Type: String

Valid Values: `S3`

Required: No

**errorMessage**

Error message if destination processing failed.

Type: String

Required: No

**processedIdentifier**

The identifier of the processed result at the destination.

Type: String

Required: No

**status**

The processing status of the destination delivery.

Type: String

Valid Values: `IN_PROGRESS | CLIENT_ERROR | FAILED | COMPLETE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/scheduledquerydestination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/scheduledquerydestination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/scheduledquerydestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3TableIntegrationSource

ScheduledQuerySummary

All content copied from https://docs.aws.amazon.com/.
