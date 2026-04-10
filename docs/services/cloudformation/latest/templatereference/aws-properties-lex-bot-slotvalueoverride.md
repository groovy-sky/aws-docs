This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotValueOverride

The slot values that Amazon Lex uses when it sets slot
values in a dialog step.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Shape" : String,
  "Value" : SlotValue,
  "Values" : [ SlotValueOverride, ... ]
}

```

### YAML

```yaml

  Shape: String
  Value:
    SlotValue
  Values:
    - SlotValueOverride

```

## Properties

`Shape`

When the shape value is `List`, it indicates that the
`values` field contains a list of slot values. When the
value is `Scalar`, it indicates that the `value`
field contains a single value.

_Required_: No

_Type_: String

_Allowed values_: `Scalar | List`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The current value of the slot.

_Required_: No

_Type_: [SlotValue](aws-properties-lex-bot-slotvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

A list of one or more values that the user provided for the slot.
For example, for a slot that elicits pizza toppings, the values
might be "pepperoni" and "pineapple."

_Required_: No

_Type_: Array of [SlotValueOverride](aws-properties-lex-bot-slotvalueoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotValueElicitationSetting

SlotValueOverrideMap

All content copied from https://docs.aws.amazon.com/.
