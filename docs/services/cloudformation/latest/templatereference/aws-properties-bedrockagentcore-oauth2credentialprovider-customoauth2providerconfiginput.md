This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider CustomOauth2ProviderConfigInput

Input configuration for a custom OAuth2 provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ClientSecret" : String,
  "OauthDiscovery" : Oauth2Discovery
}

```

### YAML

```yaml

  ClientId: String
  ClientSecret: String
  OauthDiscovery:
    Oauth2Discovery

```

## Properties

`ClientId`

The client ID for the custom OAuth2 provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret for the custom OAuth2 provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OauthDiscovery`

The OAuth2 discovery information for the custom provider.

_Required_: Yes

_Type_: [Oauth2Discovery](aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2discovery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientSecretArn

GithubOauth2ProviderConfigInput

All content copied from https://docs.aws.amazon.com/.
