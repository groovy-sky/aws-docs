This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject StandardOutputConfiguration

The project's standard output configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Audio" : AudioStandardOutputConfiguration,
  "Document" : DocumentStandardOutputConfiguration,
  "Image" : ImageStandardOutputConfiguration,
  "Video" : VideoStandardOutputConfiguration
}

```

### YAML

```yaml

  Audio:
    AudioStandardOutputConfiguration
  Document:
    DocumentStandardOutputConfiguration
  Image:
    ImageStandardOutputConfiguration
  Video:
    VideoStandardOutputConfiguration

```

## Properties

`Audio`

Settings for processing audio.

_Required_: No

_Type_: [AudioStandardOutputConfiguration](aws-properties-bedrock-dataautomationproject-audiostandardoutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Document`

Settings for processing documents.

_Required_: No

_Type_: [DocumentStandardOutputConfiguration](aws-properties-bedrock-dataautomationproject-documentstandardoutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

Settings for processing images.

_Required_: No

_Type_: [ImageStandardOutputConfiguration](aws-properties-bedrock-dataautomationproject-imagestandardoutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Video`

Settings for processing video.

_Required_: No

_Type_: [VideoStandardOutputConfiguration](aws-properties-bedrock-dataautomationproject-videostandardoutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SplitterConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
