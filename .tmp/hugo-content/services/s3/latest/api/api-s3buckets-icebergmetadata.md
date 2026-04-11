# IcebergMetadata

Contains details about the metadata for an Iceberg table.

## Contents

**partitionSpec**

The partition specification for the Iceberg table. Partitioning organizes data into separate files based on the values of one or more fields, which can improve query performance by reducing the amount of data scanned. Each partition field applies a transform (such as identity, year, month, or bucket) to a single field.

Type: [IcebergPartitionSpec](api-s3buckets-icebergpartitionspec.md) object

Required: No

**properties**

A map of custom configuration properties for the Iceberg table.

Type: String to string map

Required: No

**schema**

The schema for an Iceberg table. Use this property to define table schemas with primitive types only. For schemas that include nested or complex types such as `struct`, `list`, or `map`, use `schemaV2` instead.

Type: [IcebergSchema](api-s3buckets-icebergschema.md) object

Required: No

**schemaV2**

The schema for an Iceberg table using the V2 format. Use this property to define table schemas that include nested or complex data types such as `struct`, `list`, or `map`, in addition to primitive types. For schemas with only primitive types, you can use either `schema` or `schemaV2`.

Type: [IcebergSchemaV2](api-s3buckets-icebergschemav2.md) object

Required: No

**writeOrder**

The sort order for the Iceberg table. Sort order defines how data is sorted within data files, which can improve query performance by enabling more efficient data skipping and filtering.

Type: [IcebergSortOrder](api-s3buckets-icebergsortorder.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/icebergmetadata.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/icebergmetadata.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/icebergmetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergCompactionSettings

IcebergPartitionField

All content copied from https://docs.aws.amazon.com/.
