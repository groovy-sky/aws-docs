# S3UpdateObjectEncryptionSSEKMS

If `SSEKMS` is specified for `UpdateObjectEncryption`,
this data type specifies the AWS KMS key Amazon Resource Name (ARN) to use
and whether to use an S3 Bucket Key for server-side encryption using
AWS Key Management Service (AWS KMS) keys (SSE-KMS).

## Contents

**KMSKeyArn**

Specifies the AWS KMS key Amazon Resource Name (ARN) to use for the
updated server-side encryption type. Required if `UpdateObjectEncryption`
specifies `SSEKMS`.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws[a-zA-Z0-9-]*:kms:[a-z0-9-]+:[0-9]{12}:key/[a-zA-Z0-9-]+`

Required: Yes

**BucketKeyEnabled**

Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption
with server-side encryption using AWS Key Management Service (AWS KMS) keys (SSE-KMS). If this
value isn't specified, it defaults to `false`. Setting this value
to `true` causes Amazon S3 to use an S3 Bucket Key for update object
encryption with SSE-KMS.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/s3updateobjectencryptionssekms.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/s3updateobjectencryptionssekms.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/s3updateobjectencryptionssekms.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3UpdateObjectEncryptionOperation

Scope

All content copied from https://docs.aws.amazon.com/.
