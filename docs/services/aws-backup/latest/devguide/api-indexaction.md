# IndexAction

This is an optional array within a BackupRule.

IndexAction consists of one ResourceTypes.

## Contents

**ResourceTypes**

0 or 1 index action will be accepted for each BackupRule.

Valid values:

- `EBS` for Amazon Elastic Block Store

- `S3` for Amazon Simple Storage Service (Amazon S3)

Type: Array of strings

Pattern: `^[a-zA-Z0-9\-\_\.]{1,50}$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/indexaction.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/indexaction.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/indexaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameworkControl

IndexedRecoveryPoint

All content copied from https://docs.aws.amazon.com/.
