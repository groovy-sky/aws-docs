This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SpeechModelConfig

Configuration settings that define which speech-to-text model to use for processing speech input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeepgramConfig" : DeepgramSpeechModelConfig
}

```

### YAML

```yaml

  DeepgramConfig:
    DeepgramSpeechModelConfig

```

## Properties

`DeepgramConfig`

Configuration settings for using Deepgram as the speech-to-text provider.

_Required_: No

_Type_: [DeepgramSpeechModelConfig](aws-properties-lex-bot-deepgramspeechmodelconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpeechFoundationModel

SpeechRecognitionSettings

All content copied from https://docs.aws.amazon.com/.
