This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator

Specifies an evaluator for Amazon Bedrock AgentCore. An evaluator assesses agent quality using LLM-as-a-Judge configurations to measure and improve agent performance.

For more information, see [Evaluate agent quality with Amazon Bedrock AgentCore](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/evaluators.html).

See the **Properties** section below for descriptions of both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::Evaluator",
  "Properties" : {
      "Description" : String,
      "EvaluatorConfig" : EvaluatorConfig,
      "EvaluatorName" : String,
      "Level" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::Evaluator
Properties:
  Description: String
  EvaluatorConfig:
    EvaluatorConfig
  EvaluatorName: String
  Level: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the evaluator.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluatorConfig`

The configuration of the evaluator, including LLM-as-a-Judge settings for custom evaluators.

_Required_: Yes

_Type_: [EvaluatorConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-evaluator-evaluatorconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluatorName`

The name of the evaluator.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Level`

The evaluation level ( `TOOL_CALL`, `TRACE`, or
`SESSION`) that determines the scope of evaluation.

_Required_: Yes

_Type_: String

_Allowed values_: `TOOL_CALL | TRACE | SESSION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the evaluator.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-evaluator-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the evaluator. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:evaluator/EXAMPLE12345`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the evaluator was created.

`EvaluatorArn`

The Amazon Resource Name (ARN) of the evaluator.

`EvaluatorId`

The unique identifier of the evaluator.

`Status`

The current status of the evaluator.

`UpdatedAt`

The timestamp when the evaluator was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

BedrockEvaluatorModelConfig
