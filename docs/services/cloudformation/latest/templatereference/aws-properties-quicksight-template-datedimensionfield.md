This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DateDimensionField

The dimension type field with date type columns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "DateGranularity" : String,
  "FieldId" : String,
  "FormatConfiguration" : DateTimeFormatConfiguration,
  "HierarchyId" : String
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  DateGranularity: String
  FieldId: String
  FormatConfiguration:
    DateTimeFormatConfiguration
  HierarchyId: String

```

## Properties

`Column`

The column that is used in the `DateDimensionField`.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-template-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateGranularity`

The date granularity of the `DateDimensionField`. Choose one of the following options:

- `YEAR`

- `QUARTER`

- `MONTH`

- `WEEK`

- `DAY`

- `HOUR`

- `MINUTE`

- `SECOND`

- `MILLISECOND`

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The custom field ID.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatConfiguration`

The format configuration of the field.

_Required_: No

_Type_: [DateTimeFormatConfiguration](aws-properties-quicksight-template-datetimeformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyId`

The custom hierarchy ID.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateAxisOptions

DateMeasureField

All content copied from https://docs.aws.amazon.com/.
