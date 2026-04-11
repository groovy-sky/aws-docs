This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service MCPServerAuthorizationConfig

The authorization configuration for a custom MCP server. Specify either OAuth client credentials or an API
key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKey" : ApiKeyDetails,
  "OAuthClientCredentials" : MCPServerOAuthClientCredentialsConfig
}

```

### YAML

```yaml

  ApiKey:
    ApiKeyDetails
  OAuthClientCredentials:
    MCPServerOAuthClientCredentialsConfig

```

## Properties

`ApiKey`

The API key details for authenticating with the MCP server.

_Required_: No

_Type_: [ApiKeyDetails](aws-properties-devopsagent-service-apikeydetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OAuthClientCredentials`

The OAuth client credentials for authenticating with the MCP server.

_Required_: No

_Type_: [MCPServerOAuthClientCredentialsConfig](aws-properties-devopsagent-service-mcpserveroauthclientcredentialsconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitLabDetails

MCPServerDetails

All content copied from https://docs.aws.amazon.com/.
