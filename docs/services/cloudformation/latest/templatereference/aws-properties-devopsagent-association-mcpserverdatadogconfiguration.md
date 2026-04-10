This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association MCPServerDatadogConfiguration

Configuration for Datadog MCP server integration. Defines the server name, endpoint URL, optional description,
and webhook update settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "EnableWebhookUpdates" : Boolean,
  "Endpoint" : String,
  "Name" : String
}

```

### YAML

```yaml

  Description: String
  EnableWebhookUpdates: Boolean
  Endpoint: String
  Name: String

```

## Properties

`Description`

The description of the MCP server.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableWebhookUpdates`

When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events
from the service.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The MCP server endpoint URL. Must be an HTTPS URL.

_Required_: Yes

_Type_: String

_Pattern_: `^https://[a-zA-Z0-9.-]+(?::[0-9]+)?(?:/.*)?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the MCP server. The name must match the pattern `^[a-zA-Z0-9_-]+$`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MCPServerConfiguration

MCPServerNewRelicConfiguration

All content copied from https://docs.aws.amazon.com/.
