This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::DataTableAttribute Validation

Defines validation rules for data table attribute values. Based on JSON Schema Draft 2020-12 with additional
Connect-specific validations. Validation rules ensure data integrity and consistency across the data table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enum" : Enum,
  "ExclusiveMaximum" : Number,
  "ExclusiveMinimum" : Number,
  "Maximum" : Number,
  "MaxLength" : Integer,
  "MaxValues" : Integer,
  "Minimum" : Number,
  "MinLength" : Integer,
  "MinValues" : Integer,
  "MultipleOf" : Number
}

```

### YAML

```yaml

  Enum:
    Enum
  ExclusiveMaximum: Number
  ExclusiveMinimum: Number
  Maximum: Number
  MaxLength: Integer
  MaxValues: Integer
  Minimum: Number
  MinLength: Integer
  MinValues: Integer
  MultipleOf: Number

```

## Properties

`Enum`

Defines enumeration constraints for attribute values. Can specify a list of allowed values and whether custom
values are permitted beyond the enumerated list.

_Required_: No

_Type_: [Enum](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-datatableattribute-enum.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusiveMaximum`

The largest exclusive numeric value for NUMBER value type. Can be provided alongside Maximum where both operate
independently. Must be greater than ExclusiveMinimum and Minimum. Applies to NUMBER and values within
NUMBER\_LIST.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusiveMinimum`

The smallest exclusive numeric value for NUMBER value type. Can be provided alongside Minimum where both operate
independently. Must be less than ExclusiveMaximum and Maximum. Applies to NUMBER and values within
NUMBER\_LIST.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Maximum`

The largest inclusive numeric value for NUMBER value type. Can be provided alongside ExclusiveMaximum where both
operate independently. Must be greater than or equal to Minimum and greater than ExclusiveMinimum. Applies to NUMBER
and values within NUMBER\_LIST.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLength`

The maximum number of characters a text value can contain. Applies to TEXT value type and values within a
TEXT\_LIST. Must be greater than or equal to MinLength.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxValues`

The maximum number of values in a list. Must be an integer greater than or equal to 0 and greater than or equal
to MinValues. Applies to all list types.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Minimum`

The smallest inclusive numeric value for NUMBER value type. Cannot be provided when ExclusiveMinimum is also
provided. Must be less than or equal to Maximum and less than ExclusiveMaximum. Applies to NUMBER and values within
NUMBER\_LIST.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinLength`

The minimum number of characters a text value can contain. Applies to TEXT value type and values within a
TEXT\_LIST. Must be less than or equal to MaxLength.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinValues`

The minimum number of values in a list. Must be an integer greater than or equal to 0 and less than or equal to
MaxValues. Applies to all list types.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultipleOf`

Specifies that numeric values must be multiples of this number. Must be greater than 0. The result of dividing a
value by this multiple must result in an integer. Applies to NUMBER and values within NUMBER\_LIST.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LockVersion

AWS::Connect::DataTableRecord
