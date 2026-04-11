# ScanJobCreator

Contains identifying information about the creation of a scan job, including the backup plan and rule that initiated the scan.

## Contents

**BackupPlanArn**

An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for example,
`arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50`.

Type: String

Required: Yes

**BackupPlanId**

The ID of the backup plan.

Type: String

Required: Yes

**BackupPlanVersion**

Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most 1,024 bytes
long. Version IDs cannot be edited.

Type: String

Required: Yes

**BackupRuleId**

Uniquely identifies the backup rule that initiated the scan job.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/scanjobcreator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/scanjobcreator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/scanjobcreator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScanJob

ScanJobSummary

All content copied from https://docs.aws.amazon.com/.
