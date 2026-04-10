This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy Metric

Represents a specific metric.

`Metric` is a property of the [AWS::AutoScaling::ScalingPolicy MetricStat](../userguide/aws-properties-autoscaling-scalingpolicy-metricstat.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ MetricDimension, ... ],
  "MetricName" : String,
  "Namespace" : String
}

```

### YAML

```yaml

  Dimensions:
    - MetricDimension
  MetricName: String
  Namespace: String

```

## Properties

`Dimensions`

The dimensions for the metric. For the list of available dimensions, see the AWS
documentation available from the table in [AWS\
services that publish CloudWatch metrics](../../../amazoncloudwatch/latest/monitoring/aws-services-cloudwatch-metrics.md) in the _Amazon CloudWatch User_
_Guide_.

Conditional: If you published your metric with dimensions, you must specify the same
dimensions in your scaling policy.

_Required_: No

_Type_: Array of [MetricDimension](aws-properties-autoscaling-scalingpolicy-metricdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric. For more information, see the table in [AWS\
services that publish CloudWatch metrics](../../../amazoncloudwatch/latest/monitoring/aws-services-cloudwatch-metrics.md) in the _Amazon CloudWatch User_
_Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomizedMetricSpecification

MetricDataQuery

All content copied from https://docs.aws.amazon.com/.
