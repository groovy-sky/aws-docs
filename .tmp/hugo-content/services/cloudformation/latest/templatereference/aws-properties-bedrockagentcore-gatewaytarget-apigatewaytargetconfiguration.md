This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ApiGatewayTargetConfiguration

The configuration for an Amazon API Gateway target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiGatewayToolConfiguration" : ApiGatewayToolConfiguration,
  "RestApiId" : String,
  "Stage" : String
}

```

### YAML

```yaml

  ApiGatewayToolConfiguration:
    ApiGatewayToolConfiguration
  RestApiId: String
  Stage: String

```

## Properties

`ApiGatewayToolConfiguration`

The configuration for defining REST API tool filters and overrides for the gateway target.

_Required_: Yes

_Type_: [ApiGatewayToolConfiguration](aws-properties-bedrockagentcore-gatewaytarget-apigatewaytoolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestApiId`

The ID of the API Gateway REST API.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

The ID of the stage of the REST API to add as a target.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::GatewayTarget

ApiGatewayToolConfiguration

All content copied from https://docs.aws.amazon.com/.
