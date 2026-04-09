# ParseVPC

Use this processor to parse Amazon VPC vended logs, extract fields, and and
convert them into a JSON format. This processor always processes the entire log event
message.

This processor doesn't support custom log formats, such as NAT gateway logs. For more
information about custom log formats in Amazon VPC, see [parseVPC](../../../../services/vpc/latest/userguide/flow-logs-records-examples.md#flow-log-example-tcp-flag) For more information about this processor including examples, see [parseVPC](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseVPC).

###### Important

If you use this processor, it must be the first processor in your transformer.

## Contents

**source**

Omit this parameter and the whole log message will be processed by this processor. No
other value than `@message` is allowed for `source`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/parsevpc.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/parsevpc.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/parsevpc.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseToOCSF

ParseWAF

All content copied from https://docs.aws.amazon.com/.
