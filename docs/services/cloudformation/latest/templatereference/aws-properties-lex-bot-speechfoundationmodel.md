This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SpeechFoundationModel

Configuration for a foundation model used for speech synthesis and recognition capabilities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModelArn" : String,
  "VoiceId" : String
}

```

### YAML

```yaml

  ModelArn: String
  VoiceId: String

```

## Properties

`ModelArn`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VoiceId`

The identifier of the voice to use for speech synthesis with the foundation model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Specifications

SpeechModelConfig

All content copied from https://docs.aws.amazon.com/.
