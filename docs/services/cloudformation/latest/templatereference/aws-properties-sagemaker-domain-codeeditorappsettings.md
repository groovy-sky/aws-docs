This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain CodeEditorAppSettings

The Code Editor application settings.

For more information about Code Editor, see [Get started with Code \
Editor in Amazon SageMaker](../../../sagemaker/latest/dg/code-editor.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppLifecycleManagement" : AppLifecycleManagement,
  "BuiltInLifecycleConfigArn" : String,
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
  CustomImages:
    - CustomImage
  DefaultResourceSpec:
    ResourceSpec
  LifecycleConfigArns:
    - String

```

## Properties

`AppLifecycleManagement`

Settings that are used to configure and manage the lifecycle of CodeEditor
applications.

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

`CustomImages`

A list of custom SageMaker images that are configured to run as a Code Editor app.

_Required_: No

_Type_: Array of [CustomImage](aws-properties-sagemaker-domain-customimage.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultResourceSpec`

The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the Code
Editor app.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-domain-resourcespec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleConfigArns`

The Amazon Resource Name (ARN) of the Code Editor application
lifecycle configuration.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppLifecycleManagement

CodeRepository

All content copied from https://docs.aws.amazon.com/.
