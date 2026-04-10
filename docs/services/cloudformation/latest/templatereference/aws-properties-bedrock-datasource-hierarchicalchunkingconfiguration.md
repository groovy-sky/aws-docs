This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource HierarchicalChunkingConfiguration

Settings for hierarchical document chunking for a data source. Hierarchical chunking splits documents
into layers of chunks where the first layer contains large chunks, and the second layer contains smaller
chunks derived from the first layer.

You configure the number of tokens to overlap, or repeat across adjacent chunks. For example,
if you set overlap tokens to 60, the last 60 tokens in the first chunk are also included at the beginning of
the second chunk. For each layer, you must also configure the maximum number of tokens in a chunk.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LevelConfigurations" : [ HierarchicalChunkingLevelConfiguration, ... ],
  "OverlapTokens" : Integer
}

```

### YAML

```yaml

  LevelConfigurations:
    - HierarchicalChunkingLevelConfiguration
  OverlapTokens: Integer

```

## Properties

`LevelConfigurations`

Token settings for each layer.

_Required_: Yes

_Type_: Array of [HierarchicalChunkingLevelConfiguration](aws-properties-bedrock-datasource-hierarchicalchunkinglevelconfiguration.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OverlapTokens`

The number of tokens to repeat across chunks in the same layer.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FixedSizeChunkingConfiguration

HierarchicalChunkingLevelConfiguration

All content copied from https://docs.aws.amazon.com/.
