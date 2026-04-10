This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget McpTargetConfiguration

The Model Context Protocol (MCP) configuration for a target. This structure defines how the gateway uses MCP to communicate with the target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiGateway" : ApiGatewayTargetConfiguration,
  "Lambda" : McpLambdaTargetConfiguration,
  "McpServer" : McpServerTargetConfiguration,
  "OpenApiSchema" : ApiSchemaConfiguration,
  "SmithyModel" : ApiSchemaConfiguration
}

```

### YAML

```yaml

  ApiGateway:
    ApiGatewayTargetConfiguration
  Lambda:
    McpLambdaTargetConfiguration
  McpServer:
    McpServerTargetConfiguration
  OpenApiSchema:
    ApiSchemaConfiguration
  SmithyModel:
    ApiSchemaConfiguration

```

## Properties

`ApiGateway`

The configuration for an Amazon API Gateway target.

_Required_: No

_Type_: [ApiGatewayTargetConfiguration](aws-properties-bedrockagentcore-gatewaytarget-apigatewaytargetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Lambda`

The Lambda configuration for the Model Context Protocol target. This configuration defines how the gateway uses a Lambda function to communicate with the target.

_Required_: No

_Type_: [McpLambdaTargetConfiguration](aws-properties-bedrockagentcore-gatewaytarget-mcplambdatargetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`McpServer`

The configuration for an MCP server target.

_Required_: No

_Type_: [McpServerTargetConfiguration](aws-properties-bedrockagentcore-gatewaytarget-mcpservertargetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenApiSchema`

The OpenAPI schema for the Model Context Protocol target. This schema defines the API structure of the target.

_Required_: No

_Type_: [ApiSchemaConfiguration](aws-properties-bedrockagentcore-gatewaytarget-apischemaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmithyModel`

The Smithy model for the Model Context Protocol target. This model defines the API structure of the target using the Smithy specification.

_Required_: No

_Type_: [ApiSchemaConfiguration](aws-properties-bedrockagentcore-gatewaytarget-apischemaconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

McpServerTargetConfiguration

MetadataConfiguration

All content copied from https://docs.aws.amazon.com/.
