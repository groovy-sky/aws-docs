# PutInputVector

The attributes of a vector to add to a vector index.

## Contents

**data**

The vector data of the vector.

Vector dimensions must match the dimension count that's configured for the vector
index.

- For the `cosine` distance metric, zero vectors (vectors containing all
zeros) aren't allowed.

- For both `cosine` and `euclidean` distance metrics, vector
data must contain only valid floating-point values. Invalid values such as NaN (Not a
Number) or Infinity aren't allowed.

Type: [VectorData](api-s3vectorbuckets-vectordata.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

**key**

The name of the vector. The key uniquely identifies the vector in a vector index.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**metadata**

Metadata about the vector. All metadata entries undergo validation to ensure they meet
the format requirements for size and data types.

Type: JSON value

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3vectors-2025-07-15/putinputvector.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3vectors-2025-07-15/putinputvector.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3vectors-2025-07-15/putinputvector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataConfiguration

QueryOutputVector

All content copied from https://docs.aws.amazon.com/.
