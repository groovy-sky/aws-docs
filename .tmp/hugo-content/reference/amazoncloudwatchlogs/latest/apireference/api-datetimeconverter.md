# DateTimeConverter

This processor converts a datetime string into a format that you specify.

For more information about this processor including examples, see [datetimeConverter](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-datetimeConverter) in the _CloudWatch Logs User_
_Guide_.

## Contents

**matchPatterns**

A list of patterns to match against the `source` field.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Length Constraints: Minimum length of 1.

Required: Yes

**source**

The key to apply the date conversion to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**target**

The JSON field to store the result in.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**locale**

The locale of the source field. If you omit this, the default of `locale.ROOT`
is used.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**sourceTimezone**

The time zone of the source field. If you omit this, the default used is the UTC
zone.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**targetFormat**

The datetime format to use for the converted data in the target field.

If you omit this, the default of ` yyyy-MM-dd'T'HH:mm:ss.SSS'Z` is used.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**targetTimezone**

The time zone of the target field. If you omit this, the default used is the UTC
zone.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/datetimeconverter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/datetimeconverter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/datetimeconverter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceFilter

DeleteKeys

All content copied from https://docs.aws.amazon.com/.
