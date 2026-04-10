This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate DataSources

Describes the data sources for the experiment report.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchDashboards" : [ CloudWatchDashboard, ... ]
}

```

### YAML

```yaml

  CloudWatchDashboards:
    - CloudWatchDashboard

```

## Properties

`CloudWatchDashboards`

The CloudWatch dashboards to include as data sources in the experiment report.

_Required_: No

_Type_: Array of [CloudWatchDashboard](aws-properties-fis-experimenttemplate-cloudwatchdashboard.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogsConfiguration

ExperimentReportS3Configuration

All content copied from https://docs.aws.amazon.com/.
