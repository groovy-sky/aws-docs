This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Prompt PromptAgentResource

Contains specifications for an Amazon Bedrock agent with which to use the prompt. For more information, see [Create a prompt using Prompt management](../../../bedrock/latest/userguide/prompt-management-create.md) and [Automate tasks in your application using conversational agents](../../../bedrock/latest/userguide/agents.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentIdentifier" : String
}

```

### YAML

```yaml

  AgentIdentifier: String

```

## Properties

`AgentIdentifier`

The ARN of the agent with which to use the prompt.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:agent-alias/[0-9a-zA-Z]{10}/[0-9a-zA-Z]{10}$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Message

PromptGenAiResource

All content copied from https://docs.aws.amazon.com/.
