This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SubSlotTypeComposition

Subslot type composition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SlotTypeId" : String,
  "SlotTypeName" : String
}

```

### YAML

```yaml

  Name: String
  SlotTypeId: String
  SlotTypeName: String

```

## Properties

`Name`

Name of a constituent sub slot inside a composite slot.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotTypeId`

The unique identifier assigned to a slot type.
This refers to either a built-in slot type or the unique slotTypeId of a custom slot type.

_Required_: No

_Type_: String

_Pattern_: `^((AMAZON\.)[a-zA-Z_]+?|[0-9a-zA-Z]+)$`

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotTypeName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubSlotSetting

SubSlotValueElicitationSetting

All content copied from https://docs.aws.amazon.com/.
