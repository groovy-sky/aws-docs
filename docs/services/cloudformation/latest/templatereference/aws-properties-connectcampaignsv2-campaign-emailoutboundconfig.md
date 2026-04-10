This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign EmailOutboundConfig

The outbound configuration for email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectSourceEmailAddress" : String,
  "SourceEmailAddressDisplayName" : String,
  "WisdomTemplateArn" : String
}

```

### YAML

```yaml

  ConnectSourceEmailAddress: String
  SourceEmailAddressDisplayName: String
  WisdomTemplateArn: String

```

## Properties

`ConnectSourceEmailAddress`

The Amazon Connect source email address.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w-\.\+]+@([\w-]+\.)+[\w-]{2,4}$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceEmailAddressDisplayName`

The display name for the Amazon Connect source email address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

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

EmailChannelSubtypeConfig

EmailOutboundMode

All content copied from https://docs.aws.amazon.com/.
