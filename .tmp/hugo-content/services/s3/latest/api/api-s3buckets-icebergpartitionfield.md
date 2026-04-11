# IcebergPartitionField

Defines a single partition field in an Iceberg partition specification.

## Contents

**name**

The name for this partition field. This name is used in the partitioned file paths.

Type: String

Required: Yes

**source-id**

The ID of the source schema field to partition by. This must reference a valid field ID from the table schema.

Type: Integer

Required: Yes

**transform**

The partition transform to apply to the source field. Supported transforms include `identity`, `year`, `month`, `day`, `hour`, `bucket`, and `truncate`. For more information, see the [Apache Iceberg partition transforms documentation](https://iceberg.apache.org/spec).

Type: String

Required: Yes

**field-id**

An optional unique identifier for this partition field. If not specified, S3 Tables automatically assigns a field ID.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/icebergpartitionfield.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/icebergpartitionfield.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/icebergpartitionfield.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergMetadata

IcebergPartitionSpec

All content copied from https://docs.aws.amazon.com/.
