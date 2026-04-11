This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel VideoSelector

Information about the video to extract from the input. An input can contain only one video selector.

The parent of this entity is InputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorSpace" : String,
  "ColorSpaceSettings" : VideoSelectorColorSpaceSettings,
  "ColorSpaceUsage" : String,
  "SelectorSettings" : VideoSelectorSettings
}

```

### YAML

```yaml

  ColorSpace: String
  ColorSpaceSettings:
    VideoSelectorColorSpaceSettings
  ColorSpaceUsage: String
  SelectorSettings:
    VideoSelectorSettings

```

## Properties

`ColorSpace`

Specifies the color space of an input. This setting works in tandem with colorSpaceConversion to determine if
MediaLive will perform any conversion.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorSpaceSettings`

Settings to configure color space settings in the incoming video.

_Required_: No

_Type_: [VideoSelectorColorSpaceSettings](aws-properties-medialive-channel-videoselectorcolorspacesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorSpaceUsage`

Applies only if colorSpace is a value other than Follow. This field controls how the value in the colorSpace
field is used. Fallback means that when the input does include color space data, that data is used, but when the
input has no color space data, the value in colorSpace is used. Choose fallback if your input is sometimes missing
color space data, but when it does have color space data, that data is correct. Force means to always use the value
in colorSpace. Choose force if your input usually has no color space data or might have unreliable color space
data.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectorSettings`

Information about the video to select from the content.

_Required_: No

_Type_: [VideoSelectorSettings](aws-properties-medialive-channel-videoselectorsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VideoDescription

VideoSelectorColorSpaceSettings

All content copied from https://docs.aws.amazon.com/.
