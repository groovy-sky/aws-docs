This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ModelMetrics

Contains metrics captured from a model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bias" : Bias,
  "Explainability" : Explainability,
  "ModelDataQuality" : ModelDataQuality,
  "ModelQuality" : ModelQuality
}

```

### YAML

```yaml

  Bias:
    Bias
  Explainability:
    Explainability
  ModelDataQuality:
    ModelDataQuality
  ModelQuality:
    ModelQuality

```

## Properties

`Bias`

Metrics that measure bias in a model.

_Required_: No

_Type_: [Bias](aws-properties-sagemaker-modelpackage-bias.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Explainability`

Metrics that help explain a model.

_Required_: No

_Type_: [Explainability](aws-properties-sagemaker-modelpackage-explainability.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelDataQuality`

Metrics that measure the quality of the input data for a model.

_Required_: No

_Type_: [ModelDataQuality](aws-properties-sagemaker-modelpackage-modeldataquality.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelQuality`

Metrics that measure the quality of a model.

_Required_: No

_Type_: [ModelQuality](aws-properties-sagemaker-modelpackage-modelquality.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelInput

ModelPackageContainerDefinition

All content copied from https://docs.aws.amazon.com/.
