This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject DocumentOverrideConfiguration

Additional settings for a project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModalityProcessing" : ModalityProcessingConfiguration,
  "SensitiveDataConfiguration" : SensitiveDataConfiguration,
  "Splitter" : SplitterConfiguration
}

```

### YAML

```yaml

  ModalityProcessing:
    ModalityProcessingConfiguration
  SensitiveDataConfiguration:
    SensitiveDataConfiguration
  Splitter:
    SplitterConfiguration

```

## Properties

`ModalityProcessing`

Sets modality processing for document files. All modalities are enabled by default.

_Required_: No

_Type_: [ModalityProcessingConfiguration](aws-properties-bedrock-dataautomationproject-modalityprocessingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SensitiveDataConfiguration`

Configuration for sensitive data detection and redaction for document files.

_Required_: No

_Type_: [SensitiveDataConfiguration](aws-properties-bedrock-dataautomationproject-sensitivedataconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Splitter`

Whether document splitter is enabled for a project.

_Required_: No

_Type_: [SplitterConfiguration](aws-properties-bedrock-dataautomationproject-splitterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentOutputTextFormat

DocumentStandardExtraction

All content copied from https://docs.aws.amazon.com/.
