# ScheduledInstanceRecurrence

Describes the recurring schedule for a Scheduled Instance.

## Contents

**frequency**

The frequency ( `Daily`, `Weekly`, or `Monthly`).

Type: String

Required: No

**interval**

The interval quantity. The interval unit depends on the value of `frequency`. For example, every 2
weeks or every 2 months.

Type: Integer

Required: No

**OccurrenceDaySet.N**

The days. For a monthly schedule, this is one or more days of the month (1-31). For a weekly schedule, this is one or more days of the week (1-7, where 1 is Sunday).

Type: Array of integers

Required: No

**occurrenceRelativeToEnd**

Indicates whether the occurrence is relative to the end of the specified week or month.

Type: Boolean

Required: No

**occurrenceUnit**

The unit for `occurrenceDaySet` ( `DayOfWeek` or `DayOfMonth`).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/scheduledinstancerecurrence.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/scheduledinstancerecurrence.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/scheduledinstancerecurrence.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ScheduledInstanceAvailability

ScheduledInstanceRecurrenceRequest
