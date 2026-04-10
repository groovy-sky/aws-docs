This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase AudioConfiguration

Configuration settings for processing audio content in multimodal knowledge bases.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SegmentationConfiguration" : AudioSegmentationConfiguration
}

```

### YAML

```yaml

  SegmentationConfiguration:
    AudioSegmentationConfiguration

```

## Properties

`SegmentationConfiguration`

Configuration for segmenting audio content during processing.

_Required_: Yes

_Type_: [AudioSegmentationConfiguration](aws-properties-bedrock-knowledgebase-audiosegmentationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::KnowledgeBase

AudioSegmentationConfiguration

All content copied from https://docs.aws.amazon.com/.
