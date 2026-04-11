This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider IncludedOauth2ProviderConfigInput

Configuration settings for connecting to a supported OAuth2 provider. This includes
client credentials and OAuth2 discovery information for providers that have built-in
support.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationEndpoint" : String,
  "ClientId" : String,
  "ClientSecret" : String,
  "Issuer" : String,
  "TokenEndpoint" : String
}

```

### YAML

```yaml

  AuthorizationEndpoint: String
  ClientId: String
  ClientSecret: String
  Issuer: String
  TokenEndpoint: String

```

## Properties

`AuthorizationEndpoint`

OAuth2 authorization endpoint for your isolated OAuth2 application tenant. This is where
users are redirected to authenticate and authorize access to their resources.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The client ID for the supported OAuth2 provider. This identifier is assigned by the
OAuth2 provider when you register your application.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret for the supported OAuth2 provider. This secret is assigned by the
OAuth2 provider and used along with the client ID to authenticate your application.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

Token issuer of your isolated OAuth2 application tenant. This URL identifies the
authorization server that issues tokens for this provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenEndpoint`

OAuth2 token endpoint for your isolated OAuth2 application tenant. This is where
authorization codes are exchanged for access tokens.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GoogleOauth2ProviderConfigInput

LinkedinOauth2ProviderConfigInput

All content copied from https://docs.aws.amazon.com/.
