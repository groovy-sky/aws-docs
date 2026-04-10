This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign SmsChannelSubtypeConfig

The configuration for the SMS channel subtype.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Capacity" : Number,
  "DefaultOutboundConfig" : SmsOutboundConfig,
  "OutboundMode" : SmsOutboundMode
}

```

### YAML

```yaml

  Capacity: Number
  DefaultOutboundConfig:
    SmsOutboundConfig
  OutboundMode:
    SmsOutboundMode

```

## Properties

`Capacity`

The allocation of SMS capacity between multiple running outbound campaigns.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOutboundConfig`

The default SMS outbound configuration of an outbound campaign.

_Required_: Yes

_Type_: [SmsOutboundConfig](aws-properties-connectcampaignsv2-campaign-smsoutboundconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundMode`

The outbound mode of SMS for an outbound campaign.

_Required_: Yes

_Type_: [SmsOutboundMode](aws-properties-connectcampaignsv2-campaign-smsoutboundmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Schedule

SmsOutboundConfig

All content copied from https://docs.aws.amazon.com/.
