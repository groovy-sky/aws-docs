This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign MessageConfiguration

Specifies the message configuration settings for a campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ADMMessage" : Message,
  "APNSMessage" : Message,
  "BaiduMessage" : Message,
  "CustomMessage" : CampaignCustomMessage,
  "DefaultMessage" : Message,
  "EmailMessage" : CampaignEmailMessage,
  "GCMMessage" : Message,
  "InAppMessage" : CampaignInAppMessage,
  "SMSMessage" : CampaignSmsMessage
}

```

### YAML

```yaml

  ADMMessage:
    Message
  APNSMessage:
    Message
  BaiduMessage:
    Message
  CustomMessage:
    CampaignCustomMessage
  DefaultMessage:
    Message
  EmailMessage:
    CampaignEmailMessage
  GCMMessage:
    Message
  InAppMessage:
    CampaignInAppMessage
  SMSMessage:
    CampaignSmsMessage

```

## Properties

`ADMMessage`

The message that the campaign sends through the ADM (Amazon Device Messaging)
channel. If specified, this message overrides the default message.

_Required_: No

_Type_: [Message](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-message.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`APNSMessage`

The message that the campaign sends through the APNs (Apple Push Notification
service) channel. If specified, this message overrides the default
message.

_Required_: No

_Type_: [Message](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-message.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaiduMessage`

The message that the campaign sends through the Baidu (Baidu Cloud Push)
channel. If specified, this message overrides the default message.

_Required_: No

_Type_: [Message](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-message.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomMessage`

The message that the campaign sends through a custom channel, as specified by
the delivery configuration ( `CustomDeliveryConfiguration`) settings
for the campaign. If specified, this message overrides the default
message.

_Required_: No

_Type_: [CampaignCustomMessage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-campaigncustommessage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultMessage`

The default message that the campaign sends through all the channels that are
configured for the campaign.

_Required_: No

_Type_: [Message](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-message.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailMessage`

The message that the campaign sends through the email channel. If specified,
this message overrides the default message.

###### Note

The maximum email message size is 200 KB. You can use email templates to send larger email messages.

_Required_: No

_Type_: [CampaignEmailMessage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-campaignemailmessage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GCMMessage`

The message that the campaign sends through the GCM channel, which enables
Amazon Pinpoint to send push notifications through the Firebase Cloud
Messaging (FCM), formerly Google Cloud Messaging (GCM), service. If specified,
this message overrides the default message.

_Required_: No

_Type_: [Message](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-message.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InAppMessage`

The default message for the in-app messaging channel. This message overrides
the default message ( `DefaultMessage`).

_Required_: No

_Type_: [CampaignInAppMessage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-campaigninappmessage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SMSMessage`

The message that the campaign sends through the SMS channel. If specified,
this message overrides the default message.

_Required_: No

_Type_: [CampaignSmsMessage](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-campaignsmsmessage.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Message

OverrideButtonConfiguration
