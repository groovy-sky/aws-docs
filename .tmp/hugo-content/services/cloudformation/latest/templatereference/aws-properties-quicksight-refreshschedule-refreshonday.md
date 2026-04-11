This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::RefreshSchedule RefreshOnDay

The day that you want yout dataset to refresh.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DayOfMonth" : String,
  "DayOfWeek" : String
}

```

### YAML

```yaml

  DayOfMonth: String
  DayOfWeek: String

```

## Properties

`DayOfMonth`

The day of the month that you want your dataset to refresh. This value is required for monthly refresh
intervals.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DayOfWeek`

The day of the week that you want to schedule the refresh on. This value is required for weekly and monthly
refresh intervals.

_Required_: No

_Type_: String

_Allowed values_: `SUNDAY | MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::QuickSight::RefreshSchedule

RefreshScheduleMap

All content copied from https://docs.aws.amazon.com/.
