This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign TemplateConfiguration

Specifies the message template to use for the message, for each type of
channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmailTemplate" : Template,
  "PushTemplate" : Template,
  "SMSTemplate" : Template,
  "VoiceTemplate" : Template
}

```

### YAML

```yaml

  EmailTemplate:
    Template
  PushTemplate:
    Template
  SMSTemplate:
    Template
  VoiceTemplate:
    Template

```

## Properties

`EmailTemplate`

The email template to use for the message.

_Required_: No

_Type_: [Template](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-template.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PushTemplate`

The push notification template to use for the message.

_Required_: No

_Type_: [Template](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-template.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SMSTemplate`

The SMS template to use for the message.

_Required_: No

_Type_: [Template](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-template.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VoiceTemplate`

The voice template to use for the message. This object isn't supported for
campaigns.

_Required_: No

_Type_: [Template](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-template.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Template

WriteTreatmentResource
