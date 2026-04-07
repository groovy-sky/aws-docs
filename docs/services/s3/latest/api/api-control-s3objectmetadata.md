# S3ObjectMetadata

## Contents

**CacheControl**

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ContentDisposition**

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ContentEncoding**

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ContentLanguage**

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ContentLength**

_This member has been deprecated._

Type: Long

Valid Range: Minimum value of 0.

Required: No

**ContentMD5**

_This member has been deprecated._

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ContentType**

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**HttpExpiresDate**

Type: Timestamp

Required: No

**RequesterCharged**

_This member has been deprecated._

Type: Boolean

Required: No

**SSEAlgorithm**

The server-side encryption algorithm used when storing objects in Amazon S3.

**Directory buckets** \- For directory buckets, there are only two supported options for server-side encryption: server-side encryption with Amazon S3 managed keys (SSE-S3) ( `AES256`) and server-side encryption with AWS KMS keys (SSE-KMS) ( `KMS`). For more
information, see [Protecting data with server-side encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html) in the _Amazon S3 User Guide_. For [the Copy operation in Batch Operations](../userguide/directory-buckets-objects-batch-ops.md), see [S3CopyObjectOperation](api-control-s3copyobjectoperation.md).

Type: String

Valid Values: `AES256 | KMS`

Required: No

**UserMetadata**

Type: String to string map

Map Entries: Maximum number of 8192 items.

Key Length Constraints: Minimum length of 1. Maximum length of 1024.

Value Length Constraints: Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/S3ObjectMetadata)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/S3ObjectMetadata)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/S3ObjectMetadata)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3ObjectLockLegalHold

S3ObjectOwner
