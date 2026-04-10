This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingMetric

Represents a specific metric for a target tracking scaling policy for Application Auto
Scaling.

Metric is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingMetricStat](../userguide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricstat.md) property
type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ TargetTrackingMetricDimension, ... ],
  "MetricName" : String,
  "Namespace" : String
}

```

### YAML

```yaml

  Dimensions:
    - TargetTrackingMetricDimension
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

_Type_: Array of [TargetTrackingMetricDimension](aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric. For more information, see the table in [AWS\
services that publish CloudWatch metrics](../../../amazoncloudwatch/latest/monitoring/aws-services-cloudwatch-metrics.md) in the _Amazon CloudWatch User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StepScalingPolicyConfiguration

TargetTrackingMetricDataQuery

All content copied from https://docs.aws.amazon.com/.
