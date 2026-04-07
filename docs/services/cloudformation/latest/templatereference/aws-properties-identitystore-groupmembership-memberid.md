This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IdentityStore::GroupMembership MemberId

An object that contains the identifier of a group member. Setting the
`UserID` field to the specific identifier for a user indicates that the user
is a member of the group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UserId" : String
}

```

### YAML

```yaml

  UserId: String

```

## Properties

`UserId`

An object containing the identifiers of resources that can be members.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`

_Minimum_: `1`

_Maximum_: `47`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IdentityStore::GroupMembership

Next
