This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataAutomationProject TranscriptConfiguration

Configuration for transcript options. This option allows you to enable speaker labeling and channel labeling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelLabeling" : ChannelLabelingConfiguration,
  "SpeakerLabeling" : SpeakerLabelingConfiguration
}

```

### YAML

```yaml

  ChannelLabeling:
    ChannelLabelingConfiguration
  SpeakerLabeling:
    SpeakerLabelingConfiguration

```

## Properties

`ChannelLabeling`

Enables channel labeling. Each audio channel will be labeled with a number, and the transcript will indicate which channel is being used.

_Required_: No

_Type_: [ChannelLabelingConfiguration](aws-properties-bedrock-dataautomationproject-channellabelingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpeakerLabeling`

Enables speaker labeling. Each speaker within a transcript will recieve a number, and the transcript will note which speaker is talking.

_Required_: No

_Type_: [SpeakerLabelingConfiguration](aws-properties-bedrock-dataautomationproject-speakerlabelingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VideoBoundingBox

All content copied from https://docs.aws.amazon.com/.
