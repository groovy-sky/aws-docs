This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway

Amazon Bedrock AgentCore Gateway provides a unified connectivity layer between agents
and the tools and resources they need to interact with.

For more information about creating a gateway, see [Set up an Amazon\
Bedrock AgentCore gateway](../../../bedrock-agentcore/latest/devguide/gateway-building.md).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::Gateway",
  "Properties" : {
      "AuthorizerConfiguration" : AuthorizerConfiguration,
      "AuthorizerType" : String,
      "Description" : String,
      "ExceptionLevel" : String,
      "InterceptorConfigurations" : [ GatewayInterceptorConfiguration, ... ],
      "KmsKeyArn" : String,
      "Name" : String,
      "PolicyEngineConfiguration" : GatewayPolicyEngineConfiguration,
      "ProtocolConfiguration" : GatewayProtocolConfiguration,
      "ProtocolType" : String,
      "RoleArn" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::Gateway
Properties:
  AuthorizerConfiguration:
    AuthorizerConfiguration
  AuthorizerType: String
  Description: String
  ExceptionLevel: String
  InterceptorConfigurations:
    - GatewayInterceptorConfiguration
  KmsKeyArn: String
  Name: String
  PolicyEngineConfiguration:
    GatewayPolicyEngineConfiguration
  ProtocolConfiguration:
    GatewayProtocolConfiguration
  ProtocolType: String
  RoleArn: String
  Tags:
    Key: Value

```

## Properties

`AuthorizerConfiguration`

Represents inbound authorization configuration options used to authenticate incoming requests.

_Required_: No

_Type_: [AuthorizerConfiguration](aws-properties-bedrockagentcore-gateway-authorizerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerType`

The type of authorizer used by the gateway.

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOM_JWT | AWS_IAM | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the gateway.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExceptionLevel`

The exception level for the gateway.

_Required_: No

_Type_: String

_Allowed values_: `DEBUG`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterceptorConfigurations`

A list of configuration settings for a gateway interceptor. Gateway interceptors allow custom code to be invoked during gateway invocations.

_Required_: No

_Type_: Array of [GatewayInterceptorConfiguration](aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The KMS key ARN for the gateway.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9-]{1,20}:kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyEngineConfiguration`

Property description not available.

_Required_: No

_Type_: [GatewayPolicyEngineConfiguration](aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolConfiguration`

The protocol configuration for the gateway target.

_Required_: No

_Type_: [GatewayProtocolConfiguration](aws-properties-bedrockagentcore-gateway-gatewayprotocolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolType`

The protocol type used by the gateway.

_Required_: Yes

_Type_: String

_Allowed values_: `MCP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that provides permissions for the gateway to access AWS services.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-]{1,20}:iam::([0-9]{12})?:role/.+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the gateway.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the gateway identifier. For example:

`my-gateway-a1b2c3d4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The date and time at which the target was created.

`GatewayArn`

The Amazon Resource Name (ARN) of the gateway target.

`GatewayIdentifier`

The unique identifier of the gateway.

`GatewayUrl`

The URL endpoint for the gateway.

`Status`

The status for the gateway.

`StatusReasons`

The status reasons for the target status.

`UpdatedAt`

The date and time at which the target was updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AuthorizerConfiguration

All content copied from https://docs.aws.amazon.com/.
