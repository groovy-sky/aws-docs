This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeSourceManagedStreamingKafkaParameters

The parameters for using an MSK stream as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchSize" : Integer,
  "ConsumerGroupID" : String,
  "Credentials" : MSKAccessCredentials,
  "MaximumBatchingWindowInSeconds" : Integer,
  "StartingPosition" : String,
  "TopicName" : String
}

```

### YAML

```yaml

  BatchSize: Integer
  ConsumerGroupID: String
  Credentials:
    MSKAccessCredentials
  MaximumBatchingWindowInSeconds: Integer
  StartingPosition: String
  TopicName: String

```

## Properties

`BatchSize`

The maximum number of records to include in each batch.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConsumerGroupID`

The name of the destination queue to consume.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-\/*:_+=.@-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Credentials`

The credentials needed to access the resource.

_Required_: No

_Type_: [MSKAccessCredentials](aws-properties-pipes-pipe-mskaccesscredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBatchingWindowInSeconds`

The maximum length of a time to wait for events.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartingPosition`

The position in a stream from which to start reading.

_Required_: No

_Type_: String

_Allowed values_: `TRIM_HORIZON | LATEST`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicName`

The name of the topic that the pipe will read from.

_Required_: Yes

_Type_: String

_Pattern_: `^[^.]([a-zA-Z0-9\-_.]+)$`

_Minimum_: `1`

_Maximum_: `249`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeSourceKinesisStreamParameters

PipeSourceParameters

All content copied from https://docs.aws.amazon.com/.
