This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot TextLogSetting

Defines settings to enable text conversation logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : TextLogDestination,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Destination:
    TextLogDestination
  Enabled: Boolean

```

## Properties

`Destination`

Specifies the Amazon CloudWatch Logs destination log group for
conversation text logs.

_Required_: Yes

_Type_: [TextLogDestination](aws-properties-lex-bot-textlogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Determines whether conversation logs should be stored for an
alias.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextLogDestination

UnifiedSpeechSettings

All content copied from https://docs.aws.amazon.com/.
