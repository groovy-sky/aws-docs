This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::OAuth2CredentialProvider Oauth2Discovery

Contains the discovery information for an OAuth2 provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationServerMetadata" : Oauth2AuthorizationServerMetadata,
  "DiscoveryUrl" : String
}

```

### YAML

```yaml

  AuthorizationServerMetadata:
    Oauth2AuthorizationServerMetadata
  DiscoveryUrl: String

```

## Properties

`AuthorizationServerMetadata`

The authorization server metadata for the OAuth2 provider.

_Required_: No

_Type_: [Oauth2AuthorizationServerMetadata](aws-properties-bedrockagentcore-oauth2credentialprovider-oauth2authorizationservermetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DiscoveryUrl`

The discovery URL for the OAuth2 provider.

_Required_: No

_Type_: String

_Pattern_: `^.+/\.well-known/openid-configuration$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Oauth2AuthorizationServerMetadata

Oauth2ProviderConfigInput

All content copied from https://docs.aws.amazon.com/.
