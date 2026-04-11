This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Broker LdapServerMetadata

Optional. The metadata of the LDAP server used to authenticate and authorize
connections to the broker. Does not apply to RabbitMQ brokers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hosts" : [ String, ... ],
  "RoleBase" : String,
  "RoleName" : String,
  "RoleSearchMatching" : String,
  "RoleSearchSubtree" : Boolean,
  "ServiceAccountPassword" : String,
  "ServiceAccountUsername" : String,
  "UserBase" : String,
  "UserRoleName" : String,
  "UserSearchMatching" : String,
  "UserSearchSubtree" : Boolean
}

```

### YAML

```yaml

  Hosts:
    - String
  RoleBase: String
  RoleName: String
  RoleSearchMatching: String
  RoleSearchSubtree: Boolean
  ServiceAccountPassword: String
  ServiceAccountUsername: String
  UserBase: String
  UserRoleName: String
  UserSearchMatching: String
  UserSearchSubtree: Boolean

```

## Properties

`Hosts`

Property description not available.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleBase`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleSearchMatching`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleSearchSubtree`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccountPassword`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccountUsername`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserBase`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserRoleName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserSearchMatching`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserSearchSubtree`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionOptions

LogList

All content copied from https://docs.aws.amazon.com/.
