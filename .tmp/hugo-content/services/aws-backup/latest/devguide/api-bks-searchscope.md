# SearchScope

The search scope is all backup
properties input into a search.

## Contents

**BackupResourceTypes**

The resource types included in a search.

Eligible resource types include S3 and EBS.

Type: Array of strings

Array Members: Fixed number of 1 item.

Valid Values: `S3 | EBS`

Required: Yes

**BackupResourceArns**

The Amazon Resource Name (ARN) that uniquely identifies
the backup resources.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**BackupResourceCreationTime**

This is the time a backup resource was created.

Type: [BackupCreationTimeFilter](api-bks-backupcreationtimefilter.md) object

Required: No

**BackupResourceTags**

These are one or more tags on the backup (recovery
point).

Type: String to string map

Required: No

**SourceResourceArns**

The Amazon Resource Name (ARN) that uniquely identifies
the source resources.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backupsearch-2018-05-10/searchscope.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backupsearch-2018-05-10/searchscope.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backupsearch-2018-05-10/searchscope.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SearchJobSummary

SearchScopeSummary

All content copied from https://docs.aws.amazon.com/.
