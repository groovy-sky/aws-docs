This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotAlias BotAliasLocaleSettingsItem

Specifies settings that are unique to a locale. For example, you can
use different Lambda function depending on the bot's locale.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BotAliasLocaleSetting" : BotAliasLocaleSettings,
  "LocaleId" : String
}

```

### YAML

```yaml

  BotAliasLocaleSetting:
    BotAliasLocaleSettings
  LocaleId: String

```

## Properties

`BotAliasLocaleSetting`

Specifies settings that are unique to a locale.

_Required_: Yes

_Type_: [BotAliasLocaleSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lex-botalias-botaliaslocalesettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocaleId`

The unique identifier of the locale.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BotAliasLocaleSettings

CloudWatchLogGroupLogDestination
