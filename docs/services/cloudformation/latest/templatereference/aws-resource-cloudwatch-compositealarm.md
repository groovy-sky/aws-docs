This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::CompositeAlarm

The `AWS::CloudWatch::CompositeAlarm` type creates or updates a composite alarm. When you create
a composite alarm, you specify a rule expression for the alarm that takes into
account the alarm states of other alarms that you have created. The composite alarm goes into ALARM state
only if all conditions of the rule are met.

The alarms specified in a composite alarm's rule expression can include metric alarms and other composite alarms.

Using composite alarms can reduce alarm noise. You can create multiple metric alarms, and also create a composite alarm and set
up alerts only for the composite alarm. For example, you could create a composite alarm that goes into ALARM state
only when more than one of the underlying metric alarms are in ALARM state.

When this operation creates an alarm, the alarm state is immediately set to INSUFFICIENT\_DATA. The alarm is then evaluated and
its state is set appropriately. Any actions associated with the new state are then executed. For a composite alarm, this initial
time after creation is the only time that the alarm can be in INSUFFICIENT\_DATA state.

When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudWatch::CompositeAlarm",
  "Properties" : {
      "ActionsEnabled" : Boolean,
      "ActionsSuppressor" : String,
      "ActionsSuppressorExtensionPeriod" : Integer,
      "ActionsSuppressorWaitPeriod" : Integer,
      "AlarmActions" : [ String, ... ],
      "AlarmDescription" : String,
      "AlarmName" : String,
      "AlarmRule" : String,
      "InsufficientDataActions" : [ String, ... ],
      "OKActions" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudWatch::CompositeAlarm
Properties:
  ActionsEnabled: Boolean
  ActionsSuppressor: String
  ActionsSuppressorExtensionPeriod: Integer
  ActionsSuppressorWaitPeriod: Integer
  AlarmActions:
    - String
  AlarmDescription: String
  AlarmName: String
  AlarmRule: String
  InsufficientDataActions:
    - String
  OKActions:
    - String
  Tags:
    - Tag

```

## Properties

`ActionsEnabled`

Indicates whether actions should be executed during any changes to the alarm state of the composite alarm. The default is TRUE.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ActionsSuppressor`

Actions will be suppressed
if the suppressor alarm is
in the `ALARM` state.
`ActionsSuppressor` can be an AlarmName or an Amazon Resource Name (ARN)
from an existing alarm.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ActionsSuppressorExtensionPeriod`

The maximum time
in seconds
that the composite alarm waits
after suppressor alarm goes out
of the `ALARM` state.
After this time,
the composite alarm performs its actions.

###### Important

`ExtensionPeriod`
is required only
when `ActionsSuppressor` is specified.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ActionsSuppressorWaitPeriod`

The maximum time
in seconds
that the composite alarm waits
for the suppressor alarm
to go
into the `ALARM` state.
After this time,
the composite alarm performs its actions.

###### Important

`WaitPeriod`
is required only
when `ActionsSuppressor` is specified.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmActions`

The actions to execute when this alarm transitions to the ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
For more information about creating alarms and the actions
that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the
_Amazon CloudWatch API Reference_.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmDescription`

The description for the composite alarm.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmName`

The name for the composite alarm. This name must be unique within your AWS account.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AlarmRule`

An expression that specifies which other alarms are to be evaluated to determine this composite alarm's state. For each
alarm that you reference, you designate a function that specifies whether that alarm needs to be in ALARM state, OK state,
or INSUFFICIENT\_DATA state. You can use operators (AND, OR and NOT) to combine multiple functions in a
single expression. You can use parenthesis to logically group the functions in your expression.

You can use either alarm names or ARNs to reference the other alarms that are to be evaluated.

Functions can include the following:

- ALARM("alarm-name or alarm-ARN") is TRUE if the named alarm is in ALARM state.

- OK("alarm-name or alarm-ARN") is TRUE if the named alarm is in OK state.

- INSUFFICIENT\_DATA("alarm-name or alarm-ARN") is TRUE if the named alarm is in INSUFFICIENT\_DATA state.

- TRUE always evaluates to TRUE.

- FALSE always evaluates to FALSE.

TRUE and FALSE are useful for testing a complex AlarmRule structure, and for testing your alarm actions.

For more information about `AlarmRule` syntax, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the
_Amazon CloudWatch API Reference_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsufficientDataActions`

The actions to execute when this alarm transitions to the INSUFFICIENT\_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
For more information about creating alarms and the actions
that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the
_Amazon CloudWatch API Reference_.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OKActions`

The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
For more information about creating alarms and the actions
that you can specify, see [PutCompositeAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutCompositeAlarm.html) in the
_Amazon CloudWatch API Reference_.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1024 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate
tags with the alarm when you create the alarm, you must have the `cloudwatch:TagResource` permission.

Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a
user permission to access or change only resources with certain tag values.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudwatch-compositealarm-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the alarm name, such as `MyCompositeAlarm`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the composite alarm, such as `arn:aws:cloudwatch:us-west-2:123456789012:alarm/CompositeAlarmName`.

## Examples

### Composite alarm based on two metric alarms and another composite alarm

This example creates composite alarms named "HighResourceUsage" and "DeploymentInProgress", and also creates metrics alarms named
"HighCPUUsage" and "HighMemoryUsage". "DeploymentInProgress" is an alarm that must be manually set to TRUE or FALSE. The "HighResourceUsage"
alarm goes into ALARM state only if both "HighCPUUsage" and "HighMemoryUsage" are in ALARM state, and if "DeploymentInProgress" is FALSE. Only
"HighResourceUsage" has the alarm action of notifying SNS. This reduces alarm noise, so that you are alerted only if both CPU usage and memory usage
are high, and a deployment is not currently in progress.

#### JSON

```json

"Resources": {
    "HighResourceUsage": {
        "Type": "AWS::CloudWatch::CompositeAlarm",
        "Properties": {
            "AlarmName": "HighResourceUsage",
            "AlarmRule": "(ALARM(HighCPUUsage) OR ALARM(HighMemoryUsage)) AND NOT ALARM(DeploymentInProgress)",
            "AlarmActions": [ {"arn:aws:sns:us-west-2:123456789012:my-sns-topic"} },
            "AlarmDescription": "Indicates that the system resource usage is high while no known
            deployment is in progress"
        },
        "DependsOn": [
            "DeploymentInProgress",
            "HighCPUUsage",
            "HighMemoryUsage"
        ]
    },
    "DeploymentInProgress": {
        "Type": "AWS::CloudWatch::CompositeAlarm",
        "Properties": {
            "AlarmName": "DeploymentInProgress",
            "AlarmRule": "FALSE",
            "AlarmDescription": "Manually updated to TRUE/FALSE to disable other alarms"
        }
    },
    "HighCPUUsage": {
        "Type": "AWS::CloudWatch::Alarm",
        "Properties": {
            "AlarmDescription": "CPUusageishigh",
            "AlarmName": "HighCPUUsage",
            "ComparisonOperator": "GreaterThanThreshold",
            "EvaluationPeriods": 1,
            "MetricName": "CPUUsage",
            "Namespace": "CustomNamespace",
            "Period": 60,
            "Statistic": "Average",
            "Threshold": 70,
            "TreatMissingData": "notBreaching"
        }
    },
    "HighMemoryUsage": {
        "Type": "AWS::CloudWatch::Alarm",
        "Properties": {
            "AlarmDescription": "Memoryusageishigh",
            "AlarmName": "HighMemoryUsage",
            "ComparisonOperator": "GreaterThanThreshold",
            "EvaluationPeriods": 1,
            "MetricName": "MemoryUsage",
            "Namespace": "CustomNamespace",
            "Period": 60,
            "Statistic": "Average",
            "Threshold": 65,
            "TreatMissingData": "breaching"
        }
    }
}
```

#### YAML

```yaml

Resources:
  HighResourceUsage:
    Type: AWS::CloudWatch::CompositeAlarm
    Properties:
      AlarmName: HighResourceUsage
      AlarmRule: !Sub "(ALARM(${HighCPUUsage}) OR ALARM(${HighMemoryUsage})) AND NOT ALARM(${DeploymentInProgress})"
      AlarmActions: [
         arn:aws:sns:us-west-2:123456789012:my-sns-topic
      ]
      AlarmDescription: Indicates that the system resource usage is high while no known deployment is in progress
    DependsOn:
    - DeploymentInProgress
    - HighCPUUsage
    - HighMemoryUsage
  DeploymentInProgress:
    Type: AWS::CloudWatch::CompositeAlarm
    Properties:
      AlarmName: DeploymentInProgress
      AlarmRule:
      AlarmDescription: Manually updated to TRUE/FALSE to disable other alarms
  HighCPUUsage:
    Type: AWS::CloudWatch::Alarm
    Properties:
      AlarmDescription: CPU usage is high
      AlarmName: HighCPUUsage
      ComparisonOperator: GreaterThanThreshold
      EvaluationPeriods: 1
      MetricName: CPUUsage
      Namespace: CustomNamespace
      Period: 60
      Statistic: Average
      Threshold: 70
      TreatMissingData: notBreaching
  HighMemoryUsage:
    Type: AWS::CloudWatch::Alarm
    Properties:
      AlarmDescription: Memory usage is high
      AlarmName: HighMemoryUsage
      ComparisonOperator: GreaterThanThreshold
      EvaluationPeriods: 1
      MetricName: MemoryUsage
      Namespace: CustomNamespace
      Period: 60
      Statistic: Average
      Threshold: 65
      TreatMissingData: breaching
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SingleMetricAnomalyDetector

Tag
