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

Type: [CSVInput](api-csvinput.md) data type

Required: No

**JSON**

Specifies JSON as object's input serialization format.

Type: [JSONInput](api-jsoninput.md) data type

Required: No

**Parquet**

Specifies Parquet as object's input serialization format.

Type: [ParquetInput](api-parquetinput.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/inputserialization.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/inputserialization.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/inputserialization.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Initiator

IntelligentTieringAndOperator

All content copied from https://docs.aws.amazon.com/.
