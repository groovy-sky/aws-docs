This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CategoricalDimensionField

The dimension type field with categorical type columns..

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "FieldId" : String,
  "FormatConfiguration" : StringFormatConfiguration,
  "HierarchyId" : String
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  FieldId: String
  FormatConfiguration:
    StringFormatConfiguration
  HierarchyId: String

```

## Properties

`Column`

The column that is used in the `CategoricalDimensionField`.

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

_Type_: [StringFormatConfiguration](aws-properties-quicksight-analysis-stringformatconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyId`

The custom hierarchy ID.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CascadingControlSource

CategoricalMeasureField

All content copied from https://docs.aws.amazon.com/.
