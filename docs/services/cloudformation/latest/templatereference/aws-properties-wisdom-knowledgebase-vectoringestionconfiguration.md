This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase VectorIngestionConfiguration

Contains details about how to ingest the documents in a data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChunkingConfiguration" : ChunkingConfiguration,
  "ParsingConfiguration" : ParsingConfiguration
}

```

### YAML

```yaml

  ChunkingConfiguration:
    ChunkingConfiguration
  ParsingConfiguration:
    ParsingConfiguration

```

## Properties

`ChunkingConfiguration`

Details about how to chunk the documents in the data source. A chunk refers to an excerpt from a data source
that is returned when the knowledge base that it belongs to is queried.

_Required_: No

_Type_: [ChunkingConfiguration](aws-properties-wisdom-knowledgebase-chunkingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParsingConfiguration`

A custom parser for data source documents.

_Required_: No

_Type_: [ParsingConfiguration](aws-properties-wisdom-knowledgebase-parsingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UrlConfiguration

WebCrawlerConfiguration

All content copied from https://docs.aws.amazon.com/.
