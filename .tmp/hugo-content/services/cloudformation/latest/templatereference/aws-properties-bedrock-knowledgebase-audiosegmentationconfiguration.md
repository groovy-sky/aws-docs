This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase AudioSegmentationConfiguration

Configuration for segmenting audio content during multimodal knowledge base ingestion. Determines how audio files are divided into chunks for processing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FixedLengthDuration" : Integer
}

```

### YAML

```yaml

  FixedLengthDuration: Integer

```

## Properties

`FixedLengthDuration`

The duration in seconds for each audio segment. Audio files will be divided into chunks of this length for processing.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioConfiguration

BedrockEmbeddingModelConfiguration

All content copied from https://docs.aws.amazon.com/.
