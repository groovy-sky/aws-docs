This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::UserGroup

For Valkey 7.2 and onwards, or Redis OSS 6.0 and onwards: Creates a user group. For more information, see [Using Role Based Access Control\
(RBAC)](../../../amazonelasticache/latest/dg/clusters-rbac.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::UserGroup",
  "Properties" : {
      "Engine" : String,
      "Tags" : [ Tag, ... ],
      "UserGroupId" : String,
      "UserIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::UserGroup
Properties:
  Engine: String
  Tags:
    - Tag
  UserGroupId: String
  UserIds:
    - String

```

## Properties

`Engine`

The current supported values are valkey and redis.

_Required_: Yes

_Type_: String

_Allowed values_: `redis | valkey`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticache-usergroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserGroupId`

The ID of the user group.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\\-]*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserIds`

The list of user IDs that belong to the user group. A user named `default` must be included.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref`
returns the resource name.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the user group.

`Status`

Indicates user group status. Can be "creating", "active", "modifying",
"deleting".

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
