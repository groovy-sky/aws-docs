This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider LinkedinOauth2ProviderConfigInput

Configuration settings for connecting to LinkedIn services using OAuth2 authentication.
This includes the client credentials required to authenticate with LinkedIn's OAuth2
authorization server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ClientSecret" : String
}

```

### YAML

```yaml

  ClientId: String
  ClientSecret: String

```

## Properties

`ClientId`

The client ID for the LinkedIn OAuth2 provider. This identifier is assigned by LinkedIn
when you register your application.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret for the LinkedIn OAuth2 provider. This secret is assigned by LinkedIn
and used along with the client ID to authenticate your application.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IncludedOauth2ProviderConfigInput

MicrosoftOauth2ProviderConfigInput

All content copied from https://docs.aws.amazon.com/.
