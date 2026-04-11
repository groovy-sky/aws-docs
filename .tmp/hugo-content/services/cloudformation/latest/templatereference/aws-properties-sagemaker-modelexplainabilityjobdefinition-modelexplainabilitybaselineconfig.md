This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelExplainabilityJobDefinition ModelExplainabilityBaselineConfig

The configuration for a baseline model explainability job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseliningJobName" : String,
  "ConstraintsResource" : ConstraintsResource
}

```

### YAML

```yaml

  BaseliningJobName: String
  ConstraintsResource:
    ConstraintsResource

```

## Properties

`BaseliningJobName`

The name of the baseline model explainability job.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConstraintsResource`

The constraints resource for a model explainability job.

_Required_: No

_Type_: [ConstraintsResource](aws-properties-sagemaker-modelexplainabilityjobdefinition-constraintsresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelExplainabilityAppSpecification

ModelExplainabilityJobInput

All content copied from https://docs.aws.amazon.com/.
