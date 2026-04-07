This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImagePipeline ImageTestsConfiguration

When you create an image or container recipe with Image Builder, you can add the build or
test components that your image pipeline uses to create the final image. You must
have at least one build component to create a recipe, but test components are not required.
Your pipeline runs tests after it builds the image, to ensure that the target image is
functional and can be used reliably for launching Amazon EC2 instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageTestsEnabled" : Boolean,
  "TimeoutMinutes" : Integer
}

```

### YAML

```yaml

  ImageTestsEnabled: Boolean
  TimeoutMinutes: Integer

```

## Properties

`ImageTestsEnabled`

Defines if tests should be executed when building this image. For example,
`true` or `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutMinutes`

The maximum time in minutes that tests are permitted to run.

###### Note

The timeout property is not currently active. This value is
ignored.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ImageScanningConfiguration

PipelineLoggingConfiguration
