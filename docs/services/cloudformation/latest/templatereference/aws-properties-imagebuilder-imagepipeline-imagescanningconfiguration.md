This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImagePipeline ImageScanningConfiguration

Contains settings for Image Builder image resource and container image scans.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EcrConfiguration" : EcrConfiguration,
  "ImageScanningEnabled" : Boolean
}

```

### YAML

```yaml

  EcrConfiguration:
    EcrConfiguration
  ImageScanningEnabled: Boolean

```

## Properties

`EcrConfiguration`

Contains Amazon ECR settings for vulnerability scans.

_Required_: No

_Type_: [EcrConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagepipeline-ecrconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageScanningEnabled`

A setting that indicates whether Image Builder keeps a snapshot of the vulnerability scans that
Amazon Inspector runs against the build instance when you create a new image.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EcrConfiguration

ImageTestsConfiguration
