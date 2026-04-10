This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot AdvancedRecognitionSetting

Provides settings that enable advanced recognition settings for slot values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioRecognitionStrategy" : String
}

```

### YAML

```yaml

  AudioRecognitionStrategy: String

```

## Properties

`AudioRecognitionStrategy`

Enables using the slot values as a custom vocabulary for recognizing user utterances.

_Required_: No

_Type_: String

_Allowed values_: `UseSlotValuesAsCustomVocabulary`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lex::Bot

AllowedInputTypes

All content copied from https://docs.aws.amazon.com/.
