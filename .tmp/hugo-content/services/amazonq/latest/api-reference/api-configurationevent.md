# ConfigurationEvent

A configuration event activated by an end user request to select a specific chat
mode.

## Contents

**attributeFilter**

Enables filtering of responses based on document attributes or metadata fields.

Type: [AttributeFilter](api-attributefilter.md) object

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

For more information, see [Admin controls and guardrails](../qbusiness-ug/guardrails.md), [Plugins](../qbusiness-ug/plugins.md),
and [Conversation settings](../business-use-dg/using-web-experience.md#chat-source-scope).

Type: String

Valid Values: `RETRIEVAL_MODE | CREATOR_MODE | PLUGIN_MODE`

Required: No

**chatModeConfiguration**

Configuration information for Amazon Q Business conversation modes.

For more information, see [Admin controls and guardrails](../qbusiness-ug/guardrails.md) and [Conversation settings](../business-use-dg/using-web-experience.md#chat-source-scope).

Type: [ChatModeConfiguration](api-chatmodeconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/configurationevent.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/configurationevent.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/configurationevent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChatResponseConfigurationDetail

ContentBlockerRule

All content copied from https://docs.aws.amazon.com/.
