This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateLogConfiguration

Specifies the configuration for experiment logging.

For more information, see [Experiment logging](../../../fis/latest/userguide/monitoring-logging.md)
in the _AWS Fault Injection Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsConfiguration" : CloudWatchLogsConfiguration,
  "LogSchemaVersion" : Integer,
  "S3Configuration" : S3Configuration
}

```

### YAML

```yaml

  CloudWatchLogsConfiguration:
    CloudWatchLogsConfiguration
  LogSchemaVersion: Integer
  S3Configuration:
    S3Configuration

```

## Properties

`CloudWatchLogsConfiguration`

The configuration for experiment logging to CloudWatch Logs.

_Required_: No

_Type_: [CloudWatchLogsConfiguration](aws-properties-fis-experimenttemplate-cloudwatchlogsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogSchemaVersion`

The schema version.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Configuration`

The configuration for experiment logging to Amazon S3.

_Required_: No

_Type_: [S3Configuration](aws-properties-fis-experimenttemplate-s3configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExperimentTemplateExperimentReportConfiguration

ExperimentTemplateStopCondition

All content copied from https://docs.aws.amazon.com/.
