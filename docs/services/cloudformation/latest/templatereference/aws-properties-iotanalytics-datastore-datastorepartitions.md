This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore DatastorePartitions

Information about the partition dimensions in a data store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Partitions" : [ DatastorePartition, ... ]
}

```

### YAML

```yaml

  Partitions:
    - DatastorePartition

```

## Properties

`Partitions`

A list of partition dimensions in a data store.

_Required_: No

_Type_: Array of [DatastorePartition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-datastore-datastorepartition.html)

_Minimum_: `0`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DatastorePartition

DatastoreStorage
