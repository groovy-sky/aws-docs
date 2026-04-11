This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ItemsLimitConfiguration

The limit configuration of the visual display for an axis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ItemsLimit" : Number,
  "OtherCategories" : String
}

```

### YAML

```yaml

  ItemsLimit: Number
  OtherCategories: String

```

## Properties

`ItemsLimit`

The limit on how many items of a field are showed in the chart. For
example, the number of slices that are displayed in a pie chart.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OtherCategories`

The `Show
                other` of an axis in the chart. Choose one of the following options:

- `INCLUDE`

- `EXCLUDE`

_Required_: No

_Type_: String

_Allowed values_: `INCLUDE | EXCLUDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntegerValueWhenUnsetConfiguration

KPIActualValueConditionalFormatting

All content copied from https://docs.aws.amazon.com/.
