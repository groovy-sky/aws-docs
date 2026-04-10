This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Image ImageTestsConfiguration

When you create an image or container recipe with Image Builder, you can add the build or
test components that are used to create the final image. You must have at least one build
component to create a recipe, but test components are not required. If you have added tests,
they run after the image is created, to ensure that the target image is functional and can
be used reliably for launching Amazon EC2 instances.

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

Determines if tests should run after building the image. Image Builder defaults to enable tests
to run following the image build, before image distribution.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeoutMinutes`

The maximum time in minutes that tests are permitted to run.

###### Note

The timeout property is not currently active. This value is
ignored.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `1440`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageScanningConfiguration

LatestVersion

All content copied from https://docs.aws.amazon.com/.
