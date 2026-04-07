This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::Identity MailFromAttributes

A list of attributes that are associated with a MAIL FROM domain.

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

The action that Amazon Pinpoint to takes if it can't read the required MX record for a custom
MAIL FROM domain. When you set this value to `UseDefaultValue`, Amazon Pinpoint uses
_amazonses.com_ as the MAIL FROM domain. When you set this value
to `RejectMessage`, Amazon Pinpoint returns a `MailFromDomainNotVerified`
error, and doesn't attempt to deliver the email.

These behaviors are taken when the custom MAIL FROM domain configuration is in the
`Pending`, `Failed`, and `TemporaryFailure`
states.

_Required_: No

_Type_: String

_Allowed values_: `USE_DEFAULT_VALUE | REJECT_MESSAGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailFromDomain`

The name of a domain that an email identity uses as a custom MAIL FROM domain.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::PinpointEmail::Identity

Tags
