This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MediaPackageV2GroupSettings

The `MediaPackageV2GroupSettings` property type specifies Property description not available. for an [AWS::MediaLive::Channel](aws-resource-medialive-channel.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalDestinations" : [ MediaPackageAdditionalDestinations, ... ],
  "CaptionLanguageMappings" : [ CaptionLanguageMapping, ... ],
  "Id3Behavior" : String,
  "KlvBehavior" : String,
  "NielsenId3Behavior" : String,
  "Scte35Type" : String,
  "SegmentLength" : Integer,
  "SegmentLengthUnits" : String,
  "TimedMetadataId3Frame" : String,
  "TimedMetadataId3Period" : Integer,
  "TimedMetadataPassthrough" : String
}

```

### YAML

```yaml

  AdditionalDestinations:
    - MediaPackageAdditionalDestinations
  CaptionLanguageMappings:
    - CaptionLanguageMapping
  Id3Behavior: String
  KlvBehavior: String
  NielsenId3Behavior: String
  Scte35Type: String
  SegmentLength: Integer
  SegmentLengthUnits: String
  TimedMetadataId3Frame: String
  TimedMetadataId3Period: Integer
  TimedMetadataPassthrough: String

```

## Properties

`AdditionalDestinations`

Property description not available.

_Required_: No

_Type_: [Array](aws-properties-medialive-channel-additionaldestinations.md) of [MediaPackageAdditionalDestinations](aws-properties-medialive-channel-mediapackageadditionaldestinations.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionLanguageMappings`

Property description not available.

_Required_: No

_Type_: Array of [CaptionLanguageMapping](aws-properties-medialive-channel-captionlanguagemapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id3Behavior`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KlvBehavior`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenId3Behavior`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35Type`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentLength`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentLengthUnits`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataId3Frame`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataId3Period`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimedMetadataPassthrough`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaPackageV2DestinationSettings

MotionGraphicsConfiguration

All content copied from https://docs.aws.amazon.com/.
