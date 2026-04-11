This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage DriftCheckBias

Represents the drift check bias baselines that can be used when the model monitor is
set using the model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigFile" : FileSource,
  "PostTrainingConstraints" : MetricsSource,
  "PreTrainingConstraints" : MetricsSource
}

```

### YAML

```yaml

  ConfigFile:
    FileSource
  PostTrainingConstraints:
    MetricsSource
  PreTrainingConstraints:
    MetricsSource

```

## Properties

`ConfigFile`

The bias config file for a model.

_Required_: No

_Type_: [FileSource](aws-properties-sagemaker-modelpackage-filesource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PostTrainingConstraints`

The post-training constraints.

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreTrainingConstraints`

The pre-training constraints.

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DriftCheckBaselines

DriftCheckExplainability

All content copied from https://docs.aws.amazon.com/.
