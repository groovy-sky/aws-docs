This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ControlTower::EnabledControl EnabledControlParameter

A set of parameters that configure the behavior of the enabled control. Expressed as a key/value pair, where `Key` is of type `String` and `Value` is of type `Document`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Value:
    - String

```

## Properties

`Key`

The key of a key/value pair. It is of type `string`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a key/value pair. It can be of type `array`, `string`, `number`, `object`, or `boolean`. \[Note: The _Type_ field that follows may show a single type such as Number, which is only one possible type.\]

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ControlTower::EnabledControl

Tag

All content copied from https://docs.aws.amazon.com/.
