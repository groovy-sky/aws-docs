This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint Segment

The segment configuration, including the segment name, duration, and other configuration values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encryption" : Encryption,
  "IncludeIframeOnlyStreams" : Boolean,
  "Scte" : Scte,
  "SegmentDurationSeconds" : Integer,
  "SegmentName" : String,
  "TsIncludeDvbSubtitles" : Boolean,
  "TsUseAudioRenditionGroup" : Boolean
}

```

### YAML

```yaml

  Encryption:
    Encryption
  IncludeIframeOnlyStreams: Boolean
  Scte:
    Scte
  SegmentDurationSeconds: Integer
  SegmentName: String
  TsIncludeDvbSubtitles: Boolean
  TsUseAudioRenditionGroup: Boolean

```

## Properties

`Encryption`

Whether to use encryption for the segment.

_Required_: No

_Type_: [Encryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-encryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeIframeOnlyStreams`

Whether the segment includes I-frame-only streams.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte`

The SCTE-35 configuration associated with the segment.

_Required_: No

_Type_: [Scte](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-scte.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

The duration of the segment, in seconds.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentName`

The name of the segment associated with the origin endpoint.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TsIncludeDvbSubtitles`

Whether the segment includes DVB subtitles.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TsUseAudioRenditionGroup`

Whether the segment is an audio rendition group.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ScteHls

SpekeKeyProvider
