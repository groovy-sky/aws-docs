This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource OpenIdConnectTokenSelection

The token type that you want to process from your OIDC identity provider. Your policy
store can process either identity (ID) or access tokens from a given OIDC identity
source.

This data type is part of a [OpenIdConnectConfiguration](../../../../reference/verifiedpermissions/latest/apireference/api-openidconnectconfiguration.md) structure, which is a parameter of [CreateIdentitySource](../../../../reference/verifiedpermissions/latest/apireference/api-createidentitysource.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessTokenOnly" : OpenIdConnectAccessTokenConfiguration,
  "IdentityTokenOnly" : OpenIdConnectIdentityTokenConfiguration
}

```

### YAML

```yaml

  AccessTokenOnly:
    OpenIdConnectAccessTokenConfiguration
  IdentityTokenOnly:
    OpenIdConnectIdentityTokenConfiguration

```

## Properties

`AccessTokenOnly`

The OIDC configuration for processing access tokens. Contains allowed audience claims,
for example `https://auth.example.com`, and the claim that you want to map to the
principal, for example `sub`.

_Required_: No

_Type_: [OpenIdConnectAccessTokenConfiguration](aws-properties-verifiedpermissions-identitysource-openidconnectaccesstokenconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityTokenOnly`

The OIDC configuration for processing identity (ID) tokens. Contains allowed client ID
claims, for example `1example23456789`, and the claim that you want to map to
the principal, for example `sub`.

_Required_: No

_Type_: [OpenIdConnectIdentityTokenConfiguration](aws-properties-verifiedpermissions-identitysource-openidconnectidentitytokenconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenIdConnectIdentityTokenConfiguration

AWS::VerifiedPermissions::Policy

All content copied from https://docs.aws.amazon.com/.
