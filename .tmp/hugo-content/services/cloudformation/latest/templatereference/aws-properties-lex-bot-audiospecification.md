This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot AudioSpecification

Specifies the audio input specifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTimeoutMs" : Integer,
  "MaxLengthMs" : Integer
}

```

### YAML

```yaml

  EndTimeoutMs: Integer
  MaxLengthMs: Integer

```

## Properties

`EndTimeoutMs`

Time for which a bot waits after the customer stops speaking to assume the
utterance is finished.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLengthMs`

Time for how long Amazon Lex waits before speech input is truncated and the speech
is returned to application.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioLogSetting

BedrockAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
