# S3ItemFilter

This contains arrays of objects, which may include
ObjectKeys, Sizes, CreationTimes, VersionIds, and/or
Etags.

## Contents

**CreationTimes**

You can include 1 to 10 values.

If one value is included, the results will
return only items that match the value.

If more than one value is included, the
results will return all items that match any of the
values.

Type: Array of [TimeCondition](api-bks-timecondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**ETags**

You can include 1 to 10 values.

If one value is included, the results will
return only items that match the value.

If more than one value is included, the
results will return all items that match any of the
values.

Type: Array of [StringCondition](api-bks-stringcondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**ObjectKeys**

You can include 1 to 10 values.

If one value is included, the results will
return only items that match the value.

If more than one value is included, the
results will return all items that match any of the
values.

Type: Array of [StringCondition](api-bks-stringcondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**Sizes**

You can include 1 to 10 values.

If one value is included, the results will
return only items that match the value.

If more than one value is included, the
results will return all items that match any of the
values.

Type: Array of [LongCondition](api-bks-longcondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**VersionIds**

You can include 1 to 10 values.

If one value is included, the results will
return only items that match the value.

If more than one value is included, the
results will return all items that match any of the
values.

Type: Array of [StringCondition](api-bks-stringcondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backupsearch-2018-05-10/s3itemfilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backupsearch-2018-05-10/s3itemfilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backupsearch-2018-05-10/s3itemfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3ExportSpecification

S3ResultItem

All content copied from https://docs.aws.amazon.com/.
