This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space JupyterServerAppSettings

The JupyterServer app settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultResourceSpec" : ResourceSpec,
  "LifecycleConfigArns" : [ String, ... ]
}

```

### YAML

```yaml

  DefaultResourceSpec:
    ResourceSpec
  LifecycleConfigArns:
    - String

```

## Properties

`DefaultResourceSpec`

The default instance type and the Amazon Resource Name (ARN) of the default SageMaker AI image used by the JupyterServer app. If you use the
`LifecycleConfigArns` parameter, then this parameter is also required.

_Required_: No

_Type_: [ResourceSpec](aws-properties-sagemaker-space-resourcespec.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleConfigArns`

The Amazon Resource Name (ARN) of the Lifecycle Configurations attached to the
JupyterServerApp. If you use this parameter, the `DefaultResourceSpec` parameter is
also required.

###### Note

To remove a Lifecycle Config, you must set `LifecycleConfigArns` to an empty
list.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FSxLustreFileSystem

KernelGatewayAppSettings

All content copied from https://docs.aws.amazon.com/.
