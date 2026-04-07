This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource VectorIngestionConfiguration

Contains details about how to ingest the documents in a data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChunkingConfiguration" : ChunkingConfiguration,
  "ContextEnrichmentConfiguration" : ContextEnrichmentConfiguration,
  "CustomTransformationConfiguration" : CustomTransformationConfiguration,
  "ParsingConfiguration" : ParsingConfiguration
}

```

### YAML

```yaml

  ChunkingConfiguration:
    ChunkingConfiguration
  ContextEnrichmentConfiguration:
    ContextEnrichmentConfiguration
  CustomTransformationConfiguration:
    CustomTransformationConfiguration
  ParsingConfiguration:
    ParsingConfiguration

```

## Properties

`ChunkingConfiguration`

Details about how to chunk the documents in the data source. A _chunk_ refers to an excerpt from a data source that is returned when the knowledge base that it belongs to is queried.

_Required_: No

_Type_: [ChunkingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-chunkingconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContextEnrichmentConfiguration`

The context enrichment configuration used for ingestion of the data into the vector
store.

_Required_: No

_Type_: [ContextEnrichmentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-contextenrichmentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomTransformationConfiguration`

A custom document transformer for parsed data source documents.

_Required_: No

_Type_: [CustomTransformationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-customtransformationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParsingConfiguration`

Configurations for a parser to use for parsing documents in your data source. If you exclude this field, the default parser will be used.

_Required_: No

_Type_: [ParsingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-parsingconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UrlConfiguration

WebCrawlerConfiguration
