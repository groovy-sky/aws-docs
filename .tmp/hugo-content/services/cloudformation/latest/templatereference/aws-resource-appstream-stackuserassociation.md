This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::StackUserAssociation

The `AWS::AppStream::StackUserAssociation` resource associates the specified users with the specified stacks for Amazon WorkSpaces Applications. Users in an WorkSpaces Applications user pool cannot be assigned to stacks with fleets that are joined to an Active Directory domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppStream::StackUserAssociation",
  "Properties" : {
      "AuthenticationType" : String,
      "SendEmailNotification" : Boolean,
      "StackName" : String,
      "UserName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppStream::StackUserAssociation
Properties:
  AuthenticationType: String
  SendEmailNotification: Boolean
  StackName: String
  UserName: String

```

## Properties

`AuthenticationType`

The authentication type for the user who is associated with the stack. You must specify USERPOOL.

_Required_: Yes

_Type_: String

_Allowed values_: `API | SAML | USERPOOL | AWS_AD`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SendEmailNotification`

Specifies whether a welcome email is sent to a user after the user is created in the user pool.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StackName`

The name of the stack that is associated with the user.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserName`

The email address of the user who is associated with the stack.

###### Note

Users' email addresses are case-sensitive.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [BatchAssociateUserStack](../../../../reference/appstream2/latest/apireference/api-batchassociateuserstack.md) in the _Amazon WorkSpaces Applications API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppStream::StackFleetAssociation

AWS::AppStream::User

All content copied from https://docs.aws.amazon.com/.
