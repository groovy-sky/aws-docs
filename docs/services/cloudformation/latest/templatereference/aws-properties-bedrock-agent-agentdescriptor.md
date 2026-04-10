This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent AgentDescriptor

An agent descriptor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AliasArn" : String
}

```

### YAML

```yaml

  AliasArn: String

```

## Properties

`AliasArn`

The agent's alias ARN.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:agent-alias/[0-9a-zA-Z]{10}/[0-9a-zA-Z]{10}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AgentCollaborator

AgentKnowledgeBase

All content copied from https://docs.aws.amazon.com/.
