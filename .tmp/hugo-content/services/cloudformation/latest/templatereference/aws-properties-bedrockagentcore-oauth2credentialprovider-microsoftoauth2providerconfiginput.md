This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider MicrosoftOauth2ProviderConfigInput

Input configuration for a Microsoft OAuth2 provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ClientSecret" : String,
  "TenantId" : String
}

```

### YAML

```yaml

  ClientId: String
  ClientSecret: String
  TenantId: String

```

## Properties

`ClientId`

The client ID for the Microsoft OAuth2 provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret for the Microsoft OAuth2 provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TenantId`

The Microsoft Entra ID (formerly Azure AD) tenant ID for your organization. This
identifies the specific tenant within Microsoft's identity platform where your application
is registered.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LinkedinOauth2ProviderConfigInput

Oauth2AuthorizationServerMetadata

All content copied from https://docs.aws.amazon.com/.
