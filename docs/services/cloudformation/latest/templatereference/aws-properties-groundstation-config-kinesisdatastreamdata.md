This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config KinesisDataStreamData

Defines the configuration for delivering telemetry to an Amazon Kinesis Data Stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KinesisDataStreamArn" : String,
  "KinesisRoleArn" : String
}

```

### YAML

```yaml

  KinesisDataStreamArn: String
  KinesisRoleArn: String

```

## Properties

`KinesisDataStreamArn`

The ARN of the Amazon Kinesis Data Stream where telemetry data will be delivered, such as
`arn:aws:kinesis:us-east-2:123456789012:stream/my-telemetry-stream`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-.]{1,63}:kinesis:[-a-z0-9]{1,50}:[0-9]{12}:stream/[a-zA-Z0-9_.-]{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisRoleArn`

The ARN of an IAM role that AWS Ground Station assumes to write telemetry data to the Kinesis Data Stream.
This role must have permissions to perform `kinesis:PutRecord`, `kinesis:PutRecords`,
and `kinesis:DescribeStream` actions on the specified stream.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[^:\n]+:iam::[^:\n]+:role\/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create KinesisDataStreamData

The following example creates `KinesisDataStreamData` with the required IAM role and Kinesis stream configuration.

#### JSON

```json

{
  "KinesisDataStreamData": {
    "KinesisDataStreamArn": "arn:aws:kinesis:us-east-2:123456789012:stream/my-telemetry-stream",
    "KinesisRoleArn": "arn:aws:iam::123456789012:role/GroundStationKinesisRole"
  }
}
```

#### YAML

```yaml

KinesisDataStreamData:
  KinesisDataStreamArn: arn:aws:kinesis:us-east-2:123456789012:stream/my-telemetry-stream
  KinesisRoleArn: arn:aws:iam::123456789012:role/GroundStationKinesisRole

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrequencyBandwidth

S3RecordingConfig

All content copied from https://docs.aws.amazon.com/.
