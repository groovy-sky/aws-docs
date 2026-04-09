# VectorBucket

The attributes of a vector bucket.

## Contents

**creationTime**

Date and time when the vector bucket was created.

Type: Timestamp

Required: Yes

**vectorBucketArn**

The Amazon Resource Name (ARN) of the vector bucket.

Type: String

Pattern: `arn:aws[-a-z0-9]*:s3vectors:[a-z0-9-]+:[0-9]{12}:bucket/[a-z0-9][a-z0-9-.]{1,61}[a-z0-9]`

Required: Yes

**vectorBucketName**

The name of the vector bucket.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Required: Yes

**encryptionConfiguration**

The encryption configuration for the vector bucket.

Type: [EncryptionConfiguration](api-s3vectorbuckets-encryptionconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3vectors-2025-07-15/vectorbucket.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3vectors-2025-07-15/vectorbucket.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3vectors-2025-07-15/vectorbucket.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValidationExceptionField

VectorBucketSummary

All content copied from https://docs.aws.amazon.com/.
