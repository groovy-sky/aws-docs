This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig DataSourceConfig

The data source configuration specifying CloudWatch log groups and service names to monitor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogs" : CloudWatchLogsInputConfig
}

```

### YAML

```yaml

  CloudWatchLogs:
    CloudWatchLogsInputConfig

```

## Properties

`CloudWatchLogs`

The CloudWatch logs configuration for reading agent traces from log groups.

_Required_: Yes

_Type_: [CloudWatchLogsInputConfig](aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchOutputConfig

EvaluatorReference

All content copied from https://docs.aws.amazon.com/.
