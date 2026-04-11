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

_Type_: [FileSystemConfig](aws-properties-sagemaker-appimageconfig-filesystemconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KernelSpecs`

The specification of the Jupyter kernels in the image.

_Required_: Yes

_Type_: Array of [KernelSpec](aws-properties-sagemaker-appimageconfig-kernelspec.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JupyterLabAppImageConfig

KernelSpec

All content copied from https://docs.aws.amazon.com/.
