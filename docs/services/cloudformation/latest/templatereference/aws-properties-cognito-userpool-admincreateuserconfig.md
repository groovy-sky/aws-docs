This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool AdminCreateUserConfig

The settings for administrator creation of users in a user pool. Contains settings for
allowing user sign-up, customizing invitation messages to new users, and the amount of
time before temporary passwords expire.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowAdminCreateUserOnly" : Boolean,
  "InviteMessageTemplate" : InviteMessageTemplate,
  "UnusedAccountValidityDays" : Integer
}

```

### YAML

```yaml

  AllowAdminCreateUserOnly: Boolean
  InviteMessageTemplate:
    InviteMessageTemplate
  UnusedAccountValidityDays: Integer

```

## Properties

`AllowAdminCreateUserOnly`

The setting for allowing self-service sign-up. When `true`, only
administrators can create new user profiles. When `false`, users can register
themselves and create a new user profile with the `SignUp` operation.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InviteMessageTemplate`

The template for the welcome message to new users. This template must include the
`{####}` temporary password placeholder if you are creating users with
passwords. If your users don't have passwords, you can omit the placeholder.

See also [Customizing User Invitation Messages](../../../cognito/latest/developerguide/cognito-user-pool-settings-message-customizations.md#cognito-user-pool-settings-user-invitation-message-customization).

_Required_: No

_Type_: [InviteMessageTemplate](aws-properties-cognito-userpool-invitemessagetemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnusedAccountValidityDays`

This parameter is no longer in use.

The password expiration limit in days for administrator-created users. When this time
expires, the user can't sign in with their temporary password. To reset the account
after that time limit, you must call `AdminCreateUser` again, specifying
`RESEND` for the `MessageAction` parameter.

The default value for this parameter is 7.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `365`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccountRecoverySetting

AdvancedSecurityAdditionalFlows

All content copied from https://docs.aws.amazon.com/.
