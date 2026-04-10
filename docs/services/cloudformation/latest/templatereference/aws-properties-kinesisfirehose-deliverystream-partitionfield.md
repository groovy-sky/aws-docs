This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream PartitionField

Represents a single field in a `PartitionSpec`.

Amazon Data Firehose is in preview release and is subject to change.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceName" : String
}

```

### YAML

```yaml

  SourceName: String

```

## Properties

`SourceName`

The column name to be configured in partition spec.

Amazon Data Firehose is in preview release and is subject to change.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParquetSerDe

PartitionSpec

All content copied from https://docs.aws.amazon.com/.
