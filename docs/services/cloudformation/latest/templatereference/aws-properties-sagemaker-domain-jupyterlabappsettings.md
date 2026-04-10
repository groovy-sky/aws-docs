This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain JupyterLabAppSettings

The settings for the JupyterLab application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppLifecycleManagement" : AppLifecycleManagement,
  "BuiltInLifecycleConfigArn" : String,
  "CodeRepositories" : [ CodeRepository, ... ],
  "CustomImages" : [ CustomImage, ... ],
  "DefaultResourceSpec" : ResourceSpec,
  "LifecycleConfigArns" : [ String, ... ]
}

```

### YAML

```yaml

  AppLifecycleManagement:
    AppLifecycleManagement
  BuiltInLifecycleConfigArn: String
  CodeRepositories:
    - CodeRepository
  CustomImages:
    - CustomImage
  DefaultResourceSpec:
    ResourceSpec
  LifecycleConfigArns:
    - String

```

## Properties

`AppLifecycleManagement`

Indicates whether idle shutdown is activated for JupyterLab applications.

_Required_: No

_Type_: [AppLifecycleManagement](aws-properties-sagemaker-domain-applifecyclemanagement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BuiltInLifecycleConfigArn`

The lifecycle configuration that runs before the default lifecycle configuration. It can
override changes made in the default lifecycle configuration.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:studio-lifecycle-config/.*|None)$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeRepositories`

A list of Git repositories that SageMaker automatically displays to users for cloning in the JupyterLab application.

_Required_: No

_Type_: Array of [CodeRepository](aws-properties-sagemaker-domain-coderepository.md)

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomImages`

A list of custom SageMaker images that are configured to run as a JupyterLab app.

_Required_: No

_Type_: Array of [CustomImage](aws-properties-sagemaker-domain-customimage.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultResourceSpec`

The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the
JupyterLab app.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-domain-resourcespec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleConfigArns`

The Amazon Resource Name (ARN) of the lifecycle configurations attached to the user profile or domain. To remove a lifecycle config, you must set `LifecycleConfigArns` to an empty list.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdleSettings

JupyterServerAppSettings

All content copied from https://docs.aws.amazon.com/.
