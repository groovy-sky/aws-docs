This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase VideoConfiguration

Configuration settings for processing video content in multimodal knowledge bases.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SegmentationConfiguration" : VideoSegmentationConfiguration
}

```

### YAML

```yaml

  SegmentationConfiguration:
    VideoSegmentationConfiguration

```

## Properties

`SegmentationConfiguration`

Configuration for segmenting video content during processing.

_Required_: Yes

_Type_: [VideoSegmentationConfiguration](aws-properties-bedrock-knowledgebase-videosegmentationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorKnowledgeBaseConfiguration

VideoSegmentationConfiguration

All content copied from https://docs.aws.amazon.com/.
