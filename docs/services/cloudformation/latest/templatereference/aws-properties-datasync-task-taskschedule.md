This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::Task TaskSchedule

Configures your AWS DataSync task to run on a [schedule](../../../datasync/latest/userguide/task-scheduling.md)
(at a minimum interval of 1 hour).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScheduleExpression" : String,
  "Status" : String
}

```

### YAML

```yaml

  ScheduleExpression: String
  Status: String

```

## Properties

`ScheduleExpression`

Specifies your task schedule by using a cron or rate expression.

Use cron expressions for task schedules that run on a specific time and day. For example,
the following cron expression creates a task schedule that runs at 8 AM on the first Wednesday
of every month:

`cron(0 8 * * 3#1)`

Use rate expressions for task schedules that run on a regular interval. For example, the
following rate expression creates a task schedule that runs every 12 hours:

`rate(12 hours)`

For information about cron and rate expression syntax, see the [_Amazon EventBridge User Guide_](../../../eventbridge/latest/userguide/eb-scheduled-rule-pattern.md).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\ \_\*\?\,\|\^\-\/\#\s\(\)\+]*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Specifies whether to enable or disable your task schedule. Your schedule is enabled by
default, but there can be situations where you need to disable it. For example, you might need
to pause a recurring transfer to fix an issue with your task or perform maintenance on your
storage system.

DataSync might disable your schedule automatically if your task fails repeatedly
with the same error. For more information, see [TaskScheduleDetails](../../../datasync/latest/userguide/api-taskscheduledetails.md).

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TaskReportConfigDestinationS3

Transferred

All content copied from https://docs.aws.amazon.com/.
