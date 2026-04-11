This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CaptionDestinationSettings

The configuration of one captions encode in the output.

The parent of this entity is CaptionDescription.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AribDestinationSettings" : Json,
  "BurnInDestinationSettings" : BurnInDestinationSettings,
  "DvbSubDestinationSettings" : DvbSubDestinationSettings,
  "EbuTtDDestinationSettings" : EbuTtDDestinationSettings,
  "EmbeddedDestinationSettings" : Json,
  "EmbeddedPlusScte20DestinationSettings" : Json,
  "RtmpCaptionInfoDestinationSettings" : Json,
  "Scte20PlusEmbeddedDestinationSettings" : Json,
  "Scte27DestinationSettings" : Json,
  "SmpteTtDestinationSettings" : Json,
  "TeletextDestinationSettings" : Json,
  "TtmlDestinationSettings" : TtmlDestinationSettings,
  "WebvttDestinationSettings" : WebvttDestinationSettings
}

```

### YAML

```yaml

  AribDestinationSettings: Json
  BurnInDestinationSettings:
    BurnInDestinationSettings
  DvbSubDestinationSettings:
    DvbSubDestinationSettings
  EbuTtDDestinationSettings:
    EbuTtDDestinationSettings
  EmbeddedDestinationSettings: Json
  EmbeddedPlusScte20DestinationSettings: Json
  RtmpCaptionInfoDestinationSettings: Json
  Scte20PlusEmbeddedDestinationSettings: Json
  Scte27DestinationSettings: Json
  SmpteTtDestinationSettings: Json
  TeletextDestinationSettings: Json
  TtmlDestinationSettings:
    TtmlDestinationSettings
  WebvttDestinationSettings:
    WebvttDestinationSettings

```

## Properties

`AribDestinationSettings`

The configuration of one ARIB captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BurnInDestinationSettings`

The configuration of one burn-in captions encode in the output.

_Required_: No

_Type_: [BurnInDestinationSettings](aws-properties-medialive-channel-burnindestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DvbSubDestinationSettings`

The configuration of one DVB Sub captions encode in the output.

_Required_: No

_Type_: [DvbSubDestinationSettings](aws-properties-medialive-channel-dvbsubdestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EbuTtDDestinationSettings`

Settings for EBU-TT captions in the output.

_Required_: No

_Type_: [EbuTtDDestinationSettings](aws-properties-medialive-channel-ebuttddestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmbeddedDestinationSettings`

The configuration of one embedded captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmbeddedPlusScte20DestinationSettings`

The configuration of one embedded plus SCTE-20 captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RtmpCaptionInfoDestinationSettings`

The configuration of one RTMPCaptionInfo captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte20PlusEmbeddedDestinationSettings`

The configuration of one SCTE20 plus embedded captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte27DestinationSettings`

The configuration of one SCTE-27 captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmpteTtDestinationSettings`

The configuration of one SMPTE-TT captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeletextDestinationSettings`

The configuration of one Teletext captions encode in the output.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TtmlDestinationSettings`

The configuration of one TTML captions encode in the output.

_Required_: No

_Type_: [TtmlDestinationSettings](aws-properties-medialive-channel-ttmldestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebvttDestinationSettings`

The configuration of one WebVTT captions encode in the output.

_Required_: No

_Type_: [WebvttDestinationSettings](aws-properties-medialive-channel-webvttdestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptionDescription

CaptionLanguageMapping

All content copied from https://docs.aws.amazon.com/.
