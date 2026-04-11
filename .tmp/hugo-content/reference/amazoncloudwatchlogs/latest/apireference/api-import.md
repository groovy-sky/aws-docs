# Import

An import job to move data from CloudTrail Event Data Store to CloudWatch.

## Contents

**creationTime**

The timestamp when the import task was created, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**errorMessage**

Error message related to any failed imports

Type: String

Required: No

**importDestinationArn**

The ARN of the managed CloudWatch Logs log group where the events are being imported to.

Type: String

Required: No

**importFilter**

The filter criteria used for this import task.

Type: [ImportFilter](api-importfilter.md) object

Required: No

**importId**

The unique identifier of the import task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\-a-zA-Z0-9]+`

Required: No

**importSourceArn**

The ARN of the CloudTrail Lake Event Data Store being imported from.

Type: String

Required: No

**importStatistics**

Statistics about the import progress

Type: [ImportStatistics](api-importstatistics.md) object

Required: No

**importStatus**

The current status of the import task. Valid values are IN\_PROGRESS, CANCELLED, COMPLETED and FAILED.

Type: String

Valid Values: `IN_PROGRESS | CANCELLED | COMPLETED | FAILED`

Required: No

**lastUpdatedTime**

The timestamp when the import task was last updated, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.

Type: Long

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/import.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/import.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/import.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GroupingIdentifier

ImportBatch

All content copied from https://docs.aws.amazon.com/.
