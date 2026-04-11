This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider Oauth2AuthorizationServerMetadata

Contains the authorization server metadata for an OAuth2 provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationEndpoint" : String,
  "Issuer" : String,
  "ResponseTypes" : [ String, ... ],
  "TokenEndpoint" : String
}

```

### YAML

```yaml

  AuthorizationEndpoint: String
  Issuer: String
  ResponseTypes:
    - String
  TokenEndpoint: String

```

## Properties

`AuthorizationEndpoint`

The authorization endpoint URL for the OAuth2 authorization server.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The issuer URL for the OAuth2 authorization server.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseTypes`

The supported response types for the OAuth2 authorization server.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenEndpoint`

The token endpoint URL for the OAuth2 authorization server.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MicrosoftOauth2ProviderConfigInput

Oauth2Discovery

All content copied from https://docs.aws.amazon.com/.
