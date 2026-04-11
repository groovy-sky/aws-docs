This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot CompositeSlotTypeSetting

A composite slot is a combination of two or more slots
that capture multiple pieces of information in a single user input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SubSlots" : [ SubSlotTypeComposition, ... ]
}

```

### YAML

```yaml

  SubSlots:
    - SubSlotTypeComposition

```

## Properties

`SubSlots`

Subslots in the composite slot.

_Required_: No

_Type_: Array of [SubSlotTypeComposition](aws-properties-lex-bot-subslottypecomposition.md)

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeHookSpecification

Condition

All content copied from https://docs.aws.amazon.com/.
