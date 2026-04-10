This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource ChunkingConfiguration

Details about how to chunk the documents in the data source. A _chunk_ refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.

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

Knowledge base can split your source data into chunks. A _chunk_ refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried. You have the following options for chunking your data. If you opt for `NONE`, then you may want to pre-process your files by splitting them up such that each file corresponds to a chunk.

- `FIXED_SIZE` – Amazon Bedrock splits your source data into chunks of the approximate size that you set in the `fixedSizeChunkingConfiguration`.

- `HIERARCHICAL` – Split documents into layers of chunks where the first layer contains large chunks, and the second layer contains smaller chunks derived from the first layer.

- `SEMANTIC` – Split documents into chunks based on groups of similar content derived with natural language processing.

- `NONE` – Amazon Bedrock treats each file as one chunk. If you choose this option, you may want to pre-process your documents by splitting them into separate files.

_Required_: Yes

_Type_: String

_Allowed values_: `FIXED_SIZE | NONE | HIERARCHICAL | SEMANTIC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FixedSizeChunkingConfiguration`

Configurations for when you choose fixed-size chunking. If you set the `chunkingStrategy` as `NONE`, exclude this field.

_Required_: No

_Type_: [FixedSizeChunkingConfiguration](aws-properties-bedrock-datasource-fixedsizechunkingconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HierarchicalChunkingConfiguration`

Settings for hierarchical document chunking for a data source. Hierarchical chunking splits documents
into layers of chunks where the first layer contains large chunks, and the second layer contains smaller
chunks derived from the first layer.

_Required_: No

_Type_: [HierarchicalChunkingConfiguration](aws-properties-bedrock-datasource-hierarchicalchunkingconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SemanticChunkingConfiguration`

Settings for semantic document chunking for a data source. Semantic chunking splits
a document into into smaller documents based on groups of similar content derived from the text
with natural language processing.

_Required_: No

_Type_: [SemanticChunkingConfiguration](aws-properties-bedrock-datasource-semanticchunkingconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockFoundationModelContextEnrichmentConfiguration

ConfluenceCrawlerConfiguration

All content copied from https://docs.aws.amazon.com/.
