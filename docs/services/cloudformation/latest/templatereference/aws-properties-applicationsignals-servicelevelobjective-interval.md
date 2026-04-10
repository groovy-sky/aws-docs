This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective Interval

The time period used to evaluate the SLO. It can be either a calendar interval or rolling interval.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CalendarInterval" : CalendarInterval,
  "RollingInterval" : RollingInterval
}

```

### YAML

```yaml

  CalendarInterval:
    CalendarInterval
  RollingInterval:
    RollingInterval

```

## Properties

`CalendarInterval`

If the interval is a calendar interval, this structure contains the interval specifications.

_Required_: No

_Type_: [CalendarInterval](aws-properties-applicationsignals-servicelevelobjective-calendarinterval.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollingInterval`

If the interval is a rolling interval, this structure contains the interval specifications.

_Required_: No

_Type_: [RollingInterval](aws-properties-applicationsignals-servicelevelobjective-rollinginterval.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Goal

Metric

All content copied from https://docs.aws.amazon.com/.
