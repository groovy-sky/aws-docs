This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SheetImageScalingConfiguration

Determines how the image is scaled

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScalingType" : String
}

```

### YAML

```yaml

  ScalingType: String

```

## Properties

`ScalingType`

The scaling option to use when fitting the image inside the container.

Valid values are defined as follows:

- `SCALE_TO_WIDTH`: The image takes up the entire width of the container. The image aspect ratio is preserved.

- `SCALE_TO_HEIGHT`: The image takes up the entire height of the container. The image aspect ratio is preserved.

- `SCALE_TO_CONTAINER`: The image takes up the entire width and height of the container. The image aspect ratio is not preserved.

- `SCALE_NONE`: The image is displayed in its original size and is not scaled to the container.

_Required_: No

_Type_: String

_Allowed values_: `SCALE_TO_WIDTH | SCALE_TO_HEIGHT | SCALE_TO_CONTAINER | SCALE_NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetImage

SheetImageSource

All content copied from https://docs.aws.amazon.com/.
