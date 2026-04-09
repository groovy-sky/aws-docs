# RestoreTestingPlanForUpdate

This contains metadata about a restore testing plan.

## Contents

**RecoveryPointSelection**

Required: `Algorithm`; `RecoveryPointTypes`;
`IncludeVaults` ( _one or more_).

Optional: _SelectionWindowDays_ ( _'30' if_
_not specified_); `ExcludeVaults` (defaults to empty
list if not listed).

Type: [RestoreTestingRecoveryPointSelection](api-restoretestingrecoverypointselection.md) object

Required: No

**ScheduleExpression**

A CRON expression in specified timezone when a restore testing plan is executed. When no
CRON expression is provided, AWS Backup will use the default expression
`cron(0 5 ? * * *)`.

Type: String

Required: No

**ScheduleExpressionTimezone**

Optional. This is the timezone in which the schedule
expression is set. By default, ScheduleExpressions are in UTC.
You can modify this to a specified timezone.

Type: String

Required: No

**StartWindowHours**

Defaults to 24 hours.

A value in hours after a restore test is scheduled before a
job will be canceled if it doesn't start successfully. This value
is optional. If this value is included, this parameter has a
maximum value of 168 hours (one week).

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/restoretestingplanforupdate.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/restoretestingplanforupdate.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/restoretestingplanforupdate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestoreTestingPlanForList

RestoreTestingRecoveryPointSelection

All content copied from https://docs.aws.amazon.com/.
