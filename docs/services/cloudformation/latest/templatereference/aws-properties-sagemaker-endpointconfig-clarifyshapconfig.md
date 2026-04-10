This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig ClarifyShapConfig

The configuration for SHAP analysis using SageMaker Clarify Explainer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NumberOfSamples" : Integer,
  "Seed" : Integer,
  "ShapBaselineConfig" : ClarifyShapBaselineConfig,
  "TextConfig" : ClarifyTextConfig,
  "UseLogit" : Boolean
}

```

### YAML

```yaml

  NumberOfSamples: Integer
  Seed: Integer
  ShapBaselineConfig:
    ClarifyShapBaselineConfig
  TextConfig:
    ClarifyTextConfig
  UseLogit: Boolean

```

## Properties

`NumberOfSamples`

The number of samples to be used for analysis by the Kernal SHAP algorithm.

###### Note

The number of samples determines the size of the synthetic dataset, which has an
impact on latency of explainability requests. For more information, see the
**Synthetic data** of [Configure and create an endpoint](../../../sagemaker/latest/dg/clarify-online-explainability-create-endpoint.md).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Seed`

The starting value used to initialize the random number generator in the explainer.
Provide a value for this parameter to obtain a deterministic SHAP result.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShapBaselineConfig`

The configuration for the SHAP baseline of the Kernal SHAP algorithm.

_Required_: Yes

_Type_: [ClarifyShapBaselineConfig](aws-properties-sagemaker-endpointconfig-clarifyshapbaselineconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TextConfig`

A parameter that indicates if text features are treated as text and explanations are
provided for individual units of text. Required for natural language processing (NLP)
explainability only.

_Required_: No

_Type_: [ClarifyTextConfig](aws-properties-sagemaker-endpointconfig-clarifytextconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UseLogit`

A Boolean toggle to indicate if you want to use the logit function (true) or log-odds
units (false) for model predictions. Defaults to false.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClarifyShapBaselineConfig

ClarifyTextConfig

All content copied from https://docs.aws.amazon.com/.
