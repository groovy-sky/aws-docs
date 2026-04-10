This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AlarmMuteRule Rule

The configuration that defines when and how long alarms should be muted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Schedule" : Schedule
}

```

### YAML

```yaml

  Schedule:
    Schedule

```

## Properties

`Schedule`

Defines the schedule configuration for an alarm mute rule.

The rule contains a schedule that specifies when and how long alarms should be muted. The schedule can be a recurring pattern using cron expressions or a one-time mute window using at expressions.

_Required_: Yes

_Type_: [Schedule](aws-properties-cloudwatch-alarmmuterule-schedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MuteTargets

Schedule

All content copied from https://docs.aws.amazon.com/.
