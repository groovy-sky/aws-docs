This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel WebvttDestinationSettings

The configuration of Web VTT captions in the output.

The parent of this entity is CaptionDestinationSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StyleControl" : String
}

```

### YAML

```yaml

  StyleControl: String

```

## Properties

`StyleControl`

Controls whether the color and position of the source captions is passed through to the WebVTT output captions. PASSTHROUGH - Valid only if the source captions are EMBEDDED or TELETEXT. NO\_STYLE\_DATA - Don't pass through the style. The output captions will not contain any font styling information.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WavSettings

AWS::MediaLive::ChannelPlacementGroup

All content copied from https://docs.aws.amazon.com/.
