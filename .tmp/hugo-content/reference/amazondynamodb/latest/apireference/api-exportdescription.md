# ExportDescription

Represents the properties of the exported table.

## Contents

###### Note

In the following list, the required parameters are described first.

**BilledSizeBytes**

The billable size of the table export.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**ClientToken**

The client token that was provided for the export task. A client token makes calls to
`ExportTableToPointInTimeInput` idempotent, meaning that multiple
identical calls have the same effect as one single call.

Type: String

Pattern: `^[^\$]+$`

Required: No

**EndTime**

The time at which the export task completed.

Type: Timestamp

Required: No

**ExportArn**

The Amazon Resource Name (ARN) of the table export.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**ExportFormat**

The format of the exported data. Valid values for `ExportFormat` are
`DYNAMODB_JSON` or `ION`.

Type: String

Valid Values: `DYNAMODB_JSON | ION`

Required: No

**ExportManifest**

The name of the manifest file for the export task.

Type: String

Required: No

**ExportStatus**

Export can be in one of the following states: IN\_PROGRESS, COMPLETED, or
FAILED.

Type: String

Valid Values: `IN_PROGRESS | COMPLETED | FAILED`

Required: No

**ExportTime**

Point in time from which table data was exported.

Type: Timestamp

Required: No

**ExportType**

The type of export that was performed. Valid values are `FULL_EXPORT` or
`INCREMENTAL_EXPORT`.

Type: String

Valid Values: `FULL_EXPORT | INCREMENTAL_EXPORT`

Required: No

**FailureCode**

Status code for the result of the failed export.

Type: String

Required: No

**FailureMessage**

Export failure reason description.

Type: String

Required: No

**IncrementalExportSpecification**

Optional object containing the parameters specific to an incremental export.

Type: [IncrementalExportSpecification](api-incrementalexportspecification.md) object

Required: No

**ItemCount**

The number of items exported.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**S3Bucket**

The name of the Amazon S3 bucket containing the export.

Type: String

Length Constraints: Maximum length of 255.

Pattern: `^[a-z0-9A-Z]+[\.\-\w]*[a-z0-9A-Z]+$`

Required: No

**S3BucketOwner**

The ID of the AWS account that owns the bucket containing the
export.

Type: String

Pattern: `[0-9]{12}`

Required: No

**S3Prefix**

The Amazon S3 bucket prefix used as the file name and path of the exported
snapshot.

Type: String

Length Constraints: Maximum length of 1024.

Required: No

**S3SseAlgorithm**

Type of encryption used on the bucket where export data is stored. Valid values for
`S3SseAlgorithm` are:

- `AES256` \- server-side encryption with Amazon S3 managed
keys

- `KMS` \- server-side encryption with AWS KMS managed
keys

Type: String

Valid Values: `AES256 | KMS`

Required: No

**S3SseKmsKeyId**

The ID of the AWS KMS managed key used to encrypt the S3 bucket where
export data is stored (if applicable).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**StartTime**

The time at which the export task began.

Type: Timestamp

Required: No

**TableArn**

The Amazon Resource Name (ARN) of the table that was exported.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**TableId**

Unique ID of the table that was exported.

Type: String

Pattern: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/exportdescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/exportdescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/exportdescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpectedAttributeValue

ExportSummary

All content copied from https://docs.aws.amazon.com/.
