This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalingPolicy PredictiveScalingPolicyConfiguration

Represents a predictive scaling policy configuration. Predictive scaling is supported on Amazon ECS services.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCapacityBreachBehavior" : String,
  "MaxCapacityBuffer" : Integer,
  "MetricSpecifications" : [ PredictiveScalingMetricSpecification, ... ],
  "Mode" : String,
  "SchedulingBufferTime" : Integer
}

```

### YAML

```yaml

  MaxCapacityBreachBehavior: String
  MaxCapacityBuffer: Integer
  MetricSpecifications:
    - PredictiveScalingMetricSpecification
  Mode: String
  SchedulingBufferTime: Integer

```

## Properties

`MaxCapacityBreachBehavior`

Defines the behavior that should be applied if the forecast capacity approaches or
exceeds the maximum capacity. Defaults to
`HonorMaxCapacity` if not specified.

_Required_: No

_Type_: String

_Allowed values_: `HonorMaxCapacity | IncreaseMaxCapacity`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacityBuffer`

The size of the capacity buffer to use when the forecast capacity is close to or
exceeds the maximum capacity. The value is specified as a percentage relative to the
forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer,
such that if the forecast capacity is 50, and the maximum capacity is 40, then the
effective maximum capacity is 55.

Required if the `MaxCapacityBreachBehavior` property is set to
`IncreaseMaxCapacity`, and cannot be used otherwise.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricSpecifications`

This structure includes the metrics and target utilization to use for predictive scaling.

This is an array, but we currently only support a single metric specification. That
is, you can specify a target value and a single metric pair, or a target value and one
scaling metric and one load metric.

_Required_: Yes

_Type_: Array of [PredictiveScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-applicationautoscaling-scalingpolicy-predictivescalingmetricspecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The predictive scaling mode. Defaults to `ForecastOnly` if not specified.

_Required_: No

_Type_: String

_Allowed values_: `ForecastOnly | ForecastAndScale`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchedulingBufferTime`

The amount of time, in seconds, that the start time can be advanced.

The value must be less than the forecast interval duration of 3600 seconds (60
minutes). Defaults to 300 seconds if not specified.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PredictiveScalingMetricStat

PredictiveScalingPredefinedLoadMetric
