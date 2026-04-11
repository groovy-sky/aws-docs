This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessTrustProvider NativeApplicationOidcOptions

Describes the OpenID Connect (OIDC) options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationEndpoint" : String,
  "ClientId" : String,
  "ClientSecret" : String,
  "Issuer" : String,
  "PublicSigningKeyEndpoint" : String,
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
  PublicSigningKeyEndpoint: String
  Scope: String
  TokenEndpoint: String
  UserInfoEndpoint: String

```

## Properties

`AuthorizationEndpoint`

The authorization endpoint of the IdP.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The OAuth 2.0 client identifier.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The OAuth 2.0 client secret.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The OIDC issuer identifier of the IdP.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicSigningKeyEndpoint`

The public signing key endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The set of user claims to be requested from the IdP.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenEndpoint`

The token endpoint of the IdP.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserInfoEndpoint`

The user info endpoint of the IdP.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeviceOptions

OidcOptions

All content copied from https://docs.aws.amazon.com/.
