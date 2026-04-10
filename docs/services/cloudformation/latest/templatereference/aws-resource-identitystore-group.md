This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IdentityStore::Group

A group object, which contains a specified group’s metadata and attributes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IdentityStore::Group",
  "Properties" : {
      "Description" : String,
      "DisplayName" : String,
      "IdentityStoreId" : String
    }
}

```

### YAML

```yaml

Type: AWS::IdentityStore::Group
Properties:
  Description: String
  DisplayName: String
  IdentityStoreId: String

```

## Properties

`Description`

A string containing the description of the group.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}\p{M}\p{S}\p{N}\p{P}\t\n\r  　]+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name value for the group. The length limit is 1,024 characters. This value
can consist of letters, accented characters, symbols, numbers, punctuation, tab, new line,
carriage return, space, and nonbreaking space in this attribute. This value is specified at
the time the group is created and stored as an attribute of the group object in the
identity store.

Prefix search supports a maximum of 1,000 characters for the string.

_Required_: Yes

_Type_: String

_Pattern_: `^[\p{L}\p{M}\p{S}\p{N}\p{P}\t\n\r  ]+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityStoreId`

The globally unique identifier for the identity store.

_Required_: Yes

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$|^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the `Physical ID` of the resource created
which is the format `GroupID|IdentityStoreId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GroupId`

The identifier of the newly created group in the identity store.

## Examples

### Declare a Identity Store Resource

The following example shows how to create a Identity Store `Group`
resource:

#### JSON

```json

{
  "Type": "AWS::IdentityStore::Group",
  "Properties": {
    "Description": {
      "description": "Group for developers",
    },
    "DisplayName": {
      "description": "Developers",
    },
    "IdentityStoreId": {
      "description": "d-1111111111",

    }
}
```

#### YAML

```yaml

Type: AWS::IdentityStore::Group
    Properties:
       Description: "Group for Developers"
       DisplayName: "Developers"
       IdentityStoreId: "d-1111111111a"

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Identity Store

AWS::IdentityStore::GroupMembership

All content copied from https://docs.aws.amazon.com/.
