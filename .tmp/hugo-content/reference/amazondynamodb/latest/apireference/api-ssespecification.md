# SSESpecification

Represents the settings used to enable server-side encryption.

## Contents

###### Note

In the following list, the required parameters are described first.

**Enabled**

Indicates whether server-side encryption is done using an AWS managed
key or an AWS owned key. If enabled (true), server-side encryption type
is set to `KMS` and an AWS managed key is used (AWS KMS charges apply). If disabled (false) or not specified, server-side
encryption is set to AWS owned key.

Type: Boolean

Required: No

**KMSMasterKeyId**

The AWS KMS key that should be used for the AWS KMS encryption.
To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN.
Note that you should only provide this parameter if the key is different from the
default DynamoDB key `alias/aws/dynamodb`.

Type: String

Required: No

**SSEType**

Server-side encryption type. The only supported value is:

- `KMS` \- Server-side encryption that uses AWS Key Management Service. The
key is stored in your account and is managed by AWS KMS (AWS KMS charges apply).

Type: String

Valid Values: `AES256 | KMS`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/ssespecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/ssespecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/ssespecification.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SSEDescription

StreamSpecification

All content copied from https://docs.aws.amazon.com/.
