# Initiator

Container element that identifies who initiated the multipart upload.

## Contents

**DisplayName**

###### Note

This functionality is not supported for directory buckets.

Type: String

Required: No

**ID**

If the principal is an AWS account, it provides the Canonical User ID. If the principal is an
IAM User, it provides a user ARN value.

###### Note

**Directory buckets** \- If the principal is an AWS account,
it provides the AWS account ID. If the principal is an IAM User, it provides a user ARN
value.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/initiator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/initiator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/initiator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IndexDocument

InputSerialization

All content copied from https://docs.aws.amazon.com/.
