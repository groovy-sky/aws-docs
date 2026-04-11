This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IdentityStore::GroupMembership

Creates a relationship between a member and a group. The following identifiers must be
specified: `GroupId`, `IdentityStoreId`, and
`MemberId`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IdentityStore::GroupMembership",
  "Properties" : {
      "GroupId" : String,
      "IdentityStoreId" : String,
      "MemberId" : MemberId
    }
}

```

### YAML

```yaml

Type: AWS::IdentityStore::GroupMembership
Properties:
  GroupId: String
  IdentityStoreId: String
  MemberId:
    MemberId

```

## Properties

`GroupId`

The identifier for a group in the identity store.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$`

_Minimum_: `1`

_Maximum_: `47`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityStoreId`

The globally unique identifier for the identity store.

_Required_: Yes

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$|^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemberId`

An object containing the identifier of a group member. Setting the `MemberId`'s
`UserId` field to a specific User's ID indicates that user is a member of the group.

_Required_: Yes

_Type_: [MemberId](aws-properties-identitystore-groupmembership-memberid.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the `Physical ID` of the resource created
which is the format `GroupMembershipID|IdentityStoreId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`MembershipId`

The identifier for a `GroupMembership` in the identity store.

## Examples

### Declare a Identity Store Resource

The following example shows how to create a Identity Store `Group
                  Membership` resource:

#### JSON

```json

{
  "Type": "AWS::IdentityStore::GroupMembership",
  "Properties": {
    "GroupId": {
      "description": "g-1111111111-2222-3333-4444-5555-6666666656",
    },
    "IdentityStoreId": {
      "description": "d-1111111111",
    },
    "MemberId": {
      "description": "m-1111111112-2222-3333-4444-5555-6666666656",
    }
  }
```

#### YAML

```yaml

Type: AWS::IdentityStore::GroupMembership
    Properties:
      GroupId: "g-1111111111-2222-3333-4444-5555-6666666656"
      IdentityStoreId: "d-1111111111"
      MemberId: "m-1111111112-2222-3333-4444-5555-6666666656"

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IdentityStore::Group

MemberId

All content copied from https://docs.aws.amazon.com/.
