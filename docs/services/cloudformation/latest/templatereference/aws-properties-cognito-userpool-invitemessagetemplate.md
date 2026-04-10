This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool InviteMessageTemplate

The template for the welcome message to new users. This template must include the
`{####}` temporary password placeholder if you are creating users with
passwords. If your users don't have passwords, you can omit the placeholder.

See also [Customizing User Invitation Messages](../../../cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.md#cognito-user-pool-settings-user-invitation-message-customization).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EmailMessage" : String,
  "EmailSubject" : String,
  "SMSMessage" : String
}

```

### YAML

```yaml

  EmailMessage: String
  EmailSubject: String
  SMSMessage: String

```

## Properties

`EmailMessage`

The message template for email messages. EmailMessage is allowed only if [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*`

_Minimum_: `6`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailSubject`

The subject line for email messages. EmailSubject is allowed only if [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is DEVELOPER.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s]+`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SMSMessage`

The message template for SMS messages.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Minimum_: `6`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InboundFederation

LambdaConfig

All content copied from https://docs.aws.amazon.com/.
