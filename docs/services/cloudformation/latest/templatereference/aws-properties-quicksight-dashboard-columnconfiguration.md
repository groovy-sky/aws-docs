This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ColumnConfiguration

The general configuration of a column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorsConfiguration" : ColorsConfiguration,
  "Column" : ColumnIdentifier,
  "FormatConfiguration" : FormatConfiguration,
  "Role" : String
}

```

### YAML

```yaml

  ColorsConfiguration:
    ColorsConfiguration
  Column:
    ColumnIdentifier
  FormatConfiguration:
    FormatConfiguration
  Role: String

```

## Properties

`ColorsConfiguration`

The color configurations of the column.

_Required_: No

_Type_: [ColorsConfiguration](aws-properties-quicksight-dashboard-colorsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Column`

The column.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatConfiguration`

The format configuration of a column.

_Required_: No

_Type_: [FormatConfiguration](aws-properties-quicksight-dashboard-formatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The role of the column.

_Required_: No

_Type_: String

_Allowed values_: `DIMENSION | MEASURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColorsConfiguration

ColumnHierarchy

All content copied from https://docs.aws.amazon.com/.
