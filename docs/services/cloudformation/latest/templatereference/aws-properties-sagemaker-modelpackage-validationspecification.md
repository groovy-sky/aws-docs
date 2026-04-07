This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ValidationSpecification

Specifies batch transform jobs that SageMaker runs to validate your model package.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ValidationProfiles" : [ ValidationProfile, ... ],
  "ValidationRole" : String
}

```

### YAML

```yaml

  ValidationProfiles:
    - ValidationProfile
  ValidationRole: String

```

## Properties

`ValidationProfiles`

An array of `ModelPackageValidationProfile` objects, each of which
specifies a batch transform job that SageMaker runs to validate your model package.

_Required_: Yes

_Type_: Array of [ValidationProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-modelpackage-validationprofile.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationRole`

The IAM roles to be used for the validation of the model package.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ValidationProfile

AWS::SageMaker::ModelPackageGroup
