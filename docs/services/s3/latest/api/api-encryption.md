# Encryption

Contains the type of server-side encryption used.

## Contents

**EncryptionType**

The server-side encryption algorithm used when storing job results in Amazon S3 (for example, AES256,
`aws:kms`).

Type: String

Valid Values: `AES256 | aws:fsx | aws:kms | aws:kms:dsse`

Required: Yes

**KMSContext**

If the encryption type is `aws:kms`, this optional value can be used to specify the
encryption context for the restore results.

Type: String

Required: No

**KMSKeyId**

If the encryption type is `aws:kms`, this optional value specifies the ID of the
symmetric encryption customer managed key to use for encryption of job results. Amazon S3 only supports symmetric
encryption KMS keys. For more information, see [Asymmetric keys in AWS KMS](../../../kms/latest/developerguide/symmetric-asymmetric.md) in the
_AWS Key Management Service Developer Guide_.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/encryption.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/encryption.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/encryption.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DestinationResult

EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
