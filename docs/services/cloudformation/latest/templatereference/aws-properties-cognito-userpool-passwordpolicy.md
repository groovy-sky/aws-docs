This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPool PasswordPolicy

The password policy settings for a user pool, including complexity, history, and
length requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MinimumLength" : Integer,
  "PasswordHistorySize" : Integer,
  "RequireLowercase" : Boolean,
  "RequireNumbers" : Boolean,
  "RequireSymbols" : Boolean,
  "RequireUppercase" : Boolean,
  "TemporaryPasswordValidityDays" : Integer
}

```

### YAML

```yaml

  MinimumLength: Integer
  PasswordHistorySize: Integer
  RequireLowercase: Boolean
  RequireNumbers: Boolean
  RequireSymbols: Boolean
  RequireUppercase: Boolean
  TemporaryPasswordValidityDays: Integer

```

## Properties

`MinimumLength`

The minimum length of the password in the policy that you have set. This value can't
be less than 6.

_Required_: No

_Type_: Integer

_Minimum_: `6`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PasswordHistorySize`

The number of previous passwords that you want Amazon Cognito to restrict each user
from reusing. Users can't set a password that matches any of `n` previous
passwords, where `n` is the value of `PasswordHistorySize`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `24`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireLowercase`

The requirement in a password policy that users must include at least one lowercase
letter in their password.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireNumbers`

The requirement in a password policy that users must include at least one number in
their password.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireSymbols`

The requirement in a password policy that users must include at least one symbol in
their password.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequireUppercase`

The requirement in a password policy that users must include at least one uppercase
letter in their password.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemporaryPasswordValidityDays`

The number of days a temporary password is valid in the password policy. If the user
doesn't sign in during this time, an administrator must reset their password. Defaults
to `7`. If you submit a value of `0`, Amazon Cognito treats it as a null
value and sets `TemporaryPasswordValidityDays` to its default value.

###### Note

When you set `TemporaryPasswordValidityDays` for a user pool, you can
no longer set a value for the legacy `UnusedAccountValidityDays`
parameter in that user pool.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `365`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NumberAttributeConstraints

Policies

All content copied from https://docs.aws.amazon.com/.
