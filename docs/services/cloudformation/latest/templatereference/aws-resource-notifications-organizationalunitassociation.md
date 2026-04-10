This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Notifications::OrganizationalUnitAssociation

The `AWS::Notifications::OrganizationalUnitAssociation` resource Property description not available. for Notifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Notifications::OrganizationalUnitAssociation",
  "Properties" : {
      "NotificationConfigurationArn" : String,
      "OrganizationalUnitId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Notifications::OrganizationalUnitAssociation
Properties:
  NotificationConfigurationArn: String
  OrganizationalUnitId: String

```

## Properties

`NotificationConfigurationArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]{3,10}:notifications::[0-9]{12}:configuration/[a-z0-9]{27}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrganizationalUnitId`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^(r-[0-9a-z]{4,32})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationHubStatusSummary

Next

All content copied from https://docs.aws.amazon.com/.
