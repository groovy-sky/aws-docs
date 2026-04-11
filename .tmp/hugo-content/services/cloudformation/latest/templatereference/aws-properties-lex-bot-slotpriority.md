This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotPriority

Sets the priority that Amazon Lex should use when eliciting slot values
from a user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Priority" : Integer,
  "SlotName" : String
}

```

### YAML

```yaml

  Priority: Integer
  SlotName: String

```

## Properties

`Priority`

The priority that Amazon Lex should apply to the slot.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotName`

The name of the slot.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotDefaultValueSpecification

SlotResolutionImprovementSpecification

All content copied from https://docs.aws.amazon.com/.
