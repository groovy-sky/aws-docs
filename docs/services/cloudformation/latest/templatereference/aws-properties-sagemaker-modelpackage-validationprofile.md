This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ValidationProfile

Contains data, such as the inputs and targeted instance types that are used in the
process of validating the model package.

The data provided in the validation profile is made available to your buyers on
AWS Marketplace.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProfileName" : String,
  "TransformJobDefinition" : TransformJobDefinition
}

```

### YAML

```yaml

  ProfileName: String
  TransformJobDefinition:
    TransformJobDefinition

```

## Properties

`ProfileName`

The name of the profile for the model package.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransformJobDefinition`

The `TransformJobDefinition` object that describes the transform job used
for the validation of the model package.

_Required_: Yes

_Type_: [TransformJobDefinition](aws-properties-sagemaker-modelpackage-transformjobdefinition.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformResources

ValidationSpecification

All content copied from https://docs.aws.amazon.com/.
