This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::PackagingConfiguration CmafPackage

Parameters for a packaging configuration that uses Common Media Application Format (CMAF) packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encryption" : CmafEncryption,
  "HlsManifests" : [ HlsManifest, ... ],
  "IncludeEncoderConfigurationInSegments" : Boolean,
  "SegmentDurationSeconds" : Integer
}

```

### YAML

```yaml

  Encryption:
    CmafEncryption
  HlsManifests:
    - HlsManifest
  IncludeEncoderConfigurationInSegments: Boolean
  SegmentDurationSeconds: Integer

```

## Properties

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [CmafEncryption](aws-properties-mediapackage-packagingconfiguration-cmafencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsManifests`

A list of HLS manifest configurations that are available from this endpoint.

_Required_: Yes

_Type_: Array of [HlsManifest](aws-properties-mediapackage-packagingconfiguration-hlsmanifest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeEncoderConfigurationInSegments`

When includeEncoderConfigurationInSegments is set to true, AWS Elemental MediaPackage places your encoder's Sequence Parameter Set (SPS), Picture Parameter Set (PPS), and Video Parameter Set (VPS) metadata in every video segment instead of in the init fragment. This lets you use different SPS/PPS/VPS settings for your assets during content playback.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each segment. Actual segments are rounded to the nearest multiple of the source fragment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CmafEncryption

DashEncryption

All content copied from https://docs.aws.amazon.com/.
