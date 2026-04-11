# EBSItemFilter

This contains arrays of objects, which may include
CreationTimes time condition objects, FilePaths
string objects, LastModificationTimes time
condition objects,

## Contents

**CreationTimes**

You can include 1 to 10 values.

If one is included, the results will
return only items that match.

If more than one is included, the
results will return all items that match any of
the included values.

Type: Array of [TimeCondition](api-bks-timecondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**FilePaths**

You can include 1 to 10 values.

If one file path is included, the results will
return only items that match the file path.

If more than one file path is included, the
results will return all items that match any of the
file paths.

Type: Array of [StringCondition](api-bks-stringcondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**LastModificationTimes**

You can include 1 to 10 values.

If one is included, the results will
return only items that match.

If more than one is included, the
results will return all items that match any of
the included values.

Type: Array of [TimeCondition](api-bks-timecondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

**Sizes**

You can include 1 to 10 values.

If one is included, the results will
return only items that match.

If more than one is included, the
results will return all items that match any of
the included values.

Type: Array of [LongCondition](api-bks-longcondition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backupsearch-2018-05-10/ebsitemfilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backupsearch-2018-05-10/ebsitemfilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backupsearch-2018-05-10/ebsitemfilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CurrentSearchProgress

EBSResultItem

All content copied from https://docs.aws.amazon.com/.
