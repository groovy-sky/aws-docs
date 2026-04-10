This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage DriftCheckExplainability

Represents the drift check explainability baselines that can be used when the model
monitor is set using the model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigFile" : FileSource,
  "Constraints" : MetricsSource
}

```

### YAML

```yaml

  ConfigFile:
    FileSource
  Constraints:
    MetricsSource

```

## Properties

`ConfigFile`

The explainability config file for the model.

_Required_: No

_Type_: [FileSource](aws-properties-sagemaker-modelpackage-filesource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Constraints`

The drift check explainability constraints.

_Required_: No

_Type_: [MetricsSource](aws-properties-sagemaker-modelpackage-metricssource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DriftCheckBias

DriftCheckModelDataQuality

All content copied from https://docs.aws.amazon.com/.
