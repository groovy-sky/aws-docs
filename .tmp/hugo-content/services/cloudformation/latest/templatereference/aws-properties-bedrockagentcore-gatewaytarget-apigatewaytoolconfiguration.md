This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ApiGatewayToolConfiguration

The configuration for defining REST API tool filters and overrides for the gateway target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ToolFilters" : [ ApiGatewayToolFilter, ... ],
  "ToolOverrides" : [ ApiGatewayToolOverride, ... ]
}

```

### YAML

```yaml

  ToolFilters:
    - ApiGatewayToolFilter
  ToolOverrides:
    - ApiGatewayToolOverride

```

## Properties

`ToolFilters`

A list of path and method patterns to expose as tools using metadata from the REST API's OpenAPI specification.

_Required_: Yes

_Type_: Array of [ApiGatewayToolFilter](aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolOverrides`

A list of explicit tool definitions with optional custom names and descriptions.

_Required_: No

_Type_: Array of [ApiGatewayToolOverride](aws-properties-bedrockagentcore-gatewaytarget-apigatewaytooloverride.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiGatewayTargetConfiguration

ApiGatewayToolFilter

All content copied from https://docs.aws.amazon.com/.
