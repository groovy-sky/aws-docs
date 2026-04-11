This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ConversationLogSettings

Configures conversation logging that saves audio, text, and metadata
for the conversations with your users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioLogSettings" : [ AudioLogSetting, ... ],
  "TextLogSettings" : [ TextLogSetting, ... ]
}

```

### YAML

```yaml

  AudioLogSettings:
    - AudioLogSetting
  TextLogSettings:
    - TextLogSetting

```

## Properties

`AudioLogSettings`

The Amazon S3 settings for logging audio to an S3 bucket.

_Required_: No

_Type_: Array of [AudioLogSetting](aws-properties-lex-bot-audiologsetting.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextLogSettings`

The Amazon CloudWatch Logs settings for logging text and metadata.

_Required_: No

_Type_: Array of [TextLogSetting](aws-properties-lex-bot-textlogsetting.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionalSpecification

CustomPayload

All content copied from https://docs.aws.amazon.com/.
