This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeLogConfiguration

Represents the configuration settings for the logs to which this pipe should report events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudwatchLogsLogDestination" : CloudwatchLogsLogDestination,
  "FirehoseLogDestination" : FirehoseLogDestination,
  "IncludeExecutionData" : [ String, ... ],
  "Level" : String,
  "S3LogDestination" : S3LogDestination
}

```

### YAML

```yaml

  CloudwatchLogsLogDestination:
    CloudwatchLogsLogDestination
  FirehoseLogDestination:
    FirehoseLogDestination
  IncludeExecutionData:
    - String
  Level: String
  S3LogDestination:
    S3LogDestination

```

## Properties

`CloudwatchLogsLogDestination`

The logging configuration settings for the pipe.

_Required_: No

_Type_: [CloudwatchLogsLogDestination](aws-properties-pipes-pipe-cloudwatchlogslogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirehoseLogDestination`

The Amazon Data Firehose logging configuration settings for the pipe.

_Required_: No

_Type_: [FirehoseLogDestination](aws-properties-pipes-pipe-firehoselogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeExecutionData`

Whether the execution data (specifically, the `payload`,
`awsRequest`, and `awsResponse` fields) is included in the log
messages for this pipe.

This applies to all log destinations for the pipe.

For more information, see [Including execution data in logs](../../../eventbridge/latest/userguide/eb-pipes-logs.md#eb-pipes-logs-execution-data) in the _Amazon EventBridge User_
_Guide_.

_Allowed values:_ `ALL`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Level`

The level of logging detail to include. This applies to all log destinations for the pipe.

_Required_: No

_Type_: String

_Allowed values_: `OFF | ERROR | INFO | TRACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3LogDestination`

The Amazon S3 logging configuration settings for the pipe.

_Required_: No

_Type_: [S3LogDestination](aws-properties-pipes-pipe-s3logdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeEnrichmentParameters

PipeSourceActiveMQBrokerParameters

All content copied from https://docs.aws.amazon.com/.
