This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::RefreshSchedule RefreshScheduleMap

A summary of a configured refresh schedule for a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RefreshType" : String,
  "ScheduleFrequency" : ScheduleFrequency,
  "ScheduleId" : String,
  "StartAfterDateTime" : String
}

```

### YAML

```yaml

  RefreshType: String
  ScheduleFrequency:
    ScheduleFrequency
  ScheduleId: String
  StartAfterDateTime: String

```

## Properties

`RefreshType`

The type of refresh that a dataset undergoes. Valid values are as follows:

- `FULL_REFRESH`: A complete refresh of a dataset.

- `INCREMENTAL_REFRESH`: A partial refresh of some rows of a dataset, based on the time window
specified.

For more information on full and incremental refreshes, see [Refreshing SPICE data](../../../quicksight/latest/user/refreshing-imported-data.md) in the _Quick User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `FULL_REFRESH | INCREMENTAL_REFRESH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleFrequency`

The frequency for the refresh schedule.

_Required_: No

_Type_: [ScheduleFrequency](aws-properties-quicksight-refreshschedule-schedulefrequency.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleId`

An identifier for the refresh schedule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartAfterDateTime`

Time after which the refresh schedule can be started, expressed in `YYYY-MM-DDTHH:MM:SS`
format.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RefreshOnDay

ScheduleFrequency

All content copied from https://docs.aws.amazon.com/.
