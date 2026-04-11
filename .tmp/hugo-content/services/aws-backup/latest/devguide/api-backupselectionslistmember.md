# BackupSelectionsListMember

Contains metadata about a `BackupSelection` object.

## Contents

**BackupPlanId**

Uniquely identifies a backup plan.

Type: String

Required: No

**CreationDate**

The date and time a backup plan is created, in Unix format and Coordinated Universal
Time (UTC). The value of `CreationDate` is accurate to milliseconds. For
example, the value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087
AM.

Type: Timestamp

Required: No

**CreatorRequestId**

A unique string that identifies the request and allows failed requests to be retried
without the risk of running the operation twice. This parameter is optional.

If used, this parameter must contain 1 to 50 alphanumeric or '-\_.' characters.

Type: String

Required: No

**IamRoleArn**

Specifies the IAM role Amazon Resource Name (ARN) to create the target recovery point;
for example, `arn:aws:iam::123456789012:role/S3Access`.

Type: String

Required: No

**SelectionId**

Uniquely identifies a request to assign a set of resources to a backup plan.

Type: String

Required: No

**SelectionName**

The display name of a resource selection document.

Type: String

Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/backupselectionslistmember.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/backupselectionslistmember.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/backupselectionslistmember.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupSelection

BackupVaultListMember

All content copied from https://docs.aws.amazon.com/.
