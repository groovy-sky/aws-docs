# Index

The attributes of a vector index.

## Contents

**creationTime**

Date and time when the vector index was created.

Type: Timestamp

Required: Yes

**dataType**

The data type of the vectors inserted into the vector index.

Type: String

Valid Values: `float32`

Required: Yes

**dimension**

The number of values in the vectors that are inserted into the vector index.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 4096.

Required: Yes

**distanceMetric**

The distance metric to be used for similarity search.

Type: String

Valid Values: `euclidean | cosine`

Required: Yes

**indexArn**

The Amazon Resource Name (ARN) of the vector index.

Type: String

Pattern: `arn:aws[-a-z0-9]*:s3vectors:[a-z0-9-]+:[0-9]{12}:bucket/[a-z0-9][a-z0-9-.]{1,61}[a-z0-9]/index/[a-z0-9][a-z0-9-.]{1,61}[a-z0-9]`

Required: Yes

**indexName**

The name of the vector index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Required: Yes

**vectorBucketName**

The name of the vector bucket that contains the vector index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Required: Yes

**encryptionConfiguration**

The encryption configuration for a vector index. By default, if you don't specify, all new vectors in the vector index will use the encryption configuration of the vector bucket.

Type: [EncryptionConfiguration](api-s3vectorbuckets-encryptionconfiguration.md) object

Required: No

**metadataConfiguration**

The metadata configuration for the vector index.

Type: [MetadataConfiguration](api-s3vectorbuckets-metadataconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3vectors-2025-07-15/index.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3vectors-2025-07-15/index.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3vectors-2025-07-15/index.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetOutputVector

IndexSummary

All content copied from https://docs.aws.amazon.com/.
