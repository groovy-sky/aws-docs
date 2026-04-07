This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::AppImageConfig CodeEditorAppImageConfig

The configuration for the file system and kernels in a SageMaker image running as a Code Editor app.
The `FileSystemConfig` object is not supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerConfig" : ContainerConfig
}

```

### YAML

```yaml

  ContainerConfig:
    ContainerConfig

```

## Properties

`ContainerConfig`

The container configuration for the Code Editor application image.

_Required_: No

_Type_: [ContainerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-appimageconfig-containerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SageMaker::AppImageConfig

ContainerConfig
