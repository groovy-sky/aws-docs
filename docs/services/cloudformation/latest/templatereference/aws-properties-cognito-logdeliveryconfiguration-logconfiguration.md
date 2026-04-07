This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::LogDeliveryConfiguration LogConfiguration

The configuration of user event logs to an external AWS service like
Amazon Data Firehose, Amazon S3, or Amazon CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsConfiguration" : CloudWatchLogsConfiguration,
  "EventSource" : String,
  "FirehoseConfiguration" : FirehoseConfiguration,
  "LogLevel" : String,
  "S3Configuration" : S3Configuration
}

```

### YAML

```yaml

  CloudWatchLogsConfiguration:
    CloudWatchLogsConfiguration
  EventSource: String
  FirehoseConfiguration:
    FirehoseConfiguration
  LogLevel: String
  S3Configuration:
    S3Configuration

```

## Properties

`CloudWatchLogsConfiguration`

Configuration for the CloudWatch log group destination of user pool detailed activity
logging, or of user activity log export with advanced security features.

_Required_: No

_Type_: [CloudWatchLogsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-logdeliveryconfiguration-cloudwatchlogsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventSource`

The source of events that your user pool sends for logging. To send error-level logs
about user notification activity, set to `userNotification`. To send
info-level logs about threat-protection user activity in user pools with the Plus
feature plan, set to `userAuthEvents`.

_Required_: No

_Type_: String

_Allowed values_: `userNotification | userAuthEvents`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirehoseConfiguration`

Configuration for the Amazon Data Firehose stream destination of user activity log export with
threat protection.

_Required_: No

_Type_: [FirehoseConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-logdeliveryconfiguration-firehoseconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLevel`

The `errorlevel` selection of logs that a user pool sends for detailed
activity logging. To send `userNotification` activity with [information\
about message delivery](https://docs.aws.amazon.com/cognito/latest/developerguide/exporting-quotas-and-usage.html), choose `ERROR` with
`CloudWatchLogsConfiguration`. To send `userAuthEvents`
activity with user logs from threat protection with the Plus feature plan, choose
`INFO` with one of `CloudWatchLogsConfiguration`,
`FirehoseConfiguration`, or `S3Configuration`.

_Required_: No

_Type_: String

_Allowed values_: `ERROR | INFO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

Configuration for the Amazon S3 bucket destination of user activity log export with threat
protection.

_Required_: No

_Type_: [S3Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-logdeliveryconfiguration-s3configuration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FirehoseConfiguration

S3Configuration
