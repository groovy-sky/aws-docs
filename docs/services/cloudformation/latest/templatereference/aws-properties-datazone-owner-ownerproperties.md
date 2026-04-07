This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Owner OwnerProperties

The properties of a domain unit's owner.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Group" : OwnerGroupProperties,
  "User" : OwnerUserProperties
}

```

### YAML

```yaml

  Group:
    OwnerGroupProperties
  User:
    OwnerUserProperties

```

## Properties

`Group`

Specifies that the domain unit owner is a group.

_Required_: No

_Type_: [OwnerGroupProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-owner-ownergroupproperties.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`User`

Specifies that the domain unit owner is a user.

_Required_: No

_Type_: [OwnerUserProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-owner-owneruserproperties.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OwnerGroupProperties

OwnerUserProperties
