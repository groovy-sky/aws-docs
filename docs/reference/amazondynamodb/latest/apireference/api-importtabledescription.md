# ImportTableDescription

Represents the properties of the table being imported into.

## Contents

###### Note

In the following list, the required parameters are described first.

**ClientToken**

The client token that was provided for the import task. Reusing the client token on
retry makes a call to `ImportTable` idempotent.

Type: String

Pattern: `^[^\$]+$`

Required: No

**CloudWatchLogGroupArn**

The Amazon Resource Number (ARN) of the Cloudwatch Log Group associated with the
target table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**EndTime**

The time at which the creation of the table associated with this import task
completed.

Type: Timestamp

Required: No

**ErrorCount**

The number of errors occurred on importing the source file into the target table.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**FailureCode**

The error code corresponding to the failure that the import job ran into during
execution.

Type: String

Required: No

**FailureMessage**

The error message corresponding to the failure that the import job ran into during
execution.

Type: String

Required: No

**ImportArn**

The Amazon Resource Number (ARN) corresponding to the import request.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**ImportedItemCount**

The number of items successfully imported into the new table.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**ImportStatus**

The status of the import.

Type: String

Valid Values: `IN_PROGRESS | COMPLETED | CANCELLING | CANCELLED | FAILED`

Required: No

**InputCompressionType**

The compression options for the data that has been imported into the target table.
The values are NONE, GZIP, or ZSTD.

Type: String

Valid Values: `GZIP | ZSTD | NONE`

Required: No

**InputFormat**

The format of the source data going into the target table.

Type: String

Valid Values: `DYNAMODB_JSON | ION | CSV`

Required: No

**InputFormatOptions**

The format options for the data that was imported into the target table. There is one
value, CsvOption.

Type: [InputFormatOptions](api-inputformatoptions.md) object

Required: No

**ProcessedItemCount**

The total number of items processed from the source file.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**ProcessedSizeBytes**

The total size of data processed from the source file, in Bytes.

Type: Long

Required: No

**S3BucketSource**

Values for the S3 bucket the source file is imported from. Includes bucket name
(required), key prefix (optional) and bucket account owner ID (optional).

Type: [S3BucketSource](api-s3bucketsource.md) object

Required: No

**StartTime**

The time when this import task started.

Type: Timestamp

Required: No

**TableArn**

The Amazon Resource Number (ARN) of the table being imported into.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**TableCreationParameters**

The parameters for the new table that is being imported into.

Type: [TableCreationParameters](api-tablecreationparameters.md) object

Required: No

**TableId**

The table id corresponding to the table created by import table process.

Type: String

Pattern: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/importtabledescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/importtabledescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/importtabledescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportSummary

IncrementalExportSpecification

All content copied from https://docs.aws.amazon.com/.
