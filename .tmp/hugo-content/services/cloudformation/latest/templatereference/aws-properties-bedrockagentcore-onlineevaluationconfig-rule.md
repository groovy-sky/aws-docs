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

_Type_: Array of [Filter](aws-properties-bedrockagentcore-onlineevaluationconfig-filter.md)

_Minimum_: `0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamplingConfig`

The sampling configuration that determines what percentage of agent traces to evaluate.

_Required_: Yes

_Type_: [SamplingConfig](aws-properties-bedrockagentcore-onlineevaluationconfig-samplingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionConfig`

The session configuration that defines timeout settings for detecting when agent sessions are complete and ready for evaluation.

_Required_: No

_Type_: [SessionConfig](aws-properties-bedrockagentcore-onlineevaluationconfig-sessionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputConfig

SamplingConfig

All content copied from https://docs.aws.amazon.com/.
