This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::Identity

Specifies an identity to use for sending email through Amazon Pinpoint. In Amazon
Pinpoint, an _identity_ is an email address or domain that you use when you
send email. Before you can use Amazon Pinpoint to send an email from an identity, you first
have to verify it. By verifying an identity, you demonstrate that you're the owner of the
address or domain, and that you've given Amazon Pinpoint permission to send email from that
identity.

When you verify an email address, Amazon Pinpoint sends an email to the address. Your
email address is verified as soon as you follow the link in the verification email.

When you verify a domain, this operation provides a set of DKIM tokens, which you can
convert into CNAME tokens. You add these CNAME tokens to the DNS configuration for your
domain. Your domain is verified when Amazon Pinpoint detects these records in the DNS
configuration for your domain. It usually takes around 72 hours to complete the domain
verification process.

###### Important

When you use CloudFormation to specify an identity, CloudFormation might indicate that
the identity was created successfully. However, you have to verify the identity before you
can use it to send email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PinpointEmail::Identity",
  "Properties" : {
      "DkimSigningEnabled" : Boolean,
      "FeedbackForwardingEnabled" : Boolean,
      "MailFromAttributes" : MailFromAttributes,
      "Name" : String,
      "Tags" : [ Tags, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::PinpointEmail::Identity
Properties:
  DkimSigningEnabled: Boolean
  FeedbackForwardingEnabled: Boolean
  MailFromAttributes:
    MailFromAttributes
  Name: String
  Tags:
    - Tags

```

## Properties

`DkimSigningEnabled`

For domain identities, this attribute is used to enable or disable DomainKeys Identified
Mail (DKIM) signing for the domain.

If the value is `true`, then the messages that you send from the domain are
signed using both the DKIM keys for your domain, as well as the keys for the
`amazonses.com` domain. If the value is `false`, then the messages
that you send are only signed using the DKIM keys for the `amazonses.com`
domain.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FeedbackForwardingEnabled`

Used to enable or disable feedback forwarding for an identity. This setting determines
what happens when an identity is used to send an email that results in a bounce or complaint
event.

When you enable feedback forwarding, Amazon Pinpoint sends you email notifications when
bounce or complaint events occur. Amazon Pinpoint sends this notification to the address that
you specified in the Return-Path header of the original email.

When you disable feedback forwarding, Amazon Pinpoint sends notifications through other
mechanisms, such as by notifying an Amazon SNS topic. You're required to have a method of
tracking bounces and complaints. If you haven't set up another mechanism for receiving bounce
or complaint notifications, Amazon Pinpoint sends an email notification when these events
occur (even if this setting is disabled).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailFromAttributes`

Used to enable or disable the custom Mail-From domain configuration for an email
identity.

_Required_: No

_Type_: [MailFromAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-identity-mailfromattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The address or domain of the identity, such as _sender@example.com_ or
_example.co.uk_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An object that defines the tags (keys and values) that you want to associate with the
email identity.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-identity-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-identity-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myEmailIdentity" }`

For the Amazon Pinpoint identity `myEmailIdentity`, Ref returns the name of the
identity (the email address or domain name).

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IdentityDNSRecordName1`

The host name for the first token that you have to add to the DNS configuration for your
domain.

For more information, see [Verifying a Domain](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html#channels-email-manage-verify-domain) in the Amazon Pinpoint User Guide.

`IdentityDNSRecordName2`

The host name for the second token that you have to add to the DNS configuration for your
domain.

`IdentityDNSRecordName3`

The host name for the third token that you have to add to the DNS configuration for your
domain.

`IdentityDNSRecordValue1`

The record value for the first token that you have to add to the DNS configuration for
your domain.

`IdentityDNSRecordValue2`

The record value for the second token that you have to add to the DNS configuration for
your domain.

`IdentityDNSRecordValue3`

The record value for the third token that you have to add to the DNS configuration for
your domain.

## See also

- [Verifying Email\
Identities](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html) in the _Amazon Pinpoint User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tags

MailFromAttributes
