This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Alarm

The `AWS::CloudWatch::Alarm` type specifies an alarm and associates it with the specified metric or metric math expression.

When this operation creates an alarm, the alarm state is immediately set to
`INSUFFICIENT_DATA`. The alarm is then evaluated and its state is set
appropriately. Any actions associated with the new state are then executed.

When you update an existing alarm, its state is left unchanged, but the
update completely overwrites the previous configuration of the alarm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudWatch::Alarm",
  "Properties" : {
      "ActionsEnabled" : Boolean,
      "AlarmActions" : [ String, ... ],
      "AlarmDescription" : String,
      "AlarmName" : String,
      "ComparisonOperator" : String,
      "DatapointsToAlarm" : Integer,
      "Dimensions" : [ Dimension, ... ],
      "EvaluateLowSampleCountPercentile" : String,
      "EvaluationPeriods" : Integer,
      "ExtendedStatistic" : String,
      "InsufficientDataActions" : [ String, ... ],
      "MetricName" : String,
      "Metrics" : [ MetricDataQuery, ... ],
      "Namespace" : String,
      "OKActions" : [ String, ... ],
      "Period" : Integer,
      "Statistic" : String,
      "Tags" : [ Tag, ... ],
      "Threshold" : Number,
      "ThresholdMetricId" : String,
      "TreatMissingData" : String,
      "Unit" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudWatch::Alarm
Properties:
  ActionsEnabled: Boolean
  AlarmActions:
    - String
  AlarmDescription: String
  AlarmName: String
  ComparisonOperator: String
  DatapointsToAlarm: Integer
  Dimensions:
    - Dimension
  EvaluateLowSampleCountPercentile: String
  EvaluationPeriods: Integer
  ExtendedStatistic: String
  InsufficientDataActions:
    - String
  MetricName: String
  Metrics:
    - MetricDataQuery
  Namespace: String
  OKActions:
    - String
  Period: Integer
  Statistic: String
  Tags:
    - Tag
  Threshold: Number
  ThresholdMetricId: String
  TreatMissingData: String
  Unit: String

```

## Properties

`ActionsEnabled`

Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmActions`

The list of actions to execute when this alarm transitions into an ALARM state from any other state.

Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions
that you can specify, see [PutMetricAlarm](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricalarm.md) in the
_Amazon CloudWatch API Reference_.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmDescription`

The description of the alarm.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmName`

The name of the alarm. If you don't specify a name, CloudFormation generates a unique physical ID and uses that ID for the alarm name.

###### Important

If you specify a name, you cannot perform updates that require replacement of this resource. You can
perform updates that require no or some interruption. If you must replace the resource, specify a new name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComparisonOperator`

The arithmetic operation to use when comparing the specified
statistic and threshold. The specified statistic value is used as the first operand.

_Required_: No

_Type_: String

_Allowed values_: `GreaterThanOrEqualToThreshold | GreaterThanThreshold | LessThanThreshold | LessThanOrEqualToThreshold | LessThanLowerOrGreaterThanUpperThreshold | LessThanLowerThreshold | GreaterThanUpperThreshold`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatapointsToAlarm`

The number of datapoints that must be breaching to trigger the alarm.
This is used only if you are setting an "M out of N" alarm. In that case, this value is the M,
and the value that you set for `EvaluationPeriods` is the N value.
For more information, see [Evaluating\
an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the _Amazon CloudWatch User Guide_.

If you omit this parameter, CloudWatch uses the same value here that you set
for `EvaluationPeriods`, and the alarm goes to alarm state if that many consecutive
periods are breaching.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimensions`

The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't
specify `Dimensions`. Instead, you use `Metrics`.

_Required_: No

_Type_: Array of [Dimension](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-alarm-dimension.html)

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluateLowSampleCountPercentile`

Used only for alarms based on percentiles. If `ignore`, the alarm state
does not change during periods with too few data points to be statistically significant.
If `evaluate` or this parameter is not used, the alarm is always evaluated
and possibly changes state no matter how many data points are available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationPeriods`

The number of periods over which data is compared to the specified threshold.
If you are setting an alarm that requires that a number of consecutive data points be
breaching to trigger the alarm, this value specifies that number. If you
are setting an "M out of N" alarm, this value is the N, and `DatapointsToAlarm`
is the M.

For more information, see [Evaluating\
an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the _Amazon CloudWatch User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtendedStatistic`

The percentile statistic for the metric associated with the alarm. Specify a value between
p0.0 and p100.

For an alarm based on a metric, you must specify either `Statistic`
or `ExtendedStatistic` but not both.

For an alarm based on a math expression, you can't
specify `ExtendedStatistic`. Instead, you use `Metrics`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsufficientDataActions`

The actions to execute when this alarm transitions to the
`INSUFFICIENT_DATA` state from any other state. Each action is specified
as an Amazon Resource Name (ARN).

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric associated with the alarm. This is required for an alarm based on a
metric. For an alarm based on a math expression, you use `Metrics` instead and you can't
specify `MetricName`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metrics`

An array that enables you to create an alarm based on the result of a metric math expression. Each
item in the array either retrieves a metric or performs a math expression.

If you specify the `Metrics` parameter, you cannot specify `MetricName`, `Dimensions`,
`Period`, `Namespace`, `Statistic`, `ExtendedStatistic`, or `Unit`.

_Required_: No

_Type_: Array of [MetricDataQuery](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-alarm-metricdataquery.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The namespace of the metric associated with the alarm. This is required for an alarm based on a
metric. For an alarm based on a math expression, you can't
specify `Namespace` and you use `Metrics` instead.

For a list of namespaces for metrics from AWS services, see
[AWS Services That Publish CloudWatchMetrics.](../../../amazoncloudwatch/latest/monitoring/aws-services-cloudwatch-metrics.md)

_Required_: No

_Type_: String

_Pattern_: `[^:].*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OKActions`

The actions to execute when this alarm transitions to the `OK` state
from any other state. Each action is specified as an Amazon Resource Name
(ARN).

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The period, in seconds, over which the statistic is applied. This is required for an alarm based on a
metric. Valid values are 10, 20, 30, 60, and any multiple of 60.

For an alarm based on a math expression, you can't
specify `Period`, and instead you use the `Metrics` parameter.

_Minimum:_ 10

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The statistic for the metric associated with the alarm, other than percentile.
For percentile statistics, use `ExtendedStatistic`.

For an alarm based on a metric, you must specify either `Statistic`
or `ExtendedStatistic` but not both.

For an alarm based on a math expression, you can't
specify `Statistic`. Instead, you use `Metrics`.

_Required_: No

_Type_: String

_Allowed values_: `SampleCount | Average | Sum | Minimum | Maximum`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate
tags with the alarm when you create the alarm, you must have the `cloudwatch:TagResource` permission.

Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a
user permission to access or change only resources with certain tag values.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-alarm-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The value to compare with the specified statistic.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThresholdMetricId`

In an alarm based on an anomaly detection model, this is the ID of the
`ANOMALY_DETECTION_BAND` function used as the threshold for the
alarm.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatMissingData`

Sets how this alarm is to handle missing data points. Valid values are `breaching`, `notBreaching`, `ignore`,
and `missing`. For more information, see [Configuring How CloudWatchAlarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the
_Amazon CloudWatchUser Guide_.

If you omit this parameter, the default behavior of `missing` is used.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of the metric associated with the alarm. Specify this only if you are creating an alarm
based on a single metric. Do not specify this if you are specifying a `Metrics` array.

You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes,
Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count,
Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second,
Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.

_Required_: No

_Type_: String

_Allowed values_: `Seconds | Microseconds | Milliseconds | Bytes | Kilobytes | Megabytes | Gigabytes | Terabytes | Bits | Kilobits | Megabits | Gigabits | Terabits | Percent | Count | Bytes/Second | Kilobytes/Second | Megabytes/Second | Gigabytes/Second | Terabytes/Second | Bits/Second | Kilobits/Second | Megabits/Second | Gigabits/Second | Terabits/Second | Count/Second | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the alarm name, such as `TestAlarm`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the CloudWatchalarm, such as
`arn:aws:cloudwatch:us-west-2:123456789012:alarm:myCloudWatchAlarm-CPUAlarm-UXMMZK36R55Z`.

## Examples

### Alarm Based on an Anomaly Detector

This example creates an alarm that is based on an anomaly detector.

#### JSON

```json

"Resources": {
    "LambdaInvocationsAnomalyDetector": {
       "Type": "AWS::CloudWatch::AnomalyDetector",
       "Properties": {
          "MetricName": "Invocations",
          "Namespace": "AWS/Lambda",
          "Stat": "Sum"
       }
    },
    "LambdaInvocationsAlarm": {
       "Type": "AWS::CloudWatch::Alarm",
       "Properties": {
          "AlarmDescription": "Lambda invocations",
          "AlarmName": "LambdaInvocationsAlarm",
          "ComparisonOperator": "LessThanLowerOrGreaterThanUpperThreshold",
          "EvaluationPeriods": 1,
          "Metrics": [
             {
                "Expression": "ANOMALY_DETECTION_BAND(m1, 2)",
                "Id": "ad1"
             },
             {
                "Id": "m1",
                "MetricStat": {
                   "Metric": {
                      "MetricName": "Invocations",
                      "Namespace": "AWS/Lambda"
                   },
                   "Period": 86400,
                   "Stat": "Sum"
                }
             }
          ],
          "ThresholdMetricId": "ad1",
          "TreatMissingData": "breaching"
       }
    }
 }
```

#### YAML

```yaml

Resources:
  LambdaInvocationsAnomalyDetector:
    Type: AWS::CloudWatch::AnomalyDetector
    Properties:
      MetricName: Invocations
      Namespace: AWS/Lambda
      Stat: Sum

  LambdaInvocationsAlarm:
    Type: AWS::CloudWatch::Alarm
    Properties:
      AlarmDescription: Lambda invocations
      AlarmName: LambdaInvocationsAlarm
      ComparisonOperator: LessThanLowerOrGreaterThanUpperThreshold
      EvaluationPeriods: 1
      Metrics:
      - Expression: ANOMALY_DETECTION_BAND(m1, 2)
        Id: ad1
      - Id: m1
        MetricStat:
          Metric:
            MetricName: Invocations
            Namespace: AWS/Lambda
          Period: !!int 86400
          Stat: Sum
      ThresholdMetricId: ad1
      TreatMissingData: breaching

```

## See also

- [Amazon CloudWatch Template Snippets](../userguide/quickref-cloudwatch.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatch

Dimension
