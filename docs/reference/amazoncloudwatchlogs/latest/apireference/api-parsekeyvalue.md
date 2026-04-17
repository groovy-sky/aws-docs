---
title: "ParseKeyValue"
---

# ParseKeyValue

This processor parses a specified field in the original log event into key-value pairs.

For more information about this processor including examples, see [parseKeyValue](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md#CloudWatch-Logs-Transformation-parseKeyValue) in the _CloudWatch Logs User Guide_.

## Contents

**destination**

The destination field to put the extracted key-value pairs into

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**fieldDelimiter**

The field delimiter string that is used between key-value pairs in the original log
events. If you omit this, the ampersand `&` character is used.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**keyPrefix**

If you want to add a prefix to all transformed keys, specify it here.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**keyValueDelimiter**

The delimiter string to use between the key and value in each pair in the transformed log
event.

If you omit this, the equal `=` character is used.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**nonMatchValue**

A value to insert into the value field in the result, when a key-value pair is not
successfully split.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

**overwriteIfExists**

Specifies whether to overwrite the value if the destination key already exists. If you
omit this, the default is `false`.

Type: Boolean

Required: No

**source**

Path to the field in the log event that will be parsed. Use dot notation to access child
fields. For example, `store.book`

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/ParseKeyValue)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/ParseKeyValue)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/ParseKeyValue)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseJSON

ParsePostgres

All content copied from https://docs.aws.amazon.com/.
