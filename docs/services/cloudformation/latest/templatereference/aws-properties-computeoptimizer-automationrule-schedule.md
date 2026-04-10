This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ComputeOptimizer::AutomationRule Schedule

Configuration for scheduling when automation rules should execute, including timing and execution windows.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionWindowInMinutes" : Integer,
  "ScheduleExpression" : String,
  "ScheduleExpressionTimezone" : String
}

```

### YAML

```yaml

  ExecutionWindowInMinutes: Integer
  ScheduleExpression: String
  ScheduleExpressionTimezone: String

```

## Properties

`ExecutionWindowInMinutes`

The time window in minutes during which the automation rule can start implementing recommended actions.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

The expression that defines when the schedule runs. `cron` expression is supported. A `cron` expression consists of six fields separated by white spaces: ( `minutes` `hours` `day_of_month` `month` `day_of_week` `year`)

###### Note

You can schedule rules to run at most once per day. Your cron expression must use specific values (not wildcards) for the minutes and hours fields. For example: ( `30 12 * * *`) runs daily at 12:30 PM UTC.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpressionTimezone`

The timezone to use when interpreting the schedule expression.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceTagsCriteriaCondition

StringCriteriaCondition

All content copied from https://docs.aws.amazon.com/.
