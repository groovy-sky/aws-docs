This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream PartitionSpec

Represents how to produce partition data for a table. Partition data is produced by
transforming columns in a table. Each column transform is represented by a named
`PartitionField`.

Here is an example of the schema in JSON.

`"partitionSpec": { "identity": [ {"sourceName": "column1"}, {"sourceName":
            "column2"}, {"sourceName": "column3"} ] }`

Amazon Data Firehose is in preview release and is subject to change.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Identity" : [ PartitionField, ... ]
}

```

### YAML

```yaml

  Identity:
    - PartitionField

```

## Properties

`Identity`

List of identity [transforms](https://iceberg.apache.org/spec) that performs an identity transformation. The transform takes the
source value, and does not modify it. Result type is the source type.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: No

_Type_: Array of [PartitionField](aws-properties-kinesisfirehose-deliverystream-partitionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PartitionField

ProcessingConfiguration

All content copied from https://docs.aws.amazon.com/.
