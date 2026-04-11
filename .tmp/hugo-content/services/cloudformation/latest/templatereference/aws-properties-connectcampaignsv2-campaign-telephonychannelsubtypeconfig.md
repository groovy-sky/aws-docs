This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign TelephonyChannelSubtypeConfig

The configuration for the telephony channel subtype.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Capacity" : Number,
  "ConnectQueueId" : String,
  "DefaultOutboundConfig" : TelephonyOutboundConfig,
  "OutboundMode" : TelephonyOutboundMode
}

```

### YAML

```yaml

  Capacity: Number
  ConnectQueueId: String
  DefaultOutboundConfig:
    TelephonyOutboundConfig
  OutboundMode:
    TelephonyOutboundMode

```

## Properties

`Capacity`

The allocation of telephony capacity between multiple running outbound campaigns.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectQueueId`

The identifier of the Amazon Connect queue associated with telephony outbound requests of an outbound campaign.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOutboundConfig`

The default telephony outbound configuration of an outbound campaign.

_Required_: Yes

_Type_: [TelephonyOutboundConfig](aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundMode`

The outbound mode of telephony for an outbound campaign.

_Required_: Yes

_Type_: [TelephonyOutboundMode](aws-properties-connectcampaignsv2-campaign-telephonyoutboundmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TelephonyOutboundConfig

All content copied from https://docs.aws.amazon.com/.
