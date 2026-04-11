# RestoreAccessBackupVaultListMember

Contains information about a restore access backup vault.

## Contents

**ApprovalDate**

The date and time when the restore access backup vault was approved.

Type: Timestamp

Required: No

**CreationDate**

The date and time when the restore access backup vault was created.

Type: Timestamp

Required: No

**LatestRevokeRequest**

Information about the latest request to revoke access to this backup vault.

Type: [LatestRevokeRequest](api-latestrevokerequest.md) object

Required: No

**RestoreAccessBackupVaultArn**

The ARN of the restore access backup vault.

Type: String

Required: No

**VaultState**

The current state of the restore access backup vault.

Type: String

Valid Values: `CREATING | AVAILABLE | FAILED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/restoreaccessbackupvaultlistmember.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/restoreaccessbackupvaultlistmember.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/restoreaccessbackupvaultlistmember.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceSelection

RestoreJobCreator

All content copied from https://docs.aws.amazon.com/.
