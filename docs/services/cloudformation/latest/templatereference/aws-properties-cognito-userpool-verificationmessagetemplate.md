This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool VerificationMessageTemplate

The template for the verification message that your user pool delivers to users who
set an email address or phone number attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultEmailOption" : String,
  "EmailMessage" : String,
  "EmailMessageByLink" : String,
  "EmailSubject" : String,
  "EmailSubjectByLink" : String,
  "SmsMessage" : String
}

```

### YAML

```yaml

  DefaultEmailOption: String
  EmailMessage: String
  EmailMessageByLink: String
  EmailSubject: String
  EmailSubjectByLink: String
  SmsMessage: String

```

## Properties

`DefaultEmailOption`

The configuration of verification emails to contain a clickable link or a verification
code.

For link, your template body must contain link text in the format `{##Click
                here##}`. "Click here" in the example is a customizable string. For code, your
template body must contain a code placeholder in the format `{####}`.

_Required_: No

_Type_: String

_Allowed values_: `CONFIRM_WITH_LINK | CONFIRM_WITH_CODE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailMessage`

The template for email messages that Amazon Cognito sends to your users. You can set an
`EmailMessage` template only if the value of [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`. When your [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`, your user pool sends email
messages with your own Amazon SES configuration.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*`

_Minimum_: `6`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailMessageByLink`

The email message template for sending a confirmation link to the user. You can set an
`EmailMessageByLink` template only if the value of [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`. When your [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`, your user pool sends email
messages with your own Amazon SES configuration.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{##[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*##\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*`

_Minimum_: `6`

_Maximum_: `20000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailSubject`

The subject line for the email message template. You can set an
`EmailSubject` template only if the value of [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`. When your [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`, your user pool sends email
messages with your own Amazon SES configuration.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s]+`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailSubjectByLink`

The subject line for the email message template for sending a confirmation link to the
user. You can set an `EmailSubjectByLink` template only if the value of
[EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`. When your [EmailSendingAccount](../../../../reference/cognito-user-identity-pools/latest/apireference/api-emailconfigurationtype.md#CognitoUserPools-Type-EmailConfigurationType-EmailSendingAccount) is `DEVELOPER`, your user pool sends email
messages with your own Amazon SES configuration.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}\s]+`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmsMessage`

The template for SMS messages that Amazon Cognito sends to your users.

_Required_: No

_Type_: String

_Pattern_: `.*\{####\}.*`

_Minimum_: `6`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserPoolAddOns

AWS::Cognito::UserPoolClient

All content copied from https://docs.aws.amazon.com/.
