This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetKinesisStreamParameters

The parameters for using a Kinesis stream as a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PartitionKey" : String
}

```

### YAML

```yaml

  PartitionKey: String

```

## Properties

`PartitionKey`

Determines which shard in the stream the data record is assigned to. Partition keys are
Unicode strings with a maximum length limit of 256 characters for each key. Amazon Kinesis Data Streams uses the partition key as input to a hash function that maps the
partition key and associated data to a specific shard. Specifically, an MD5 hash function
is used to map partition keys to 128-bit integer values and to map associated data records
to shards. As a result of this hashing mechanism, all data records with the same partition
key map to the same shard within the stream.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetHttpParameters

PipeTargetLambdaFunctionParameters

All content copied from https://docs.aws.amazon.com/.
