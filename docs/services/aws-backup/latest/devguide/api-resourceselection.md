# ResourceSelection

This contains metadata about resource selection for tiering configurations.

You can specify up to 5 different resource selections per tiering configuration.
Data moved to lower-cost tier remains there until deletion (one-way transition).

## Contents

**Resources**

An array of strings that either contains ARNs of the associated resources
or contains a wildcard `*` to specify all resources.
You can specify up to 100 specific resources per tiering configuration.

Type: Array of strings

Required: Yes

**ResourceType**

The type of AWS resource; for example, `S3` for Amazon S3.
For tiering configurations, this is currently limited to `S3`.

Type: String

Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

Required: Yes

**TieringDownSettingsInDays**

The number of days after creation within a backup vault that an object can transition to the low cost warm storage tier.
Must be a positive integer between 60 and 36500 days.

Type: Integer

Valid Range: Minimum value of 60. Maximum value of 36500.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/resourceselection.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/resourceselection.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/resourceselection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReportSetting

RestoreAccessBackupVaultListMember

All content copied from https://docs.aws.amazon.com/.
