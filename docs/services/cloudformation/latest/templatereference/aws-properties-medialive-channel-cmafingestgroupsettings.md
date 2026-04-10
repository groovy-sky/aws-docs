This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CmafIngestGroupSettings

The `CmafIngestGroupSettings` property type specifies Property description not available. for an [AWS::MediaLive::Channel](aws-resource-medialive-channel.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalDestinations" : [ AdditionalDestinations, ... ],
  "CaptionLanguageMappings" : [ CmafIngestCaptionLanguageMapping, ... ],
  "Destination" : OutputLocationRef,
  "Id3Behavior" : String,
  "Id3NameModifier" : String,
  "KlvBehavior" : String,
  "KlvNameModifier" : String,
  "NielsenId3Behavior" : String,
  "NielsenId3NameModifier" : String,
  "Scte35NameModifier" : String,
  "Scte35Type" : String,
  "SegmentLength" : Integer,
  "SegmentLengthUnits" : String,
  "SendDelayMs" : Integer,
  "TimedMetadataId3Frame" : String,
  "TimedMetadataId3Period" : Integer,
  "TimedMetadataPassthrough" : String
}

```

### YAML

```yaml

  AdditionalDestinations:
    - AdditionalDestinations
  CaptionLanguageMappings:
    - CmafIngestCaptionLanguageMapping
  Destination:
    OutputLocationRef
  Id3Behavior: String
  Id3NameModifier: String
  KlvBehavior: String
  KlvNameModifier: String
  NielsenId3Behavior: String
  NielsenId3NameModifier: String
  Scte35NameModifier: String
  Scte35Type: String
  SegmentLength: Integer
  SegmentLengthUnits: String
  SendDelayMs: Integer
  TimedMetadataId3Frame: String
  TimedMetadataId3Period: Integer
  TimedMetadataPassthrough: String

```

## Properties

`AdditionalDestinations`

Property description not available.

_Required_: No

_Type_: [Array](aws-properties-medialive-channel-additionaldestinations.md) of [AdditionalDestinations](aws-properties-medialive-channel-additionaldestinations.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptionLanguageMappings`

Property description not available.

_Required_: No

_Type_: Array of [CmafIngestCaptionLanguageMapping](aws-properties-medialive-channel-cmafingestcaptionlanguagemapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

Property description not available.

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id3Behavior`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id3NameModifier`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KlvBehavior`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KlvNameModifier`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenId3Behavior`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenId3NameModifier`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35NameModifier`

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

`SendDelayMs`

Property description not available.

_Required_: No

_Type_: Integer

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

CmafIngestCaptionLanguageMapping

CmafIngestOutputSettings

All content copied from https://docs.aws.amazon.com/.
