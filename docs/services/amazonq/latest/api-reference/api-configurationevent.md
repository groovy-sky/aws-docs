# ConfigurationEvent

A configuration event activated by an end user request to select a specific chat
mode.

## Contents

**attributeFilter**

Enables filtering of responses based on document attributes or metadata fields.

Type: [AttributeFilter](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_AttributeFilter.html) object

Required: No

**chatMode**

The chat modes available to an Amazon Q Business end user.

- `RETRIEVAL_MODE` \- The default chat mode for an
Amazon Q Business application. When this mode is enabled, Amazon Q Business generates
responses only from data sources connected to an Amazon Q Business
application.

- `CREATOR_MODE` \- By selecting this mode, users can choose to
generate responses only from the LLM knowledge, without consulting connected
data sources, for a chat request.

- `PLUGIN_MODE` \- By selecting this mode, users can choose to
use plugins in chat.

For more information, see [Admin controls and guardrails](../qbusiness-ug/guardrails.md), [Plugins](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/plugins.html),
and [Conversation settings](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/using-web-experience.html#chat-source-scope).

Type: String

Valid Values: `RETRIEVAL_MODE | CREATOR_MODE | PLUGIN_MODE`

Required: No

**chatModeConfiguration**

Configuration information for Amazon Q Business conversation modes.

For more information, see [Admin controls and guardrails](../qbusiness-ug/guardrails.md) and [Conversation settings](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/using-web-experience.html#chat-source-scope).

Type: [ChatModeConfiguration](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_ChatModeConfiguration.html) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/ConfigurationEvent)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/ConfigurationEvent)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/ConfigurationEvent)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ChatResponseConfigurationDetail

ContentBlockerRule
