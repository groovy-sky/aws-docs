This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessTrustProvider OidcOptions

Describes the options for an OpenID Connect-compatible user-identity trust
provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationEndpoint" : String,
  "ClientId" : String,
  "ClientSecret" : String,
  "Issuer" : String,
  "Scope" : String,
  "TokenEndpoint" : String,
  "UserInfoEndpoint" : String
}

```

### YAML

```yaml

  AuthorizationEndpoint: String
  ClientId: String
  ClientSecret: String
  Issuer: String
  Scope: String
  TokenEndpoint: String
  UserInfoEndpoint: String

```

## Properties

`AuthorizationEndpoint`

The OIDC authorization endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The client identifier.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The OIDC issuer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The OpenID Connect (OIDC) scope specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenEndpoint`

The OIDC token endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserInfoEndpoint`

The OIDC user info endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NativeApplicationOidcOptions

SseSpecification

All content copied from https://docs.aws.amazon.com/.
