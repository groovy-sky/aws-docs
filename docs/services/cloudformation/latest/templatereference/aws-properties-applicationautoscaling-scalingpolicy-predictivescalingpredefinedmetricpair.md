This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy PredictiveScalingPredefinedMetricPair

Represents a metric pair for a predictive scaling policy.

The following predefined metrics are available for predictive scaling:

- `ECSServiceAverageCPUUtilization`

- `ECSServiceAverageMemoryUtilization`

- `ECSServiceCPUUtilization`

- `ECSServiceMemoryUtilization`

- `ECSServiceTotalCPUUtilization`

- `ECSServiceTotalMemoryUtilization`

- `ALBRequestCount`

- `ALBRequestCountPerTarget`

- `TotalALBRequestCount`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PredefinedMetricType" : String,
  "ResourceLabel" : String
}

```

### YAML

```yaml

  PredefinedMetricType: String
  ResourceLabel: String

```

## Properties

`PredefinedMetricType`

Indicates which metrics to use. There are two different types of metrics for each
metric type: one is a load metric and one is a scaling metric.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLabel`

A label that uniquely identifies a specific target group from which to determine
the total and average request count.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingPredefinedLoadMetric

PredictiveScalingPredefinedScalingMetric

All content copied from https://docs.aws.amazon.com/.
