This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow PromptFlowNodeSourceConfiguration

Contains configurations for a prompt and whether it is from Prompt management or defined inline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Inline" : PromptFlowNodeInlineConfiguration,
  "Resource" : PromptFlowNodeResourceConfiguration
}

```

### YAML

```yaml

  Inline:
    PromptFlowNodeInlineConfiguration
  Resource:
    PromptFlowNodeResourceConfiguration

```

## Properties

`Inline`

Contains configurations for a prompt that is defined inline

_Required_: No

_Type_: [PromptFlowNodeInlineConfiguration](aws-properties-bedrock-flow-promptflownodeinlineconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resource`

Contains configurations for a prompt from Prompt management.

_Required_: No

_Type_: [PromptFlowNodeResourceConfiguration](aws-properties-bedrock-flow-promptflownoderesourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptFlowNodeResourceConfiguration

PromptInferenceConfiguration

All content copied from https://docs.aws.amazon.com/.
