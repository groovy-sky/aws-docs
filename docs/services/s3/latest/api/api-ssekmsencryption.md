# SSEKMSEncryption

If `SSEKMS` is specified for `ObjectEncryption`, this data type specifies
the AWS KMS key Amazon Resource Name (ARN) to use and whether to use an S3 Bucket Key for
server-side encryption using AWS Key Management Service (AWS KMS) keys (SSE-KMS).

## Contents

**KMSKeyArn**

Specifies the AWS KMS key Amazon Resource Name (ARN) to use for the updated server-side encryption
type. Required if `ObjectEncryption` specifies `SSEKMS`.

###### Note

You must specify the full AWS KMS key ARN. The KMS key ID and KMS key alias aren't
supported.

Pattern: ( `arn:aws[-a-z0-9]*:kms:[-a-z0-9]*:[0-9]{12}:key/.+`)

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws[a-zA-Z0-9-]*:kms:[a-z0-9-]+:[0-9]{12}:key/[a-zA-Z0-9-]+`

Required: Yes

**BucketKeyEnabled**

Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption with server-side encryption
using AWS Key Management Service (AWS KMS) keys (SSE-KMS). If this value isn't specified, it defaults to `false`.
Setting this value to `true` causes Amazon S3 to use an S3 Bucket Key for object encryption with
SSE-KMS. For more information, see
[Using Amazon S3 Bucket Keys](../userguide/bucket-key.md) in the _Amazon S3 User Guide_.

Valid Values: `true` \| `false`

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/ssekmsencryption.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/ssekmsencryption.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/ssekmsencryption.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SseKmsEncryptedObjects

SSES3

All content copied from https://docs.aws.amazon.com/.
