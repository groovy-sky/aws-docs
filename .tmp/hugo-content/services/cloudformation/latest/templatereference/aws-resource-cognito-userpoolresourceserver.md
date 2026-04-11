This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolResourceServer

The `AWS::Cognito::UserPoolResourceServer` resource creates a new OAuth2.0
resource server and defines custom scopes in it.

###### Note

If you don't specify a value for a parameter, Amazon Cognito sets it to a default
value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolResourceServer",
  "Properties" : {
      "Identifier" : String,
      "Name" : String,
      "Scopes" : [ ResourceServerScopeType, ... ],
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolResourceServer
Properties:
  Identifier: String
  Name: String
  Scopes:
    - ResourceServerScopeType
  UserPoolId: String

```

## Properties

`Identifier`

A unique resource server identifier for the resource server. The identifier can be an
API friendly name like `solar-system-data`. You can also set an API URL like
`https://solar-system-data-api.example.com` as your identifier.

Amazon Cognito represents scopes in the access token in the format
`$resource-server-identifier/$scope`. Longer scope-identifier strings
increase the size of your access tokens.

_Required_: Yes

_Type_: String

_Pattern_: `[\x21\x23-\x5B\x5D-\x7E]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A friendly name for the resource server.

_Required_: Yes

_Type_: String

_Pattern_: `[\w\s+=,.@-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scopes`

A list of scopes. Each scope is a map with keys `ScopeName` and
`ScopeDescription`.

_Required_: No

_Type_: Array of [ResourceServerScopeType](aws-properties-cognito-userpoolresourceserver-resourceserverscopetype.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool where you want to create a resource server.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns physicalResourceId, which is the resource server
identifier “Identifier". For example:

`{ "Ref": "yourResourceServerIdentifier" }`

For the Amazon Cognito resource server `yourResourceServerIdentifier`, Ref
returns the name of the resource server.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating a new resource server for a user pool

The following example creates a resource server "Name" with the identifier
"Identifier" in the referenced user pool.

#### JSON

```json

{
  "UserPoolResourceServer": {
    "Type": "AWS::Cognito::UserPoolResourceServer",
    "Properties": {
      "UserPoolId": {
        "Ref": "UserPool"
      },
      "Identifier": "Identifier",
      "Name": "Name",
      "Scopes": [{
        "ScopeName": "ScopeName1",
        "ScopeDescription": "description"
      }, {
        "ScopeName": "ScopeName2",
        "ScopeDescription": "description"
      }]
    }
  }
}
```

#### YAML

```yaml

UserPoolResourceServer:
  Type: AWS::Cognito::UserPoolResourceServer
  Properties:
    UserPoolId: !Ref UserPool
    Identifier: "Identifier"
    Name: "Name"
    Scopes:
     - ScopeName: "ScopeName1"
       ScopeDescription: "description"
     - ScopeName: "ScopeName2"
       ScopeDescription: "description"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::UserPoolIdentityProvider

ResourceServerScopeType

All content copied from https://docs.aws.amazon.com/.
