This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage DriftCheckBaselines

Represents the drift check baselines that can be used when the model monitor is set
using the model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bias" : DriftCheckBias,
  "Explainability" : DriftCheckExplainability,
  "ModelDataQuality" : DriftCheckModelDataQuality,
  "ModelQuality" : DriftCheckModelQuality
}

```

### YAML

```yaml

  Bias:
    DriftCheckBias
  Explainability:
    DriftCheckExplainability
  ModelDataQuality:
    DriftCheckModelDataQuality
  ModelQuality:
    DriftCheckModelQuality

```

## Properties

`Bias`

Represents the drift check bias baselines that can be used when the model monitor is
set using the model package.

_Required_: No

_Type_: [DriftCheckBias](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-driftcheckbias.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Explainability`

Represents the drift check explainability baselines that can be used when the model
monitor is set using the model package.

_Required_: No

_Type_: [DriftCheckExplainability](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-driftcheckexplainability.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelDataQuality`

Represents the drift check model data quality baselines that can be used when the
model monitor is set using the model package.

_Required_: No

_Type_: [DriftCheckModelDataQuality](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-driftcheckmodeldataquality.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelQuality`

Represents the drift check model quality baselines that can be used when the model
monitor is set using the model package.

_Required_: No

_Type_: [DriftCheckModelQuality](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-driftcheckmodelquality.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataSource

DriftCheckBias
