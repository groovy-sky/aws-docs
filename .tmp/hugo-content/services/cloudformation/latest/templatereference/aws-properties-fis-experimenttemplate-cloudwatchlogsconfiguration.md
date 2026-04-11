This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate CloudWatchLogsConfiguration

Specifies the configuration for experiment logging to CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupArn" : String
}

```

### YAML

```yaml

  LogGroupArn: String

```

## Properties

`LogGroupArn`

The Amazon Resource Name (ARN) of the destination Amazon CloudWatch Logs log group.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchDashboard

DataSources

All content copied from https://docs.aws.amazon.com/.
