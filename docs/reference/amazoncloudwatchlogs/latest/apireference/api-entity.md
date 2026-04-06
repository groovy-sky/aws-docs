# Entity

The entity associated with the log events in a `PutLogEvents` call.

## Contents

**attributes**

Additional attributes of the entity that are not used to specify the identity of the
entity. A list of key-value pairs.

For details about how to use the attributes, see [How to add\
related information to telemetry](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/adding-your-own-related-telemetry.html) in the _CloudWatch User_
_Guide_.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 10 items.

Key Length Constraints: Minimum length of 1. Maximum length of 256.

Value Length Constraints: Minimum length of 1. Maximum length of 512.

Required: No

**keyAttributes**

The attributes of the entity which identify the specific entity, as a list of key-value
pairs. Entities with the same `keyAttributes` are considered to be the same
entity.

There are five allowed attributes (key names): `Type`,
`ResourceType`, `Identifier` `Name`, and `Environment`.

For details about how to use the key attributes, see [How to add\
related information to telemetry](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/adding-your-own-related-telemetry.html) in the _CloudWatch User_
_Guide_.

Type: String to string map

Map Entries: Minimum number of 2 items. Maximum number of 4 items.

Key Length Constraints: Minimum length of 1. Maximum length of 32.

Value Length Constraints: Minimum length of 1. Maximum length of 512.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/Entity)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/Entity)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/Entity)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DestinationConfiguration

ExportTask
