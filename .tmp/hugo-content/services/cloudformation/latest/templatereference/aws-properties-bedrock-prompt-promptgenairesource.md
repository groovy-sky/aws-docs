This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt PromptGenAiResource

Contains specifications for a generative AI resource with which to use the prompt. For more information, see [Create a prompt using Prompt management](../../../bedrock/latest/userguide/prompt-management-create.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Agent" : PromptAgentResource
}

```

### YAML

```yaml

  Agent:
    PromptAgentResource

```

## Properties

`Agent`

Specifies an Amazon Bedrock agent with which to use the prompt.

_Required_: Yes

_Type_: [PromptAgentResource](aws-properties-bedrock-prompt-promptagentresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptAgentResource

PromptInferenceConfiguration

All content copied from https://docs.aws.amazon.com/.
