This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow MediaStream

A media stream represents one component of your content, such as video, audio, or ancillary data. After you add a media stream to your flow, you can associate it with sources and outputs that use the ST 2110 JPEG XS or CDI protocol.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : MediaStreamAttributes,
  "ClockRate" : Integer,
  "Description" : String,
  "Fmt" : Integer,
  "MediaStreamId" : Integer,
  "MediaStreamName" : String,
  "MediaStreamType" : String,
  "Tags" : [ Tag, ... ],
  "VideoFormat" : String
}

```

### YAML

```yaml

  Attributes:
    MediaStreamAttributes
  ClockRate: Integer
  Description: String
  Fmt: Integer
  MediaStreamId: Integer
  MediaStreamName: String
  MediaStreamType: String
  Tags:
    - Tag
  VideoFormat: String

```

## Properties

`Attributes`

Attributes that are related to the media stream.

_Required_: No

_Type_: [MediaStreamAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flow-mediastreamattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClockRate`

The sample rate for the stream. This value is measured in Hz.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description that can help you quickly identify what your media stream is used for.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fmt`

The format type number (sometimes referred to as RTP payload type) of the media stream. MediaConnect assigns this value to the media stream. For ST 2110 JPEG XS outputs, you need to provide this value to the receiver.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreamId`

A unique identifier for the media stream.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreamName`

A name that helps you distinguish one media stream from another.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreamType`

The type of media stream.

_Required_: Yes

_Type_: String

_Allowed values_: `video | audio | ancillary-data`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-flow-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoFormat`

The resolution of the video.

_Required_: No

_Type_: String

_Allowed values_: `2160p | 1080p | 1080i | 720p | 480p`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Maintenance

MediaStreamAttributes
