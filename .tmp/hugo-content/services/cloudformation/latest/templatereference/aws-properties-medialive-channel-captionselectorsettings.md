This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel CaptionSelectorSettings

Captions Selector Settings

The parent of this entity is CaptionSelector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AncillarySourceSettings" : AncillarySourceSettings,
  "AribSourceSettings" : Json,
  "DvbSubSourceSettings" : DvbSubSourceSettings,
  "EmbeddedSourceSettings" : EmbeddedSourceSettings,
  "Scte20SourceSettings" : Scte20SourceSettings,
  "Scte27SourceSettings" : Scte27SourceSettings,
  "TeletextSourceSettings" : TeletextSourceSettings
}

```

### YAML

```yaml

  AncillarySourceSettings:
    AncillarySourceSettings
  AribSourceSettings: Json
  DvbSubSourceSettings:
    DvbSubSourceSettings
  EmbeddedSourceSettings:
    EmbeddedSourceSettings
  Scte20SourceSettings:
    Scte20SourceSettings
  Scte27SourceSettings:
    Scte27SourceSettings
  TeletextSourceSettings:
    TeletextSourceSettings

```

## Properties

`AncillarySourceSettings`

Information about the ancillary captions to extract from the input.

_Required_: No

_Type_: [AncillarySourceSettings](aws-properties-medialive-channel-ancillarysourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AribSourceSettings`

Information about the ARIB captions to extract from the input.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DvbSubSourceSettings`

Information about the DVB Sub captions to extract from the input.

_Required_: No

_Type_: [DvbSubSourceSettings](aws-properties-medialive-channel-dvbsubsourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmbeddedSourceSettings`

Information about the embedded captions to extract from the input.

_Required_: No

_Type_: [EmbeddedSourceSettings](aws-properties-medialive-channel-embeddedsourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte20SourceSettings`

Information about the SCTE-20 captions to extract from the input.

_Required_: No

_Type_: [Scte20SourceSettings](aws-properties-medialive-channel-scte20sourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte27SourceSettings`

Information about the SCTE-27 captions to extract from the input.

_Required_: No

_Type_: [Scte27SourceSettings](aws-properties-medialive-channel-scte27sourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeletextSourceSettings`

Information about the Teletext captions to extract from the input.

_Required_: No

_Type_: [TeletextSourceSettings](aws-properties-medialive-channel-teletextsourcesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptionSelector

CdiInputSpecification

All content copied from https://docs.aws.amazon.com/.
