This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot TestBotAliasSettings

Specifies configuration settings for the alias used to test the bot.
If the `TestBotAliasSettings` property is not specified, the
settings are configured with default values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BotAliasLocaleSettings" : [ BotAliasLocaleSettingsItem, ... ],
  "ConversationLogSettings" : ConversationLogSettings,
  "Description" : String,
  "SentimentAnalysisSettings" : SentimentAnalysisSettings
}

```

### YAML

```yaml

  BotAliasLocaleSettings:
    - BotAliasLocaleSettingsItem
  ConversationLogSettings:
    ConversationLogSettings
  Description: String
  SentimentAnalysisSettings:
    SentimentAnalysisSettings

```

## Properties

`BotAliasLocaleSettings`

Specifies settings that are unique to a locale. For example, you can
use a different Lambda function depending on the bot's
locale.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lex-bot-botaliaslocalesettings.html) of [BotAliasLocaleSettingsItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lex-bot-botaliaslocalesettingsitem.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConversationLogSettings`

Specifies settings for conversation logs that save audio, text, and
metadata information for conversations with your users.

_Required_: No

_Type_: [ConversationLogSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lex-bot-conversationlogsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Specifies a description for the test bot alias.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SentimentAnalysisSettings`

Specifies whether Amazon Lex will use Amazon Comprehend to
detect the sentiment of user utterances.

_Required_: No

_Type_: [SentimentAnalysisSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lex-bot-sentimentanalysissettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

TextInputSpecification
