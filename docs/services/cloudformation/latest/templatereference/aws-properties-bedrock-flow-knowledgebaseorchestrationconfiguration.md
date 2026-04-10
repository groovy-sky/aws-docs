This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow KnowledgeBaseOrchestrationConfiguration

Configures how the knowledge base orchestrates the retrieval and generation process, allowing for customization of prompts, inference parameters, and performance settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalModelRequestFields" : Json,
  "InferenceConfig" : PromptInferenceConfiguration,
  "PerformanceConfig" : PerformanceConfiguration,
  "PromptTemplate" : KnowledgeBasePromptTemplate
}

```

### YAML

```yaml

  AdditionalModelRequestFields: Json
  InferenceConfig:
    PromptInferenceConfiguration
  PerformanceConfig:
    PerformanceConfiguration
  PromptTemplate:
    KnowledgeBasePromptTemplate

```

## Properties

`AdditionalModelRequestFields`

The additional model-specific request parameters as key-value pairs to be included in the request to the foundation model.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceConfig`

Contains inference configurations for the prompt.

_Required_: No

_Type_: [PromptInferenceConfiguration](aws-properties-bedrock-flow-promptinferenceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceConfig`

The performance configuration options for the knowledge base retrieval and generation process.

_Required_: No

_Type_: [PerformanceConfiguration](aws-properties-bedrock-flow-performanceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptTemplate`

A custom prompt template for orchestrating the retrieval and generation process.

_Required_: No

_Type_: [KnowledgeBasePromptTemplate](aws-properties-bedrock-flow-knowledgebaseprompttemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KnowledgeBaseFlowNodeConfiguration

KnowledgeBasePromptTemplate

All content copied from https://docs.aws.amazon.com/.
