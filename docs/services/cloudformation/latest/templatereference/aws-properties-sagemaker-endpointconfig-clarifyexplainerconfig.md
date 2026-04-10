This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig ClarifyExplainerConfig

The configuration parameters for the SageMaker Clarify explainer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableExplanations" : String,
  "InferenceConfig" : ClarifyInferenceConfig,
  "ShapConfig" : ClarifyShapConfig
}

```

### YAML

```yaml

  EnableExplanations: String
  InferenceConfig:
    ClarifyInferenceConfig
  ShapConfig:
    ClarifyShapConfig

```

## Properties

`EnableExplanations`

A JMESPath boolean expression used to filter which records to explain. Explanations
are activated by default. See [`EnableExplanations`](../../../sagemaker/latest/dg/clarify-online-explainability-create-endpoint.md#clarify-online-explainability-create-endpoint-enable) for additional information.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InferenceConfig`

The inference configuration parameter for the model container.

_Required_: No

_Type_: [ClarifyInferenceConfig](aws-properties-sagemaker-endpointconfig-clarifyinferenceconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShapConfig`

The configuration for SHAP analysis.

_Required_: Yes

_Type_: [ClarifyShapConfig](aws-properties-sagemaker-endpointconfig-clarifyshapconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptureOption

ClarifyInferenceConfig

All content copied from https://docs.aws.amazon.com/.
