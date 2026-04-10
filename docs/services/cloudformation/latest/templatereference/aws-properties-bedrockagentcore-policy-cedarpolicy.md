This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Policy CedarPolicy

A Cedar policy statement within the AgentCore Policy system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Statement" : String
}

```

### YAML

```yaml

  Statement: String

```

## Properties

`Statement`

The Cedar policy statement that defines the authorization logic.

_Required_: Yes

_Type_: String

_Minimum_: `35`

_Maximum_: `153600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::Policy

PolicyDefinition

All content copied from https://docs.aws.amazon.com/.
