This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard YAxisOptions

The options that are available for a single Y axis in a chart.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "YAxis" : String
}

```

### YAML

```yaml

  YAxis: String

```

## Properties

`YAxis`

The Y axis type to be used in the chart.

If you choose `PRIMARY_Y_AXIS`, the primary Y Axis is located on the leftmost vertical axis of the chart.

_Required_: Yes

_Type_: String

_Allowed values_: `PRIMARY_Y_AXIS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WordCloudVisual

AWS::QuickSight::DataSet

All content copied from https://docs.aws.amazon.com/.
