This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MWAA::Environment ModuleLoggingConfiguration

Defines the type of logs to send for the Apache Airflow log type (e.g. `DagProcessingLogs`).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogGroupArn" : String,
  "Enabled" : Boolean,
  "LogLevel" : String
}

```

### YAML

```yaml

  CloudWatchLogGroupArn: String
  Enabled: Boolean
  LogLevel: String

```

## Properties

`CloudWatchLogGroupArn`

The ARN of the CloudWatch Logs log group for each type of Apache Airflow log type that you have enabled.

###### Note

`CloudWatchLogGroupArn` is available only as a return value, accessible when specified as an attribute in the
[`Fn:GetAtt`](../userguide/aws-resource-mwaa-environment.md#aws-resource-mwaa-environment-return-values)
intrinsic function. Any value you provide for `CloudWatchLogGroupArn` is discarded by Amazon MWAA.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b)(-[a-z]+)?:logs:[a-z0-9\-]+:\d{12}:log-group:\w+`

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether to enable the Apache Airflow log type (e.g. `DagProcessingLogs`) in CloudWatch Logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLevel`

Defines the Apache Airflow logs to send for the log type (e.g. `DagProcessingLogs`) to CloudWatch Logs.

_Required_: No

_Type_: String

_Allowed values_: `CRITICAL | ERROR | WARNING | INFO | DEBUG`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingConfiguration

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
