---
title: "ResultItem"
---

# ResultItem

This is an object representing the item
returned in the results of a search for a specific
resource type.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**EBSResultItem**

These are items returned in the search results
of an Amazon EBS search.

Type: [EBSResultItem](api-bks-ebsresultitem.md) object

Required: No

**S3ResultItem**

These are items returned in the search results
of an Amazon S3 search.

Type: [S3ResultItem](api-bks-s3resultitem.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/backupsearch-2018-05-10/ResultItem)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/backupsearch-2018-05-10/ResultItem)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/backupsearch-2018-05-10/ResultItem)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LongCondition

S3ExportSpecification

All content copied from https://docs.aws.amazon.com/.
