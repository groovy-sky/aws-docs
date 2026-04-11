This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DataLabelType

The option that determines the data label type.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataPathLabelType" : DataPathLabelType,
  "FieldLabelType" : FieldLabelType,
  "MaximumLabelType" : MaximumLabelType,
  "MinimumLabelType" : MinimumLabelType,
  "RangeEndsLabelType" : RangeEndsLabelType
}

```

### YAML

```yaml

  DataPathLabelType:
    DataPathLabelType
  FieldLabelType:
    FieldLabelType
  MaximumLabelType:
    MaximumLabelType
  MinimumLabelType:
    MinimumLabelType
  RangeEndsLabelType:
    RangeEndsLabelType

```

## Properties

`DataPathLabelType`

The option that specifies individual data values for labels.

_Required_: No

_Type_: [DataPathLabelType](aws-properties-quicksight-dashboard-datapathlabeltype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldLabelType`

Determines the label configuration for the entire field.

_Required_: No

_Type_: [FieldLabelType](aws-properties-quicksight-dashboard-fieldlabeltype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumLabelType`

Determines the label configuration for the maximum value in a visual.

_Required_: No

_Type_: [MaximumLabelType](aws-properties-quicksight-dashboard-maximumlabeltype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumLabelType`

Determines the label configuration for the minimum value in a visual.

_Required_: No

_Type_: [MinimumLabelType](aws-properties-quicksight-dashboard-minimumlabeltype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeEndsLabelType`

Determines the label configuration for range end value in a visual.

_Required_: No

_Type_: [RangeEndsLabelType](aws-properties-quicksight-dashboard-rangeendslabeltype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataLabelOptions

DataPathColor

All content copied from https://docs.aws.amazon.com/.
