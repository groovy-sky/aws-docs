This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel H265ColorSpaceSettings

H265 Color Space Settings

The parent of this entity is H265Settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorSpacePassthroughSettings" : Json,
  "DolbyVision81Settings" : Json,
  "Hdr10Settings" : Hdr10Settings,
  "Hlg2020Settings" : Json,
  "Rec601Settings" : Json,
  "Rec709Settings" : Json
}

```

### YAML

```yaml

  ColorSpacePassthroughSettings: Json
  DolbyVision81Settings: Json
  Hdr10Settings:
    Hdr10Settings
  Hlg2020Settings: Json
  Rec601Settings: Json
  Rec709Settings: Json

```

## Properties

`ColorSpacePassthroughSettings`

Passthrough applies no color space conversion to the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DolbyVision81Settings`

Property description not available.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Hdr10Settings`

Settings to configure the handling of HDR10 color space.

_Required_: No

_Type_: [Hdr10Settings](aws-properties-medialive-channel-hdr10settings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Hlg2020Settings`

Property description not available.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rec601Settings`

Settings to configure the handling of Rec601 color space.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rec709Settings`

Settings to configure the handling of Rec709 color space.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

H264Settings

H265FilterSettings

All content copied from https://docs.aws.amazon.com/.
