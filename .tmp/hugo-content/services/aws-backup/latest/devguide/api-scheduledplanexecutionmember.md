# ScheduledPlanExecutionMember

Contains information about a scheduled backup plan execution, including the execution time, rule type, and associated rule identifier.

## Contents

**ExecutionTime**

The timestamp when the backup is scheduled to run, in Unix format and Coordinated Universal Time (UTC). The value is accurate to milliseconds.

Type: Timestamp

Required: No

**RuleExecutionType**

The type of backup rule execution. Valid values are `CONTINUOUS` (point-in-time recovery), `SNAPSHOTS` (snapshot backups), or `CONTINUOUS_AND_SNAPSHOTS` (both types combined).

Type: String

Valid Values: `CONTINUOUS | SNAPSHOTS | CONTINUOUS_AND_SNAPSHOTS`

Required: No

**RuleId**

The unique identifier of the backup rule that will execute at the scheduled time.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/scheduledplanexecutionmember.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/scheduledplanexecutionmember.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/scheduledplanexecutionmember.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScanSetting

TieringConfiguration

All content copied from https://docs.aws.amazon.com/.
