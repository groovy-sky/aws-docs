This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource OpenIdConnectConfiguration

Contains configuration details of an OpenID Connect (OIDC) identity provider, or
identity source, that Verified Permissions can use to generate entities from authenticated identities. It
specifies the issuer URL, token type that you want to use, and policy store entity
details.

This data type is part of a [Configuration](../../../../reference/verifiedpermissions/latest/apireference/api-configuration.md) structure, which
is a parameter to [CreateIdentitySource](../../../../reference/verifiedpermissions/latest/apireference/api-createidentitysource.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EntityIdPrefix" : String,
  "GroupConfiguration" : OpenIdConnectGroupConfiguration,
  "Issuer" : String,
  "TokenSelection" : OpenIdConnectTokenSelection
}

```

### YAML

```yaml

  EntityIdPrefix: String
  GroupConfiguration:
    OpenIdConnectGroupConfiguration
  Issuer: String
  TokenSelection:
    OpenIdConnectTokenSelection

```

## Properties

`EntityIdPrefix`

A descriptive string that you want to prefix to user entities from your OIDC identity
provider. For example, if you set an `entityIdPrefix` of
`MyOIDCProvider`, you can reference principals in your policies in the format
`MyCorp::User::MyOIDCProvider|Carlos`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupConfiguration`

The claim in OIDC identity provider tokens that indicates a user's group membership, and
the entity type that you want to map it to. For example, this object can map the contents
of a `groups` claim to `MyCorp::UserGroup`.

_Required_: No

_Type_: [OpenIdConnectGroupConfiguration](aws-properties-verifiedpermissions-identitysource-openidconnectgroupconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The issuer URL of an OIDC identity provider. This URL must have an OIDC discovery
endpoint at the path `.well-known/openid-configuration`.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenSelection`

The token type that you want to process from your OIDC identity provider. Your policy
store can process either identity (ID) or access tokens from a given OIDC identity
source.

_Required_: Yes

_Type_: [OpenIdConnectTokenSelection](aws-properties-verifiedpermissions-identitysource-openidconnecttokenselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenIdConnectAccessTokenConfiguration

OpenIdConnectGroupConfiguration

All content copied from https://docs.aws.amazon.com/.
