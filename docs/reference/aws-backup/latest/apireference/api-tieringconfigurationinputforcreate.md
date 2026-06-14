---
title: "TieringConfigurationInputForCreate"
---

# TieringConfigurationInputForCreate

This contains metadata about a tiering configuration for create operations.

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/backup-2018-11-15/TieringConfigurationInputForCreate)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/backup-2018-11-15/TieringConfigurationInputForCreate)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/backup-2018-11-15/TieringConfigurationInputForCreate)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TieringConfiguration

TieringConfigurationInputForUpdate

All content copied from https://docs.aws.amazon.com/.
