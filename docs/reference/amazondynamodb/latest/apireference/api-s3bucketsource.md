# S3BucketSource

The S3 bucket that is being imported from.

## Contents

###### Note

In the following list, the required parameters are described first.

**S3Bucket**

The S3 bucket that is being imported from.

Type: String

Length Constraints: Maximum length of 255.

Pattern: `^[a-z0-9A-Z]+[\.\-\w]*[a-z0-9A-Z]+$`

Required: Yes

**S3BucketOwner**

The account number of the S3 bucket that is being imported from. If the bucket is
owned by the requester this is optional.

Type: String

Pattern: `[0-9]{12}`

Required: No

**S3KeyPrefix**

The key prefix shared by all S3 Objects that are being imported.

Type: String

Length Constraints: Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/s3bucketsource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/s3bucketsource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/s3bucketsource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestoreSummary

SourceTableDetails

All content copied from https://docs.aws.amazon.com/.
