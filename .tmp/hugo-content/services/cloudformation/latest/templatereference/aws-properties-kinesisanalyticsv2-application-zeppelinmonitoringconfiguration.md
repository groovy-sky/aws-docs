This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ZeppelinMonitoringConfiguration

Describes configuration parameters for Amazon CloudWatch logging for a Kinesis Data
Analytics Studio notebook. For more information about CloudWatch logging, see [Monitoring](../../../managed-flink/latest/java/monitoring-overview.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogLevel" : String
}

```

### YAML

```yaml

  LogLevel: String

```

## Properties

`LogLevel`

The verbosity of the CloudWatch Logs for an application. You can set it to
`INFO`, `WARN`, `ERROR`, or
`DEBUG`.

_Required_: No

_Type_: String

_Allowed values_: `DEBUG | INFO | WARN | ERROR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ZeppelinApplicationConfiguration

AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption

All content copied from https://docs.aws.amazon.com/.
