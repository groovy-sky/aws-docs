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

_Type_: [EmailChannelSubtypeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sms`

The configuration of the SMS channel subtype.

_Required_: No

_Type_: [SmsChannelSubtypeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Telephony`

The configuration of the telephony channel subtype.

_Required_: No

_Type_: [TelephonyChannelSubtypeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhatsApp`

The configuration of the WhatsApp channel subtype.

_Required_: No

_Type_: [WhatsAppChannelSubtypeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnswerMachineDetectionConfig

CommunicationLimit
