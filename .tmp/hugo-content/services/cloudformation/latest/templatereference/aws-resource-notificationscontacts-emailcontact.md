This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NotificationsContacts::EmailContact

Configures email contacts for AWS User Notifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NotificationsContacts::EmailContact",
  "Properties" : {
      "EmailAddress" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NotificationsContacts::EmailContact
Properties:
  EmailAddress: String
  Name: String
  Tags:
    - Tag

```

## Properties

`EmailAddress`

The email address of the contact. The activation and notification emails are sent
here.

_Required_: Yes

_Type_: String

_Pattern_: `^(.+)@(.+)$`

_Minimum_: `6`

_Maximum_: `254`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the contact.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-.~]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags to apply to the email contact.

_Required_: No

_Type_: Array of [Tag](aws-properties-notificationscontacts-emailcontact-tag.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the Amazon Resource Name (ARN) of the configuration created.

### Fn::GetAtt

`Arn`

Returns the ARN of the contact.

`EmailContact.Address`

Returns the email address that the contact points to.

`EmailContact.Arn`

Returns the ARN of the contact.

`EmailContact.CreationTime`

Returns the creation time of the resource.

`EmailContact.Name`

Returns the name of the contact.

`EmailContact.Status`

Returns the status of the contact.

`EmailContact.UpdateTime`

Returns the time the resource was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationsContacts

EmailContact

All content copied from https://docs.aws.amazon.com/.
