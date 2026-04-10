This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationAutoScaling::ScalableTarget ScheduledAction

`ScheduledAction` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](../userguide/aws-resource-applicationautoscaling-scalabletarget.md) resource that specifies a scheduled
action for a scalable target.

For more information, see [Scheduled scaling](../../../autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.md) in the _Application Auto Scaling User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTime" : String,
  "ScalableTargetAction" : ScalableTargetAction,
  "Schedule" : String,
  "ScheduledActionName" : String,
  "StartTime" : String,
  "Timezone" : String
}

```

### YAML

```yaml

  EndTime: String
  ScalableTargetAction:
    ScalableTargetAction
  Schedule: String
  ScheduledActionName: String
  StartTime: String
  Timezone: String

```

## Properties

`EndTime`

The date and time that the action is scheduled to end, in UTC.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalableTargetAction`

The new minimum and maximum capacity. You can set both values or just one. At the
scheduled time, if the current capacity is below the minimum capacity, Application Auto Scaling scales out
to the minimum capacity. If the current capacity is above the maximum capacity, Application Auto Scaling
scales in to the maximum capacity.

_Required_: No

_Type_: [ScalableTargetAction](aws-properties-applicationautoscaling-scalabletarget-scalabletargetaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule for this action. The following formats are supported:

- At expressions -
" `at(yyyy-mm-ddThh:mm:ss)`"

- Rate expressions - " `rate(valueunit)`"

- Cron expressions - " `cron(fields)`"

At expressions are useful for one-time schedules. Cron expressions are useful for
scheduled actions that run periodically at a specified date and time, and rate expressions are
useful for scheduled actions that run at a regular interval.

At and cron expressions use Universal Coordinated Time (UTC) by default.

The cron format consists of six fields separated by white spaces: \[Minutes\] \[Hours\]
\[Day\_of\_Month\] \[Month\] \[Day\_of\_Week\] \[Year\].

For rate expressions, _value_ is a positive integer and
_unit_ is `minute` \| `minutes` \| `hour`
\| `hours` \| `day` \| `days`.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduledActionName`

The name of the scheduled action. This name must be unique among all other scheduled
actions on the specified scalable target.

_Required_: Yes

_Type_: String

_Pattern_: `(?!((^[ ]+.*)|(.*([\u0000-\u001f]|[\u007f-\u009f]|[:/|])+.*)|(.*[ ]+$))).+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The date and time that the action is scheduled to begin, in UTC.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timezone`

The time zone used when referring to the date and time of a scheduled action, when the
scheduled action uses an at or cron expression.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Configure Application Auto Scaling resources](../userguide/quickref-application-auto-scaling.md).

- [Getting started](../../../autoscaling/application/userguide/getting-started.md)
in the _Application Auto Scaling User Guide_

- [Schedule recurring scaling actions using cron expressions](../../../autoscaling/application/userguide/scheduled-scaling-using-cron-expressions.md) in the
_Application Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScalableTargetAction

SuspendedState

All content copied from https://docs.aws.amazon.com/.
