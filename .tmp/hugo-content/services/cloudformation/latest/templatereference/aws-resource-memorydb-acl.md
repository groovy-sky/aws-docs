This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MemoryDB::ACL

Specifies an Access Control List. For more information, see [Authenticating users with Access\
Contol Lists (ACLs)](../../../memorydb/latest/devguide/clusters-acls.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MemoryDB::ACL",
  "Properties" : {
      "ACLName" : String,
      "Tags" : [ Tag, ... ],
      "UserNames" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MemoryDB::ACL
Properties:
  ACLName: String
  Tags:
    - Tag
  UserNames:
    - String

```

## Properties

`ACLName`

The name of the Access Control List. This value is stored as a lowercase string.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\\-]*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-memorydb-acl-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserNames`

The list of users that belong to the Access Control List.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Arn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ARN of the Access Control List,
such as `arn:aws:memorydb:us-east-1:123456789012:acl/my-acl`

`Status`

Indicates ACL status.

_Valid values_: `creating` \| `active` \|
`modifying` \| `deleting`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon MemoryDB

Tag

All content copied from https://docs.aws.amazon.com/.
