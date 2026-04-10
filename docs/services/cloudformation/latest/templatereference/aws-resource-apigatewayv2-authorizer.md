This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Authorizer

The `AWS::ApiGatewayV2::Authorizer` resource creates an authorizer for a
WebSocket API or an HTTP API. To learn more, see [Controlling and managing access to a WebSocket API in API Gateway](../../../apigateway/latest/developerguide/apigateway-websocket-api-control-access.md) and
[Controlling and\
managing access to an HTTP API in API Gateway](../../../apigateway/latest/developerguide/http-api-access-control.md) in the _API Gateway_
_Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::Authorizer",
  "Properties" : {
      "ApiId" : String,
      "AuthorizerCredentialsArn" : String,
      "AuthorizerPayloadFormatVersion" : String,
      "AuthorizerResultTtlInSeconds" : Integer,
      "AuthorizerType" : String,
      "AuthorizerUri" : String,
      "EnableSimpleResponses" : Boolean,
      "IdentitySource" : [ String, ... ],
      "IdentityValidationExpression" : String,
      "JwtConfiguration" : JWTConfiguration,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::Authorizer
Properties:
  ApiId: String
  AuthorizerCredentialsArn: String
  AuthorizerPayloadFormatVersion: String
  AuthorizerResultTtlInSeconds: Integer
  AuthorizerType: String
  AuthorizerUri: String
  EnableSimpleResponses: Boolean
  IdentitySource:
    - String
  IdentityValidationExpression: String
  JwtConfiguration:
    JWTConfiguration
  Name: String

```

## Properties

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthorizerCredentialsArn`

Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for `REQUEST` authorizers.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerPayloadFormatVersion`

Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are `1.0` and `2.0`. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerResultTtlInSeconds`

The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization
caching is disabled. If it is greater than 0, API Gateway caches authorizer
responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerType`

The authorizer type. Specify `REQUEST` for a Lambda function using incoming request parameters. Specify `JWT` to use JSON Web Tokens (supported only for HTTP APIs).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerUri`

The authorizer's Uniform Resource Identifier (URI). For `REQUEST`
authorizers, this must be a well-formed Lambda function URI, for example,
`arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations`.
In general, the URI has this form:
`arn:aws:apigateway:{region}:lambda:path/{service_api}`, where _{region}_ is the same as the region hosting the
Lambda function, path indicates that the remaining substring in the URI should be
treated as the path to the resource, including the initial `/`. For
Lambda functions, this is usually of the form
`/2015-03-31/functions/[FunctionARN]/invocations`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableSimpleResponses`

Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentitySource`

The identity source for which authorization is requested.

For a `REQUEST` authorizer, this is optional. The value is a set of
one or more mapping expressions of the specified request parameters. The
identity source can be headers, query string parameters, stage variables, and
context parameters. For example, if an Auth header and a Name query string
parameter are defined as identity sources, this value is
route.request.header.Auth, route.request.querystring.Name for WebSocket APIs.
For HTTP APIs, use selection expressions prefixed with `$`, for
example, `$request.header.Auth`,
`$request.querystring.Name`. These parameters are used to perform
runtime validation for Lambda-based authorizers by verifying all of the
identity-related request parameters are present in the request, not null, and
non-empty. Only when this is true does the authorizer invoke the authorizer
Lambda function. Otherwise, it returns a 401 Unauthorized response without
calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](../../../apigateway/latest/developerguide/http-api-lambda-authorizer.md).

For `JWT`, a single entry that specifies where to extract the JSON
Web Token (JWT) from inbound requests. Currently only header-based and query
parameter-based selections are supported, for example
`$request.header.Authorization`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityValidationExpression`

This parameter is not used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JwtConfiguration`

The `JWTConfiguration` property specifies the configuration of a JWT
authorizer. Required for the `JWT` authorizer type. Supported only for
HTTP APIs.

_Required_: No

_Type_: [JWTConfiguration](aws-properties-apigatewayv2-authorizer-jwtconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the authorizer.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the authorizer's ID, such as
`abcde1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AuthorizerId`

The authorizer ID.

## Examples

### Authorizer creation example

The following example creates a Lambda `authorizer` resource using the Lambda function `AuthorizerFunction`
for the `MyApi` API.

#### JSON

```json

{
    "Authorizer": {
        "Type": "AWS::ApiGatewayV2::Authorizer",
        "Properties": {
            "Name": "LambdaAuthorizer",
            "ApiId": {
                "Ref": "MyApi"
            },
            "AuthorizerType": "REQUEST",
            "AuthorizerCredentialsArn": "Arn",
            "AuthorizerUri": {
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
                              "AuthorizerFunction",
                              "Arn"
                            ]
                        },
                        "/invocations"
                    ]
                ]
            },
            "AuthorizerResultTtlInSeconds": 500,
            "IdentitySource": [
                "route.request.header.Auth"
            ]
        }
    }
}
```

#### YAML

```yaml

Authorizer:
  Type: 'AWS::ApiGatewayV2::Authorizer'
  Properties:
    Name: LambdaAuthorizer
    ApiId: !Ref MyApi
    AuthorizerType: REQUEST
    AuthorizerCredentialsArn: Arn
    AuthorizerUri: !Join
      - ''
      - - 'arn:'
        - !Ref 'AWS::Partition'
        - ':apigateway:'
        - !Ref 'AWS::Region'
        - ':lambda:path/2015-03-31/functions/'
        - !GetAtt AuthorizerFunction.Arn
        - /invocations
    AuthorizerResultTtlInSeconds: 500
    IdentitySource:
      - route.request.header.Auth
```

## See also

- [CreateAuthorizer](../../../apigatewayv2/latest/api-reference/apis-apiid-authorizers.md#CreateAuthorizer) in the _Amazon API_
_Gateway Version 2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGatewayV2::ApiMapping

JWTConfiguration

All content copied from https://docs.aws.amazon.com/.
