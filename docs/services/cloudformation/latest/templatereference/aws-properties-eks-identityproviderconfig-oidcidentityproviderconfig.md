This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::IdentityProviderConfig OidcIdentityProviderConfig

An object representing the configuration for an OpenID Connect (OIDC) identity provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "GroupsClaim" : String,
  "GroupsPrefix" : String,
  "IssuerUrl" : String,
  "RequiredClaims" : [ RequiredClaim, ... ],
  "UsernameClaim" : String,
  "UsernamePrefix" : String
}

```

### YAML

```yaml

  ClientId: String
  GroupsClaim: String
  GroupsPrefix: String
  IssuerUrl: String
  RequiredClaims:
    - RequiredClaim
  UsernameClaim: String
  UsernamePrefix: String

```

## Properties

`ClientId`

This is also known as _audience_. The ID of the client application
that makes authentication requests to the OIDC identity provider.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupsClaim`

The JSON web token (JWT) claim that the provider uses to return your groups.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupsPrefix`

The prefix that is prepended to group claims to prevent clashes with existing names
(such as `system:` groups). For example, the value ` oidc:` creates
group names like `oidc:engineering` and `oidc:infra`. The prefix
can't contain `system:`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IssuerUrl`

The URL of the OIDC identity provider that allows the API server to discover public
signing keys for verifying tokens.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RequiredClaims`

The key-value pairs that describe required claims in the identity token. If set, each
claim is verified to be present in the token with a matching value.

_Required_: No

_Type_: Array of [RequiredClaim](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-identityproviderconfig-requiredclaim.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UsernameClaim`

The JSON Web token (JWT) claim that is used as the username.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UsernamePrefix`

The prefix that is prepended to username claims to prevent clashes with existing
names. The prefix can't contain `system:`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EKS::IdentityProviderConfig

RequiredClaim
