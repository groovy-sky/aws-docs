# SubstituteStringEntry

This object defines one log field key that will be replaced using the [substituteString](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-substituteString) processor.

## Contents

**from**

The regular expression string to be replaced. Special regex characters such as \[ and \]
must be escaped using \\\ when using double quotes and with \ when using single quotes. For
more information, see [Class Pattern](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/regex/Pattern.html) on the Oracle web site.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**source**

The key to modify

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**to**

The string to be substituted for each match of `from`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/substitutestringentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/substitutestringentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/substitutestringentry.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubstituteString

SuppressionPeriod

All content copied from https://docs.aws.amazon.com/.
