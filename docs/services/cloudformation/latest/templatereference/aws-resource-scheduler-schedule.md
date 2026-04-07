This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule

A _schedule_ is the main resource you create, configure, and manage using Amazon EventBridge Scheduler.

Every schedule has a _schedule expression_ that determines when, and with what frequency, the schedule runs. EventBridge Scheduler
supports three types of schedules: rate, cron, and one-time schedules. For more information about different schedule types,
see [Schedule types](https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html) in the _EventBridge Scheduler User Guide_.

When you create a schedule, you configure a target for the schedule to invoke. A target is an API operation that EventBridge Scheduler calls on your behalf
every time your schedule runs. EventBridge Scheduler supports two types of targets: _templated_ targets invoke common API operations across
a core groups of services, and customizeable _universal_ targets that you can use to call more than 6,000 operations
across over 270 services. For more information about configuring targets, see
[Managing targets](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets.html) in the _EventBridge Scheduler User Guide_.

For more information about managing schedules, changing the schedule state, setting up flexible time windows, and configuring a dead-letter queue for a schedule, see
[Managing a schedule](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule.html) in the _EventBridge Scheduler User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Scheduler::Schedule",
  "Properties" : {
      "Description" : String,
      "EndDate" : String,
      "FlexibleTimeWindow" : FlexibleTimeWindow,
      "GroupName" : String,
      "KmsKeyArn" : String,
      "Name" : String,
      "ScheduleExpression" : String,
      "ScheduleExpressionTimezone" : String,
      "StartDate" : String,
      "State" : String,
      "Target" : Target
    }
}

```

### YAML

```yaml

Type: AWS::Scheduler::Schedule
Properties:
  Description: String
  EndDate: String
  FlexibleTimeWindow:
    FlexibleTimeWindow
  GroupName: String
  KmsKeyArn: String
  Name: String
  ScheduleExpression: String
  ScheduleExpressionTimezone: String
  StartDate: String
  State: String
  Target:
    Target

```

## Properties

`Description`

The description you specify for the schedule.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndDate`

The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the `EndDate` you specify.
EventBridge Scheduler ignores `EndDate` for one-time schedules.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlexibleTimeWindow`

Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.

_Required_: Yes

_Type_: [FlexibleTimeWindow](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-scheduler-schedule-flexibletimewindow.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupName`

The name of the schedule group associated with this schedule.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z-_.]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The Amazon Resource Name (ARN) for the customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt your data.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:kms:[a-z0-9\-]+:\d{12}:(key|alias)\/[0-9a-zA-Z-_]*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the schedule.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z-_.]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScheduleExpression`

The expression that defines when the schedule runs. The following formats are supported.

- `at` expression - `at(yyyy-mm-ddThh:mm:ss)`

- `rate` expression - `rate(value unit)`

- `cron` expression - `cron(fields)`

You can use `at` expressions to create one-time schedules that invoke a target once, at the time and in the time zone, that you specify.
You can use `rate` and `cron` expressions to create recurring schedules. Rate-based schedules are useful when you want to invoke a target
at regular intervals, such as every 15 minutes or every five days. Cron-based schedules are useful when you want to invoke a target periodically at a specific time,
such as at 8:00 am (UTC+0) every 1st day of the month.

A `cron` expression consists of six fields separated by white spaces: `(minutes hours day_of_month month day_of_week year)`.

A `rate` expression consists of a _value_ as a positive integer, and a _unit_ with the following options:
`minute` \| `minutes` \| `hour` \| `hours` \| `day` \| `days`

For more information and examples, see [Schedule types on EventBridge Scheduler](https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html) in the _EventBridge Scheduler User Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpressionTimezone`

The timezone in which the scheduling expression is evaluated.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartDate`

The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the `StartDate` you specify.
EventBridge Scheduler ignores `StartDate` for one-time schedules.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

Specifies whether the schedule is enabled or disabled.

_Allowed Values_: `ENABLED` \| `DISABLED`

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The schedule's target details.

_Required_: Yes

_Type_: [Target](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-scheduler-schedule-target.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `Name` attribute of theschedule.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes that `Fn::GetAtt` returns.
For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Arn`

The Amazon Resource Name (ARN) for the Amazon EventBridge Scheduler schedule.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon EventBridge Scheduler

AwsVpcConfiguration
