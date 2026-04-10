This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject OverrideConfiguration

Additional settings for a project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Audio" : AudioOverrideConfiguration,
  "Document" : DocumentOverrideConfiguration,
  "Image" : ImageOverrideConfiguration,
  "ModalityRouting" : ModalityRoutingConfiguration,
  "Video" : VideoOverrideConfiguration
}

```

### YAML

```yaml

  Audio:
    AudioOverrideConfiguration
  Document:
    DocumentOverrideConfiguration
  Image:
    ImageOverrideConfiguration
  ModalityRouting:
    ModalityRoutingConfiguration
  Video:
    VideoOverrideConfiguration

```

## Properties

`Audio`

This element declares whether your project will process audio files.

_Required_: No

_Type_: [AudioOverrideConfiguration](aws-properties-bedrock-dataautomationproject-audiooverrideconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Document`

Additional settings for a project.

_Required_: No

_Type_: [DocumentOverrideConfiguration](aws-properties-bedrock-dataautomationproject-documentoverrideconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

This element declares whether your project will process image files.

_Required_: No

_Type_: [ImageOverrideConfiguration](aws-properties-bedrock-dataautomationproject-imageoverrideconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModalityRouting`

Lets you set which modalities certain file types are processed as.

_Required_: No

_Type_: [ModalityRoutingConfiguration](aws-properties-bedrock-dataautomationproject-modalityroutingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Video`

This element declares whether your project will process video files.

_Required_: No

_Type_: [VideoOverrideConfiguration](aws-properties-bedrock-dataautomationproject-videooverrideconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModalityRoutingConfiguration

PIIEntitiesConfiguration

All content copied from https://docs.aws.amazon.com/.
