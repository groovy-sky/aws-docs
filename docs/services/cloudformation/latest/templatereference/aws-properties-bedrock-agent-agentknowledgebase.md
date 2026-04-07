This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent AgentKnowledgeBase

Contains details about a knowledge base that is associated with an agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "KnowledgeBaseId" : String,
  "KnowledgeBaseState" : String
}

```

### YAML

```yaml

  Description: String
  KnowledgeBaseId: String
  KnowledgeBaseState: String

```

## Properties

`Description`

The description of the association between the agent and the knowledge base.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBaseId`

The unique identifier of the association between the agent and the knowledge base.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z]{10}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBaseState`

Specifies whether to use the knowledge base or not when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AgentDescriptor

APISchema
