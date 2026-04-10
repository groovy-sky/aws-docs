This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy PredictiveScalingPredefinedLoadMetric

Describes a load metric for a predictive scaling policy.

When returned in the output of `DescribePolicies`, it indicates that a
predictive scaling policy uses individually specified load and scaling metrics instead
of a metric pair.

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

The metric type.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceLabel`

A label that uniquely identifies a target group.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveScalingPolicyConfiguration

PredictiveScalingPredefinedMetricPair

All content copied from https://docs.aws.amazon.com/.
