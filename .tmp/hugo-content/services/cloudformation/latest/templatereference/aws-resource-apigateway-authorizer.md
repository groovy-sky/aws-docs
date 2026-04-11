This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Authorizer

The `AWS::ApiGateway::Authorizer` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled. API Gateway activates the authorizer when a client calls those methods.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::Authorizer",
  "Properties" : {
      "AuthorizerCredentials" : String,
      "AuthorizerResultTtlInSeconds" : Integer,
      "AuthorizerUri" : String,
      "AuthType" : String,
      "IdentitySource" : String,
      "IdentityValidationExpression" : String,
      "Name" : String,
      "ProviderARNs" : [ String, ... ],
      "RestApiId" : String,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::Authorizer
Properties:
  AuthorizerCredentials: String
  AuthorizerResultTtlInSeconds: Integer
  AuthorizerUri: String
  AuthType: String
  IdentitySource: String
  IdentityValidationExpression: String
  Name: String
  ProviderARNs:
    - String
  RestApiId: String
  Type: String

```

## Properties

`AuthorizerCredentials`

Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerResultTtlInSeconds`

The TTL in seconds of cached authorizer results. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway will cache authorizer responses. If this field is not set, the default value is 300. The maximum value is 3600, or 1 hour.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerUri`

Specifies the authorizer's Uniform Resource Identifier (URI). For `TOKEN` or `REQUEST` authorizers, this must be a well-formed Lambda function URI, for example, `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations`. In general, the URI has this form `arn:aws:apigateway:{region}:lambda:path/{service_api}`, where `{region}` is the same as the region hosting the Lambda function, `path` indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial `/`. For Lambda functions, this is usually of the form `/2015-03-31/functions/[FunctionARN]/invocations`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthType`

Optional customer-defined field, used in OpenAPI imports and exports without functional impact.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentitySource`

The identity source for which authorization is requested. For a `TOKEN` or
`COGNITO_USER_POOLS` authorizer, this is required and specifies the request
header mapping expression for the custom header holding the authorization token submitted by
the client. For example, if the token header name is `Auth`, the header mapping
expression is `method.request.header.Auth`. For the `REQUEST`
authorizer, this is required when authorization caching is enabled. The value is a
comma-separated string of one or more mapping expressions of the specified request parameters.
For example, if an `Auth` header, a `Name` query string parameter are
defined as identity sources, this value is `method.request.header.Auth,
        method.request.querystring.Name`. These parameters will be used to derive the
authorization caching key and to perform runtime validation of the `REQUEST`
authorizer by verifying all of the identity-related request parameters are present, not null
and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda
function, otherwise, it returns a 401 Unauthorized response without calling the Lambda
function. The valid value is a string of comma-separated mapping expressions of the specified
request parameters. When the authorization caching is not enabled, this property is
optional.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityValidationExpression`

A validation expression for the incoming identity token. For `TOKEN` authorizers, this value is a regular expression. For `COGNITO_USER_POOLS` authorizers, API Gateway will match the `aud` field of the incoming token from the client against the specified regular expression. It will invoke the authorizer's Lambda function when there is a match. Otherwise, it will return a 401 Unauthorized response without calling the Lambda function. The validation expression does not apply to the `REQUEST` authorizer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the authorizer.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderARNs`

A list of the Amazon Cognito user pool ARNs for the `COGNITO_USER_POOLS` authorizer. Each element is of this format: `arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}`. For a `TOKEN` or `REQUEST` authorizer, this is not defined.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The authorizer type. Valid values are `TOKEN` for a Lambda function using a single authorization token submitted in a custom header, `REQUEST` for a Lambda function using incoming request parameters, and `COGNITO_USER_POOLS` for using an Amazon Cognito user pool.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the authorizer's ID, such as `abcde1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AuthorizerId`

The ID for the authorizer. For example: `abc123`.

## Examples

### Create authorizer

The following examples create a custom authorizer that is an AWS Lambda function.

#### JSON

```json

{
    "Authorizer": {
        "Type": "AWS::ApiGateway::Authorizer",
        "Properties": {
            "AuthorizerCredentials": {
                "Fn::GetAtt": [
                    "LambdaInvocationRole",
                    "Arn"
                ]
            },
            "AuthorizerResultTtlInSeconds": "300",
            "AuthorizerUri": {
                "Fn::Join": [
                    "",
                    [
                        "arn:aws:apigateway:",
                        {
                            "Ref": "AWS::Region"
                        },
                        ":lambda:path/2015-03-31/functions/",
                        {
                            "Fn::GetAtt": [
                                "LambdaAuthorizer",
                                "Arn"
                            ]
                        },
                        "/invocations"
                    ]
                ]
            },
            "Type": "TOKEN",
            "IdentitySource": "method.request.header.Auth",
            "Name": "DefaultAuthorizer",
            "RestApiId": {
                "Ref": "RestApi"
            }
        }
    }
}
```

#### YAML

```yaml

Authorizer:
  Type: 'AWS::ApiGateway::Authorizer'
  Properties:
    AuthorizerCredentials: !GetAtt
      - LambdaInvocationRole
      - Arn
    AuthorizerResultTtlInSeconds: '300'
    AuthorizerUri: !Join
      - ''
      - - 'arn:aws:apigateway:'
        - !Ref 'AWS::Region'
        - ':lambda:path/2015-03-31/functions/'
        - !GetAtt
          - LambdaAuthorizer
          - Arn
        - /invocations
    Type: TOKEN
    IdentitySource: method.request.header.Auth
    Name: DefaultAuthorizer
    RestApiId: !Ref RestApi

```

## See also

- [authorizer:create](../../../apigateway/latest/api/api-createauthorizer.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ApiGateway::BasePathMapping

All content copied from https://docs.aws.amazon.com/.
