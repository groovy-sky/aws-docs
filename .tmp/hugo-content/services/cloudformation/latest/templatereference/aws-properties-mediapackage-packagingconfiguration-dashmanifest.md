This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration DashManifest

Parameters for a DASH manifest.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManifestLayout" : String,
  "ManifestName" : String,
  "MinBufferTimeSeconds" : Integer,
  "Profile" : String,
  "ScteMarkersSource" : String,
  "StreamSelection" : StreamSelection
}

```

### YAML

```yaml

  ManifestLayout: String
  ManifestName: String
  MinBufferTimeSeconds: Integer
  Profile: String
  ScteMarkersSource: String
  StreamSelection:
    StreamSelection

```

## Properties

`ManifestLayout`

Determines the position of some tags in the Media Presentation Description (MPD). When set to `FULL`, elements like `SegmentTemplate` and
`ContentProtection` are included in each `Representation`. When set to `COMPACT`, duplicate elements are combined and presented at the
AdaptationSet level.

_Required_: No

_Type_: String

_Allowed values_: `FULL | COMPACT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestName`

A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinBufferTimeSeconds`

Minimum amount of content (measured in seconds) that a player must keep available in the buffer.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Profile`

The DASH profile type. When set to `HBBTV_1_5`, the content is compliant with HbbTV 1.5.

_Required_: No

_Type_: String

_Allowed values_: `NONE | HBBTV_1_5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScteMarkersSource`

The source of scte markers used.

Value description:

- `SEGMENTS` \- The scte markers are sourced from the segments of the ingested content.

- `MANIFEST` \- the scte markers are sourced from the manifest of the ingested content.
The MANIFEST value is compatible with source HLS playlists using the SCTE-35 Enhanced syntax ( `EXT-OATCLS-SCTE35` tags).
SCTE-35 Elemental and SCTE-35 Daterange syntaxes are not supported with this option.

_Required_: No

_Type_: String

_Allowed values_: `SEGMENTS | MANIFEST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamSelection`

Limitations for outputs from the endpoint, based on the video bitrate.

_Required_: No

_Type_: [StreamSelection](aws-properties-mediapackage-packagingconfiguration-streamselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashEncryption

DashPackage

All content copied from https://docs.aws.amazon.com/.
