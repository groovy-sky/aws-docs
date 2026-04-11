This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace LoggingDestination

The logging destination in an Amazon Managed Service for Prometheus workspace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogDestination,
  "Filters" : LoggingFilter
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogDestination
  Filters:
    LoggingFilter

```

## Properties

`CloudWatchLogs`

Configuration details for logging to CloudWatch Logs.

_Required_: Yes

_Type_: [CloudWatchLogDestination](aws-properties-aps-workspace-cloudwatchlogdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filters`

Filtering criteria that determine which queries are logged.

_Required_: Yes

_Type_: [LoggingFilter](aws-properties-aps-workspace-loggingfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoggingConfiguration

LoggingFilter

All content copied from https://docs.aws.amazon.com/.
