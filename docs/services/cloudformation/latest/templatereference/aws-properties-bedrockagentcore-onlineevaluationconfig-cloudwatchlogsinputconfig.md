This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig CloudWatchLogsInputConfig

The CloudWatch logs configuration for reading agent traces from log groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupNames" : [ String, ... ],
  "ServiceNames" : [ String, ... ]
}

```

### YAML

```yaml

  LogGroupNames:
    - String
  ServiceNames:
    - String

```

## Properties

`LogGroupNames`

The list of CloudWatch log group names to monitor for agent traces.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `512 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNames`

The list of service names to filter traces within the specified log groups. Used to identify relevant agent sessions.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::OnlineEvaluationConfig

CloudWatchOutputConfig

All content copied from https://docs.aws.amazon.com/.
