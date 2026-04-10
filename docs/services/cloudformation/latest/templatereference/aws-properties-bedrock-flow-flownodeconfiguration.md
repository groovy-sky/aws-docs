This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow FlowNodeConfiguration

Contains configurations for a node in your flow. For more information, see [Node types in a flow](../../../bedrock/latest/userguide/flows-nodes.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Agent" : AgentFlowNodeConfiguration,
  "Collector" : Json,
  "Condition" : ConditionFlowNodeConfiguration,
  "InlineCode" : InlineCodeFlowNodeConfiguration,
  "Input" : Json,
  "Iterator" : Json,
  "KnowledgeBase" : KnowledgeBaseFlowNodeConfiguration,
  "LambdaFunction" : LambdaFunctionFlowNodeConfiguration,
  "Lex" : LexFlowNodeConfiguration,
  "Loop" : LoopFlowNodeConfiguration,
  "LoopController" : LoopControllerFlowNodeConfiguration,
  "LoopInput" : Json,
  "Output" : Json,
  "Prompt" : PromptFlowNodeConfiguration,
  "Retrieval" : RetrievalFlowNodeConfiguration,
  "Storage" : StorageFlowNodeConfiguration
}

```

### YAML

```yaml

  Agent:
    AgentFlowNodeConfiguration
  Collector: Json
  Condition:
    ConditionFlowNodeConfiguration
  InlineCode:
    InlineCodeFlowNodeConfiguration
  Input: Json
  Iterator: Json
  KnowledgeBase:
    KnowledgeBaseFlowNodeConfiguration
  LambdaFunction:
    LambdaFunctionFlowNodeConfiguration
  Lex:
    LexFlowNodeConfiguration
  Loop:
    LoopFlowNodeConfiguration
  LoopController:
    LoopControllerFlowNodeConfiguration
  LoopInput: Json
  Output: Json
  Prompt:
    PromptFlowNodeConfiguration
  Retrieval:
    RetrievalFlowNodeConfiguration
  Storage:
    StorageFlowNodeConfiguration

```

## Properties

`Agent`

Contains configurations for an agent node in your flow. Invokes an alias of an agent and returns the response.

_Required_: No

_Type_: [AgentFlowNodeConfiguration](aws-properties-bedrock-flow-agentflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Collector`

Contains configurations for a collector node in your flow. Collects an iteration of inputs and consolidates them into an array of outputs.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

Contains configurations for a condition node in your flow. Defines conditions that lead to different branches of the flow.

_Required_: No

_Type_: [ConditionFlowNodeConfiguration](aws-properties-bedrock-flow-conditionflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InlineCode`

Contains configurations for an inline code node in your flow. Inline code nodes let you write and execute code directly within your flow, enabling data transformations, custom logic, and integrations without needing an external Lambda function.

_Required_: No

_Type_: [InlineCodeFlowNodeConfiguration](aws-properties-bedrock-flow-inlinecodeflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Input`

Contains configurations for an input flow node in your flow. The first node in the flow. `inputs` can't be specified for this node.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Iterator`

Contains configurations for an iterator node in your flow. Takes an input that is an array and iteratively sends each item of the array as an output to the following node. The size of the array is also returned in the output.

The output flow node at the end of the flow iteration will return a response for each member of the array. To return only one response, you can include a collector node downstream from the iterator node.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBase`

Contains configurations for a knowledge base node in your flow. Queries a knowledge base and returns the retrieved results or generated response.

_Required_: No

_Type_: [KnowledgeBaseFlowNodeConfiguration](aws-properties-bedrock-flow-knowledgebaseflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunction`

Contains configurations for a Lambda function node in your flow. Invokes an AWS Lambda function.

_Required_: No

_Type_: [LambdaFunctionFlowNodeConfiguration](aws-properties-bedrock-flow-lambdafunctionflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lex`

Contains configurations for a Lex node in your flow. Invokes an Amazon Lex bot to identify the intent of the input and return the intent as the output.

_Required_: No

_Type_: [LexFlowNodeConfiguration](aws-properties-bedrock-flow-lexflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Loop`

Contains configurations for a DoWhile loop in your flow.

_Required_: No

_Type_: [LoopFlowNodeConfiguration](aws-properties-bedrock-flow-loopflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoopController`

Contains controller node configurations for a DoWhile loop in your flow.

_Required_: No

_Type_: [LoopControllerFlowNodeConfiguration](aws-properties-bedrock-flow-loopcontrollerflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoopInput`

Contains input node configurations for a DoWhile loop in your flow.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Output`

Contains configurations for an output flow node in your flow. The last node in the flow. `outputs` can't be specified for this node.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prompt`

Contains configurations for a prompt node in your flow. Runs a prompt and generates the model response as the output. You can use a prompt from Prompt management or you can configure one in this node.

_Required_: No

_Type_: [PromptFlowNodeConfiguration](aws-properties-bedrock-flow-promptflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Retrieval`

Contains configurations for a retrieval node in your flow. Retrieves data from an Amazon S3 location and returns it as the output.

_Required_: No

_Type_: [RetrievalFlowNodeConfiguration](aws-properties-bedrock-flow-retrievalflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Storage`

Contains configurations for a storage node in your flow. Stores an input in an Amazon S3 location.

_Required_: No

_Type_: [StorageFlowNodeConfiguration](aws-properties-bedrock-flow-storageflownodeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowNode

FlowNodeInput

All content copied from https://docs.aws.amazon.com/.
