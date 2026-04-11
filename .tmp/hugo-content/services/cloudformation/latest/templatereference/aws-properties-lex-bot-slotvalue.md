This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotValue

The value to set in a slot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InterpretedValue" : String
}

```

### YAML

```yaml

  InterpretedValue: String

```

## Properties

`InterpretedValue`

The value that Amazon Lex determines for the slot. The
actual value depends on the setting of the value selection strategy for
the bot. You can choose to use the value entered by the user, or you
can have Amazon Lex choose the first value in the
`resolvedValues` list.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `202`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotTypeValue

SlotValueElicitationSetting

All content copied from https://docs.aws.amazon.com/.
