This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage SourceAlgorithmSpecification

A list of algorithms that were used to create a model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceAlgorithms" : [ SourceAlgorithm, ... ]
}

```

### YAML

```yaml

  SourceAlgorithms:
    - SourceAlgorithm

```

## Properties

`SourceAlgorithms`

A list of the algorithms that were used to create a model package.

_Required_: Yes

_Type_: Array of [SourceAlgorithm](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-sourcealgorithm.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SourceAlgorithm

Tag
