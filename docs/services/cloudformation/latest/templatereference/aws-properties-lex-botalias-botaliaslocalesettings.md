This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotAlias BotAliasLocaleSettings

Specifies settings that are unique to a locale. For example, you can
use different Lambda function depending on the bot's locale.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CodeHookSpecification" : CodeHookSpecification,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  CodeHookSpecification:
    CodeHookSpecification
  Enabled: Boolean

```

## Properties

`CodeHookSpecification`

Specifies the Lambda function that should be used in the
locale.

_Required_: No

_Type_: [CodeHookSpecification](aws-properties-lex-botalias-codehookspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Determines whether the locale is enabled for the bot. If the value
is `false`, the locale isn't available for use.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioLogSetting

BotAliasLocaleSettingsItem

All content copied from https://docs.aws.amazon.com/.
