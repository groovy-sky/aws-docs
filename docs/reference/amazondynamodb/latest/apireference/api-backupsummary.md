# BackupSummary

Contains details for the backup.

## Contents

###### Note

In the following list, the required parameters are described first.

**BackupArn**

ARN associated with the backup.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**BackupCreationDateTime**

Time at which the backup was created.

Type: Timestamp

Required: No

**BackupExpiryDateTime**

Time at which the automatic on-demand backup created by DynamoDB will
expire. This `SYSTEM` on-demand backup expires automatically 35 days after
its creation.

Type: Timestamp

Required: No

**BackupName**

Name of the specified backup.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**BackupSizeBytes**

Size of the backup in bytes.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**BackupStatus**

Backup can be in one of the following states: CREATING, ACTIVE, DELETED.

Type: String

Valid Values: `CREATING | DELETED | AVAILABLE`

Required: No

**BackupType**

BackupType:

- `USER` \- You create and manage these using the on-demand backup
feature.

- `SYSTEM` \- If you delete a table with point-in-time recovery enabled,
a `SYSTEM` backup is automatically created and is retained for 35
days (at no additional cost). System backups allow you to restore the deleted
table to the state it was in just before the point of deletion.

- `AWS_BACKUP` \- On-demand backup created by you from AWS Backup service.

Type: String

Valid Values: `USER | SYSTEM | AWS_BACKUP`

Required: No

**TableArn**

ARN associated with the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**TableId**

Unique identifier for the table.

Type: String

Pattern: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

Required: No

**TableName**

Name of the table.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/backupsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/backupsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/backupsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupDetails

BatchStatementError

All content copied from https://docs.aws.amazon.com/.
