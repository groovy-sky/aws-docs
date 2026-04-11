# TypeConverter

Use this processor to convert a value type associated with the specified key to the
specified type. It's a casting processor that changes the types of the specified fields.
Values can be converted into one of the following datatypes: `integer`,
`double`, `string` and `boolean`.

For more information about this processor including examples, see [trimString](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-trimString) in the _CloudWatch Logs User Guide_.

## Contents

**entries**

An array of `TypeConverterEntry` objects, where each object contains the
information about one field to change the type of.

Type: Array of [TypeConverterEntry](api-typeconverterentry.md) objects

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/typeconverter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/typeconverter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/typeconverter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrimString

TypeConverterEntry

All content copied from https://docs.aws.amazon.com/.
