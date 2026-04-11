This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::User

The `AWS::AppStream::User` resource creates a new user in the WorkSpaces Applications user pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::User",
  "Properties" : {
      "AuthenticationType" : String,
      "FirstName" : String,
      "LastName" : String,
      "MessageAction" : String,
      "UserName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::User
Properties:
  AuthenticationType: String
  FirstName: String
  LastName: String
  MessageAction: String
  UserName: String

```

## Properties

`AuthenticationType`

The authentication type for the user. You must specify USERPOOL.

_Required_: Yes

_Type_: String

_Allowed values_: `API | SAML | USERPOOL | AWS_AD`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirstName`

The first name, or given name, of the user.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_\-\s]+$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LastName`

The last name, or surname, of the user.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_\-\s]+$`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MessageAction`

The action to take for the welcome email that is sent to a user after the user is created in the user pool. If you specify SUPPRESS, no email is sent. If you specify RESEND, do not specify the first name or last name of the user. If the value is null, the email is sent.

###### Note

The temporary password in the welcome email is valid for only 7 days. If users don’t set their passwords within 7 days, you must send them a new welcome email.

_Required_: No

_Type_: String

_Allowed values_: `SUPPRESS | RESEND`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserName`

The email address of the user.

Users' email addresses are case-sensitive. During login, if they specify an email address that doesn't use the same capitalization as the email address specified when their user pool account was created, a "user does not exist" error message displays.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [CreateUser](../../../../reference/appstream2/latest/apireference/api-createuser.md) in the _Amazon WorkSpaces Applications API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppStream::StackUserAssociation

Next

All content copied from https://docs.aws.amazon.com/.
