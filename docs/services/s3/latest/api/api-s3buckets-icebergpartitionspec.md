# IcebergPartitionSpec

Defines how data in an Iceberg table is partitioned. Partitioning helps optimize query performance by organizing data into separate files based on field values. Each partition field specifies a transform to apply to a source field.

## Contents

**fields**

The list of partition fields that define how the table data is partitioned. Each field specifies a source field and a transform to apply. This field is required if `partitionSpec` is provided.

Type: Array of [IcebergPartitionField](api-s3buckets-icebergpartitionfield.md) objects

Required: Yes

**spec-id**

The unique identifier for this partition specification. If not specified, defaults to `0`.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/icebergpartitionspec.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/icebergpartitionspec.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/icebergpartitionspec.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergPartitionField

IcebergSchema

All content copied from https://docs.aws.amazon.com/.
