This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign WhatsAppOutboundConfig

The outbound configuration for WhatsApp.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectSourcePhoneNumberArn" : String,
  "WisdomTemplateArn" : String
}

```

### YAML

```yaml

  ConnectSourcePhoneNumberArn: String
  WisdomTemplateArn: String

```

## Properties

`ConnectSourcePhoneNumberArn`

The Amazon Resource Name (ARN) of the Amazon Connect source WhatsApp phone number.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `20`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WisdomTemplateArn`

The Amazon Resource Name (ARN) of the Amazon Q in Connect template.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `20`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WhatsAppChannelSubtypeConfig

WhatsAppOutboundMode

All content copied from https://docs.aws.amazon.com/.
