This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel InputSpecification

The input specification for this channel. It specifies the key characteristics of the inputs for this channel:
the maximum bitrate, the resolution, and the codec.

This entity is at the top level in the channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Codec" : String,
  "MaximumBitrate" : String,
  "Resolution" : String
}

```

### YAML

```yaml

  Codec: String
  MaximumBitrate: String
  Resolution: String

```

## Properties

`Codec`

The codec to include in the input specification for this channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBitrate`

The maximum input bitrate for any input attached to this channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resolution`

The resolution for any input attached to this channel.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InputSettings

KeyProviderSettings
