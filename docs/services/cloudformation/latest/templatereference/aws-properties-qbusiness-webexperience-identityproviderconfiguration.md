This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::WebExperience IdentityProviderConfiguration

Provides information about the identity provider (IdP) used to authenticate end users
of an Amazon Q Business web experience.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OpenIDConnectConfiguration" : OpenIDConnectProviderConfiguration,
  "SamlConfiguration" : SamlProviderConfiguration
}

```

### YAML

```yaml

  OpenIDConnectConfiguration:
    OpenIDConnectProviderConfiguration
  SamlConfiguration:
    SamlProviderConfiguration

```

## Properties

`OpenIDConnectConfiguration`

The OIDC-compliant identity provider (IdP) used to authenticate end users of an Amazon
Q Business web experience.

_Required_: No

_Type_: [OpenIDConnectProviderConfiguration](aws-properties-qbusiness-webexperience-openidconnectproviderconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamlConfiguration`

The SAML 2.0-compliant identity provider (IdP) used to authenticate end users of an
Amazon Q Business web experience.

_Required_: No

_Type_: [SamlProviderConfiguration](aws-properties-qbusiness-webexperience-samlproviderconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomizationConfiguration

OpenIDConnectProviderConfiguration

All content copied from https://docs.aws.amazon.com/.
