This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration HlsManifest

Parameters for an HLS manifest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdMarkers" : String,
  "IncludeIframeOnlyStream" : Boolean,
  "ManifestName" : String,
  "ProgramDateTimeIntervalSeconds" : Integer,
  "RepeatExtXKey" : Boolean,
  "StreamSelection" : StreamSelection
}

```

### YAML

```yaml

  AdMarkers: String
  IncludeIframeOnlyStream: Boolean
  ManifestName: String
  ProgramDateTimeIntervalSeconds: Integer
  RepeatExtXKey: Boolean
  StreamSelection:
    StreamSelection

```

## Properties

`AdMarkers`

This setting controls ad markers in the packaged content.

Valid values:

- `NONE` \- Omits all SCTE-35 ad markers from the output.

- `PASSTHROUGH` \- Creates a copy in the output of the SCTE-35 ad markers (comments) taken directly from the input manifest.

- `SCTE35_ENHANCED` \- Generates ad markers and blackout tags in the output based on the SCTE-35 messages from the input manifest.

_Required_: No

_Type_: String

_Allowed values_: `NONE | SCTE35_ENHANCED | PASSTHROUGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeIframeOnlyStream`

Applies to stream sets with a single video track only. When enabled, the output includes an additional I-frame only stream, along with the other tracks.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestName`

A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgramDateTimeIntervalSeconds`

Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.

Irrespective of this parameter, if any ID3Timed metadata is in the HLS input, it is passed through to the HLS output.

Omit this attribute or enter `0` to indicate that the
`EXT-X-PROGRAM-DATE-TIME` tags are not included in the manifest.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepeatExtXKey`

Repeat the `EXT-X-KEY` directive for every media segment. This might result in an increase in client requests to the DRM server.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamSelection`

Video bitrate limitations for outputs from this packaging configuration.

_Required_: No

_Type_: [StreamSelection](aws-properties-mediapackage-packagingconfiguration-streamselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsEncryption

HlsPackage

All content copied from https://docs.aws.amazon.com/.
