This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Evaluator BedrockEvaluatorModelConfig

The Amazon Bedrock model configuration for evaluation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalModelRequestFields" : Json,
  "InferenceConfig" : InferenceConfiguration,
  "ModelId" : String
}

```

### YAML

```yaml

  AdditionalModelRequestFields: Json
  InferenceConfig:
    InferenceConfiguration
  ModelId: String

```

## Properties

`AdditionalModelRequestFields`

Additional model-specific request fields to customize model behavior beyond the standard inference configuration.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceConfig`

The inference configuration parameters that control model behavior during evaluation, including temperature, token limits, and sampling settings.

_Required_: No

_Type_: [InferenceConfiguration](aws-properties-bedrockagentcore-evaluator-inferenceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The identifier of the Amazon Bedrock model to use for evaluation. Must be a supported foundation model available in your region.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::Evaluator

CategoricalScaleDefinition

All content copied from https://docs.aws.amazon.com/.
