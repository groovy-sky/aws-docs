This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template AxisLabelReferenceOptions

The reference that specifies where the axis label is applied to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : ColumnIdentifier,
  "FieldId" : String
}

```

### YAML

```yaml

  Column:
    ColumnIdentifier
  FieldId: String

```

## Properties

`Column`

The column that the axis label is targeted to.

_Required_: Yes

_Type_: [ColumnIdentifier](aws-properties-quicksight-template-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The field that the axis label is targeted to.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AxisLabelOptions

AxisLinearScale

All content copied from https://docs.aws.amazon.com/.
