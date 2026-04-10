This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow VectorSearchBedrockRerankingConfiguration

Configuration for using Amazon Bedrock foundation models to rerank Knowledge Base vector search results. This enables more sophisticated relevance ranking using large language models.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetadataConfiguration" : MetadataConfigurationForReranking,
  "ModelConfiguration" : VectorSearchBedrockRerankingModelConfiguration,
  "NumberOfRerankedResults" : Number
}

```

### YAML

```yaml

  MetadataConfiguration:
    MetadataConfigurationForReranking
  ModelConfiguration:
    VectorSearchBedrockRerankingModelConfiguration
  NumberOfRerankedResults:
    Number

```

## Properties

`MetadataConfiguration`

Configuration for how document metadata should be used during the reranking process. This determines which metadata fields are included when reordering search results.

_Required_: No

_Type_: [MetadataConfigurationForReranking](aws-properties-bedrock-flow-metadataconfigurationforreranking.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelConfiguration`

Configuration for the Amazon Bedrock foundation model used for reranking. This includes the model ARN and any additional request fields required by the model.

_Required_: Yes

_Type_: [VectorSearchBedrockRerankingModelConfiguration](aws-properties-bedrock-flow-vectorsearchbedrockrerankingmodelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfRerankedResults`

The maximum number of results to rerank. This limits how many of the initial vector search results will be processed by the reranking model. A smaller number improves performance but may exclude potentially relevant results.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextPromptTemplateConfiguration

VectorSearchBedrockRerankingModelConfiguration

All content copied from https://docs.aws.amazon.com/.
