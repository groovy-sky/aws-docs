This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent AgentCollaborator

An agent collaborator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentDescriptor" : AgentDescriptor,
  "CollaborationInstruction" : String,
  "CollaboratorName" : String,
  "RelayConversationHistory" : String
}

```

### YAML

```yaml

  AgentDescriptor:
    AgentDescriptor
  CollaborationInstruction: String
  CollaboratorName: String
  RelayConversationHistory: String

```

## Properties

`AgentDescriptor`

The collaborator's agent descriptor.

_Required_: Yes

_Type_: [AgentDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-agentdescriptor.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollaborationInstruction`

The collaborator's instructions.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollaboratorName`

The collaborator's collaborator name.

_Required_: Yes

_Type_: String

_Pattern_: `([0-9a-zA-Z][_-]?){1,100}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelayConversationHistory`

The collaborator's relay conversation history.

_Required_: No

_Type_: String

_Allowed values_: `TO_COLLABORATOR | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AgentActionGroup

AgentDescriptor
