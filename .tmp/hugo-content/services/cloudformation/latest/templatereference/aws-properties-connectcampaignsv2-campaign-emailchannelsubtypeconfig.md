This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign EmailChannelSubtypeConfig

The configuration for the email channel subtype.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Capacity" : Number,
  "DefaultOutboundConfig" : EmailOutboundConfig,
  "OutboundMode" : EmailOutboundMode
}

```

### YAML

```yaml

  Capacity: Number
  DefaultOutboundConfig:
    EmailOutboundConfig
  OutboundMode:
    EmailOutboundMode

```

## Properties

`Capacity`

The allocation of email capacity between multiple running outbound campaigns.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOutboundConfig`

The default email outbound configuration of an outbound campaign.

_Required_: Yes

_Type_: [EmailOutboundConfig](aws-properties-connectcampaignsv2-campaign-emailoutboundconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundMode`

The outbound mode for email of an outbound campaign.

_Required_: Yes

_Type_: [EmailOutboundMode](aws-properties-connectcampaignsv2-campaign-emailoutboundmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DailyHour

EmailOutboundConfig

All content copied from https://docs.aws.amazon.com/.
