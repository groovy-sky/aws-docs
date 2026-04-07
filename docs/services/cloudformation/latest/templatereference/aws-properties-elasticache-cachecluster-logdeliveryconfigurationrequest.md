This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::CacheCluster LogDeliveryConfigurationRequest

Specifies the destination, format and type of the logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationDetails" : DestinationDetails,
  "DestinationType" : String,
  "LogFormat" : String,
  "LogType" : String
}

```

### YAML

```yaml

  DestinationDetails:
    DestinationDetails
  DestinationType: String
  LogFormat: String
  LogType: String

```

## Properties

`DestinationDetails`

Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.

_Required_: Yes

_Type_: [DestinationDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticache-cachecluster-destinationdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationType`

Specify either CloudWatch Logs or Kinesis Data Firehose as the destination type. Valid values are either
`cloudwatch-logs` or `kinesis-firehose`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogFormat`

Valid values are either `json` or `text`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogType`

Valid value is either `slow-log`, which refers to [slow-log](https://redis.io/commands/slowlog) or `engine-log`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KinesisFirehoseDestinationDetails

Tag
