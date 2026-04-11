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

_Type_: [DriftCheckBias](aws-properties-sagemaker-modelpackage-driftcheckbias.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Explainability`

Represents the drift check explainability baselines that can be used when the model
monitor is set using the model package.

_Required_: No

_Type_: [DriftCheckExplainability](aws-properties-sagemaker-modelpackage-driftcheckexplainability.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelDataQuality`

Represents the drift check model data quality baselines that can be used when the
model monitor is set using the model package.

_Required_: No

_Type_: [DriftCheckModelDataQuality](aws-properties-sagemaker-modelpackage-driftcheckmodeldataquality.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelQuality`

Represents the drift check model quality baselines that can be used when the model
monitor is set using the model package.

_Required_: No

_Type_: [DriftCheckModelQuality](aws-properties-sagemaker-modelpackage-driftcheckmodelquality.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSource

DriftCheckBias

All content copied from https://docs.aws.amazon.com/.
