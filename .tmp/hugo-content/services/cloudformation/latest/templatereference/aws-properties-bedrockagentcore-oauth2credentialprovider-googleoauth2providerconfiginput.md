This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider GoogleOauth2ProviderConfigInput

Input configuration for a Google OAuth2 provider.

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

The client ID for the Google OAuth2 provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret for the Google OAuth2 provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GithubOauth2ProviderConfigInput

IncludedOauth2ProviderConfigInput

All content copied from https://docs.aws.amazon.com/.
