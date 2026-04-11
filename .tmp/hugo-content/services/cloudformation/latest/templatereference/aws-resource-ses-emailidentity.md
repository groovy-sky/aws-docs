This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::EmailIdentity

Specifies an identity for using within SES. An identity is an email address or domain
that you use when you send email. Before you can use an identity to send email, you
first have to verify it. By verifying an identity, you demonstrate that you're the owner
of the identity, and that you've given Amazon SES API v2 permission to send email from
the identity.

When you verify an email address, SES sends an email to the address. Your email
address is verified as soon as you follow the link in the verification email. When you
verify a domain without specifying the `DkimSigningAttributes` properties, OR
only the `NextSigningKeyLength` property of
`DkimSigningAttributes`, this resource provides a set of CNAME token
names and values ( _DkimDNSTokenName1_,
_DkimDNSTokenValue1_, _DkimDNSTokenName2_,
_DkimDNSTokenValue2_, _DkimDNSTokenName3_,
_DkimDNSTokenValue3_) as outputs. You can then add these to the
DNS configuration for your domain. Your domain is verified when Amazon SES detects these
records in the DNS configuration for your domain. This verification method is known as
Easy DKIM.

Alternatively, you can perform the verification process by providing your own
public-private key pair. This verification method is known as Bring Your Own DKIM
(BYODKIM). To use BYODKIM, your resource must include `DkimSigningAttributes`
properties `DomainSigningSelector` and `DomainSigningPrivateKey`.
When you specify this object, you provide a selector
( `DomainSigningSelector`) (a component of the DNS record name that identifies
the public key to use for DKIM authentication) and a private key
( `DomainSigningPrivateKey`).

Additionally, you can associate an existing configuration set with the email identity
that you're verifying.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::EmailIdentity",
  "Properties" : {
      "ConfigurationSetAttributes" : ConfigurationSetAttributes,
      "DkimAttributes" : DkimAttributes,
      "DkimSigningAttributes" : DkimSigningAttributes,
      "EmailIdentity" : String,
      "FeedbackAttributes" : FeedbackAttributes,
      "MailFromAttributes" : MailFromAttributes,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::EmailIdentity
Properties:
  ConfigurationSetAttributes:
    ConfigurationSetAttributes
  DkimAttributes:
    DkimAttributes
  DkimSigningAttributes:
    DkimSigningAttributes
  EmailIdentity: String
  FeedbackAttributes:
    FeedbackAttributes
  MailFromAttributes:
    MailFromAttributes
  Tags:
    - Tag

```

## Properties

`ConfigurationSetAttributes`

Used to associate a configuration set with an email identity.

_Required_: No

_Type_: [ConfigurationSetAttributes](aws-properties-ses-emailidentity-configurationsetattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DkimAttributes`

An object that contains information about the DKIM attributes for the identity.

_Required_: No

_Type_: [DkimAttributes](aws-properties-ses-emailidentity-dkimattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DkimSigningAttributes`

If your request includes this object, Amazon SES configures the identity to use Bring Your
Own DKIM (BYODKIM) for DKIM authentication purposes, or, configures the key length to be
used for [Easy\
DKIM](../../../ses/latest/developerguide/easy-dkim.md).

You can only specify this object if the email identity is a domain, as opposed to an
address.

_Required_: No

_Type_: [DkimSigningAttributes](aws-properties-ses-emailidentity-dkimsigningattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailIdentity`

The email address or domain to verify.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeedbackAttributes`

Used to enable or disable feedback forwarding for an identity.

_Required_: No

_Type_: [FeedbackAttributes](aws-properties-ses-emailidentity-feedbackattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailFromAttributes`

Used to enable or disable the custom Mail-From domain configuration for an email
identity.

_Required_: No

_Type_: [MailFromAttributes](aws-properties-ses-emailidentity-mailfromattributes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of objects that define the tags (keys and values) to associate with the email
identity.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-emailidentity-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`DkimDNSTokenName1`

The host name for the first token that you have to add to the DNS configuration for
your domain.

`DkimDNSTokenName2`

The host name for the second token that you have to add to the DNS configuration for
your domain.

`DkimDNSTokenName3`

The host name for the third token that you have to add to the DNS configuration for
your domain.

`DkimDNSTokenValue1`

The record value for the first token that you have to add to the DNS configuration for
your domain.

`DkimDNSTokenValue2`

The record value for the second token that you have to add to the DNS configuration
for your domain.

`DkimDNSTokenValue3`

The record value for the third token that you have to add to the DNS configuration for
your domain.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ConfigurationSetAttributes

All content copied from https://docs.aws.amazon.com/.
