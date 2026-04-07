This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint HlsPackage

Parameters for Apple HLS packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdMarkers" : String,
  "AdsOnDeliveryRestrictions" : String,
  "AdTriggers" : [ String, ... ],
  "Encryption" : HlsEncryption,
  "IncludeDvbSubtitles" : Boolean,
  "IncludeIframeOnlyStream" : Boolean,
  "PlaylistType" : String,
  "PlaylistWindowSeconds" : Integer,
  "ProgramDateTimeIntervalSeconds" : Integer,
  "SegmentDurationSeconds" : Integer,
  "StreamSelection" : StreamSelection,
  "UseAudioRenditionGroup" : Boolean
}

```

### YAML

```yaml

  AdMarkers: String
  AdsOnDeliveryRestrictions: String
  AdTriggers:
    - String
  Encryption:
    HlsEncryption
  IncludeDvbSubtitles: Boolean
  IncludeIframeOnlyStream: Boolean
  PlaylistType: String
  PlaylistWindowSeconds: Integer
  ProgramDateTimeIntervalSeconds: Integer
  SegmentDurationSeconds: Integer
  StreamSelection:
    StreamSelection
  UseAudioRenditionGroup: Boolean

```

## Properties

`AdMarkers`

Controls how ad markers are included in the packaged endpoint.

Valid values:

- `NONE` \- Omits all SCTE-35 ad markers from the output.

- `PASSTHROUGH` \- Creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.

- `SCTE35_ENHANCED` \- Generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.

_Required_: No

_Type_: String

_Allowed values_: `NONE | SCTE35_ENHANCED | PASSTHROUGH | DATERANGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdsOnDeliveryRestrictions`

The flags on SCTE-35 segmentation descriptors that have to be present for AWS Elemental MediaPackage to insert ad markers in the output manifest. For information about SCTE-35 in AWS Elemental MediaPackage, see [SCTE-35 Message Options in AWS Elemental MediaPackage](https://docs.aws.amazon.com/mediapackage/latest/ug/scte.html).

_Required_: No

_Type_: String

_Allowed values_: `NONE | RESTRICTED | UNRESTRICTED | BOTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdTriggers`

Specifies the SCTE-35 message types that AWS Elemental MediaPackage treats as ad markers in the output manifest.

Valid values:

- `BREAK`

- `DISTRIBUTOR_ADVERTISEMENT`

- `DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY`

- `DISTRIBUTOR_PLACEMENT_OPPORTUNITY`

- `PROVIDER_ADVERTISEMENT`

- `PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY`

- `PROVIDER_PLACEMENT_OPPORTUNITY`

- `SPLICE_INSERT`

_Required_: No

_Type_: Array of String

_Allowed values_: `SPLICE_INSERT | BREAK | PROVIDER_ADVERTISEMENT | DISTRIBUTOR_ADVERTISEMENT | PROVIDER_PLACEMENT_OPPORTUNITY | DISTRIBUTOR_PLACEMENT_OPPORTUNITY | PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY | DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [HlsEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-hlsencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeDvbSubtitles`

When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeIframeOnlyStream`

Only applies to stream sets with a single video track. When true, the stream set includes an additional I-frame only stream, along with the other tracks. If false, this extra stream is not included.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlaylistType`

When specified as either `event` or `vod`, a corresponding `EXT-X-PLAYLIST-TYPE` entry is included in the media playlist.
Indicates if the playlist is live-to-VOD content.

_Required_: No

_Type_: String

_Allowed values_: `NONE | EVENT | VOD`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlaylistWindowSeconds`

Time window (in seconds) contained in each parent manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramDateTimeIntervalSeconds`

Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.

Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.

Omit this attribute or enter `0` to indicate that the
`EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each fragment. Actual fragments are rounded to the nearest multiple of the source fragment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamSelection`

Limitations for outputs from the endpoint, based on the video bitrate.

_Required_: No

_Type_: [StreamSelection](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackage-originendpoint-streamselection.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAudioRenditionGroup`

When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group. All other tracks in the stream can be used with any audio rendition from the group.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

HlsManifest

MssEncryption
