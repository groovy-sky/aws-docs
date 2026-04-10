This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotAlias AudioLogSetting

Settings for logging audio of conversations between Amazon Lex and a
user. You specify whether to log audio and the Amazon S3 bucket where
the audio file is stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : AudioLogDestination,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Destination:
    AudioLogDestination
  Enabled: Boolean

```

## Properties

`Destination`

The location of audio log files collected when conversation logging
is enabled for a bot.

_Required_: Yes

_Type_: [AudioLogDestination](aws-properties-lex-botalias-audiologdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Determines whether audio logging in enabled for the bot.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioLogDestination

BotAliasLocaleSettings

All content copied from https://docs.aws.amazon.com/.
