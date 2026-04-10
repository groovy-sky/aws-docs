This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel H264ColorSpaceSettings

Settings for configuring color space in an H264 video encode.

The parent of this entity is H264Settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorSpacePassthroughSettings" : Json,
  "Rec601Settings" : Json,
  "Rec709Settings" : Json
}

```

### YAML

```yaml

  ColorSpacePassthroughSettings: Json
  Rec601Settings: Json
  Rec709Settings: Json

```

## Properties

`ColorSpacePassthroughSettings`

Passthrough applies no color space conversion to the output.

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

GlobalConfiguration

H264FilterSettings

All content copied from https://docs.aws.amazon.com/.
