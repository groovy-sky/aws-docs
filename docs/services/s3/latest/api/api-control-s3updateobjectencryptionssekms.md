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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/S3UpdateObjectEncryptionSSEKMS)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/S3UpdateObjectEncryptionSSEKMS)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/S3UpdateObjectEncryptionSSEKMS)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3UpdateObjectEncryptionOperation

Scope
