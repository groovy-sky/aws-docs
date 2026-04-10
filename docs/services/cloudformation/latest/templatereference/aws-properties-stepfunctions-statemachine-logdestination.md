This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachine LogDestination

Defines a destination for `LoggingConfiguration`.

###### Note

For more information on logging with `EXPRESS` workflows, see [Logging Express\
Workflows Using CloudWatch Logs](../../../step-functions/latest/dg/cw-logs.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogsLogGroup" : CloudWatchLogsLogGroup
}

```

### YAML

```yaml

  CloudWatchLogsLogGroup:
    CloudWatchLogsLogGroup

```

## Properties

`CloudWatchLogsLogGroup`

An object describing a CloudWatch log group. For more information, see [AWS::Logs::LogGroup](../userguide/aws-resource-logs-loggroup.md) in the CloudFormation User Guide.

_Required_: No

_Type_: [CloudWatchLogsLogGroup](aws-properties-stepfunctions-statemachine-cloudwatchlogsloggroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

LoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
