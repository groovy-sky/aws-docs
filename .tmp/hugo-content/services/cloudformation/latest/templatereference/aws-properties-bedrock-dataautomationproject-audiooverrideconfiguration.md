This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject AudioOverrideConfiguration

Sets whether your project will process audio or not.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LanguageConfiguration" : AudioLanguageConfiguration,
  "ModalityProcessing" : ModalityProcessingConfiguration,
  "SensitiveDataConfiguration" : SensitiveDataConfiguration
}

```

### YAML

```yaml

  LanguageConfiguration:
    AudioLanguageConfiguration
  ModalityProcessing:
    ModalityProcessingConfiguration
  SensitiveDataConfiguration:
    SensitiveDataConfiguration

```

## Properties

`LanguageConfiguration`

The output and input language configuration for your audio.

_Required_: No

_Type_: [AudioLanguageConfiguration](aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModalityProcessing`

Sets modality processing for audio files. All modalities are enabled by default.

_Required_: No

_Type_: [ModalityProcessingConfiguration](aws-properties-bedrock-dataautomationproject-modalityprocessingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SensitiveDataConfiguration`

Configuration for sensitive data detection and redaction for audio files.

_Required_: No

_Type_: [SensitiveDataConfiguration](aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioLanguageConfiguration

AudioStandardExtraction

All content copied from https://docs.aws.amazon.com/.
