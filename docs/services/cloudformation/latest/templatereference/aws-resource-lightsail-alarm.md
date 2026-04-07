This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Alarm

The `AWS::Lightsail::Alarm` resource specifies an alarm that can be used to
monitor a single metric for one of your Lightsail resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Alarm",
  "Properties" : {
      "AlarmName" : String,
      "ComparisonOperator" : String,
      "ContactProtocols" : [ String, ... ],
      "DatapointsToAlarm" : Integer,
      "EvaluationPeriods" : Integer,
      "MetricName" : String,
      "MonitoredResourceName" : String,
      "NotificationEnabled" : Boolean,
      "NotificationTriggers" : [ String, ... ],
      "Threshold" : Number,
      "TreatMissingData" : String
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Alarm
Properties:
  AlarmName: String
  ComparisonOperator: String
  ContactProtocols:
    - String
  DatapointsToAlarm: Integer
  EvaluationPeriods: Integer
  MetricName: String
  MonitoredResourceName: String
  NotificationEnabled: Boolean
  NotificationTriggers:
    - String
  Threshold: Number
  TreatMissingData: String

```

## Properties

`AlarmName`

The name of the alarm.

_Required_: Yes

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComparisonOperator`

The arithmetic operation to use when comparing the specified statistic and
threshold.

_Required_: Yes

_Type_: String

_Allowed values_: `GreaterThanOrEqualToThreshold | GreaterThanThreshold | LessThanThreshold | LessThanOrEqualToThreshold`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContactProtocols`

The contact protocols for the alarm, such as `Email`, `SMS` (text
messaging), or both.

_Allowed Values_: `Email` \| `SMS`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatapointsToAlarm`

The number of data points within the evaluation periods that must be breaching to cause
the alarm to go to the `ALARM` state.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationPeriods`

The number of periods over which data is compared to the specified threshold.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric associated with the alarm.

_Required_: Yes

_Type_: String

_Allowed values_: `CPUUtilization | NetworkIn | NetworkOut | StatusCheckFailed | StatusCheckFailed_Instance | StatusCheckFailed_System | ClientTLSNegotiationErrorCount | HealthyHostCount | UnhealthyHostCount | HTTPCode_LB_4XX_Count | HTTPCode_LB_5XX_Count | HTTPCode_Instance_2XX_Count | HTTPCode_Instance_3XX_Count | HTTPCode_Instance_4XX_Count | HTTPCode_Instance_5XX_Count | InstanceResponseTime | RejectedConnectionCount | RequestCount | DatabaseConnections | DiskQueueDepth | FreeStorageSpace | NetworkReceiveThroughput | NetworkTransmitThroughput | BurstCapacityTime | BurstCapacityPercentage`

_Update requires_: Updates are not supported.

`MonitoredResourceName`

The name of the Lightsail resource that the alarm monitors.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`NotificationEnabled`

A Boolean value indicating whether the alarm is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationTriggers`

The alarm states that trigger a notification.

###### Note

To specify the `OK` and `INSUFFICIENT_DATA` values, you must also
specify `ContactProtocols` values. Otherwise, the `OK`
and `INSUFFICIENT_DATA` values will not take effect and the stack will
drift.

_Allowed Values_: `OK` \| `ALARM` \|
`INSUFFICIENT_DATA`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The value against which the specified statistic is compared.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatMissingData`

Specifies how the alarm handles missing data points.

An alarm can treat missing data in the following ways:

- `breaching` \- Assumes the missing data is not within the threshold. Missing
data counts towards the number of times that the metric is not within the
threshold.

- `notBreaching` \- Assumes the missing data is within the threshold. Missing
data does not count towards the number of times that the metric is not within the
threshold.

- `ignore` \- Ignores the missing data. Maintains the current alarm
state.

- `missing` \- Missing data is treated as missing.

_Required_: No

_Type_: String

_Allowed values_: `breaching | notBreaching | ignore | missing`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AlarmArn`

The Amazon Resource Name (ARN) of the alarm.

`State`

The current state of the alarm.

An alarm has the following possible states:

- `ALARM` \- The metric is outside of the defined threshold.

- `INSUFFICIENT_DATA` \- The alarm has recently started, the metric is not
available, or not enough data is available for the metric to determine the alarm
state.

- `OK` \- The metric is within the defined threshold.

## Remarks

_Notification triggers_

To specify the `OK` and `INSUFFICIENT_DATA` values of the
`NotificationTriggers` parameter, you must also specify
`ContactProtocols` values. Otherwise, the `OK` and
`INSUFFICIENT_DATA` values of the `NotificationTriggers`
parameter will not take effect and the stack will drift.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Lightsail

AWS::Lightsail::Bucket
