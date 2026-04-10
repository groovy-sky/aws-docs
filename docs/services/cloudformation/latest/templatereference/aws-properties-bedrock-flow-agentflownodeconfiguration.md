This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow AgentFlowNodeConfiguration

Defines an agent node in your flow. You specify the agent to invoke at this point in the flow. For more information, see [Node types in a flow](../../../bedrock/latest/userguide/flows-nodes.md) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentAliasArn" : String
}

```

### YAML

```yaml

  AgentAliasArn: String

```

## Properties

`AgentAliasArn`

The Amazon Resource Name (ARN) of the alias of the agent to invoke.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:agent-alias/[0-9a-zA-Z]{10}/[0-9a-zA-Z]{10}$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::Flow

ConditionFlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
