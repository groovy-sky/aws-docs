This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel BlackoutSlate

The settings for a blackout slate.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlackoutSlateImage" : InputLocation,
  "NetworkEndBlackout" : String,
  "NetworkEndBlackoutImage" : InputLocation,
  "NetworkId" : String,
  "State" : String
}

```

### YAML

```yaml

  BlackoutSlateImage:
    InputLocation
  NetworkEndBlackout: String
  NetworkEndBlackoutImage:
    InputLocation
  NetworkId: String
  State: String

```

## Properties

`BlackoutSlateImage`

The blackout slate image to be used. Keep empty for solid black. Only .bmp and .png images are supported.

_Required_: No

_Type_: [InputLocation](aws-properties-medialive-channel-inputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkEndBlackout`

Setting to enabled causes MediaLive to blackout the video, audio, and captions, and raise the "Network Blackout
Image" slate when an SCTE104/35 Network End Segmentation Descriptor is encountered. The blackout is lifted when the
Network Start Segmentation Descriptor is encountered. The Network End and Network Start descriptors must contain a
network ID that matches the value entered in Network ID.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkEndBlackoutImage`

The path to the local file to use as the Network End Blackout image. The image is scaled to fill the entire
output raster.

_Required_: No

_Type_: [InputLocation](aws-properties-medialive-channel-inputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkId`

Provides a Network ID that matches EIDR ID format (for example, "10.XXXX/XXXX-XXXX-XXXX-XXXX-XXXX-C").

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

When set to enabled, this causes video, audio, and captions to be blanked when indicated by program
metadata.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BandwidthReductionFilterSettings

BurnInDestinationSettings

All content copied from https://docs.aws.amazon.com/.
