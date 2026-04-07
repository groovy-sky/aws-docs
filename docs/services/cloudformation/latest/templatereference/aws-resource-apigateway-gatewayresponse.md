This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::GatewayResponse

The `AWS::ApiGateway::GatewayResponse` resource creates a gateway response for your API. When you delete a stack containing this resource, your custom gateway responses are reset. For more information, see [API Gateway Responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html#api-gateway-gatewayResponse-definition) in the _API Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::GatewayResponse",
  "Properties" : {
      "ResponseParameters" : {Key: Value, ...},
      "ResponseTemplates" : {Key: Value, ...},
      "ResponseType" : String,
      "RestApiId" : String,
      "StatusCode" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::GatewayResponse
Properties:
  ResponseParameters:
    Key: Value
  ResponseTemplates:
    Key: Value
  ResponseType: String
  RestApiId: String
  StatusCode: String

```

## Properties

`ResponseParameters`

Response parameters (paths, query strings and headers) of the GatewayResponse as a
string-to-string map of key-value pairs.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseTemplates`

Response templates of the GatewayResponse as a string-to-string map of key-value pairs.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseType`

The response type of the associated GatewayResponse.

_Required_: Yes

_Type_: String

_Allowed values_: `DEFAULT_4XX | DEFAULT_5XX | RESOURCE_NOT_FOUND | UNAUTHORIZED | INVALID_API_KEY | ACCESS_DENIED | AUTHORIZER_FAILURE | AUTHORIZER_CONFIGURATION_ERROR | INVALID_SIGNATURE | EXPIRED_TOKEN | MISSING_AUTHENTICATION_TOKEN | INTEGRATION_FAILURE | INTEGRATION_TIMEOUT | API_CONFIGURATION_ERROR | UNSUPPORTED_MEDIA_TYPE | BAD_REQUEST_PARAMETERS | BAD_REQUEST_BODY | REQUEST_TOO_LARGE | THROTTLED | QUOTA_EXCEEDED | WAF_FILTERED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StatusCode`

The HTTP status code for this GatewayResponse.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID for the gateway response. For example: `abc123`.

## Examples

- [404 Response](#aws-resource-apigateway-gatewayresponse--examples--404_Response)

- [Parameterized Response](#aws-resource-apigateway-gatewayresponse--examples--Parameterized_Response)

### 404 Response

The following example returns a 404 status code for resource not found instead of missing authentication token for a CORS request (applicable to unsecured/unrestricted APIs).

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "RestApi": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "Name": "myRestApi"
            }
        },
        "GatewayResponse": {
            "Type": "AWS::ApiGateway::GatewayResponse",
            "Properties": {
                "ResponseParameters": {
                    "gatewayresponse.header.Access-Control-Allow-Origin": "'*'",
                    "gatewayresponse.header.Access-Control-Allow-Headers": "'*'"
                },
                "ResponseType": "MISSING_AUTHENTICATION_TOKEN",
                "RestApiId": {
                    "Ref": "RestApi"
                },
                "StatusCode": "404"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  RestApi:
    Type: AWS::ApiGateway::RestApi
    Properties:
      Name: myRestApi
  GatewayResponse:
    Type: AWS::ApiGateway::GatewayResponse
    Properties:
      ResponseParameters:
        gatewayresponse.header.Access-Control-Allow-Origin: "'*'"
        gatewayresponse.header.Access-Control-Allow-Headers: "'*'"
      ResponseType: MISSING_AUTHENTICATION_TOKEN
      RestApiId: !Ref RestApi
      StatusCode: '404'
```

### Parameterized Response

The following example creates a response for an API based on the supplied parameters.

#### JSON

```json

{
    "Parameters": {
        "apiName": {
            "Type": "String"
        },
        "responseParameter1": {
            "Type": "String"
        },
        "responseParameter2": {
            "Type": "String"
        },
        "responseType": {
            "Type": "String"
        },
        "statusCode": {
            "Type": "String"
        }
    },
    "Resources": {
        "RestApi": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "Name": {
                    "Ref": "apiName"
                }
            }
        },
        "GatewayResponse": {
            "Type": "AWS::ApiGateway::GatewayResponse",
            "Properties": {
                "ResponseParameters": {
                    "gatewayresponse.header.k1": {
                        "Ref": "responseParameter1"
                    },
                    "gatewayresponse.header.k2": {
                        "Ref": "responseParameter2"
                    }
                },
                "ResponseType": {
                    "Ref": "responseType"
                },
                "RestApiId": {
                    "Ref": "RestApi"
                },
                "StatusCode": {
                    "Ref": "statusCode"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
    apiName :
        Type : String
    responseParameter1:
        Type : String
    responseParameter2:
        Type : String
    responseType:
        Type : String
    statusCode:
        Type : String
Resources :
    RestApi:
        Type: AWS::ApiGateway::RestApi
        Properties:
            Name: !Ref apiName
    GatewayResponse:
        Type: AWS::ApiGateway::GatewayResponse
        Properties:
            ResponseParameters:
                gatewayresponse.header.k1 : !Ref responseParameter1
                gatewayresponse.header.k2 : !Ref responseParameter2
            ResponseType: !Ref responseType
            RestApiId: !Ref RestApi
            StatusCode: !Ref statusCode
```

## See also

- [gatewayresponse:put](https://docs.aws.amazon.com/apigateway/latest/api/API_PutGatewayResponse.html) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::ApiGateway::Method
