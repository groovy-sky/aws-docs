This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment NotifyConfigurationType

The configuration for Amazon SES email messages that advanced security features sends
to a user when your adaptive authentication automated response has a
_Notify_ action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockEmail" : NotifyEmailType,
  "From" : String,
  "MfaEmail" : NotifyEmailType,
  "NoActionEmail" : NotifyEmailType,
  "ReplyTo" : String,
  "SourceArn" : String
}

```

### YAML

```yaml

  BlockEmail:
    NotifyEmailType
  From: String
  MfaEmail:
    NotifyEmailType
  NoActionEmail:
    NotifyEmailType
  ReplyTo: String
  SourceArn: String

```

## Properties

`BlockEmail`

The template for the email message that your user pool sends when a detected risk
event is blocked.

_Required_: No

_Type_: [NotifyEmailType](aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`From`

The email address that sends the email message. The address must be either
individually verified with Amazon Simple Email Service, or from a domain that has been verified with
Amazon SES.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MfaEmail`

The template for the email message that your user pool sends when MFA is challenged in
response to a detected risk.

_Required_: No

_Type_: [NotifyEmailType](aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoActionEmail`

The template for the email message that your user pool sends when no action is taken
in response to a detected risk.

_Required_: No

_Type_: [NotifyEmailType](aws-properties-cognito-userpoolriskconfigurationattachment-notifyemailtype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplyTo`

The reply-to email address of an email template. Can be an email address in the format
`admin@example.com` or `Administrator
                <admin@example.com>`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceArn`

The Amazon Resource Name (ARN) of the identity that is associated with the sending
authorization policy. This identity permits Amazon Cognito to send for the email address
specified in the `From` parameter.

_Required_: Yes

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CompromisedCredentialsRiskConfigurationType

NotifyEmailType

All content copied from https://docs.aws.amazon.com/.
