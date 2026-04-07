This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudExadataInfrastructure MaintenanceWindow

The scheduling details for the maintenance window. Patching and system updates take place during the maintenance window.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomActionTimeoutInMins" : Integer,
  "DaysOfWeek" : [ String, ... ],
  "HoursOfDay" : [ Integer, ... ],
  "IsCustomActionTimeoutEnabled" : Boolean,
  "LeadTimeInWeeks" : Integer,
  "Months" : [ String, ... ],
  "PatchingMode" : String,
  "Preference" : String,
  "WeeksOfMonth" : [ Integer, ... ]
}

```

### YAML

```yaml

  CustomActionTimeoutInMins: Integer
  DaysOfWeek:
    - String
  HoursOfDay:
    - Integer
  IsCustomActionTimeoutEnabled: Boolean
  LeadTimeInWeeks: Integer
  Months:
    - String
  PatchingMode: String
  Preference: String
  WeeksOfMonth:
    - Integer

```

## Properties

`CustomActionTimeoutInMins`

The custom action timeout in minutes for the maintenance window.

_Required_: No

_Type_: Integer

_Minimum_: `15`

_Maximum_: `120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DaysOfWeek`

The days of the week when maintenance can be performed.

_Required_: No

_Type_: Array of String

_Allowed values_: `MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY | SUNDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HoursOfDay`

The hours of the day when maintenance can be performed.

_Required_: No

_Type_: Array of Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsCustomActionTimeoutEnabled`

Indicates whether custom action timeout is enabled for the maintenance window.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LeadTimeInWeeks`

The lead time in weeks before the maintenance window.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Months`

The months when maintenance can be performed.

_Required_: No

_Type_: Array of String

_Allowed values_: `JANUARY | FEBRUARY | MARCH | APRIL | MAY | JUNE | JULY | AUGUST | SEPTEMBER | OCTOBER | NOVEMBER | DECEMBER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatchingMode`

The patching mode for the maintenance window.

_Required_: No

_Type_: String

_Allowed values_: `ROLLING | NONROLLING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Preference`

The preference for the maintenance window scheduling.

_Required_: No

_Type_: String

_Allowed values_: `NO_PREFERENCE | CUSTOM_PREFERENCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeeksOfMonth`

The weeks of the month when maintenance can be performed.

_Required_: No

_Type_: Array of Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomerContact

Tag
