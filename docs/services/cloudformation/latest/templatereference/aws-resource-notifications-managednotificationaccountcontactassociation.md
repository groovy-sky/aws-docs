This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Notifications::ManagedNotificationAccountContactAssociation

Associates an Account Management Contact with a
`ManagedNotificationConfiguration` for AWS User Notifications. For more
information about AWS User Notifications, see the [AWS User Notifications User Guide](../../../notifications/latest/userguide/what-is-service.md). For more information about Account Management Contacts, see the [AWS Account Management Reference Guide](../../../accounts/latest/reference/api-alternatecontact.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Notifications::ManagedNotificationAccountContactAssociation",
  "Properties" : {
      "ContactIdentifier" : String,
      "ManagedNotificationConfigurationArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Notifications::ManagedNotificationAccountContactAssociation
Properties:
  ContactIdentifier: String
  ManagedNotificationConfigurationArn: String

```

## Properties

`ContactIdentifier`

The unique identifier of the notification contact associated with the AWS
account. For more information about the contact types associated with an account, see the
[AWS Account Management Reference Guide](../../../accounts/latest/reference/manage-acct-update-contact-alternate.md#manage-acct-update-contact-alternate-orgs).

_Required_: Yes

_Type_: String

_Allowed values_: `ACCOUNT_PRIMARY | ACCOUNT_ALTERNATE_SECURITY | ACCOUNT_ALTERNATE_OPERATIONS | ACCOUNT_ALTERNATE_BILLING`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedNotificationConfigurationArn`

The ARN of the `ManagedNotificationConfiguration` to be associated with the
`Channel`.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]{3,10}:notifications::([0-9]{12}|):managed-notification-configuration/category/[a-zA-Z0-9\-]{3,64}/sub-category/[a-zA-Z0-9\-]{3,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the ARN of the configuration created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventRuleStatusSummary

AWS::Notifications::ManagedNotificationAdditionalChannelAssociation

All content copied from https://docs.aws.amazon.com/.
