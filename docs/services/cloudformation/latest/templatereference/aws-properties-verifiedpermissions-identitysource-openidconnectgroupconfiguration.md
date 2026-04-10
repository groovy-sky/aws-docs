This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource OpenIdConnectGroupConfiguration

The claim in OIDC identity provider tokens that indicates a user's group membership, and
the entity type that you want to map it to. For example, this object can map the contents
of a `groups` claim to `MyCorp::UserGroup`.

This data type is part of a [OpenIdConnectConfiguration](../../../../reference/verifiedpermissions/latest/apireference/api-openidconnectconfiguration.md) structure, which is a parameter of [CreateIdentitySource](../../../../reference/verifiedpermissions/latest/apireference/api-createidentitysource.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupClaim" : String,
  "GroupEntityType" : String
}

```

### YAML

```yaml

  GroupClaim: String
  GroupEntityType: String

```

## Properties

`GroupClaim`

The token claim that you want Verified Permissions to interpret as group membership. For example,
`groups`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupEntityType`

The policy store entity type that you want to map your users' group claim to. For example,
`MyCorp::UserGroup`. A group entity type is an entity that can have a user
entity type as a member.

_Required_: Yes

_Type_: String

_Pattern_: `^([_a-zA-Z][_a-zA-Z0-9]*::)*[_a-zA-Z][_a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenIdConnectConfiguration

OpenIdConnectIdentityTokenConfiguration

All content copied from https://docs.aws.amazon.com/.
