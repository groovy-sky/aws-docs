This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::EmailIdentity MailFromAttributes

Used to enable or disable the custom Mail-From domain configuration for an email
identity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BehaviorOnMxFailure" : String,
  "MailFromDomain" : String
}

```

### YAML

```yaml

  BehaviorOnMxFailure: String
  MailFromDomain: String

```

## Properties

`BehaviorOnMxFailure`

The action to take if the required MX record isn't found when you send an email. When
you set this value to `USE_DEFAULT_VALUE`, the mail is sent using
_amazonses.com_ as the MAIL FROM domain. When you set this value
to `REJECT_MESSAGE`, the Amazon SES API v2 returns a
`MailFromDomainNotVerified` error, and doesn't attempt to deliver the
email.

These behaviors are taken when the custom MAIL FROM domain configuration is in the
`Pending`, `Failed`, and `TemporaryFailure`
states.

Valid Values: `USE_DEFAULT_VALUE | REJECT_MESSAGE`

_Required_: No

_Type_: String

_Pattern_: `USE_DEFAULT_VALUE|REJECT_MESSAGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailFromDomain`

The custom MAIL FROM domain that you want the verified identity to use. The MAIL FROM
domain must meet the following criteria:

- It has to be a subdomain of the verified identity.

- It can't be used to receive email.

- It can't be used in a "From" address if the MAIL FROM domain is a destination
for feedback forwarding emails.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FeedbackAttributes

Tag

All content copied from https://docs.aws.amazon.com/.
