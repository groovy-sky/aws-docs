This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject VideoOverrideConfiguration

Sets whether your project will process videos or not.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModalityProcessing" : ModalityProcessingConfiguration,
  "SensitiveDataConfiguration" : SensitiveDataConfiguration
}

```

### YAML

```yaml

  ModalityProcessing:
    ModalityProcessingConfiguration
  SensitiveDataConfiguration:
    SensitiveDataConfiguration

```

## Properties

`ModalityProcessing`

Sets modality processing for video files. All modalities are enabled by default.

_Required_: No

_Type_: [ModalityProcessingConfiguration](aws-properties-bedrock-dataautomationproject-modalityprocessingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SensitiveDataConfiguration`

Configuration for sensitive data detection and redaction for video files.

_Required_: No

_Type_: [SensitiveDataConfiguration](aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoExtractionCategory

VideoStandardExtraction

All content copied from https://docs.aws.amazon.com/.
