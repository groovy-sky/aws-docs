This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::EnforcedGuardrailConfiguration SelectiveContentGuarding

Selective content guarding controls for enforced guardrails.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Messages" : String,
  "System" : String
}

```

### YAML

```yaml

  Messages: String
  System: String

```

## Properties

`Messages`

Selective guarding mode for user messages.

_Required_: No

_Type_: String

_Allowed values_: `SELECTIVE | COMPREHENSIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`System`

Selective guarding mode for system prompts."

_Required_: No

_Type_: String

_Allowed values_: `SELECTIVE | COMPREHENSIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelEnforcement

AWS::Bedrock::Flow

All content copied from https://docs.aws.amazon.com/.
