This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource OpenIdConnectAccessTokenConfiguration

The configuration of an OpenID Connect (OIDC) identity source for handling access token
claims. Contains the claim that you want to identify as the principal in an authorization
request, and the values of the `aud` claim, or audiences, that you want to
accept.

This data type is part of a [OpenIdConnectTokenSelection](../../../../reference/verifiedpermissions/latest/apireference/api-openidconnecttokenselection.md) structure, which is a parameter of [CreateIdentitySource](../../../../reference/verifiedpermissions/latest/apireference/api-createidentitysource.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Audiences" : [ String, ... ],
  "PrincipalIdClaim" : String
}

```

### YAML

```yaml

  Audiences:
    - String
  PrincipalIdClaim: String

```

## Properties

`Audiences`

The access token `aud` claim values that you want to accept in your policy
store. For example, `https://myapp.example.com, https://myapp2.example.com`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrincipalIdClaim`

The claim that determines the principal in OIDC access tokens. For example,
`sub`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdentitySourceConfiguration

OpenIdConnectConfiguration

All content copied from https://docs.aws.amazon.com/.
