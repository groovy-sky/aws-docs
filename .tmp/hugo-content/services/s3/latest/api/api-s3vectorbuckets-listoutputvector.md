# ListOutputVector

The attributes of a vector returned by the `ListVectors` operation.

## Contents

**key**

The name of the vector.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**data**

The vector data of the vector.

Type: [VectorData](api-s3vectorbuckets-vectordata.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**metadata**

Metadata about the vector.

Type: JSON value

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3vectors-2025-07-15/listoutputvector.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3vectors-2025-07-15/listoutputvector.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3vectors-2025-07-15/listoutputvector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IndexSummary

MetadataConfiguration

All content copied from https://docs.aws.amazon.com/.
