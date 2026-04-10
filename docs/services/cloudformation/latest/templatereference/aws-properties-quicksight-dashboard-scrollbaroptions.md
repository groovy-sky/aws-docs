This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ScrollBarOptions

The visual display options for a data zoom scroll bar.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Visibility" : String,
  "VisibleRange" : VisibleRangeOptions
}

```

### YAML

```yaml

  Visibility: String
  VisibleRange:
    VisibleRangeOptions

```

## Properties

`Visibility`

The visibility of the data zoom scroll bar.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisibleRange`

The visibility range for the data zoom scroll bar.

_Required_: No

_Type_: [VisibleRangeOptions](aws-properties-quicksight-dashboard-visiblerangeoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScatterPlotVisual

SecondaryValueOptions

All content copied from https://docs.aws.amazon.com/.
