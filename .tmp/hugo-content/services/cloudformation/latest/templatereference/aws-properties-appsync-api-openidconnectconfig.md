This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Api OpenIDConnectConfig

Describes an OpenID Connect (OIDC) configuration.

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

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The client identifier of the relying party at the OpenID identity provider. This identifier is typically
obtained when the relying party is registered with the OpenID identity provider. You can specify a regular
expression so that AWS AppSync can validate against multiple client identifiers at a time.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IatTTL`

The number of milliseconds that a token is valid after it's issued to a user.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The issuer for the OIDC configuration. The issuer returned by discovery must exactly match the value of
`iss` in the ID token.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaAuthorizerConfig

Tag

All content copied from https://docs.aws.amazon.com/.
