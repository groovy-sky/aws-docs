# ParseJSON

This processor parses log events that are in JSON format. It can extract JSON key-value
pairs and place them under a destination that you specify.

Additionally, because you must have at least one parse-type processor in a transformer,
you can use `ParseJSON` as that processor for JSON-format logs, so that you can
also apply other processors, such as mutate processors, to these logs.

For more information about this processor including examples, see [parseJSON](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseJSON) in the _CloudWatch Logs User Guide_.

## Contents

**destination**

The location to put the parsed key value pair into. If you omit this parameter, it is
placed under the root node.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**source**

Path to the field in the log event that will be parsed. Use dot notation to access child
fields. For example, `store.book`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/parsejson.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/parsejson.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/parsejson.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseCloudfront

ParseKeyValue

All content copied from https://docs.aws.amazon.com/.
