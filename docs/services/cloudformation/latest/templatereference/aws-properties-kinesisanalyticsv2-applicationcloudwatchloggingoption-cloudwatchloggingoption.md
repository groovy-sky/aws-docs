This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption CloudWatchLoggingOption

Provides a description of Amazon CloudWatch logging options, including the log stream
Amazon Resource Name (ARN).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogStreamARN" : String
}

```

### YAML

```yaml

  LogStreamARN: String

```

## Properties

`LogStreamARN`

The ARN of the CloudWatch log to receive application messages.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CloudWatchLoggingOption](../../../managed-flink/latest/apiv2/api-cloudwatchloggingoption.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption

AWS::KinesisAnalyticsV2::ApplicationOutput

All content copied from https://docs.aws.amazon.com/.
