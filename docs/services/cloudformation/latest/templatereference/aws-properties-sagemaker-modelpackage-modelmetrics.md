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

_Type_: [Bias](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-bias.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Explainability`

Metrics that help explain a model.

_Required_: No

_Type_: [Explainability](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-explainability.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelDataQuality`

Metrics that measure the quality of the input data for a model.

_Required_: No

_Type_: [ModelDataQuality](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-modeldataquality.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelQuality`

Metrics that measure the quality of a model.

_Required_: No

_Type_: [ModelQuality](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-modelquality.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModelInput

ModelPackageContainerDefinition
