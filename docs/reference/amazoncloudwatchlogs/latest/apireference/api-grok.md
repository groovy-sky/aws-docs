---
title: "Grok"
---

# Grok

This processor uses pattern matching to parse and structure unstructured data. This
processor can also extract fields from log messages.

For more information about this processor including examples, see [grok](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-configurable.md#CloudWatch-Logs-Transformation-Grok) in the _CloudWatch Logs User Guide_.

## Contents

**match**

The grok pattern to match against the log event. For a list of supported grok patterns,
see [Supported grok patterns](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-configurable.md#CloudWatch-Logs-Transformation-Grok).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Required: Yes

**source**

The path to the field in the log event that you want to parse. If you omit this value, the
whole log message is parsed.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/Grok)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/Grok)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/Grok)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetLogObjectResponseStream

GroupingIdentifier

All content copied from https://docs.aws.amazon.com/.
