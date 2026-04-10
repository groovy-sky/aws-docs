This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::KnowledgeBase BedrockEmbeddingModelConfiguration

The vector configuration details for the Bedrock embeddings model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Audio" : [ AudioConfiguration, ... ],
  "Dimensions" : Integer,
  "EmbeddingDataType" : String,
  "Video" : [ VideoConfiguration, ... ]
}

```

### YAML

```yaml

  Audio:
    - AudioConfiguration
  Dimensions: Integer
  EmbeddingDataType: String
  Video:
    - VideoConfiguration

```

## Properties

`Audio`

Configuration settings for processing audio content in multimodal knowledge bases.

_Required_: No

_Type_: Array of [AudioConfiguration](aws-properties-bedrock-knowledgebase-audioconfiguration.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Dimensions`

The dimensions details for the vector configuration used on the Bedrock embeddings model.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EmbeddingDataType`

The data type for the vectors when using a model to convert text into vector
embeddings. The model must support the specified data type for vector embeddings.
Floating-point (float32) is the default data type, and is supported by most models
for vector embeddings. See [Supported embeddings \
models](../../../bedrock/latest/userguide/knowledge-base-supported.md) for information on the available models and their vector data types.

_Required_: No

_Type_: String

_Allowed values_: `FLOAT32 | BINARY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Video`

Configuration settings for processing video content in multimodal knowledge bases.

_Required_: No

_Type_: Array of [VideoConfiguration](aws-properties-bedrock-knowledgebase-videoconfiguration.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioSegmentationConfiguration

CuratedQuery

All content copied from https://docs.aws.amazon.com/.
