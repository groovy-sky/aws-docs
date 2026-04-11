This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CustomContentConfiguration

The configuration of a `CustomContentVisual`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentType" : String,
  "ContentUrl" : String,
  "ImageScaling" : String,
  "Interactions" : VisualInteractionOptions
}

```

### YAML

```yaml

  ContentType: String
  ContentUrl: String
  ImageScaling: String
  Interactions:
    VisualInteractionOptions

```

## Properties

`ContentType`

The content type of the custom content visual. You can use this to have the visual render as an image.

_Required_: No

_Type_: String

_Allowed values_: `IMAGE | OTHER_EMBEDDED_CONTENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentUrl`

The input URL that links to the custom content that you want in the custom visual.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageScaling`

The sizing options for the size of the custom content visual. This structure is required when the `ContentType` of the visual is `'IMAGE'`.

_Required_: No

_Type_: String

_Allowed values_: `FIT_TO_HEIGHT | FIT_TO_WIDTH | DO_NOT_SCALE | SCALE_TO_VISUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interactions`

The general visual interactions setup for a visual.

_Required_: No

_Type_: [VisualInteractionOptions](aws-properties-quicksight-analysis-visualinteractionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomColor

CustomContentVisual

All content copied from https://docs.aws.amazon.com/.
