This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket TargetObjectKeyFormat

Amazon S3 key format for log objects. Only one format, PartitionedPrefix or SimplePrefix, is
allowed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PartitionedPrefix" : PartitionedPrefix,
  "SimplePrefix" : Json
}

```

### YAML

```yaml

  PartitionedPrefix:
    PartitionedPrefix
  SimplePrefix: Json

```

## Properties

`PartitionedPrefix`

Partitioned S3 key for log objects.

_Required_: No

_Type_: [PartitionedPrefix](aws-properties-s3-bucket-partitionedprefix.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SimplePrefix`

To use the simple format for S3 keys for log objects. To specify SimplePrefix format, set
SimplePrefix to {}.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagFilter

Tiering

All content copied from https://docs.aws.amazon.com/.
