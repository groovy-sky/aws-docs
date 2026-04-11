This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template CustomColor

Determines the color that's applied to a particular data value in a column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "FieldValue" : String,
  "SpecialValue" : String
}

```

### YAML

```yaml

  Color: String
  FieldValue: String
  SpecialValue: String

```

## Properties

`Color`

The color that is applied to the data value.

_Required_: Yes

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldValue`

The data value that the color is applied to.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpecialValue`

The value of a special data value.

_Required_: No

_Type_: String

_Allowed values_: `EMPTY | NULL | OTHER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomActionURLOperation

CustomContentConfiguration

All content copied from https://docs.aws.amazon.com/.
