This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase ChunkingConfiguration

Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source
that is returned when the knowledge base that it belongs to is queried.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChunkingStrategy" : String,
  "FixedSizeChunkingConfiguration" : FixedSizeChunkingConfiguration,
  "HierarchicalChunkingConfiguration" : HierarchicalChunkingConfiguration,
  "SemanticChunkingConfiguration" : SemanticChunkingConfiguration
}

```

### YAML

```yaml

  ChunkingStrategy: String
  FixedSizeChunkingConfiguration:
    FixedSizeChunkingConfiguration
  HierarchicalChunkingConfiguration:
    HierarchicalChunkingConfiguration
  SemanticChunkingConfiguration:
    SemanticChunkingConfiguration

```

## Properties

`ChunkingStrategy`

Knowledge base can split your source data into chunks. A chunk refers to an excerpt from a data source that is
returned when the knowledge base that it belongs to is queried. You have the following options for chunking your
data. If you opt for `NONE`, then you may want to pre-process your files by splitting them up such that
each file corresponds to a chunk.

_Required_: Yes

_Type_: String

_Allowed values_: `FIXED_SIZE | NONE | HIERARCHICAL | SEMANTIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FixedSizeChunkingConfiguration`

Configurations for when you choose fixed-size chunking. If you set the `chunkingStrategy` as
`NONE`, exclude this field.

_Required_: No

_Type_: [FixedSizeChunkingConfiguration](aws-properties-wisdom-knowledgebase-fixedsizechunkingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchicalChunkingConfiguration`

Settings for hierarchical document chunking for a data source. Hierarchical chunking splits documents into
layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived
from the first layer.

_Required_: No

_Type_: [HierarchicalChunkingConfiguration](aws-properties-wisdom-knowledgebase-hierarchicalchunkingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticChunkingConfiguration`

Settings for semantic document chunking for a data source. Semantic chunking splits a document into smaller
documents based on groups of similar content derived from the text with natural language processing.

_Required_: No

_Type_: [SemanticChunkingConfiguration](aws-properties-wisdom-knowledgebase-semanticchunkingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockFoundationModelConfiguration

CrawlerLimits

All content copied from https://docs.aws.amazon.com/.
