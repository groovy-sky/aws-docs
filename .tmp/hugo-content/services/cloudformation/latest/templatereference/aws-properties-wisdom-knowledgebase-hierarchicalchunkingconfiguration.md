This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase HierarchicalChunkingConfiguration

Settings for hierarchical document chunking for a data source. Hierarchical chunking splits documents into
layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived
from the first layer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LevelConfigurations" : [ HierarchicalChunkingLevelConfiguration, ... ],
  "OverlapTokens" : Number
}

```

### YAML

```yaml

  LevelConfigurations:
    - HierarchicalChunkingLevelConfiguration
  OverlapTokens: Number

```

## Properties

`LevelConfigurations`

Token settings for each layer.

_Required_: Yes

_Type_: Array of [HierarchicalChunkingLevelConfiguration](aws-properties-wisdom-knowledgebase-hierarchicalchunkinglevelconfiguration.md)

_Minimum_: `2`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverlapTokens`

The number of tokens to repeat across chunks in the same layer.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FixedSizeChunkingConfiguration

HierarchicalChunkingLevelConfiguration

All content copied from https://docs.aws.amazon.com/.
