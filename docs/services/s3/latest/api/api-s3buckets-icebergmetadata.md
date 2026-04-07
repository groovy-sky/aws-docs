# IcebergMetadata

Contains details about the metadata for an Iceberg table.

## Contents

**partitionSpec**

The partition specification for the Iceberg table. Partitioning organizes data into separate files based on the values of one or more fields, which can improve query performance by reducing the amount of data scanned. Each partition field applies a transform (such as identity, year, month, or bucket) to a single field.

Type: [IcebergPartitionSpec](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_IcebergPartitionSpec.html) object

Required: No

**properties**

A map of custom configuration properties for the Iceberg table.

Type: String to string map

Required: No

**schema**

The schema for an Iceberg table. Use this property to define table schemas with primitive types only. For schemas that include nested or complex types such as `struct`, `list`, or `map`, use `schemaV2` instead.

Type: [IcebergSchema](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_IcebergSchema.html) object

Required: No

**schemaV2**

The schema for an Iceberg table using the V2 format. Use this property to define table schemas that include nested or complex data types such as `struct`, `list`, or `map`, in addition to primitive types. For schemas with only primitive types, you can use either `schema` or `schemaV2`.

Type: [IcebergSchemaV2](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_IcebergSchemaV2.html) object

Required: No

**writeOrder**

The sort order for the Iceberg table. Sort order defines how data is sorted within data files, which can improve query performance by enabling more efficient data skipping and filtering.

Type: [IcebergSortOrder](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_IcebergSortOrder.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/IcebergMetadata)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/IcebergMetadata)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/IcebergMetadata)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IcebergCompactionSettings

IcebergPartitionField
