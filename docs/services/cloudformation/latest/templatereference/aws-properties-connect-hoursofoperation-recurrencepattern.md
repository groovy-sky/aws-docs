This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::HoursOfOperation RecurrencePattern

Specifies the detailed pattern for event recurrence. Use this to define complex scheduling rules such as "every 2nd Tuesday of the month" or "every 3 months on the 15th".

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ByMonth" : [ Integer, ... ],
  "ByMonthDay" : [ Integer, ... ],
  "ByWeekdayOccurrence" : [ Integer, ... ],
  "Frequency" : String,
  "Interval" : Integer
}

```

### YAML

```yaml

  ByMonth:
    - Integer
  ByMonthDay:
    - Integer
  ByWeekdayOccurrence:
    - Integer
  Frequency: String
  Interval: Integer

```

## Properties

`ByMonth`

Specifies which month the event should occur in (1-12, where 1=January, 12=December). Used with YEARLY frequency to schedule events in specific month.

Note: It does not accept multiple values in the same list

_Required_: No

_Type_: Array of Integer

_Minimum_: `1`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ByMonthDay`

Specifies which day of the month the event should occur on (1-31). Used with MONTHLY or YEARLY frequency to schedule events on specific date within a month.

Examples:
\[15\] for events on the 15th of each month,
\[-1\] for events on the last day of month.

Note: It does not accept multiple values in the same list. If a specified day doesn't exist in a particular month (e.g., day 31 in February), the event will be skipped for that month. This field cannot be used simultaneously with ByWeekdayOccurrence as they represent different scheduling approaches (specific dates vs. relative weekday positions).

_Required_: No

_Type_: Array of Integer

_Minimum_: `-1`

_Maximum_: `31`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ByWeekdayOccurrence`

Specifies which occurrence of a weekday within the month the event should occur on. Must be used with MONTHLY or YEARLY frequency.

Example: 2 corresponds to second occurrence of the weekday in the month. -1 corresponds to last occurrence of the weekday in the month

The weekday itself is specified separately in the HoursOfOperationConfig. Example: To schedule the recurring event for the 2nd Thursday of April every year, set ByWeekdayOccurrence=\[2\], Day=THURSDAY, ByMonth=\[4\], Frequency: YEARLY and INTERVAL=1.

_Required_: No

_Type_: Array of Integer

_Minimum_: `0 | -1`

_Maximum_: `1 | 4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Frequency`

Defines how often the pattern repeats. This is the base unit for the recurrence schedule and works in conjunction with the Interval field to determine the exact repetition sequence.

_Required_: No

_Type_: String

_Allowed values_: `WEEKLY | MONTHLY | YEARLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

Specifies the number of frequency units between each occurrence. Must be a positive integer.

Examples: To repeat every week, set Interval=1 with WEEKLY frequency. To repeat every two months, set Interval=2 with MONTHLY frequency.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecurrenceConfig

Tag

All content copied from https://docs.aws.amazon.com/.
