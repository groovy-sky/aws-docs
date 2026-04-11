This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan CustomizedLoadMetricSpecification

`CustomizedLoadMetricSpecification` is a subproperty of [ScalingInstruction](../userguide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.md) that specifies a customized load metric for predictive
scaling to use with a scaling plan.

For predictive scaling to work with a customized load metric specification, AWS Auto Scaling needs access to the `Sum` and `Average` statistics
that CloudWatch computes from metric data.

When you choose a load metric, make sure that the required `Sum` and
`Average` statistics for your metric are available in CloudWatch and that
they provide relevant data for predictive scaling. The `Sum` statistic must
represent the total load on the resource, and the `Average` statistic must
represent the average load per capacity unit of the resource. For example, there is a
metric that counts the number of requests processed by your Auto Scaling group. If the
`Sum` statistic represents the total request count processed by the group,
then the `Average` statistic for the specified metric must represent the average
request count processed by each instance of the group.

If you publish your own metrics, you can aggregate the data points at a given interval
and then publish the aggregated data points to CloudWatch. Before AWS Auto Scaling
generates the forecast, it sums up all the metric data points that occurred within each
hour to match the granularity period that is used in the forecast (60 minutes).

For information about terminology, available metrics, or how to publish new metrics, see
[Amazon CloudWatch\
Concepts](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md) in the _Amazon CloudWatch User Guide_.

After creating your scaling plan, you can use the AWS Auto Scaling console to
visualize forecasts for the specified metric. For more information, see [View\
scaling information for a resource](../../../autoscaling/plans/userguide/gs-create-scaling-plan.md#gs-view-resource) in the _Scaling Plans User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : [ MetricDimension, ... ],
  "MetricName" : String,
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
  Namespace: String
  Statistic: String
  Unit: String

```

## Properties

`Dimensions`

The dimensions of the metric.

Conditional: If you published your metric with dimensions, you must specify the same
dimensions in your customized load metric specification.

_Required_: No

_Type_: Array of [MetricDimension](aws-properties-autoscalingplans-scalingplan-metricdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The statistic of the metric.

_Allowed Values_: `Sum`

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of the metric.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Scaling Plans User Guide](../../../autoscaling/plans/userguide/what-is-a-scaling-plan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationSource

CustomizedScalingMetricSpecification

All content copied from https://docs.aws.amazon.com/.
