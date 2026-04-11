This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeSourceKinesisStreamParameters

The parameters for using a Kinesis stream as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchSize" : Integer,
  "DeadLetterConfig" : DeadLetterConfig,
  "MaximumBatchingWindowInSeconds" : Integer,
  "MaximumRecordAgeInSeconds" : Integer,
  "MaximumRetryAttempts" : Integer,
  "OnPartialBatchItemFailure" : String,
  "ParallelizationFactor" : Integer,
  "StartingPosition" : String,
  "StartingPositionTimestamp" : String
}

```

### YAML

```yaml

  BatchSize: Integer
  DeadLetterConfig:
    DeadLetterConfig
  MaximumBatchingWindowInSeconds: Integer
  MaximumRecordAgeInSeconds: Integer
  MaximumRetryAttempts: Integer
  OnPartialBatchItemFailure: String
  ParallelizationFactor: Integer
  StartingPosition: String
  StartingPositionTimestamp: String

```

## Properties

`BatchSize`

The maximum number of records to include in each batch.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeadLetterConfig`

Define the target queue to send dead-letter queue events to.

_Required_: No

_Type_: [DeadLetterConfig](aws-properties-pipes-pipe-deadletterconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBatchingWindowInSeconds`

The maximum length of a time to wait for events.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRecordAgeInSeconds`

Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite.
When the value is set to infinite, EventBridge never discards old records.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `604800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRetryAttempts`

Discard records after the specified number of retries. The default value is -1, which sets the maximum number of
retries to infinite. When MaximumRetryAttempts is infinite, EventBridge retries failed records until the record expires in the event source.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnPartialBatchItemFailure`

Define how to handle item process failures. `AUTOMATIC_BISECT` halves each batch and retry each half
until all the records are processed or there is one failed message left in the batch.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC_BISECT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelizationFactor`

The number of batches to process concurrently from each shard. The default value is 1.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartingPosition`

The position in a stream from which to start reading.

_Required_: Yes

_Type_: String

_Allowed values_: `TRIM_HORIZON | LATEST | AT_TIMESTAMP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartingPositionTimestamp`

With `StartingPosition` set to `AT_TIMESTAMP`, the time from which
to start reading, in Unix time seconds.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeSourceDynamoDBStreamParameters

PipeSourceManagedStreamingKafkaParameters

All content copied from https://docs.aws.amazon.com/.
