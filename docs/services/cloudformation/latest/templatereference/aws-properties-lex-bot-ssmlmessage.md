This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SSMLMessage

Defines a Speech Synthesis Markup Language (SSML) prompt.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : String
}

```

### YAML

```yaml

  Value: String

```

## Properties

`Value`

The SSML text that defines the prompt.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpeechRecognitionSettings

StillWaitingResponseSpecification

All content copied from https://docs.aws.amazon.com/.
