# StorageLensGroup

A custom grouping of objects that include filters for prefixes, suffixes, object tags,
object size, or object age. You can create an S3 Storage Lens group that includes a single
filter or multiple filter conditions. To specify multiple filter conditions, you use
`AND` or `OR` logical operators.

## Contents

**Filter**

Sets the criteria for the Storage Lens group data that is displayed. For multiple filter
conditions, the `AND` or `OR` logical operator is used.

Type: [StorageLensGroupFilter](api-control-storagelensgroupfilter.md) data type

Required: Yes

**Name**

Contains the name of the Storage Lens group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[a-zA-Z0-9\-\_]+`

Required: Yes

**StorageLensGroupArn**

Contains the Amazon Resource Name (ARN) of the Storage Lens group. This property is
read-only.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 1024.

Pattern: `arn:[a-z\-]+:s3:[a-z0-9\-]+:\d{12}:storage\-lens\-group\/.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/storagelensgroup.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/storagelensgroup.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/storagelensgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensExpandedPrefixesDataExport

StorageLensGroupAndOperator

All content copied from https://docs.aws.amazon.com/.
