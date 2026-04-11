This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SpeechRecognitionSettings

Settings that control how Amazon Lex processes and recognizes speech input from users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SpeechModelConfig" : SpeechModelConfig,
  "SpeechModelPreference" : String
}

```

### YAML

```yaml

  SpeechModelConfig:
    SpeechModelConfig
  SpeechModelPreference: String

```

## Properties

`SpeechModelConfig`

Configuration settings for the selected speech-to-text model.

_Required_: No

_Type_: [SpeechModelConfig](aws-properties-lex-bot-speechmodelconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpeechModelPreference`

The speech-to-text model to use.

_Required_: No

_Type_: String

_Allowed values_: `Standard | Neural | Deepgram`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpeechModelConfig

SSMLMessage

All content copied from https://docs.aws.amazon.com/.
