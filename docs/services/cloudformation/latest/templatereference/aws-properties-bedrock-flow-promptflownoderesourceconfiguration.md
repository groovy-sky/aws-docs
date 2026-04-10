This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow PromptFlowNodeResourceConfiguration

Contains configurations for a prompt from Prompt management to use in a node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PromptArn" : String
}

```

### YAML

```yaml

  PromptArn: String

```

## Properties

`PromptArn`

The Amazon Resource Name (ARN) of the prompt from Prompt management.

_Required_: Yes

_Type_: String

_Pattern_: `^(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:prompt/[0-9a-zA-Z]{10}(?::[0-9]{1,5})?)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptFlowNodeInlineConfiguration

PromptFlowNodeSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
