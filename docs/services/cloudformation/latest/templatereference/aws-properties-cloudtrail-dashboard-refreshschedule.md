This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Dashboard RefreshSchedule

The schedule for a dashboard refresh.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Frequency" : Frequency,
  "Status" : String,
  "TimeOfDay" : String
}

```

### YAML

```yaml

  Frequency:
    Frequency
  Status: String
  TimeOfDay: String

```

## Properties

`Frequency`

The frequency at which you want the dashboard refreshed.

_Required_: No

_Type_: [Frequency](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-dashboard-frequency.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Specifies whether the refresh schedule is enabled. Set the value to `ENABLED` to enable the refresh schedule, or to `DISABLED` to turn off the refresh schedule.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeOfDay`

The time of day in UTC to run the schedule; for hourly only refer to minutes; default is 00:00.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{2}:[0-9]{2}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Frequency

Tag
