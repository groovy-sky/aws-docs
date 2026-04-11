This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Hdr10Settings

Hdr10 Settings

The parents of this entity are H265ColorSpaceSettings (for color space settings in the output) and
VideoSelectorColorSpaceSettings (for color space settings in the input).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCll" : Integer,
  "MaxFall" : Integer
}

```

### YAML

```yaml

  MaxCll: Integer
  MaxFall: Integer

```

## Properties

`MaxCll`

Maximum Content Light Level
An integer metadata value defining the maximum light level, in nits,
of any single pixel within an encoded HDR video stream or file.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxFall`

Maximum Frame Average Light Level
An integer metadata value defining the maximum average light level, in nits,
for any single frame within an encoded HDR video stream or file.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

H265Settings

HlsAkamaiSettings

All content copied from https://docs.aws.amazon.com/.
