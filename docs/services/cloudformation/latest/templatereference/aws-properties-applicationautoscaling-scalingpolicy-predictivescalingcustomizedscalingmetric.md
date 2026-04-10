This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy PredictiveScalingCustomizedScalingMetric

The `PredictiveScalingCustomizedScalingMetric` property type specifies Property description not available. for an [AWS::ApplicationAutoScaling::ScalingPolicy](aws-resource-applicationautoscaling-scalingpolicy.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricDataQueries" : [ PredictiveScalingMetricDataQuery, ... ]
}

```

### YAML

```yaml

  MetricDataQueries:
    - PredictiveScalingMetricDataQuery

```

## Properties

`MetricDataQueries`

One or more metric data queries to provide data points for a metric specification.

_Required_: Yes

_Type_: Array of [PredictiveScalingMetricDataQuery](aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricdataquery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingCustomizedLoadMetric

PredictiveScalingMetric

All content copied from https://docs.aws.amazon.com/.
