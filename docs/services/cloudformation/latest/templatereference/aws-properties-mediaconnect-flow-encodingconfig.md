This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow EncodingConfig

The encoding configuration to apply to the NDI® source when transcoding it to a transport stream for downstream distribution.
You can choose between several predefined encoding profiles based on common use cases.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncodingProfile" : String,
  "VideoMaxBitrate" : Integer
}

```

### YAML

```yaml

  EncodingProfile: String
  VideoMaxBitrate: Integer

```

## Properties

`EncodingProfile`

The encoding profile to use when transcoding the NDI source content to a transport stream. You can change this value while the flow is running.

_Required_: No

_Type_: String

_Allowed values_: `DISTRIBUTION_H264_DEFAULT | CONTRIBUTION_H264_DEFAULT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoMaxBitrate`

The maximum video bitrate to use when transcoding the NDI source to a transport stream.
This parameter enables you to override the default video bitrate within the encoding profile's supported range.

The supported range is 10,000,000 - 50,000,000 bits per second (bps). If you don't specify a value, MediaConnect uses the default value of 20,000,000 bps.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BlackFrames

Encryption
