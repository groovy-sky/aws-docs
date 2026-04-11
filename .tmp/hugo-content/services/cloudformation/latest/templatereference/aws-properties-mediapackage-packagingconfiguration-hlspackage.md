This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration HlsPackage

Parameters for a packaging configuration that uses HTTP Live Streaming (HLS) packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encryption" : HlsEncryption,
  "HlsManifests" : [ HlsManifest, ... ],
  "IncludeDvbSubtitles" : Boolean,
  "SegmentDurationSeconds" : Integer,
  "UseAudioRenditionGroup" : Boolean
}

```

### YAML

```yaml

  Encryption:
    HlsEncryption
  HlsManifests:
    - HlsManifest
  IncludeDvbSubtitles: Boolean
  SegmentDurationSeconds: Integer
  UseAudioRenditionGroup: Boolean

```

## Properties

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [HlsEncryption](aws-properties-mediapackage-packagingconfiguration-hlsencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsManifests`

A list of HLS manifest configurations that are available from this endpoint.

_Required_: Yes

_Type_: Array of [HlsManifest](aws-properties-mediapackage-packagingconfiguration-hlsmanifest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeDvbSubtitles`

When enabled, MediaPackage passes through digital video broadcasting (DVB) subtitles into the output.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each fragment. Actual fragments are rounded to the nearest multiple of the source fragment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAudioRenditionGroup`

When true, AWS Elemental MediaPackage bundles all audio tracks in a rendition group. All other tracks in the stream can be used with any audio rendition from the group.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HlsManifest

MssEncryption

All content copied from https://docs.aws.amazon.com/.
