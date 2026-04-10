This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AutomatedReasoningPolicy PolicyDefinitionRule

A rule within the policy definition that defines logical constraints.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlternateExpression" : String,
  "Expression" : String,
  "Id" : String
}

```

### YAML

```yaml

  AlternateExpression: String
  Expression: String
  Id: String

```

## Properties

`AlternateExpression`

An alternative expression for the policy rule.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The logical expression that defines the rule.

_Required_: Yes

_Type_: String

_Pattern_: `^[\s\S]+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The unique identifier for the policy definition rule.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z][0-9A-Z]{11}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PolicyDefinition

PolicyDefinitionType

All content copied from https://docs.aws.amazon.com/.
