# TieringConfiguration

This contains metadata about a tiering configuration.

## Contents

**BackupVaultName**

The name of the backup vault where the tiering configuration applies.
Use `*` to apply to all backup vaults.

Type: String

Pattern: `^(\*|[a-zA-Z0-9\-\_]{2,50})$`

Required: Yes

**ResourceSelection**

An array of resource selection objects that specify which resources
are included in the tiering configuration and their tiering settings.

Type: Array of [ResourceSelection](api-resourceselection.md) objects

Required: Yes

**TieringConfigurationName**

The unique name of the tiering configuration. This cannot be changed
after creation, and it must consist of only alphanumeric characters and underscores.

Type: String

Pattern: `^[a-zA-Z0-9_]{1,200}$`

Required: Yes

**CreationTime**

The date and time a tiering configuration was created, in Unix format
and Coordinated Universal Time (UTC). The value of `CreationTime`
is accurate to milliseconds. For example, the value 1516925490.087 represents
Friday, January 26, 2018 12:11:30.087AM.

Type: Timestamp

Required: No

**CreatorRequestId**

This is a unique string that identifies the request and
allows failed requests to be retried without the risk of running
the operation twice.

Type: String

Required: No

**LastUpdatedTime**

The date and time a tiering configuration was updated, in Unix format
and Coordinated Universal Time (UTC). The value of `LastUpdatedTime`
is accurate to milliseconds. For example, the value 1516925490.087 represents
Friday, January 26, 2018 12:11:30.087AM.

Type: Timestamp

Required: No

**TieringConfigurationArn**

An Amazon Resource Name (ARN) that uniquely identifies the
tiering configuration.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/tieringconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/tieringconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/tieringconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScheduledPlanExecutionMember

TieringConfigurationInputForCreate

All content copied from https://docs.aws.amazon.com/.
