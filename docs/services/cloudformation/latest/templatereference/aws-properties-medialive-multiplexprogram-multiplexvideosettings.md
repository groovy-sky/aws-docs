This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplexprogram MultiplexVideoSettings

The video configuration for each program in a multiplex.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConstantBitrate" : Integer,
  "StatmuxSettings" : MultiplexStatmuxVideoSettings
}

```

### YAML

```yaml

  ConstantBitrate: Integer
  StatmuxSettings:
    MultiplexStatmuxVideoSettings

```

## Properties

`ConstantBitrate`

The constant bitrate configuration for the video encode.
When this field is defined, StatmuxSettings must be undefined.

_Required_: No

_Type_: Integer

_Minimum_: `100000`

_Maximum_: `100000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatmuxSettings`

Statmux rate control settings.
When this field is defined, ConstantBitrate must be undefined.

_Required_: No

_Type_: [MultiplexStatmuxVideoSettings](aws-properties-medialive-multiplexprogram-multiplexstatmuxvideosettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexStatmuxVideoSettings

AWS::MediaLive::Network

All content copied from https://docs.aws.amazon.com/.
