This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::User LoginProfile

Creates a password for the specified user, giving the user the ability to access AWS services through the AWS Management Console. For more information about
managing passwords, see [Managing Passwords](../../../iam/latest/userguide/using-managinglogins.md) in the _IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Password" : String,
  "PasswordResetRequired" : Boolean
}

```

### YAML

```yaml

  Password: String
  PasswordResetRequired: Boolean

```

## Properties

`Password`

The user's password.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PasswordResetRequired`

Specifies whether the user is required to set a new password on next sign-in.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LoginProfile](../../../../reference/iam/latest/apireference/api-loginprofile.md) in the _AWS Identity and Access Management API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IAM::User

Policy

All content copied from https://docs.aws.amazon.com/.
