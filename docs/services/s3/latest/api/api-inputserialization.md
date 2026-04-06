# InputSerialization

Describes the serialization format of the object.

## Contents

**CompressionType**

Specifies object's compression format. Valid values: NONE, GZIP, BZIP2. Default Value: NONE.

Type: String

Valid Values: `NONE | GZIP | BZIP2`

Required: No

**CSV**

Describes the serialization of a CSV-encoded object.

Type: [CSVInput](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CSVInput.html) data type

Required: No

**JSON**

Specifies JSON as object's input serialization format.

Type: [JSONInput](https://docs.aws.amazon.com/AmazonS3/latest/API/API_JSONInput.html) data type

Required: No

**Parquet**

Specifies Parquet as object's input serialization format.

Type: [ParquetInput](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ParquetInput.html) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/InputSerialization)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/InputSerialization)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/InputSerialization)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Initiator

IntelligentTieringAndOperator
