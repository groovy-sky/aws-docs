This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile MetricToRetain

The metric you want to retain. Dimensions are optional.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExportMetric" : Boolean,
  "Metric" : String,
  "MetricDimension" : MetricDimension
}

```

### YAML

```yaml

  ExportMetric: Boolean
  Metric: String
  MetricDimension:
    MetricDimension

```

## Properties

`ExportMetric`

The value indicates exporting metrics related to the `MetricToRetain` when
it's true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metric`

A standard of measurement.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricDimension`

The dimension of the metric.

_Required_: No

_Type_: [MetricDimension](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-securityprofile-metricdimension.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetricsExportConfig

MetricValue
