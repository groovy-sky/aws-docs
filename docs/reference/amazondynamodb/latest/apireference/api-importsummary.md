# ImportSummary

Summary information about the source file for the import.

## Contents

###### Note

In the following list, the required parameters are described first.

**CloudWatchLogGroupArn**

The Amazon Resource Number (ARN) of the Cloudwatch Log Group associated with this
import task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**EndTime**

The time at which this import task ended. (Does this include the successful complete
creation of the table it was imported to?)

Type: Timestamp

Required: No

**ImportArn**

The Amazon Resource Number (ARN) corresponding to the import request.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**ImportStatus**

The status of the import operation.

Type: String

Valid Values: `IN_PROGRESS | COMPLETED | CANCELLING | CANCELLED | FAILED`

Required: No

**InputFormat**

The format of the source data. Valid values are `CSV`,
`DYNAMODB_JSON` or `ION`.

Type: String

Valid Values: `DYNAMODB_JSON | ION | CSV`

Required: No

**S3BucketSource**

The path and S3 bucket of the source file that is being imported. This includes the
S3Bucket (required), S3KeyPrefix (optional) and S3BucketOwner (optional if the bucket is
owned by the requester).

Type: [S3BucketSource](api-s3bucketsource.md) object

Required: No

**StartTime**

The time at which this import task began.

Type: Timestamp

Required: No

**TableArn**

The Amazon Resource Number (ARN) of the table being imported into.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/importsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/importsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/importsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalTableWitnessGroupUpdate

ImportTableDescription

All content copied from https://docs.aws.amazon.com/.
