This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ReplicationGroup DestinationDetails

Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsDetails" : CloudWatchLogsDestinationDetails,
  "KinesisFirehoseDetails" : KinesisFirehoseDestinationDetails
}

```

### YAML

```yaml

  CloudWatchLogsDetails:
    CloudWatchLogsDestinationDetails
  KinesisFirehoseDetails:
    KinesisFirehoseDestinationDetails

```

## Properties

`CloudWatchLogsDetails`

The configuration details of the CloudWatch Logs destination. Note that this field is marked as required but
only if CloudWatch Logs was chosen as the destination.

_Required_: No

_Type_: [CloudWatchLogsDestinationDetails](aws-properties-elasticache-replicationgroup-cloudwatchlogsdestinationdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KinesisFirehoseDetails`

The configuration details of the Kinesis Data Firehose destination. Note that this field is marked as required
but only if Kinesis Data Firehose was chosen as the destination.

_Required_: No

_Type_: [KinesisFirehoseDestinationDetails](aws-properties-elasticache-replicationgroup-kinesisfirehosedestinationdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogsDestinationDetails

KinesisFirehoseDestinationDetails

All content copied from https://docs.aws.amazon.com/.
