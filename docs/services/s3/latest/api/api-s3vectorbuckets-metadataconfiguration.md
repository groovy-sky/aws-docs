# MetadataConfiguration

The metadata configuration for a vector index.

## Contents

**nonFilterableMetadataKeys**

Non-filterable metadata keys allow you to enrich vectors with additional context during
storage and retrieval. Unlike default metadata keys, these keys can’t be used as query
filters. Non-filterable metadata keys can be retrieved but can’t be searched, queried, or
filtered. You can access non-filterable metadata keys of your vectors after finding the
vectors. For more information about non-filterable metadata keys, see [Vectors](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-vectors.html) and [Limitations and\
restrictions](../userguide/s3-vectors-limitations.md) in the _Amazon S3 User Guide_.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Length Constraints: Minimum length of 1. Maximum length of 63.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3vectors-2025-07-15/MetadataConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3vectors-2025-07-15/MetadataConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3vectors-2025-07-15/MetadataConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListOutputVector

PutInputVector
