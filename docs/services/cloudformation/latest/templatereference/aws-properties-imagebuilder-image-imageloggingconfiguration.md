This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Image ImageLoggingConfiguration

The logging configuration that's defined for the image. Image Builder uses the defined settings
to direct execution log output during image creation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogGroupName" : String
}

```

### YAML

```yaml

  LogGroupName: String

```

## Properties

`LogGroupName`

The log group name that Image Builder uses for image creation. If not specified, the log group
name defaults to `/aws/imagebuilder/image-name`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_/\.]{1,512}$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EcrConfiguration

ImagePipelineExecutionSettings
