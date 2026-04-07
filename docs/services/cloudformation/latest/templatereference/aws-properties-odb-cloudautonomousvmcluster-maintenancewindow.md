This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudAutonomousVmCluster MaintenanceWindow

The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DaysOfWeek" : [ String, ... ],
  "HoursOfDay" : [ Integer, ... ],
  "LeadTimeInWeeks" : Integer,
  "Months" : [ String, ... ],
  "Preference" : String,
  "WeeksOfMonth" : [ Integer, ... ]
}

```

### YAML

```yaml

  DaysOfWeek:
    - String
  HoursOfDay:
    - Integer
  LeadTimeInWeeks: Integer
  Months:
    - String
  Preference: String
  WeeksOfMonth:
    - Integer

```

## Properties

`DaysOfWeek`

The days of the week when maintenance can be performed.

_Required_: No

_Type_: Array of String

_Allowed values_: `MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY | SUNDAY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HoursOfDay`

The hours of the day when maintenance can be performed.

_Required_: No

_Type_: Array of Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LeadTimeInWeeks`

The lead time in weeks before the maintenance window.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Months`

The months when maintenance can be performed.

_Required_: No

_Type_: Array of String

_Allowed values_: `JANUARY | FEBRUARY | MARCH | APRIL | MAY | JUNE | JULY | AUGUST | SEPTEMBER | OCTOBER | NOVEMBER | DECEMBER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Preference`

The preference for the maintenance window scheduling.

_Required_: No

_Type_: String

_Allowed values_: `NO_PREFERENCE | CUSTOM_PREFERENCE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WeeksOfMonth`

The weeks of the month when maintenance can be performed.

_Required_: No

_Type_: Array of Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IamRole

Tag
