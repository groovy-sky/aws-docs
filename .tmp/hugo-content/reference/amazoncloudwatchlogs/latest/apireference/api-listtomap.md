# ListToMap

This processor takes a list of objects that contain key fields, and converts them into a
map of target keys.

For more information about this processor including examples, see [listToMap](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-listToMap) in the _CloudWatch Logs User Guide_.

## Contents

**key**

The key of the field to be extracted as keys in the generated map

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**source**

The key in the log event that has a list of objects that will be converted to a
map.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**flatten**

A Boolean value to indicate whether the list will be flattened into single items. Specify
`true` to flatten the list. The default is `false`

Type: Boolean

Required: No

**flattenedElement**

If you set `flatten` to `true`, use `flattenedElement` to
specify which element, `first` or `last`, to keep.

You must specify this parameter if `flatten` is `true`

Type: String

Valid Values: `first | last`

Required: No

**target**

The key of the field that will hold the generated map

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**valueKey**

If this is specified, the values that you specify in this parameter will be extracted from
the `source` objects and put into the values of the generated map. Otherwise,
original objects in the source list will be put into the values of the generated map.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/listtomap.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/listtomap.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/listtomap.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntegrationSummary

LiveTailSessionLogEvent

All content copied from https://docs.aws.amazon.com/.
