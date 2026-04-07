This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::AppImageConfig KernelGatewayImageConfig

The configuration for the file system and kernels in a SageMaker AI image running as a
KernelGateway app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileSystemConfig" : FileSystemConfig,
  "KernelSpecs" : [ KernelSpec, ... ]
}

```

### YAML

```yaml

  FileSystemConfig:
    FileSystemConfig
  KernelSpecs:
    - KernelSpec

```

## Properties

`FileSystemConfig`

The Amazon Elastic File System storage configuration for a SageMaker AI image.

_Required_: No

_Type_: [FileSystemConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-appimageconfig-filesystemconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KernelSpecs`

The specification of the Jupyter kernels in the image.

_Required_: Yes

_Type_: Array of [KernelSpec](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-appimageconfig-kernelspec.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JupyterLabAppImageConfig

KernelSpec
