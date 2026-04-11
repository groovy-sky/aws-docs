This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::RefreshSchedule ScheduleFrequency

The frequency for the refresh schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Interval" : String,
  "RefreshOnDay" : RefreshOnDay,
  "TimeOfTheDay" : String,
  "TimeZone" : String
}

```

### YAML

```yaml

  Interval: String
  RefreshOnDay:
    RefreshOnDay
  TimeOfTheDay: String
  TimeZone: String

```

## Properties

`Interval`

The interval between scheduled refreshes. Valid values are as follows:

- `MINUTE15`: The dataset refreshes every 15 minutes. This value is only supported for incremental
refreshes. This interval can only be used for one schedule per dataset.

- `MINUTE30`: The dataset refreshes every 30 minutes. This value is only supported for incremental
refreshes. This interval can only be used for one schedule per dataset.

- `HOURLY`: The dataset refreshes every hour. This interval can only be used for one schedule per
dataset.

- `DAILY`: The dataset refreshes every day.

- `WEEKLY`: The dataset refreshes every week.

- `MONTHLY`: The dataset refreshes every month.

_Required_: No

_Type_: String

_Allowed values_: `MINUTE15 | MINUTE30 | HOURLY | DAILY | WEEKLY | MONTHLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshOnDay`

The day of the week that you want to schedule the refresh on. This value is required for weekly and monthly
refresh intervals.

_Required_: No

_Type_: [RefreshOnDay](aws-properties-quicksight-refreshschedule-refreshonday.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeOfTheDay`

The time of day that you want the dataset to refresh. This value is expressed in HH:MM format. This field is not
required for schedules that refresh hourly.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The timezone that you want the refresh schedule to use. The timezone ID must match a corresponding ID found on
`java.util.time.getAvailableIDs()`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RefreshScheduleMap

AWS::QuickSight::Template

All content copied from https://docs.aws.amazon.com/.
