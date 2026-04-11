This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint LowLatencyHlsManifestConfiguration

Specify a low-latency HTTP live streaming (LL-HLS) manifest configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChildManifestName" : String,
  "FilterConfiguration" : FilterConfiguration,
  "ManifestName" : String,
  "ManifestWindowSeconds" : Integer,
  "ProgramDateTimeIntervalSeconds" : Integer,
  "ScteHls" : ScteHls,
  "StartTag" : StartTag,
  "Url" : String,
  "UrlEncodeChildManifest" : Boolean
}

```

### YAML

```yaml

  ChildManifestName: String
  FilterConfiguration:
    FilterConfiguration
  ManifestName: String
  ManifestWindowSeconds: Integer
  ProgramDateTimeIntervalSeconds: Integer
  ScteHls:
    ScteHls
  StartTag:
    StartTag
  Url: String
  UrlEncodeChildManifest: Boolean

```

## Properties

`ChildManifestName`

The name of the child manifest associated with the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterConfiguration`

Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.

_Required_: No

_Type_: [FilterConfiguration](aws-properties-mediapackagev2-originendpoint-filterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestName`

A short string that's appended to the endpoint URL. The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, `index`. MediaPackage automatically inserts the format extension, such as `.m3u8`. You can't use the same manifest name if you use HLS manifest and low-latency HLS manifest. The `manifestName` on the `HLSManifest` object overrides the `manifestName` you provided on the `originEndpoint` object.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestWindowSeconds`

The total duration (in seconds) of the manifest's content.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramDateTimeIntervalSeconds`

Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify. If you don't enter an interval, `EXT-X-PROGRAM-DATE-TIME` tags aren't included in the manifest.
The tags sync the stream to the wall clock so that viewers can seek to a specific time in the playback timeline on the player.

Irrespective of this parameter, if any `ID3Timed` metadata is in the HLS input, MediaPackage passes through that metadata to the HLS output.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScteHls`

The SCTE-35 HLS configuration associated with the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.

_Required_: No

_Type_: [ScteHls](aws-properties-mediapackagev2-originendpoint-sctehls.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTag`

To insert an EXT-X-START tag in your HLS playlist, specify a StartTag configuration object with a valid TimeOffset. When you do, you can also optionally specify whether to include a PRECISE value in the EXT-X-START tag.

_Required_: No

_Type_: [StartTag](aws-properties-mediapackagev2-originendpoint-starttag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL of the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UrlEncodeChildManifest`

When enabled, MediaPackage URL-encodes the query string for API requests for LL-HLS child manifests to comply with AWS Signature Version 4 (SigV4) signature signing protocol.
For more information, see [AWS Signature Version 4 for API requests](../../../iam/latest/userguide/reference-sigv.md) in _AWS Identity and Access Management User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsManifestConfiguration

MssManifestConfiguration

All content copied from https://docs.aws.amazon.com/.
