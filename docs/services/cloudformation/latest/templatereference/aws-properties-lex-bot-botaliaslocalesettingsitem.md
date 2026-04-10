This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot BotAliasLocaleSettingsItem

Specifies locale settings for a single locale.

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

Specifies locale settings for a locale.

_Required_: Yes

_Type_: [BotAliasLocaleSettings](aws-properties-lex-bot-botaliaslocalesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocaleId`

Specifies the locale that the settings apply to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BotAliasLocaleSettings

BotLocale

All content copied from https://docs.aws.amazon.com/.
