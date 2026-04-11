This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotAlias

###### Note

Amazon Lex V2 is the only supported version in CloudFormation.

Specifies an alias for the specified version of a bot. Use an alias
to enable you to change the version of a bot without updating
applications that use the bot.

For example, you can specify an alias called "PROD" that your
applications use to call the Amazon Lex bot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lex::BotAlias",
  "Properties" : {
      "BotAliasLocaleSettings" : [ BotAliasLocaleSettingsItem, ... ],
      "BotAliasName" : String,
      "BotAliasTags" : [ Tag, ... ],
      "BotId" : String,
      "BotVersion" : String,
      "ConversationLogSettings" : ConversationLogSettings,
      "Description" : String,
      "SentimentAnalysisSettings" : SentimentAnalysisSettings
    }
}

```

### YAML

```yaml

Type: AWS::Lex::BotAlias
Properties:
  BotAliasLocaleSettings:
    - BotAliasLocaleSettingsItem
  BotAliasName: String
  BotAliasTags:
    - Tag
  BotId: String
  BotVersion: String
  ConversationLogSettings:
    ConversationLogSettings
  Description: String
  SentimentAnalysisSettings:
    SentimentAnalysisSettings

```

## Properties

`BotAliasLocaleSettings`

Specifies settings that are unique to a locale. For example, you can
use different Lambda function depending on the bot's locale.

_Required_: No

_Type_: [Array](aws-properties-lex-botalias-botaliaslocalesettings.md) of [BotAliasLocaleSettingsItem](aws-properties-lex-botalias-botaliaslocalesettingsitem.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BotAliasName`

The name of the bot alias.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BotAliasTags`

An array of key-value pairs to apply to this resource.

You can only add tags when you specify an alias.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-lex-botalias-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BotId`

The unique identifier of the bot.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z]+$`

_Minimum_: `10`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BotVersion`

The version of the bot that the bot alias references.

_Required_: No

_Type_: String

_Pattern_: `^(DRAFT|[0-9]+)$`

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConversationLogSettings`

Specifies whether Amazon Lex logs text and audio for
conversations with the bot. When you enable conversation logs, text
logs store text input, transcripts of audio input, and associated
metadata in Amazon CloudWatch logs. Audio logs store input in
Amazon S3.

_Required_: No

_Type_: [ConversationLogSettings](aws-properties-lex-botalias-conversationlogsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the bot alias.

_Required_: No

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SentimentAnalysisSettings`

Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of
user utterances.

_Required_: No

_Type_: [SentimentAnalysisSettings](aws-properties-lex-botalias-sentimentanalysissettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the bot alias.

`BotAliasId`

The unique identifier of the bot alias.

`BotAliasStatus`

The current status of the bot alias. When the status is Available
the alias is ready for use with your bot.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WaitAndContinueSpecification

AudioLogDestination

All content copied from https://docs.aws.amazon.com/.
