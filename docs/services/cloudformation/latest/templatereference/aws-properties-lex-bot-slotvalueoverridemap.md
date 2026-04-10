This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotValueOverrideMap

Maps a slot name to the [SlotValueOverride](../../../../reference/lexv2/latest/apireference/api-slotvalueoverride.md)
object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SlotName" : String,
  "SlotValueOverride" : SlotValueOverride
}

```

### YAML

```yaml

  SlotName: String
  SlotValueOverride:
    SlotValueOverride

```

## Properties

`SlotName`

The name of the slot.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotValueOverride`

The SlotValueOverride object to which the slot name will be mapped.

_Required_: No

_Type_: [SlotValueOverride](aws-properties-lex-bot-slotvalueoverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotValueOverride

SlotValueRegexFilter

All content copied from https://docs.aws.amazon.com/.
