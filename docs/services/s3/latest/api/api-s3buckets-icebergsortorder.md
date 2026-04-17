---
title: "IcebergSortOrder"
---

# IcebergSortOrder

Defines the sort order for data within an Iceberg table. Sorting data can improve query performance by enabling more efficient data skipping.

## Contents

**fields**

The list of sort fields that define how data is sorted within files. Each field specifies a source field, sort direction, and null ordering. This field is required if `writeOrder` is provided.

Type: Array of [IcebergSortField](api-s3buckets-icebergsortfield.md) objects

Required: Yes

**order-id**

The unique identifier for this sort order. If not specified, defaults to `1`. The order ID is used by Apache Iceberg to track sort order evolution.

Type: Integer

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/IcebergSortOrder)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/IcebergSortOrder)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/IcebergSortOrder)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergSortField

IcebergUnreferencedFileRemovalSettings

All content copied from https://docs.aws.amazon.com/.
