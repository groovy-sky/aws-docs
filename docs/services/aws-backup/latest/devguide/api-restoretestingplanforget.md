# RestoreTestingPlanForGet

This contains metadata about a restore testing plan.

## Contents

**CreationTime**

The date and time that a restore testing plan was created,
in Unix format and Coordinated Universal Time (UTC). The value of
`CreationTime` is accurate to milliseconds. For example,
the value 1516925490.087 represents Friday, January 26, 2018
12:11:30.087 AM.

Type: Timestamp

Required: Yes

**RecoveryPointSelection**

The specified criteria to assign a set of resources, such as
recovery point types or backup vaults.

Type: [RestoreTestingRecoveryPointSelection](api-restoretestingrecoverypointselection.md) object

Required: Yes

**RestoreTestingPlanArn**

An Amazon Resource Name (ARN) that uniquely identifies
a restore testing plan.

Type: String

Required: Yes

**RestoreTestingPlanName**

The restore testing plan name.

Type: String

Required: Yes

**ScheduleExpression**

A CRON expression in specified timezone when a restore
testing plan is executed. When no CRON expression is provided, AWS Backup will use the default
expression `cron(0 5 ? * * *)`.

Type: String

Required: Yes

**CreatorRequestId**

This identifies the request and allows failed requests to
be retried without the risk of running the operation twice.
If the request includes a `CreatorRequestId` that
matches an existing backup plan, that plan is returned. This
parameter is optional.

If used, this parameter must
contain 1 to 50 alphanumeric or '-\_.' characters.

Type: String

Required: No

**LastExecutionTime**

The last time a restore test was run with the specified
restore testing plan. A date and time, in Unix format and
Coordinated Universal Time (UTC). The value of
`LastExecutionDate` is accurate to milliseconds.
For example, the value 1516925490.087 represents Friday,
January 26, 2018 12:11:30.087 AM.

Type: Timestamp

Required: No

**LastUpdateTime**

The date and time that the restore testing plan was updated.
This update is in Unix format and Coordinated Universal Time (UTC).
The value of `LastUpdateTime` is accurate to milliseconds.
For example, the value 1516925490.087 represents Friday,
January 26, 2018 12:11:30.087 AM.

Type: Timestamp

Required: No

**ScheduleExpressionTimezone**

Optional. This is the timezone in which the schedule
expression is set. By default, ScheduleExpressions are in UTC.
You can modify this to a specified timezone.

Type: String

Required: No

**StartWindowHours**

Defaults to 24 hours.

A value in hours after a
restore test is scheduled before a job will be canceled if it
doesn't start successfully. This value is optional. If this value
is included, this parameter has a maximum value of 168 hours
(one week).

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/restoretestingplanforget.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/restoretestingplanforget.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/restoretestingplanforget.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestoreTestingPlanForCreate

RestoreTestingPlanForList

All content copied from https://docs.aws.amazon.com/.
