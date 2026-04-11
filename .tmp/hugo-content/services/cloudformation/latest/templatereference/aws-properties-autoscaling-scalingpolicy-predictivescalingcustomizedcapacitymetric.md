This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScalingPolicy PredictiveScalingCustomizedCapacityMetric

Contains capacity metric information for the
`CustomizedCapacityMetricSpecification` property of the [AWS::AutoScaling::ScalingPolicy PredictiveScalingMetricSpecification](../userguide/aws-properties-autoscaling-scalingpolicy-predictivescalingmetricspecification.md) property
type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricDataQueries" : [ MetricDataQuery, ... ]
}

```

### YAML

```yaml

  MetricDataQueries:
    - MetricDataQuery

```

## Properties

`MetricDataQueries`

One or more metric data queries to provide the data points for a capacity metric. Use
multiple metric data queries only if you are performing a math expression on returned
data.

_Required_: Yes

_Type_: Array of [MetricDataQuery](aws-properties-autoscaling-scalingpolicy-metricdataquery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingConfiguration

PredictiveScalingCustomizedLoadMetric

All content copied from https://docs.aws.amazon.com/.
