This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::TableBucket ReplicationDestination

Specifies a destination table bucket for replication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationTableBucketARN" : String
}

```

### YAML

```yaml

  DestinationTableBucketARN: String

```

## Properties

`DestinationTableBucketARN`

The Amazon Resource Name (ARN) of the destination table bucket where tables will be replicated.

_Required_: Yes

_Type_: String

_Pattern_: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationConfiguration

ReplicationRule

All content copied from https://docs.aws.amazon.com/.
