This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Integration

The `AWS::ApiGatewayV2::Integration` resource creates an integration
for an API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::Integration",
  "Properties" : {
      "ApiId" : String,
      "ConnectionId" : String,
      "ConnectionType" : String,
      "ContentHandlingStrategy" : String,
      "CredentialsArn" : String,
      "Description" : String,
      "IntegrationMethod" : String,
      "IntegrationSubtype" : String,
      "IntegrationType" : String,
      "IntegrationUri" : String,
      "PassthroughBehavior" : String,
      "PayloadFormatVersion" : String,
      "RequestParameters" : {Key: Value, ...},
      "RequestTemplates" : {Key: Value, ...},
      "ResponseParameters" : {Key: Value, ...},
      "TemplateSelectionExpression" : String,
      "TimeoutInMillis" : Integer,
      "TlsConfig" : TlsConfig
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::Integration
Properties:
  ApiId: String
  ConnectionId: String
  ConnectionType: String
  ContentHandlingStrategy: String
  CredentialsArn: String
  Description: String
  IntegrationMethod: String
  IntegrationSubtype: String
  IntegrationType: String
  IntegrationUri: String
  PassthroughBehavior: String
  PayloadFormatVersion: String
  RequestParameters:
    Key: Value
  RequestTemplates:
    Key: Value
  ResponseParameters:
    Key: Value
  TemplateSelectionExpression: String
  TimeoutInMillis: Integer
  TlsConfig:
    TlsConfig

```

## Properties

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectionId`

The ID of the VPC link for a private integration. Supported only for HTTP APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionType`

The type of the network connection to the integration endpoint. Specify `INTERNET` for connections through the public routable internet or
`VPC_LINK` for private connections between API Gateway and resources in a VPC. The default value is `INTERNET`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentHandlingStrategy`

Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT`, with the following behaviors:

`CONVERT_TO_BINARY`: Converts a response payload from a Base64-encoded string to the corresponding binary blob.

`CONVERT_TO_TEXT`: Converts a response payload from a binary blob to a Base64-encoded string.

If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialsArn`

Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string `arn:aws:iam::*:user/*`. To use resource-based permissions on supported AWS services, don't specify this parameter.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-cn|aws-us-gov):iam::[0-9]*:(role|user|group)\/.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the integration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationMethod`

Specifies the integration's HTTP method type. For WebSocket APIs, if you use a Lambda integration, you must set the integration method to `POST`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationSubtype`

Supported only for HTTP API `AWS_PROXY` integrations. Specifies the AWS service action to invoke. To learn more, see [Integration subtype reference](../../../apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationType`

The integration type of an integration. One of the following:

`AWS`: for integrating the route or method request with an AWS service action, including the Lambda function-invoking action. With the Lambda function-invoking action, this is referred to as the Lambda custom integration. With any other AWS service action, this is known as AWS integration. Supported only for WebSocket APIs.

`AWS_PROXY`: for integrating the route or method request with a Lambda function or other AWS service action. This integration is also referred to as a Lambda proxy integration.

`HTTP`: for integrating the route or method request with an HTTP endpoint. This integration is also referred to as the HTTP custom integration. Supported only for WebSocket APIs.

`HTTP_PROXY`: for integrating the route or method request with an
HTTP endpoint, with the
client request passed through as-is. This is also referred to as HTTP proxy
integration. For HTTP API private integrations, use an `HTTP_PROXY`
integration.

`MOCK`: for integrating the route or method request with API Gateway as a "loopback" endpoint without invoking any backend. Supported only for WebSocket APIs.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationUri`

For a Lambda integration, specify the URI of a Lambda function.

For an HTTP integration, specify a fully-qualified URL.

For an HTTP API private integration, specify the ARN of an Application Load
Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. If
you specify the ARN of an AWS Cloud Map service, API Gateway uses
`DiscoverInstances` to identify resources. You can use query
parameters to target specific resources. To learn more, see [DiscoverInstances](../../../cloud-map/latest/api/api-discoverinstances.md). For private integrations, all resources must be
owned by the same AWS account.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PassthroughBehavior`

Specifies the pass-through behavior for incoming requests based on the `Content-Type` header in the request, and the available mapping templates specified as the `requestTemplates` property on the `Integration` resource. There are three valid values: `WHEN_NO_MATCH`, `WHEN_NO_TEMPLATES`, and `NEVER`. Supported only for WebSocket APIs.

`WHEN_NO_MATCH` passes the request body for unmapped content types through to the integration backend without transformation.

`NEVER` rejects unmapped content types with an `HTTP 415 Unsupported Media Type` response.

`WHEN_NO_TEMPLATES` allows pass-through when the integration has no content types mapped to templates. However, if there is at least one content type defined, unmapped content types will be rejected with the same `HTTP 415 Unsupported Media Type` response.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadFormatVersion`

Specifies the format of the payload sent to an integration. Required for HTTP
APIs. For HTTP APIs, supported values for Lambda proxy integrations are
`1.0` and `2.0`. For all other integrations,
`1.0` is the only supported value. To learn more, see
[Working with AWS Lambda proxy integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-lambda.md).

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestParameters`

For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the backend. The method request parameter value must match the pattern of `method.request.{location}.{name}`, where `
                            {location}
                        ` is `querystring`, `path`, or `header`; and `
                            {name}
                        ` must be a valid and unique method request parameter name.

For HTTP API integrations with a specified `integrationSubtype`, request parameters are a key-value map specifying parameters that are passed to `AWS_PROXY` integrations. You can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Working with AWS service integrations for HTTP APIs](../../../apigateway/latest/developerguide/http-api-develop-integrations-aws-services.md).

For HTTP API integrations without a specified `integrationSubtype` request parameters are a key-value map specifying how to transform HTTP requests before sending them to the backend. The key should follow the pattern <action>:<header\|querystring\|path>.<location> where action can be `append`, `overwrite` or ` remove`. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestTemplates`

Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseParameters`

Supported only for HTTP APIs. You use response parameters to transform the HTTP response from a backend
integration before returning the response to clients. Specify a key-value map from a selection key to response
parameters. The selection key must be a valid HTTP status code within the range of 200-599. The value is of type [`ResponseParameterList`](../userguide/aws-properties-apigatewayv2-integration-responseparameterlist.md). To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

_Required_: No

_Type_: Object of [ResponseParameterMap](aws-properties-apigatewayv2-integration-responseparametermap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateSelectionExpression`

The template selection expression for the integration. Supported only for WebSocket APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInMillis`

Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and
between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29
seconds for WebSocket APIs and 30 seconds for HTTP APIs.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TlsConfig`

The TLS configuration for a private integration. If you specify a TLS configuration, private integration traffic uses the HTTPS
protocol. Supported only for HTTP APIs.

_Required_: No

_Type_: [TlsConfig](aws-properties-apigatewayv2-integration-tlsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Integration resource ID, such as
`abcd123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`IntegrationId`

The integration ID.

## Examples

- [Integration creation example](#aws-resource-apigatewayv2-integration--examples--Integration_creation_example)

- [Integration with parameter mapping for an HTTP API](#aws-resource-apigatewayv2-integration--examples--Integration_with_parameter_mapping_for_an_HTTP_API)

### Integration creation example

The following example creates a Lambda integration for an HTTP API. For full examples, see [example CloudFormation templates](https://github.com/awsdocs/amazon-api-gateway-developer-guide/tree/main/cloudformation-templates) on GitHub.

#### JSON

```json

{
    "Integration": {
        "Type": "AWS::ApiGatewayV2::Integration",
        "Properties": {
            "ApiId": {
                "Ref": "HTTPApi"
            },
            "Description": "Lambda Integration",
            "IntegrationType": "AWS_PROXY",
            "IntegrationUri": {
                "Fn::Join": [
                    "",
                    [
                        "arn:",
                        {
                            "Ref": "AWS::Partition"
                        },
                        ":apigateway:",
                        {
                            "Ref": "AWS::Region"
                        },
                        ":lambda:path/2015-03-31/functions/",
                        {
                            "Fn::GetAtt": [
                                "MyLambdaFunction",
                                "Arn"
                            ]
                        },
                        "/invocations"
                    ]
                ]
            },
            "IntegrationMethod": "POST",
            "PayloadFormatVersion": "2.0"
        }
    }
}
```

#### YAML

```yaml

Integration:
  Type: 'AWS::ApiGatewayV2::Integration'
  Properties:
    ApiId: !Ref HTTPApi
    Description: Lambda Integration
    IntegrationType: AWS_PROXY
    IntegrationUri: !Join
      - ''
      - - 'arn:'
        - !Ref 'AWS::Partition'
        - ':apigateway:'
        - !Ref 'AWS::Region'
        - ':lambda:path/2015-03-31/functions/'
        - !GetAtt MyLambdaFunction.Arn
        - /invocations
    IntegrationMethod: POST
    PayloadFormatVersion: '2.0'

```

### Integration with parameter mapping for an HTTP API

The following example creates an integration with parameter mapping. The request parameters
add a header named `header1` to the request before it reaches the backend integration. The response
parameters add a header to the integration's response named `header2`, with the static value
`headervalue`, when the integration returns a 200 status code.

#### JSON

```json

{
  "MyIntegration": {
    "Type": "AWS::ApiGatewayV2::Integration",
    "Properties": {
      "ApiId": {
        "Ref": "MyApi"
      },
      "Description": "HTTP proxy integration",
      "IntegrationType": "HTTP_PROXY",
      "IntegrationMethod": "ANY",
      "IntegrationUri": "https://api.example.com",
      "PayloadFormatVersion": "1.0",
      "RequestParameters": {
        "append:header.header1": "$context.requestId"
      },
        "ResponseParameters": {
        "200": {
          "ResponseParameters": [
            {
              "Source": "headervalue",
              "Destination": "append:header.header2"
            }
          ]
        }
      }
    }
  }
}
```

#### YAML

```yaml

MyIntegration:
Type: AWS::ApiGatewayV2::Integration
Properties:
  ApiId: !Ref MyApi
  Description: HTTP proxy integration
  IntegrationType: HTTP_PROXY
  IntegrationMethod: ANY
  IntegrationUri: https://api.example.com
  PayloadFormatVersion: 1.0
  RequestParameters:
    "append:header.header1": "$context.requestId"
   ResponseParameters:
    "200":
      ResponseParameters:
        - Source: "headerValue"
          Destination: "append:header.header2"
```

## See also

- [CreateIntegration](../../../apigatewayv2/latest/api-reference/apis-apiid-integrations.md#CreateIntegration) in the _Amazon API_
_Gateway Version 2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MutualTlsAuthentication

ResponseParameter

All content copied from https://docs.aws.amazon.com/.
