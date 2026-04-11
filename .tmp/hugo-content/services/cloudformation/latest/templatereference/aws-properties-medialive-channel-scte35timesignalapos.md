This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel Scte35TimeSignalApos

The settings for the SCTE-35 time signal APOS mode.

The parent of this entity is AvailSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdAvailOffset" : Integer,
  "NoRegionalBlackoutFlag" : String,
  "WebDeliveryAllowedFlag" : String
}

```

### YAML

```yaml

  AdAvailOffset: Integer
  NoRegionalBlackoutFlag: String
  WebDeliveryAllowedFlag: String

```

## Properties

`AdAvailOffset`

When specified, this offset (in milliseconds) is added to the input ad avail PTS time. This applies only to
embedded SCTE 104/35 messages. It doesn't apply to OOB messages.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoRegionalBlackoutFlag`

When set to ignore, segment descriptors with noRegionalBlackoutFlag set to 0 no longer trigger blackouts or ad
avail slates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebDeliveryAllowedFlag`

When set to ignore, segment descriptors with webDeliveryAllowedFlag set to 0 no longer trigger blackouts or ad
avail slates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Scte35SpliceInsert

SrtGroupSettings

All content copied from https://docs.aws.amazon.com/.
