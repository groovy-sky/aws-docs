This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage Bias

Contains bias metrics for a model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PostTrainingReport" : MetricsSource,
  "PreTrainingReport" : MetricsSource,
  "Report" : MetricsSource
}

```

### YAML

```yaml

  PostTrainingReport:
    MetricsSource
  PreTrainingReport:
    MetricsSource
  Report:
    MetricsSource

```

## Properties

`PostTrainingReport`

The post-training bias report for a model.

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreTrainingReport`

The pre-training bias report for a model.

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Report`

The bias report for a model

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdditionalInferenceSpecificationDefinition

DataSource

All content copied from https://docs.aws.amazon.com/.
