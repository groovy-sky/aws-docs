This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SubSlotSetting

Specifications for the constituent sub slots and the
expression for the composite slot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "SlotSpecifications" : {Key: Value, ...}
}

```

### YAML

```yaml

  Expression: String
  SlotSpecifications:
    Key: Value

```

## Properties

`Expression`

The expression text for defining the constituent sub slots in the composite slot using logical AND and OR operators.

_Required_: No

_Type_: String

_Pattern_: `[0-9A-Za-z_\-\s\(\)]+`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotSpecifications`

Specifications for the constituent sub slots of a composite slot.

_Required_: No

_Type_: Object of [Specifications](aws-properties-lex-bot-specifications.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StillWaitingResponseSpecification

SubSlotTypeComposition

All content copied from https://docs.aws.amazon.com/.
