This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application MonitoringConfiguration

Describes configuration parameters for Amazon CloudWatch logging for a Java-based
Kinesis Data Analytics application. For more information about CloudWatch logging, see
[Monitoring](../../../managed-flink/latest/java/monitoring-overview.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigurationType" : String,
  "LogLevel" : String,
  "MetricsLevel" : String
}

```

### YAML

```yaml

  ConfigurationType: String
  LogLevel: String
  MetricsLevel: String

```

## Properties

`ConfigurationType`

Describes whether to use the default CloudWatch logging configuration for an application.
You must set this property to `CUSTOM` in order to set the `LogLevel` or
`MetricsLevel` parameters.

_Required_: Yes

_Type_: String

_Allowed values_: `DEFAULT | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogLevel`

Describes the verbosity of the CloudWatch Logs for an application.

_Required_: No

_Type_: String

_Allowed values_: `DEBUG | INFO | WARN | ERROR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsLevel`

Describes the granularity of the CloudWatch Logs for an application. The `Parallelism`
level is not recommended for applications with a Parallelism over 64 due to excessive costs.

_Required_: No

_Type_: String

_Allowed values_: `APPLICATION | OPERATOR | PARALLELISM | TASK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [MonitoringConfiguration](../../../managed-flink/latest/apiv2/api-monitoringconfiguration.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MavenReference

ParallelismConfiguration

All content copied from https://docs.aws.amazon.com/.
