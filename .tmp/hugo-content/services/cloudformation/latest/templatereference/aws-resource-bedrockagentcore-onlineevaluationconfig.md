This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig

Specifies an online evaluation configuration for Amazon Bedrock AgentCore. An online evaluation configuration enables continuous monitoring and assessment of agent performance in production.

For more information, see [Monitor agent performance with Amazon Bedrock AgentCore online evaluation](../../../bedrock-agentcore/latest/devguide/online-evaluation.md).

See the **Properties** section below for descriptions of both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::OnlineEvaluationConfig",
  "Properties" : {
      "DataSourceConfig" : DataSourceConfig,
      "Description" : String,
      "EvaluationExecutionRoleArn" : String,
      "Evaluators" : [ EvaluatorReference, ... ],
      "ExecutionStatus" : String,
      "OnlineEvaluationConfigName" : String,
      "Rule" : Rule,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::OnlineEvaluationConfig
Properties:
  DataSourceConfig:
    DataSourceConfig
  Description: String
  EvaluationExecutionRoleArn: String
  Evaluators:
    - EvaluatorReference
  ExecutionStatus: String
  OnlineEvaluationConfigName: String
  Rule:
    Rule
  Tags:
    - Tag

```

## Properties

`DataSourceConfig`

The data source configuration specifying CloudWatch log groups and service names to monitor.

_Required_: Yes

_Type_: [DataSourceConfig](aws-properties-bedrockagentcore-onlineevaluationconfig-datasourceconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the online evaluation configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationExecutionRoleArn`

The Amazon Resource Name (ARN) of the IAM role used for evaluation execution.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:iam::([0-9]{12})?:role/.+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Evaluators`

The list of evaluators applied during online evaluation.

_Required_: Yes

_Type_: Array of [EvaluatorReference](aws-properties-bedrockagentcore-onlineevaluationconfig-evaluatorreference.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionStatus`

The execution status indicating whether the online evaluation is currently running.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnlineEvaluationConfigName`

The name of the online evaluation configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rule`

The evaluation rule containing sampling configuration, filters, and session settings.

_Required_: Yes

_Type_: [Rule](aws-properties-bedrockagentcore-onlineevaluationconfig-rule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the online evaluation configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-bedrockagentcore-onlineevaluationconfig-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier of the online evaluation configuration.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the online evaluation configuration was created.

`OnlineEvaluationConfigArn`

The Amazon Resource Name (ARN) of the online evaluation configuration.

`OnlineEvaluationConfigId`

The unique identifier of the online evaluation configuration.

`Status`

The current status of the online evaluation configuration.

`UpdatedAt`

The timestamp when the online evaluation configuration was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CloudWatchLogsInputConfig

All content copied from https://docs.aws.amazon.com/.
