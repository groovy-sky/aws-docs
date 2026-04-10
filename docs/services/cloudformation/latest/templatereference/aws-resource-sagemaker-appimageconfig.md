This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::AppImageConfig

Creates a configuration for running a SageMaker AI image as a KernelGateway app. The
configuration specifies the Amazon Elastic File System storage volume on the image, and a list of the
kernels in the image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::AppImageConfig",
  "Properties" : {
      "AppImageConfigName" : String,
      "CodeEditorAppImageConfig" : CodeEditorAppImageConfig,
      "JupyterLabAppImageConfig" : JupyterLabAppImageConfig,
      "KernelGatewayImageConfig" : KernelGatewayImageConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::AppImageConfig
Properties:
  AppImageConfigName: String
  CodeEditorAppImageConfig:
    CodeEditorAppImageConfig
  JupyterLabAppImageConfig:
    JupyterLabAppImageConfig
  KernelGatewayImageConfig:
    KernelGatewayImageConfig
  Tags:
    - Tag

```

## Properties

`AppImageConfigName`

The name of the AppImageConfig. Must be unique to your account.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CodeEditorAppImageConfig`

The configuration for the file system and the runtime,
such as the environment variables and entry point.

_Required_: No

_Type_: [CodeEditorAppImageConfig](aws-properties-sagemaker-appimageconfig-codeeditorappimageconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JupyterLabAppImageConfig`

The configuration for the file system and the runtime, such as the environment variables and entry point.

_Required_: No

_Type_: [JupyterLabAppImageConfig](aws-properties-sagemaker-appimageconfig-jupyterlabappimageconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KernelGatewayImageConfig`

The configuration for the file system and kernels in the SageMaker AI image.

_Required_: No

_Type_: [KernelGatewayImageConfig](aws-properties-sagemaker-appimageconfig-kernelgatewayimageconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-appimageconfig-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the AppImageConfig.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AppImageConfigArn`

The Amazon Resource Name (ARN) of the AppImageConfig, such as
`arn:aws:sagemaker:us-west-2:account-id:app-image-config/my-app-image-config-name`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CodeEditorAppImageConfig

All content copied from https://docs.aws.amazon.com/.
