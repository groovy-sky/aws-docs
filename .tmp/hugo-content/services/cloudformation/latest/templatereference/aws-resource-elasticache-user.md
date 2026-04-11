This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::User

For Valkey 7.2 and onwards, or Redis OSS engine version 6.0 and onwards: Creates user. For more information, see [Using Role Based Access Control\
(RBAC)](../../../amazonelasticache/latest/dg/clusters-rbac.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::User",
  "Properties" : {
      "AccessString" : String,
      "AuthenticationMode" : AuthenticationMode,
      "Engine" : String,
      "NoPasswordRequired" : Boolean,
      "Passwords" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "UserId" : String,
      "UserName" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::User
Properties:
  AccessString:
    String
  AuthenticationMode:
    AuthenticationMode
  Engine: String
  NoPasswordRequired: Boolean
  Passwords:
    - String
  Tags:
    - Tag
  UserId: String
  UserName: String

```

## Properties

`AccessString`

Access permissions string used for this user.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationMode`

Specifies the authentication mode to use. Below is an example of the possible JSON values:

```

{
 Passwords: ["*****", "******"] // If Type is password.
}
```

_Required_: No

_Type_: [AuthenticationMode](aws-properties-elasticache-user-authenticationmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The current supported values are valkey and redis.

_Required_: Yes

_Type_: String

_Allowed values_: `redis | valkey`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoPasswordRequired`

Indicates a password is not required for this user.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Passwords`

Passwords used for this user. You can create up to two passwords for each user.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticache-user-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserId`

The ID of the user.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\\-]*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserName`

The username of the user.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref`
returns the resource name.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the user.

`Status`

Indicates the user status. Can be "active", "modifying" or "deleting".

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AuthenticationMode

All content copied from https://docs.aws.amazon.com/.
