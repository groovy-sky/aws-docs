This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SheetImage

An image that is located on a sheet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ ImageCustomAction, ... ],
  "ImageContentAltText" : String,
  "Interactions" : ImageInteractionOptions,
  "Scaling" : SheetImageScalingConfiguration,
  "SheetImageId" : String,
  "Source" : SheetImageSource,
  "Tooltip" : SheetImageTooltipConfiguration
}

```

### YAML

```yaml

  Actions:
    - ImageCustomAction
  ImageContentAltText: String
  Interactions:
    ImageInteractionOptions
  Scaling:
    SheetImageScalingConfiguration
  SheetImageId: String
  Source:
    SheetImageSource
  Tooltip:
    SheetImageTooltipConfiguration

```

## Properties

`Actions`

A list of custom actions that are configured for an image.

_Required_: No

_Type_: Array of [ImageCustomAction](aws-properties-quicksight-template-imagecustomaction.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageContentAltText`

The alt text for the image.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general image interactions setup for an image.

_Required_: No

_Type_: [ImageInteractionOptions](aws-properties-quicksight-template-imageinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scaling`

Determines how the image is scaled.

_Required_: No

_Type_: [SheetImageScalingConfiguration](aws-properties-quicksight-template-sheetimagescalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetImageId`

The ID of the sheet image.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The source of the image.

_Required_: Yes

_Type_: [SheetImageSource](aws-properties-quicksight-template-sheetimagesource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tooltip`

The tooltip to be shown when hovering over the image.

_Required_: No

_Type_: [SheetImageTooltipConfiguration](aws-properties-quicksight-template-sheetimagetooltipconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetElementRenderingRule

SheetImageScalingConfiguration

All content copied from https://docs.aws.amazon.com/.
