This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy CustomizedMetricSpecification

Contains customized metric specification information for a target tracking scaling policy
for Application Auto Scaling.

For information about the available metrics for a service, see [AWS services that publish CloudWatch metrics](../../../amazoncloudwatch/latest/monitoring/aws-services-cloudwatch-metrics.md) in the _Amazon_
_CloudWatch User Guide_.

To create your customized metric specification:

- Add values for each required parameter from CloudWatch. You can use an existing
metric, or a new metric that you create. To use your own metric, you must first publish
the metric to CloudWatch. For more information, see [Publish custom\
metrics](../../../amazoncloudwatch/latest/monitoring/publishingmetrics.md) in the _Amazon CloudWatch User Guide_.

- Choose a metric that changes proportionally with capacity. The value of the metric
should increase or decrease in inverse proportion to the number of capacity units. That
is, the value of the metric should decrease when capacity increases, and increase when
capacity decreases.

For an example of how creating new metrics can be useful, see [Scaling based on Amazon SQS](../../../autoscaling/ec2/userguide/as-using-sqs-queue.md)
in the _Amazon EC2 Auto Scaling User Guide_. This topic mentions Auto
Scaling groups, but the same scenario for Amazon SQS can apply to the target tracking scaling
policies that you create for a Spot Fleet by using Application Auto Scaling.

For more information about the CloudWatch terminology below, see [Amazon CloudWatch concepts](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md).

`CustomizedMetricSpecification` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingScalingPolicyConfiguration](../userguide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.md)
property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ MetricDimension, ... ],
  "MetricName" : String,
  "Metrics" : [ TargetTrackingMetricDataQuery, ... ],
  "Namespace" : String,
  "Statistic" : String,
  "Unit" : String
}

```

### YAML

```yaml

  Dimensions:
    - MetricDimension
  MetricName: String
  Metrics:
    - TargetTrackingMetricDataQuery
  Namespace: String
  Statistic: String
  Unit: String

```

## Properties

`Dimensions`

The dimensions of the metric.

Conditional: If you published your metric with dimensions, you must specify the same
dimensions in your scaling policy.

_Required_: No

_Type_: Array of [MetricDimension](aws-properties-applicationautoscaling-scalingpolicy-metricdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric. To get the exact metric name, namespace, and dimensions, inspect
the [Metric](../../../../reference/amazoncloudwatch/latest/apireference/api-metric.md) object that's returned by a call to [ListMetrics](../../../../reference/amazoncloudwatch/latest/apireference/api-listmetrics.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metrics`

The metrics to include in the target tracking scaling policy, as a metric data query.
This can include both raw metric and metric math expressions.

_Required_: No

_Type_: Array of [TargetTrackingMetricDataQuery](aws-properties-applicationautoscaling-scalingpolicy-targettrackingmetricdataquery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The statistic of the metric.

_Required_: No

_Type_: String

_Allowed values_: `Average | Minimum | Maximum | SampleCount | Sum`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of the metric. For a complete list of the units that CloudWatch supports, see the
[MetricDatum](../../../../reference/amazoncloudwatch/latest/apireference/api-metricdatum.md) data
type in the _Amazon CloudWatch API Reference_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md)

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApplicationAutoScaling::ScalingPolicy

MetricDimension

All content copied from https://docs.aws.amazon.com/.
