This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service MCPServerSplunkAuthorizationConfig

The authorization configuration for a Splunk MCP server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BearerToken" : BearerTokenDetails
}

```

### YAML

```yaml

  BearerToken:
    BearerTokenDetails

```

## Properties

`BearerToken`

The bearer token details for authenticating with the Splunk MCP server.

_Required_: Yes

_Type_: [BearerTokenDetails](aws-properties-devopsagent-service-bearertokendetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MCPServerOAuthClientCredentialsConfig

MCPServerSplunkDetails

All content copied from https://docs.aws.amazon.com/.
