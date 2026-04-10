This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis UnaggregatedField

The unaggregated field for a table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "FieldId" : String,
  "FormatConfiguration" : FormatConfiguration
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  FieldId: String
  FormatConfiguration:
    FormatConfiguration

```

## Properties

`Column`

The column that is used in the `UnaggregatedField`.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

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

_Type_: [FormatConfiguration](aws-properties-quicksight-analysis-formatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrendArrowOptions

UniqueValuesComputation

All content copied from https://docs.aws.amazon.com/.
