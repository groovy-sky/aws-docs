This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel AvailSettings

The settings for the ad avail setup in the output.

The parent of this entity is AvailConfiguration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Esam" : Esam,
  "Scte35SpliceInsert" : Scte35SpliceInsert,
  "Scte35TimeSignalApos" : Scte35TimeSignalApos
}

```

### YAML

```yaml

  Esam:
    Esam
  Scte35SpliceInsert:
    Scte35SpliceInsert
  Scte35TimeSignalApos:
    Scte35TimeSignalApos

```

## Properties

`Esam`

Property description not available.

_Required_: No

_Type_: [Esam](aws-properties-medialive-channel-esam.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35SpliceInsert`

The setup for SCTE-35 splice insert handling.

_Required_: No

_Type_: [Scte35SpliceInsert](aws-properties-medialive-channel-scte35spliceinsert.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scte35TimeSignalApos`

The setup for SCTE-35 time signal APOS handling.

_Required_: No

_Type_: [Scte35TimeSignalApos](aws-properties-medialive-channel-scte35timesignalapos.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AvailConfiguration

BandwidthReductionFilterSettings

All content copied from https://docs.aws.amazon.com/.
