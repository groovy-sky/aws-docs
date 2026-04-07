This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OnlineEvaluationConfig Rule

The evaluation rule containing sampling configuration, filters, and session settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filters" : [ Filter, ... ],
  "SamplingConfig" : SamplingConfig,
  "SessionConfig" : SessionConfig
}

```

### YAML

```yaml

  Filters:
    - Filter
  SamplingConfig:
    SamplingConfig
  SessionConfig:
    SessionConfig

```

## Properties

`Filters`

The list of filters that determine which agent traces should be included in the evaluation based on trace properties.

_Required_: No

_Type_: Array of [Filter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-onlineevaluationconfig-filter.html)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamplingConfig`

The sampling configuration that determines what percentage of agent traces to evaluate.

_Required_: Yes

_Type_: [SamplingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-onlineevaluationconfig-samplingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionConfig`

The session configuration that defines timeout settings for detecting when agent sessions are complete and ready for evaluation.

_Required_: No

_Type_: [SessionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-onlineevaluationconfig-sessionconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OutputConfig

SamplingConfig
