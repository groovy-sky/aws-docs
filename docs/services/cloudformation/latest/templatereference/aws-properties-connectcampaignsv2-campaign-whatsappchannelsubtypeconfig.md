This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign WhatsAppChannelSubtypeConfig

The configuration for the WhatsApp channel subtype.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Capacity" : Number,
  "DefaultOutboundConfig" : WhatsAppOutboundConfig,
  "OutboundMode" : WhatsAppOutboundMode
}

```

### YAML

```yaml

  Capacity: Number
  DefaultOutboundConfig:
    WhatsAppOutboundConfig
  OutboundMode:
    WhatsAppOutboundMode

```

## Properties

`Capacity`

The allocation of WhatsApp capacity between multiple running outbound campaigns.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOutboundConfig`

The default WhatsApp outbound configuration of an outbound campaign.

_Required_: Yes

_Type_: [WhatsAppOutboundConfig](aws-properties-connectcampaignsv2-campaign-whatsappoutboundconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundMode`

The outbound mode for WhatsApp of an outbound campaign.

_Required_: Yes

_Type_: [WhatsAppOutboundMode](aws-properties-connectcampaignsv2-campaign-whatsappoutboundmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeWindow

WhatsAppOutboundConfig

All content copied from https://docs.aws.amazon.com/.
