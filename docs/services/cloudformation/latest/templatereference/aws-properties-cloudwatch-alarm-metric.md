This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Alarm Metric

The `Metric` property type represents a specific metric. `Metric` is a property of the
[MetricStat](../userguide/aws-properties-cloudwatch-alarm-metricstat.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ Dimension, ... ],
  "MetricName" : String,
  "Namespace" : String
}

```

### YAML

```yaml

  Dimensions:
    - Dimension
  MetricName: String
  Namespace: String

```

## Properties

`Dimensions`

The metric dimensions that you want to be used for the metric that
the alarm will watch.

_Required_: No

_Type_: Array of [Dimension](aws-properties-cloudwatch-alarm-dimension.md)

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric that you want the alarm to watch. This is a required field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric that the alarm will watch.

_Required_: No

_Type_: String

_Pattern_: `[^:].*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dimension

MetricDataQuery

All content copied from https://docs.aws.amazon.com/.
