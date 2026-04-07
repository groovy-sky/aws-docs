This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::EncoderConfiguration Video

The Video property type describes a stream's video configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bitrate" : Integer,
  "Framerate" : Number,
  "Height" : Integer,
  "Width" : Integer
}

```

### YAML

```yaml

  Bitrate: Integer
  Framerate: Number
  Height: Integer
  Width: Integer

```

## Properties

`Bitrate`

Bitrate for generated output, in bps. Default: 2500000.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `8500000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Framerate`

Video frame rate, in fps. Default: 30.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Height`

Video-resolution height. Note that the maximum value is determined by width times height,
such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 720.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `1920`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Width`

Video-resolution width. Note that the maximum value is determined by width times height,
such that the maximum total pixels is 2073600 (1920x1080 or 1080x1920). Default: 1280.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `1920`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::IVS::IngestConfiguration
