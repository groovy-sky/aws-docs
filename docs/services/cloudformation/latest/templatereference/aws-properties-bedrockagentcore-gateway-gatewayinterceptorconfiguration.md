This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway GatewayInterceptorConfiguration

The configuration for an interceptor on a gateway. This structure defines settings for an interceptor that will be invoked during the invocation of the gateway.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputConfiguration" : InterceptorInputConfiguration,
  "InterceptionPoints" : [ String, ... ],
  "Interceptor" : InterceptorConfiguration
}

```

### YAML

```yaml

  InputConfiguration:
    InterceptorInputConfiguration
  InterceptionPoints:
    - String
  Interceptor:
    InterceptorConfiguration

```

## Properties

`InputConfiguration`

The configuration for the input of the interceptor. This field specifies how the input to the interceptor is constructed

_Required_: No

_Type_: [InterceptorInputConfiguration](aws-properties-bedrockagentcore-gateway-interceptorinputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterceptionPoints`

The supported points of interception. This field specifies which points during the gateway invocation to invoke the interceptor

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interceptor`

The infrastructure settings of an interceptor configuration. This structure defines how the interceptor can be invoked.

_Required_: Yes

_Type_: [InterceptorConfiguration](aws-properties-bedrockagentcore-gateway-interceptorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomJWTAuthorizerConfiguration

GatewayPolicyEngineConfiguration

All content copied from https://docs.aws.amazon.com/.
