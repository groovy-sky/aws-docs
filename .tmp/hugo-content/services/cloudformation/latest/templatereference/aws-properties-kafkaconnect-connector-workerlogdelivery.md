This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector WorkerLogDelivery

Workers can send worker logs to different destination types. This configuration
specifies the details of these destinations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogsLogDelivery,
  "Firehose" : FirehoseLogDelivery,
  "S3" : S3LogDelivery
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogsLogDelivery
  Firehose:
    FirehoseLogDelivery
  S3:
    S3LogDelivery

```

## Properties

`CloudWatchLogs`

Details about delivering logs to Amazon CloudWatch Logs.

_Required_: No

_Type_: [CloudWatchLogsLogDelivery](aws-properties-kafkaconnect-connector-cloudwatchlogslogdelivery.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Firehose`

Details about delivering logs to Amazon Kinesis Data Firehose.

_Required_: No

_Type_: [FirehoseLogDelivery](aws-properties-kafkaconnect-connector-firehoselogdelivery.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3`

Details about delivering logs to Amazon S3.

_Required_: No

_Type_: [S3LogDelivery](aws-properties-kafkaconnect-connector-s3logdelivery.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkerConfiguration

AWS::KafkaConnect::CustomPlugin

All content copied from https://docs.aws.amazon.com/.
