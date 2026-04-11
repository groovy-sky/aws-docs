This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DecimalValueWhenUnsetConfiguration

The configuration that defines the default value of a `Decimal` parameter when a value has not been set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomValue" : Number,
  "ValueWhenUnsetOption" : String
}

```

### YAML

```yaml

  CustomValue: Number
  ValueWhenUnsetOption: String

```

## Properties

`CustomValue`

A custom value that's used when the value of a parameter isn't set.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueWhenUnsetOption`

The built-in options for default values. The value can be one of the following:

- `RECOMMENDED`: The recommended value.

- `NULL`: The `NULL` value.

_Required_: No

_Type_: String

_Allowed values_: `RECOMMENDED_VALUE | NULL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecimalPlacesConfiguration

DefaultDateTimePickerControlOptions

All content copied from https://docs.aws.amazon.com/.
