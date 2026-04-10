This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage DriftCheckModelQuality

Represents the drift check model quality baselines that can be used when the model
monitor is set using the model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Constraints" : MetricsSource,
  "Statistics" : MetricsSource
}

```

### YAML

```yaml

  Constraints:
    MetricsSource
  Statistics:
    MetricsSource

```

## Properties

`Constraints`

The drift check model quality constraints.

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Statistics`

The drift check model quality statistics.

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DriftCheckModelDataQuality

Explainability

All content copied from https://docs.aws.amazon.com/.
