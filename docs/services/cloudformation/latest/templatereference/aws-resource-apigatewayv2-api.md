This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Api

The `AWS::ApiGatewayV2::Api` resource creates an API. WebSocket APIs
and HTTP APIs are supported. For more information about
WebSocket APIs, see [About WebSocket APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html) in the _API Gateway_
_Developer Guide_. For more information about HTTP APIs, see [HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api.html) in the _API Gateway Developer_
_Guide._

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::Api",
  "Properties" : {
      "ApiKeySelectionExpression" : String,
      "BasePath" : String,
      "Body" : Json,
      "BodyS3Location" : BodyS3Location,
      "CorsConfiguration" : Cors,
      "CredentialsArn" : String,
      "Description" : String,
      "DisableExecuteApiEndpoint" : Boolean,
      "DisableSchemaValidation" : Boolean,
      "FailOnWarnings" : Boolean,
      "IpAddressType" : String,
      "Name" : String,
      "ProtocolType" : String,
      "RouteKey" : String,
      "RouteSelectionExpression" : String,
      "Tags" : {Key: Value, ...},
      "Target" : String,
      "Version" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::Api
Properties:
  ApiKeySelectionExpression: String
  BasePath: String
  Body: Json
  BodyS3Location:
    BodyS3Location
  CorsConfiguration:
    Cors
  CredentialsArn: String
  Description: String
  DisableExecuteApiEndpoint: Boolean
  DisableSchemaValidation: Boolean
  FailOnWarnings: Boolean
  IpAddressType: String
  Name: String
  ProtocolType: String
  RouteKey: String
  RouteSelectionExpression: String
  Tags:
    Key: Value
  Target: String
  Version: String

```

## Properties

`ApiKeySelectionExpression`

An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasePath`

Specifies how to interpret the base path of the API during import. Valid values are `ignore`, `prepend`, and
`split`. The default value is `ignore`. To learn more, see [Set the OpenAPI basePath\
Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Body`

The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a `Body` or `BodyS3Location`. If you specify
a `Body` or `BodyS3Location`, don't specify CloudFormation resources such as `AWS::ApiGatewayV2::Authorizer` or `AWS::ApiGatewayV2::Route`.
API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.

_Required_: Conditional

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BodyS3Location`

The S3 location of an OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a `Body` or `BodyS3Location`. If you specify
a `Body` or `BodyS3Location`, don't specify CloudFormation resources such as `AWS::ApiGatewayV2::Authorizer` or `AWS::ApiGatewayV2::Route`.
API Gateway doesn't support the combination of OpenAPI and CloudFormation resources.

_Required_: Conditional

_Type_: [BodyS3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-api-bodys3location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CorsConfiguration`

A CORS configuration. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.

_Required_: No

_Type_: [Cors](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-api-cors.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialsArn`

This property is part of quick create. It specifies the credentials required for
the integration, if any. For a Lambda integration, three options are available. To
specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name
(ARN). To require that the caller's identity be passed through from the request,
specify `arn:aws:iam::*:user/*`. To use resource-based permissions on
supported AWS services, specify `null`. Currently, this property is not used for HTTP
integrations. Supported only for HTTP APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the API.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableExecuteApiEndpoint`

Specifies whether clients can invoke your API by using the default
`execute-api` endpoint. By default, clients can invoke your API
with the default https://{api\_id}.execute-api.{region}.amazonaws.com endpoint.
To require that clients use a custom domain name to invoke your API, disable the
default endpoint.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableSchemaValidation`

Avoid validating models when creating a deployment. Supported only for WebSocket APIs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailOnWarnings`

Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address types that can invoke the API. Use `ipv4` to allow only IPv4 addresses to
invoke your API, or use `dualstack` to allow both IPv4 and IPv6 addresses to invoke your API.

Don’t use IP address type for an HTTP API based on an OpenAPI specification. Instead,
specify the IP address type in the OpenAPI specification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the API. Required unless you specify an OpenAPI definition for `Body` or `S3BodyLocation`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolType`

The API protocol. Valid values are `WEBSOCKET` or `HTTP`. Required unless you specify an OpenAPI definition for `Body` or `S3BodyLocation`.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RouteKey`

This property is part of quick create. If you don't specify a
`routeKey`, a default route of `$default` is created. The
`$default` route acts as a catch-all for any request made to your API,
for a particular stage. The `$default` route key can't be modified. You
can add routes after creating the API, and you can update the route keys of
additional routes. Supported only for HTTP APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RouteSelectionExpression`

The route selection expression for the API. For HTTP APIs, the `routeSelectionExpression` must be `${request.method} ${request.path}`. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

This property is part of quick create. Quick create produces an API with an
integration, a default catch-all route, and a default stage which is configured to
automatically deploy changes. For HTTP integrations, specify a fully qualified URL.
For Lambda integrations, specify a function ARN. The type of the integration will be
HTTP\_PROXY or AWS\_PROXY, respectively. Supported only for HTTP APIs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

A version identifier for the API.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the API ID, such as `a1bcdef2gh`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApiEndpoint`

The default endpoint for an API. For example: `https://abcdef.execute-api.us-west-2.amazonaws.com`.

`ApiId`

The API identifier.

## Examples

- [API creation example](#aws-resource-apigatewayv2-api--examples--API_creation_example)

- [Quick create HTTP API](#aws-resource-apigatewayv2-api--examples--Quick_create_HTTP_API)

### API creation example

The following example creates a WebSocket `Api` resource called
`MyApi`.

#### JSON

```json

{
    "MyApi": {
      "Type": "AWS::ApiGatewayV2::Api",
      "Properties": {
        "Name": "MyApi",
        "ProtocolType": "WEBSOCKET",
        "RouteSelectionExpression": "$request.body.action",
        "ApiKeySelectionExpression": "$request.header.x-api-key"
      }
   }
}

```

#### YAML

```yaml

MyApi:
  Type: 'AWS::ApiGatewayV2::Api'
  Properties:
    Name: MyApi
    ProtocolType: WEBSOCKET
    RouteSelectionExpression: $request.body.action
    ApiKeySelectionExpression: $request.header.x-api-key

```

### Quick create HTTP API

The following example uses quick create to launch an HTTP API
`Api` resource called `HttpApi` that's
integrated with a Lambda function. Quick create produces an HTTP API
with an integration, a default catch-all route, and a default stage
which is configured to automatically deploy changes.

###### Note

To invoke a Lambda integration, API Gateway must have the required permissions. You can use a
resource-based policy or an IAM role to grant API Gateway permissions to invoke a Lambda
function. To learn more, see [AWS Lambda\
Permissions](https://docs.aws.amazon.com/lambda/latest/dg/lambda-permissions.html) in the _AWS Lambda Developer_
_Guide_.

#### JSON

```json

"HttpApi": {
    "Type": "AWS::ApiGatewayV2::Api",
    "Properties": {
        "Name": "Lambda Proxy",
        "Description": "Lambda proxy using quick create",
        "ProtocolType": "HTTP",
        "Target": "arn:aws:apigateway:{region}:lambda:path/2015-03-31/functions/arn:aws:lambda:{region}:{account-id}:function:{function-name}/invocations"
     }
}

```

#### YAML

```yaml

HttpApi:
  Type: AWS::ApiGatewayV2::Api
  Properties:
    Name: Lambda Proxy
    Description: Lambda proxy using quick create
    ProtocolType: HTTP
    Target: arn:aws:apigateway:{region}:lambda:path/2015-03-31/functions/arn:aws:lambda:{region}:{account-id}:function:{function-name}/invocations
```

## See also

- [CreateApi](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis.html#CreateApi) in the _Amazon API Gateway_
_Version 2 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon API Gateway V2

BodyS3Location
