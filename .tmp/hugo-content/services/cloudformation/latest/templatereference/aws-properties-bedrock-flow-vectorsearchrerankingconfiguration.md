This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow VectorSearchRerankingConfiguration

Configuration for reranking vector search results to improve relevance. Reranking applies additional relevance models to reorder the initial vector search results based on more sophisticated criteria.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockRerankingConfiguration" : VectorSearchBedrockRerankingConfiguration,
  "Type" : String
}

```

### YAML

```yaml

  BedrockRerankingConfiguration:
    VectorSearchBedrockRerankingConfiguration
  Type: String

```

## Properties

`BedrockRerankingConfiguration`

Configuration for using Amazon Bedrock foundation models to rerank search results. This is required when the reranking type is set to BEDROCK.

_Required_: No

_Type_: [VectorSearchBedrockRerankingConfiguration](aws-properties-bedrock-flow-vectorsearchbedrockrerankingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of reranking to apply to vector search results. Currently, the only supported value is BEDROCK, which uses Amazon Bedrock foundation models for reranking.

_Required_: Yes

_Type_: String

_Allowed values_: `BEDROCK_RERANKING_MODEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VectorSearchBedrockRerankingModelConfiguration

AWS::Bedrock::FlowAlias

All content copied from https://docs.aws.amazon.com/.
