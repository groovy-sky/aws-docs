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

_Type_: [OwnerGroupProperties](aws-properties-datazone-owner-ownergroupproperties.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`User`

Specifies that the domain unit owner is a user.

_Required_: No

_Type_: [OwnerUserProperties](aws-properties-datazone-owner-owneruserproperties.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OwnerGroupProperties

OwnerUserProperties

All content copied from https://docs.aws.amazon.com/.
