This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AudioOnlyHlsSettings

The configuration of an audio-only HLS output.

The parent of this entity is HlsSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AudioGroupId" : String,
  "AudioOnlyImage" : InputLocation,
  "AudioTrackType" : String,
  "SegmentType" : String
}

```

### YAML

```yaml

  AudioGroupId: String
  AudioOnlyImage:
    InputLocation
  AudioTrackType: String
  SegmentType: String

```

## Properties

`AudioGroupId`

Specifies the group that the audio rendition belongs to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioOnlyImage`

Used with an audio-only stream. It must be a .jpg or .png file. If given, this image is used as the cover art
for the audio-only output. Ideally, it should be formatted for an iPhone screen for two reasons. The iPhone does not
resize the image; instead, it crops a centered image on the top/bottom and left/right. Additionally, this image file
gets saved bit-for-bit into every 10-second segment file, so it increases bandwidth by {image file size} \* {segment
count} \* {user count.}.

_Required_: No

_Type_: [InputLocation](aws-properties-medialive-channel-inputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioTrackType`

Four types of audio-only tracks are supported: Audio-Only Variant Stream The client can play back this
audio-only stream instead of video in low-bandwidth scenarios. Represented as an EXT-X-STREAM-INF in the HLS
manifest. Alternate Audio, Auto Select, Default Alternate rendition that the client should try to play back by
default. Represented as an EXT-X-MEDIA in the HLS manifest with DEFAULT=YES, AUTOSELECT=YES Alternate Audio, Auto
Select, Not Default Alternate rendition that the client might try to play back by default. Represented as an
EXT-X-MEDIA in the HLS manifest with DEFAULT=NO, AUTOSELECT=YES Alternate Audio, not Auto Select Alternate rendition
that the client will not try to play back by default. Represented as an EXT-X-MEDIA in the HLS manifest with
DEFAULT=NO, AUTOSELECT=NO.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentType`

Specifies the segment type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AudioNormalizationSettings

AudioPidSelection

All content copied from https://docs.aws.amazon.com/.
