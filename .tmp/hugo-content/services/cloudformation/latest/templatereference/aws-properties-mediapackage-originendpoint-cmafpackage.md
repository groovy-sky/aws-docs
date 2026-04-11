This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackage::OriginEndpoint CmafPackage

Parameters for Common Media Application Format (CMAF) packaging.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encryption" : CmafEncryption,
  "HlsManifests" : [ HlsManifest, ... ],
  "SegmentDurationSeconds" : Integer,
  "SegmentPrefix" : String,
  "StreamSelection" : StreamSelection
}

```

### YAML

```yaml

  Encryption:
    CmafEncryption
  HlsManifests:
    - HlsManifest
  SegmentDurationSeconds: Integer
  SegmentPrefix: String
  StreamSelection:
    StreamSelection

```

## Properties

`Encryption`

Parameters for encrypting content.

_Required_: No

_Type_: [CmafEncryption](aws-properties-mediapackage-originendpoint-cmafencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HlsManifests`

A list of HLS manifest configurations that are available from this endpoint.

_Required_: No

_Type_: Array of [HlsManifest](aws-properties-mediapackage-originendpoint-hlsmanifest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDurationSeconds`

Duration (in seconds) of each segment. Actual segments are rounded to the nearest multiple of the source segment duration.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentPrefix`

An optional custom string that is prepended to the name of each segment. If not specified, the segment prefix defaults to the ChannelId.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamSelection`

Limitations for outputs from the endpoint, based on the video bitrate.

_Required_: No

_Type_: [StreamSelection](aws-properties-mediapackage-originendpoint-streamselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CmafEncryption

DashEncryption

All content copied from https://docs.aws.amazon.com/.
