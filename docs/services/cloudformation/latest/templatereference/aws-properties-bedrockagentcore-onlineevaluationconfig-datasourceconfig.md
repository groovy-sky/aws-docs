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

_Type_: [CloudWatchLogsInputConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-onlineevaluationconfig-cloudwatchlogsinputconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatchOutputConfig

EvaluatorReference
