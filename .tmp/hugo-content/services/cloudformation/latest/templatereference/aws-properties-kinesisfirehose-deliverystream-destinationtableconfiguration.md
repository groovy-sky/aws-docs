This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DestinationTableConfiguration

Describes the configuration of a destination in Apache Iceberg Tables. This section is only needed for tables where you want to update or delete data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationDatabaseName" : String,
  "DestinationTableName" : String,
  "PartitionSpec" : PartitionSpec,
  "S3ErrorOutputPrefix" : String,
  "UniqueKeys" : [ String, ... ]
}

```

### YAML

```yaml

  DestinationDatabaseName: String
  DestinationTableName: String
  PartitionSpec:
    PartitionSpec
  S3ErrorOutputPrefix: String
  UniqueKeys:
    - String

```

## Properties

`DestinationDatabaseName`

The name of the Apache Iceberg database.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationTableName`

Specifies the name of the Apache Iceberg Table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartitionSpec`

The partition spec configuration for a table that is used by automatic table
creation.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: [PartitionSpec](aws-properties-kinesisfirehose-deliverystream-partitionspec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ErrorOutputPrefix`

The table specific S3 error output prefix. All the errors that occurred while delivering to this table will be prefixed with this value in S3 destination.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UniqueKeys`

A list of unique keys for a given Apache Iceberg table. Firehose will use these for running Create, Update, or Delete operations on the given Iceberg table.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deserializer

DirectPutSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
