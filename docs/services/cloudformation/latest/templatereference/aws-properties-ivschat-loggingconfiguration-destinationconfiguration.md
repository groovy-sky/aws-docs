This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVSChat::LoggingConfiguration DestinationConfiguration

The DestinationConfiguration property type describes a location where chat logs will be stored. Each member
represents the configuration of one log destination. For logging, you define only one type of
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogsDestinationConfiguration,
  "Firehose" : FirehoseDestinationConfiguration,
  "S3" : S3DestinationConfiguration
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogsDestinationConfiguration
  Firehose:
    FirehoseDestinationConfiguration
  S3:
    S3DestinationConfiguration

```

## Properties

`CloudWatchLogs`

An Amazon CloudWatch Logs destination configuration where chat activity will be
logged.

_Required_: No

_Type_: [CloudWatchLogsDestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivschat-loggingconfiguration-cloudwatchlogsdestinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Firehose`

An Amazon Kinesis Data Firehose destination configuration where chat activity will be
logged.

_Required_: No

_Type_: [FirehoseDestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivschat-loggingconfiguration-firehosedestinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

An Amazon S3 destination configuration where chat activity will be logged.

_Required_: No

_Type_: [S3DestinationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ivschat-loggingconfiguration-s3destinationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatchLogsDestinationConfiguration

FirehoseDestinationConfiguration
