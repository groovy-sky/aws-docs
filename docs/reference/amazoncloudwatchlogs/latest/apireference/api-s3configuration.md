# S3Configuration

Configuration for Amazon S3 destination where scheduled query results are delivered.

## Contents

**destinationIdentifier**

The Amazon S3 URI where query results are delivered. Must be a valid S3 URI format.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

Required: Yes

**roleArn**

The ARN of the IAM role that grants permissions to write query results to the specified
Amazon S3 destination.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**kmsKeyId**

The Amazon Resource Name (ARN) of the KMS encryption key. Must belong to the same AWS Region
as the destination Amazon S3 bucket.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**ownerAccountId**

The AWS accountId for the bucket owning account.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^\d{12}$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/s3configuration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/s3configuration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/s3configuration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResultField

S3DeliveryConfiguration

All content copied from https://docs.aws.amazon.com/.
