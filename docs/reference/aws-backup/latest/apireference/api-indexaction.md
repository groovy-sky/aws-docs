---
title: "IndexAction"
---

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/backup-2018-11-15/IndexAction)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/backup-2018-11-15/IndexAction)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/backup-2018-11-15/IndexAction)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameworkControl

IndexedRecoveryPoint

All content copied from https://docs.aws.amazon.com/.
