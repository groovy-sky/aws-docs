This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain RSessionAppSettings

A collection of settings that apply to an `RSessionGateway` app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomImages" : [ CustomImage, ... ],
  "DefaultResourceSpec" : ResourceSpec
}

```

### YAML

```yaml

  CustomImages:
    - CustomImage
  DefaultResourceSpec:
    ResourceSpec

```

## Properties

`CustomImages`

A list of custom SageMaker AI images that are configured to run as a RSession
app.

_Required_: No

_Type_: Array of [CustomImage](aws-properties-sagemaker-domain-customimage.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultResourceSpec`

Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs
on.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-domain-resourcespec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceSpec

RStudioServerProAppSettings

All content copied from https://docs.aws.amazon.com/.
