This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow KnowledgeBaseFlowNodeConfiguration

Contains configurations for a knowledge base node in a flow. This node takes a query as the input and returns, as the output, the retrieved responses directly (as an array) or a response generated based on the retrieved responses. For more information, see [Node types in a flow](../../../bedrock/latest/userguide/flows-nodes.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GuardrailConfiguration" : GuardrailConfiguration,
  "InferenceConfiguration" : PromptInferenceConfiguration,
  "KnowledgeBaseId" : String,
  "ModelId" : String,
  "NumberOfResults" : Number,
  "OrchestrationConfiguration" : KnowledgeBaseOrchestrationConfiguration,
  "PromptTemplate" : KnowledgeBasePromptTemplate,
  "RerankingConfiguration" : VectorSearchRerankingConfiguration
}

```

### YAML

```yaml

  GuardrailConfiguration:
    GuardrailConfiguration
  InferenceConfiguration:
    PromptInferenceConfiguration
  KnowledgeBaseId: String
  ModelId: String
  NumberOfResults:
    Number
  OrchestrationConfiguration:
    KnowledgeBaseOrchestrationConfiguration
  PromptTemplate:
    KnowledgeBasePromptTemplate
  RerankingConfiguration:
    VectorSearchRerankingConfiguration

```

## Properties

`GuardrailConfiguration`

Contains configurations for a guardrail to apply during query and response generation for the knowledge base in this configuration.

_Required_: No

_Type_: [GuardrailConfiguration](aws-properties-bedrock-flow-guardrailconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceConfiguration`

Contains inference configurations for the prompt.

_Required_: No

_Type_: [PromptInferenceConfiguration](aws-properties-bedrock-flow-promptinferenceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBaseId`

The unique identifier of the knowledge base to query.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z]+$`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The unique identifier of the model or [inference profile](../../../bedrock/latest/userguide/cross-region-inference.md) to use to generate a response from the query results. Omit this field if you want to return the retrieved results as an array.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws(-[^:]{1,12})?:(bedrock|sagemaker):[a-z0-9-]{1,20}:([0-9]{12})?:([a-z-]+/)?)?([a-zA-Z0-9.-]{1,63}){0,2}(([:][a-z0-9-]{1,63}){0,2})?(/[a-z0-9]{1,12})?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfResults`

The number of results to retrieve from the knowledge base.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrchestrationConfiguration`

The configuration for orchestrating the retrieval and generation process in the knowledge base node.

_Required_: No

_Type_: [KnowledgeBaseOrchestrationConfiguration](aws-properties-bedrock-flow-knowledgebaseorchestrationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptTemplate`

A custom prompt template to use with the knowledge base for generating responses.

_Required_: No

_Type_: [KnowledgeBasePromptTemplate](aws-properties-bedrock-flow-knowledgebaseprompttemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RerankingConfiguration`

The configuration for reranking the retrieved results from the knowledge base to improve relevance.

_Required_: No

_Type_: [VectorSearchRerankingConfiguration](aws-properties-bedrock-flow-vectorsearchrerankingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InlineCodeFlowNodeConfiguration

KnowledgeBaseOrchestrationConfiguration

All content copied from https://docs.aws.amazon.com/.
