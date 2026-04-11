This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CisScanConfiguration MonthlySchedule

A monthly schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Day" : String,
  "StartTime" : Time
}

```

### YAML

```yaml

  Day: String
  StartTime:
    Time

```

## Properties

`Day`

The monthly schedule's day.

_Required_: Yes

_Type_: String

_Allowed values_: `MON | TUE | WED | THU | FRI | SAT | SUN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The monthly schedule's start time.

_Required_: Yes

_Type_: [Time](aws-properties-inspectorv2-cisscanconfiguration-time.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DailySchedule

Schedule

All content copied from https://docs.aws.amazon.com/.
