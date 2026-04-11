This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow PromptFlowNodeInlineConfiguration

Contains configurations for a prompt defined inline in the node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InferenceConfiguration" : PromptInferenceConfiguration,
  "ModelId" : String,
  "TemplateConfiguration" : PromptTemplateConfiguration,
  "TemplateType" : String
}

```

### YAML

```yaml

  InferenceConfiguration:
    PromptInferenceConfiguration
  ModelId: String
  TemplateConfiguration:
    PromptTemplateConfiguration
  TemplateType: String

```

## Properties

`InferenceConfiguration`

Contains inference configurations for the prompt.

_Required_: No

_Type_: [PromptInferenceConfiguration](aws-properties-bedrock-flow-promptinferenceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The unique identifier of the model or [inference profile](../../../bedrock/latest/userguide/cross-region-inference.md) to run inference with.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws(-[^:]{1,12})?:(bedrock|sagemaker):[a-z0-9-]{1,20}:([0-9]{12})?:([a-z-]+/)?)?([a-zA-Z0-9.-]{1,63}){0,2}(([:][a-z0-9-]{1,63}){0,2})?(/[a-z0-9]{1,12})?$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateConfiguration`

Contains a prompt and variables in the prompt that can be replaced with values at runtime.

_Required_: Yes

_Type_: [PromptTemplateConfiguration](aws-properties-bedrock-flow-prompttemplateconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateType`

The type of prompt template.

_Required_: Yes

_Type_: String

_Allowed values_: `TEXT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptFlowNodeConfiguration

PromptFlowNodeResourceConfiguration

All content copied from https://docs.aws.amazon.com/.
