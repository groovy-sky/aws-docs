# TopicConfiguration

The topic specific controls configured for an Amazon Q Business application.

## Contents

**name**

A name for your topic control configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{0,35}`

Required: Yes

**rules**

Rules defined for a topic configuration.

Type: Array of [Rule](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_Rule.html) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: Yes

**description**

A description for your topic control configuration. Use this to outline how the large
language model (LLM) should use this topic control configuration.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 350.

Pattern: `\P{C}*`

Required: No

**exampleChatMessages**

A list of example phrases that you expect the end user to use in relation to the
topic.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 5 items.

Length Constraints: Minimum length of 0. Maximum length of 350.

Pattern: `\P{C}*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/TopicConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/TopicConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/TopicConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TextSegment

UserAlias
