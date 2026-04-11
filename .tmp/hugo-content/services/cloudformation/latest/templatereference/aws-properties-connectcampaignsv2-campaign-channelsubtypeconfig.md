This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign ChannelSubtypeConfig

Contains channel subtype configuration for an outbound campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Email" : EmailChannelSubtypeConfig,
  "Sms" : SmsChannelSubtypeConfig,
  "Telephony" : TelephonyChannelSubtypeConfig,
  "WhatsApp" : WhatsAppChannelSubtypeConfig
}

```

### YAML

```yaml

  Email:
    EmailChannelSubtypeConfig
  Sms:
    SmsChannelSubtypeConfig
  Telephony:
    TelephonyChannelSubtypeConfig
  WhatsApp:
    WhatsAppChannelSubtypeConfig

```

## Properties

`Email`

The configuration of the email channel subtype.

_Required_: No

_Type_: [EmailChannelSubtypeConfig](aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sms`

The configuration of the SMS channel subtype.

_Required_: No

_Type_: [SmsChannelSubtypeConfig](aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Telephony`

The configuration of the telephony channel subtype.

_Required_: No

_Type_: [TelephonyChannelSubtypeConfig](aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhatsApp`

The configuration of the WhatsApp channel subtype.

_Required_: No

_Type_: [WhatsAppChannelSubtypeConfig](aws-properties-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnswerMachineDetectionConfig

CommunicationLimit

All content copied from https://docs.aws.amazon.com/.
