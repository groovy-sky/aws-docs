# TypeConverterEntry

This object defines one value type that will be converted using the [typeConverter](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-typeConverter) processor.

## Contents

**key**

The key with the value that is to be converted to a different type.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**type**

The type to convert the field value to. Valid values are `integer`,
`double`, `string` and `boolean`.

Type: String

Valid Values: `boolean | integer | double | string`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/typeconverterentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/typeconverterentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/typeconverterentry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TypeConverter

UpperCaseString

All content copied from https://docs.aws.amazon.com/.
