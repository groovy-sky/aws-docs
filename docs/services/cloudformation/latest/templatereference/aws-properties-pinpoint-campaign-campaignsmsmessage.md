This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign CampaignSmsMessage

Specifies the content and settings for an SMS message that's sent to
recipients of a campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Body" : String,
  "EntityId" : String,
  "MessageType" : String,
  "OriginationNumber" : String,
  "SenderId" : String,
  "TemplateId" : String
}

```

### YAML

```yaml

  Body: String
  EntityId: String
  MessageType: String
  OriginationNumber: String
  SenderId: String
  TemplateId: String

```

## Properties

`Body`

The body of the SMS message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityId`

The entity ID or Principal Entity (PE) id received from the regulatory body for
sending SMS in your country.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageType`

The SMS message type. Valid values are `TRANSACTIONAL` (for
messages that are critical or time-sensitive, such as a one-time passwords) and
`PROMOTIONAL` (for messsages that aren't critical or
time-sensitive, such as marketing messages).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginationNumber`

The long code to send the SMS message from. This value should be one of the dedicated
long codes that's assigned to your AWS account. Although it isn't
required, we recommend that you specify the long code using an E.164 format to ensure
prompt and accurate delivery of the message. For example, +12065550100.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SenderId`

The alphabetic Sender ID to display as the sender of the message on a
recipient's device. Support for sender IDs varies by country or region. To
specify a phone number as the sender, omit this parameter and use
`OriginationNumber` instead. For more information about support
for Sender ID by country, see the [Amazon Pinpoint User Guide](../../../pinpoint/latest/userguide/channels-sms-countries.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateId`

The template ID received from the regulatory body for sending SMS in your
country.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CampaignInAppMessage

CustomDeliveryConfiguration

All content copied from https://docs.aws.amazon.com/.
