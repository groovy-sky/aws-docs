This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis NumericRangeFilterValue

The value input pf the numeric range filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameter" : String,
  "StaticValue" : Number
}

```

### YAML

```yaml

  Parameter: String
  StaticValue: Number

```

## Properties

`Parameter`

The parameter that is used in the numeric range.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticValue`

The static value of the numeric range filter.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumericRangeFilter

NumericSeparatorConfiguration

All content copied from https://docs.aws.amazon.com/.
