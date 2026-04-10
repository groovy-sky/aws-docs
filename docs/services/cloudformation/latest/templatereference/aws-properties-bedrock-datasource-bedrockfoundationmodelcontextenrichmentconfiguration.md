This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource BedrockFoundationModelContextEnrichmentConfiguration

Context enrichment configuration is used to provide additional context to the RAG application
using Amazon Bedrock foundation models.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnrichmentStrategyConfiguration" : EnrichmentStrategyConfiguration,
  "ModelArn" : String
}

```

### YAML

```yaml

  EnrichmentStrategyConfiguration:
    EnrichmentStrategyConfiguration
  ModelArn: String

```

## Properties

`EnrichmentStrategyConfiguration`

The enrichment stategy used to provide additional context. For example, Neptune GraphRAG uses
Amazon Bedrock foundation models to perform chunk entity extraction.

_Required_: Yes

_Type_: [EnrichmentStrategyConfiguration](aws-properties-bedrock-datasource-enrichmentstrategyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelArn`

The Amazon Resource Name (ARN) of the model used to create vector embeddings for the knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws(-cn|-us-gov|-eusc|-iso(-[b-f])?)?:(bedrock):[a-z0-9-]{1,20}:([0-9]{12})?:([a-z-]+/)?)?([a-zA-Z0-9.-]{1,63}){0,2}(([:][a-z0-9-]{1,63}){0,2})?(/[a-z0-9]{1,12})?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BedrockFoundationModelConfiguration

ChunkingConfiguration

All content copied from https://docs.aws.amazon.com/.
