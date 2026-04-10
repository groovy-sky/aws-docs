This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile BehaviorCriteria

The criteria by which the behavior is determined to be normal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComparisonOperator" : String,
  "ConsecutiveDatapointsToAlarm" : Integer,
  "ConsecutiveDatapointsToClear" : Integer,
  "DurationSeconds" : Integer,
  "MlDetectionConfig" : MachineLearningDetectionConfig,
  "StatisticalThreshold" : StatisticalThreshold,
  "Value" : MetricValue
}

```

### YAML

```yaml

  ComparisonOperator: String
  ConsecutiveDatapointsToAlarm: Integer
  ConsecutiveDatapointsToClear: Integer
  DurationSeconds: Integer
  MlDetectionConfig:
    MachineLearningDetectionConfig
  StatisticalThreshold:
    StatisticalThreshold
  Value:
    MetricValue

```

## Properties

`ComparisonOperator`

The operator that relates the thing measured ( `metric`) to the criteria
(containing a `value` or `statisticalThreshold`). Valid operators include:

- `string-list`: `in-set` and `not-in-set`

- `number-list`: `in-set` and `not-in-set`

- `ip-address-list`: `in-cidr-set` and `not-in-cidr-set`

- `number`: `less-than`, `less-than-equals`, `greater-than`, and `greater-than-equals`

_Required_: No

_Type_: String

_Allowed values_: `less-than | less-than-equals | greater-than | greater-than-equals | in-cidr-set | not-in-cidr-set | in-port-set | not-in-port-set | in-set | not-in-set`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConsecutiveDatapointsToAlarm`

If a device is in violation of the behavior for the specified number of consecutive
datapoints, an alarm occurs. If not specified, the default is 1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConsecutiveDatapointsToClear`

If an alarm has occurred and the offending device is no longer in violation of the behavior
for the specified number of consecutive datapoints, the alarm is cleared. If not specified,
the default is 1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DurationSeconds`

Use this to specify the time duration over which the behavior is evaluated, for those criteria that
have a time dimension (for example, `NUM_MESSAGES_SENT`). For a
`statisticalThreshhold` metric comparison, measurements from all devices are
accumulated over this time duration before being used to calculate percentiles, and later,
measurements from an individual device are also accumulated over this time duration before
being given a percentile rank. Cannot be used with list-based metric datatypes.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MlDetectionConfig`

The confidence level of the detection model.

_Required_: No

_Type_: [MachineLearningDetectionConfig](aws-properties-iot-securityprofile-machinelearningdetectionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatisticalThreshold`

A statistical ranking (percentile)that
indicates a threshold value by which a behavior is determined to be in compliance or in
violation of the behavior.

_Required_: No

_Type_: [StatisticalThreshold](aws-properties-iot-securityprofile-statisticalthreshold.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to be compared with the `metric`.

_Required_: No

_Type_: [MetricValue](aws-properties-iot-securityprofile-metricvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Behavior

MachineLearningDetectionConfig

All content copied from https://docs.aws.amazon.com/.
