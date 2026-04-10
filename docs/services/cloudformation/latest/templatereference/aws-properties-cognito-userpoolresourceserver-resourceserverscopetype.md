This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolResourceServer ResourceServerScopeType

One custom scope associated with a user pool resource server. This data type is a
member of `ResourceServerScopeType`. For more information, see [Scopes, M2M, and API authorization with resource servers](../../../cognito/latest/developerguide/cognito-user-pools-define-resource-servers.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScopeDescription" : String,
  "ScopeName" : String
}

```

### YAML

```yaml

  ScopeDescription: String
  ScopeName: String

```

## Properties

`ScopeDescription`

A friendly description of a custom scope.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScopeName`

The name of the scope. Amazon Cognito renders custom scopes in the format
`resourceServerIdentifier/ScopeName`. For example, if this parameter is
`exampleScope` in the resource server with the identifier
`exampleResourceServer`, you request and receive the scope
`exampleResourceServer/exampleScope`.

_Required_: Yes

_Type_: String

_Pattern_: `[\x21\x23-\x2E\x30-\x5B\x5D-\x7E]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::UserPoolResourceServer

AWS::Cognito::UserPoolRiskConfigurationAttachment

All content copied from https://docs.aws.amazon.com/.
