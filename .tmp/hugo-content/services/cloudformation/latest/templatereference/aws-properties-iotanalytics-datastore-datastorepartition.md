This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore DatastorePartition

A single dimension to partition a data store. The dimension must be an `AttributePartition` or a `TimestampPartition`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Partition" : Partition,
  "TimestampPartition" : TimestampPartition
}

```

### YAML

```yaml

  Partition:
    Partition
  TimestampPartition:
    TimestampPartition

```

## Properties

`Partition`

A partition dimension defined by an attribute.

_Required_: No

_Type_: [Partition](aws-properties-iotanalytics-datastore-partition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimestampPartition`

A partition dimension defined by a timestamp attribute.

_Required_: No

_Type_: [TimestampPartition](aws-properties-iotanalytics-datastore-timestamppartition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomerManagedS3Storage

DatastorePartitions

All content copied from https://docs.aws.amazon.com/.
