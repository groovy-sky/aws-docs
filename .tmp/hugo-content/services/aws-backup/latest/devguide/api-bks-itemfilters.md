# ItemFilters

Item Filters represent all input item
properties specified when the search was
created.

Contains either EBSItemFilters or
S3ItemFilters

## Contents

**EBSItemFilters**

This array can contain CreationTimes,
FilePaths, LastModificationTimes, or Sizes objects.

Type: Array of [EBSItemFilter](api-bks-ebsitemfilter.md) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: No

**S3ItemFilters**

This array can contain CreationTimes, ETags,
ObjectKeys, Sizes, or VersionIds objects.

Type: Array of [S3ItemFilter](api-bks-s3itemfilter.md) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backupsearch-2018-05-10/itemfilters.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backupsearch-2018-05-10/itemfilters.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backupsearch-2018-05-10/itemfilters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportSpecification

LongCondition

All content copied from https://docs.aws.amazon.com/.
