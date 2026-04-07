This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MemoryDB::User

Specifies a MemoryDB user. For more information, see [Authenticating\
users with Access Contol Lists (ACLs)](https://docs.aws.amazon.com/memorydb/latest/devguide/clusters.acls.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MemoryDB::User",
  "Properties" : {
      "AccessString" : String,
      "AuthenticationMode" : AuthenticationMode,
      "Tags" : [ Tag, ... ],
      "UserName" : String
    }
}

```

### YAML

```yaml

Type: AWS::MemoryDB::User
Properties:
  AccessString:
    String
  AuthenticationMode:
    AuthenticationMode
  Tags:
    - Tag
  UserName: String

```

## Properties

`AccessString`

Access permissions string used for this user.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationMode`

Denotes whether the user requires a password to authenticate.

**Example:**

```json

mynewdbuser:
     Type: AWS::MemoryDB::User
     Properties:
     AccessString: on ~* &* +@all
     AuthenticationMode:
         Passwords: '1234567890123456'
         Type: password
     UserName: mynewdbuser

     AuthenticationMode:
     {
         "Passwords": ["1234567890123456"],
         "Type": "Password"
     }
```

_Required_: No

_Type_: [AuthenticationMode](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-memorydb-user-authenticationmode.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-memorydb-user-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserName`

The name of the user.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\\-]*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

`Arn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ARN of the user,
such as `arn:aws:memorydb:us-east-1:123456789012:user/user1`

`Status`

Indicates the user status.

_Valid values_: `active` \| `modifying` \|
`deleting`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AuthenticationMode
