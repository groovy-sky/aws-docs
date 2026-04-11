# IntelligentTieringConfiguration

Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.

For information about the S3 Intelligent-Tiering storage class, see [Storage class for\
automatically optimizing frequently and infrequently accessed objects](../dev/storage-class-intro.md#sc-dynamic-data-access).

## Contents

**Id**

The ID used to identify the S3 Intelligent-Tiering configuration.

Type: String

Required: Yes

**Status**

Specifies the status of the configuration.

Type: String

Valid Values: `Enabled | Disabled`

Required: Yes

**Tierings**

Specifies the S3 Intelligent-Tiering storage class tier of the configuration.

Type: Array of [Tiering](api-tiering.md) data types

Required: Yes

**Filter**

Specifies a bucket filter. The configuration only includes objects that meet the filter's
criteria.

Type: [IntelligentTieringFilter](api-intelligenttieringfilter.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/intelligenttieringconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/intelligenttieringconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/intelligenttieringconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntelligentTieringAndOperator

IntelligentTieringFilter

All content copied from https://docs.aws.amazon.com/.
