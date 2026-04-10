This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget MetadataConfiguration

Configuration for HTTP header and query parameter propagation between the gateway and target servers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedQueryParameters" : [ String, ... ],
  "AllowedRequestHeaders" : [ String, ... ],
  "AllowedResponseHeaders" : [ String, ... ]
}

```

### YAML

```yaml

  AllowedQueryParameters:
    - String
  AllowedRequestHeaders:
    - String
  AllowedResponseHeaders:
    - String

```

## Properties

`AllowedQueryParameters`

A list of URL query parameters that are allowed to be propagated from incoming gateway URL to the target.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedRequestHeaders`

A list of HTTP headers that are allowed to be propagated from incoming client requests to the target.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedResponseHeaders`

A list of HTTP headers that are allowed to be propagated from the target response back to the client.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

McpTargetConfiguration

OAuthCredentialProvider

All content copied from https://docs.aws.amazon.com/.
