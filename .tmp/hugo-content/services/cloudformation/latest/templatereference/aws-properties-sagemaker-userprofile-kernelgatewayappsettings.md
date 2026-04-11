This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::UserProfile KernelGatewayAppSettings

The KernelGateway app settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomImages" : [ CustomImage, ... ],
  "DefaultResourceSpec" : ResourceSpec,
  "LifecycleConfigArns" : [ String, ... ]
}

```

### YAML

```yaml

  CustomImages:
    - CustomImage
  DefaultResourceSpec:
    ResourceSpec
  LifecycleConfigArns:
    - String

```

## Properties

`CustomImages`

A list of custom SageMaker AI images that are configured to run as a KernelGateway
app.

The maximum number of custom images are as follows.

- On a domain level: 200

- On a space level: 5

- On a user profile level: 5

_Required_: No

_Type_: Array of [CustomImage](aws-properties-sagemaker-userprofile-customimage.md)

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultResourceSpec`

The default instance type and the Amazon Resource Name (ARN) of the default SageMaker AI image used by the KernelGateway app.

###### Note

The Amazon SageMaker AI Studio UI does not use the default instance type value set
here. The default instance type set here is used when Apps are created using the AWS CLI or CloudFormation and the instance type parameter value is not
passed.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-userprofile-resourcespec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleConfigArns`

The Amazon Resource Name (ARN) of the Lifecycle Configurations attached to the the user
profile or domain.

###### Note

To remove a Lifecycle Config, you must set `LifecycleConfigArns` to an empty
list.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JupyterServerAppSettings

ResourceSpec

All content copied from https://docs.aws.amazon.com/.
