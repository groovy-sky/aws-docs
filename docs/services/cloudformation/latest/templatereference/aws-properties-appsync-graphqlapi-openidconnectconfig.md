This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLApi OpenIDConnectConfig

The `OpenIDConnectConfig` property type specifies the optional
authorization configuration for using an OpenID Connect compliant service with your
GraphQL endpoint for an AWS AppSync GraphQL API.

`OpenIDConnectConfig` is a property of the [AWS::AppSync::GraphQLApi](../userguide/aws-resource-appsync-graphqlapi.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthTTL" : Number,
  "ClientId" : String,
  "IatTTL" : Number,
  "Issuer" : String
}

```

### YAML

```yaml

  AuthTTL: Number
  ClientId: String
  IatTTL: Number
  Issuer: String

```

## Properties

`AuthTTL`

The number of milliseconds that a token is valid after being authenticated.

_Required_: No

_Type_: Number

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ClientId`

The client identifier of the Relying party at the OpenID identity provider. This
identifier is typically obtained when the Relying party is registered with the OpenID
identity provider. You can specify a regular expression so that AWS AppSync can validate against multiple client identifiers at a time.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`IatTTL`

The number of milliseconds that a token is valid after it's issued to a user.

_Required_: No

_Type_: Number

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Issuer`

The issuer for the OIDC configuration. The issuer returned by discovery must exactly match the value of
`iss` in the ID token.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfig

Tag

All content copied from https://docs.aws.amazon.com/.
