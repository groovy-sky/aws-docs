This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelBiasJobDefinition ModelBiasBaselineConfig

The configuration for a baseline model bias job.

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

The name of the baseline model bias job.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConstraintsResource`

The constraints resource for a monitoring job.

_Required_: No

_Type_: [ConstraintsResource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelbiasjobdefinition-constraintsresource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModelBiasAppSpecification

ModelBiasJobInput
