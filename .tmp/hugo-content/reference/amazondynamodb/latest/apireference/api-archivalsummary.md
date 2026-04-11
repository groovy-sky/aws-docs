# ArchivalSummary

Contains details of a table archival operation.

## Contents

###### Note

In the following list, the required parameters are described first.

**ArchivalBackupArn**

The Amazon Resource Name (ARN) of the backup the table was archived to, when
applicable in the archival reason. If you wish to restore this backup to the same table
name, you will need to delete the original table.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**ArchivalDateTime**

The date and time when table archival was initiated by DynamoDB, in UNIX epoch time
format.

Type: Timestamp

Required: No

**ArchivalReason**

The reason DynamoDB archived the table. Currently, the only possible value is:

- `INACCESSIBLE_ENCRYPTION_CREDENTIALS` \- The table was archived due
to the table's AWS KMS key being inaccessible for more than seven
days. An On-Demand backup was created at the archival time.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/archivalsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/archivalsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/archivalsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon DynamoDB

AttributeDefinition

All content copied from https://docs.aws.amazon.com/.
