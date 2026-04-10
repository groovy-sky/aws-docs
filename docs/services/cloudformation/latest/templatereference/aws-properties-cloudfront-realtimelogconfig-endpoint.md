This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::RealtimeLogConfig EndPoint

Contains information about the Amazon Kinesis data stream where you are sending real-time
log data for this real-time log configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KinesisStreamConfig" : KinesisStreamConfig,
  "StreamType" : String
}

```

### YAML

```yaml

  KinesisStreamConfig:
    KinesisStreamConfig
  StreamType: String

```

## Properties

`KinesisStreamConfig`

Contains information about the Amazon Kinesis data stream where you are sending real-time
log data in a real-time log configuration.

_Required_: Yes

_Type_: [KinesisStreamConfig](aws-properties-cloudfront-realtimelogconfig-kinesisstreamconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamType`

The type of data stream where you are sending real-time log data. The only valid value
is `Kinesis`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::RealtimeLogConfig

KinesisStreamConfig

All content copied from https://docs.aws.amazon.com/.
