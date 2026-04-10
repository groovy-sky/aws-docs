This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::Method

The `AWS::ApiGateway::Method` resource creates API Gateway methods that define the parameters and body that clients must send in their requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::Method",
  "Properties" : {
      "ApiKeyRequired" : Boolean,
      "AuthorizationScopes" : [ String, ... ],
      "AuthorizationType" : String,
      "AuthorizerId" : String,
      "HttpMethod" : String,
      "Integration" : Integration,
      "MethodResponses" : [ MethodResponse, ... ],
      "OperationName" : String,
      "RequestModels" : {Key: Value, ...},
      "RequestParameters" : {Key: Value, ...},
      "RequestValidatorId" : String,
      "ResourceId" : String,
      "RestApiId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::Method
Properties:
  ApiKeyRequired: Boolean
  AuthorizationScopes:
    - String
  AuthorizationType: String
  AuthorizerId: String
  HttpMethod: String
  Integration:
    Integration
  MethodResponses:
    - MethodResponse
  OperationName: String
  RequestModels:
    Key: Value
  RequestParameters:
    Key: Value
  RequestValidatorId: String
  ResourceId: String
  RestApiId: String

```

## Properties

`ApiKeyRequired`

A boolean flag specifying whether a valid ApiKey is required to invoke this method.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizationScopes`

A list of authorization scopes configured on the method. The scopes are used with a `COGNITO_USER_POOLS` authorizer to authorize the method invocation. The authorization works by matching the method scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any method scopes matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the method scope is configured, the client must provide an access token instead of an identity token for authorization purposes.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizationType`

The method's authorization type. This parameter is required. For valid values, see [Method](../../../apigateway/latest/api/api-method.md) in the _API Gateway API Reference_.

###### Note

If you specify the `AuthorizerId` property, specify `CUSTOM` or `COGNITO_USER_POOLS` for this property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerId`

The identifier of an authorizer to use on this method. The method's authorization type must be `CUSTOM` or `COGNITO_USER_POOLS`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpMethod`

The method's HTTP verb.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Integration`

Represents an `HTTP`, `HTTP_PROXY`, `AWS`, `AWS_PROXY`, or Mock integration.

_Required_: No

_Type_: [Integration](aws-properties-apigateway-method-integration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MethodResponses`

Gets a method response associated with a given HTTP status code.

_Required_: No

_Type_: Array of [MethodResponse](aws-properties-apigateway-method-methodresponse.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperationName`

A human-friendly operation identifier for the method. For example, you can assign the `operationName` of `ListPets` for the `GET /pets` method in the `PetStore` example.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestModels`

A key-value map specifying data schemas, represented by Model resources, (as the mapped value) of the request payloads of given content types (as the mapping key).

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestParameters`

A key-value map defining required or optional method request parameters that can be accepted by API Gateway. A key is a method request parameter name matching the pattern of `method.request.{location}.{name}`, where `location` is `querystring`, `path`, or `header` and `name` is a valid and unique parameter name. The value associated with the key is a Boolean flag indicating whether the parameter is required ( `true`) or optional ( `false`). The method request parameter names defined here are available in Integration to be mapped to integration request parameters or templates.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestValidatorId`

The identifier of a RequestValidator for request validation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

The Resource identifier for the MethodResponse resource.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the method ID, such as `mysta-metho-01234b567890example`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Mock Method](#aws-resource-apigateway-method--examples--Mock_Method)

- [Lambda Proxy](#aws-resource-apigateway-method--examples--Lambda_Proxy)

- [Associated Request Validator](#aws-resource-apigateway-method--examples--Associated_Request_Validator)

### Mock Method

The following example creates a mock GET method for the `MyApi` API.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "Api": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "Name": "myAPI"
            }
        },
        "MockMethod": {
            "Type": "AWS::ApiGateway::Method",
            "Properties": {
                "RestApiId": {
                    "Ref": "Api"
                },
                "ResourceId": {
                    "Fn::GetAtt": [
                        "Api",
                        "RootResourceId"
                    ]
                },
                "HttpMethod": "GET",
                "AuthorizationType": "NONE",
                "RequestParameters": {
                    "method.request.header.myheader": false
                },
                "Integration": {
                    "Type": "MOCK",
                    "RequestParameters": {
                        "integration.request.header.header1": "method.request.header.myheader"
                    },
                    "RequestTemplates": {
                        "application/json": "{\"statusCode\": 200}"
                    },
                    "IntegrationResponses": [
                        {
                            "StatusCode": "200",
                            "ResponseParameters": {
                                "method.response.header.header1": "integration.response.header.header1",
                                "method.response.header.header2": "'staticvalue'"
                            }
                        }
                    ]
                },
                "MethodResponses": [
                    {
                        "StatusCode": "200",
                        "ResponseParameters": {
                            "method.response.header.header1": true,
                            "method.response.header.header2": true
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  Api:
    Type: AWS::ApiGateway::RestApi
    Properties:
      Name: myAPI
  MockMethod:
    Type: AWS::ApiGateway::Method
    Properties:
      RestApiId: !Ref Api
      ResourceId: !GetAtt Api.RootResourceId
      HttpMethod: GET
      AuthorizationType: NONE
      RequestParameters:
        method.request.header.myheader: false
      Integration:
        Type: MOCK
        RequestTemplates:
          application/json: "{\"statusCode\": 200}"
        RequestParameters:
          integration.request.header.header1: method.request.header.myheader
        IntegrationResponses:
          - StatusCode: '200'
            ResponseParameters:
              method.response.header.header1: integration.response.header.header1
              method.response.header.header2: '''staticvalue'''
      MethodResponses:
        - StatusCode: '200'
          ResponseParameters:
            method.response.header.header1: true
            method.response.header.header2: true
```

### Lambda Proxy

The following example creates a proxy resource to enable clients to call a Lambda function with a single integration setup on a catch-all ANY method. The `Uri` property specifies the Lambda function. For more information about Lambda proxy integration and a sample Lambda function, see [Create an API with Lambda Proxy Integration through a Proxy Resource](../../../apigateway/latest/developerguide/api-gateway-create-api-as-simple-proxy-for-lambda.md) in the _API Gateway Developer Guide_.

###### Note

Use the [AWS::Lambda::Permission](../userguide/aws-resource-lambda-permission.md) resource to grant API Gateway permission to invoke your Lambda function.

#### JSON

```json

{
    "ProxyResource": {
        "Type": "AWS::ApiGateway::Resource",
        "Properties": {
            "RestApiId": {
                "Ref": "LambdaSimpleProxy"
            },
            "ParentId": {
                "Fn::GetAtt": [
                    "LambdaSimpleProxy",
                    "RootResourceId"
                ]
            },
            "PathPart": "{proxy+}"
        }
    },
    "ProxyResourceANY": {
        "Type": "AWS::ApiGateway::Method",
        "Properties": {
            "RestApiId": {
                "Ref": "LambdaSimpleProxy"
            },
            "ResourceId": {
                "Ref": "ProxyResource"
            },
            "HttpMethod": "ANY",
            "AuthorizationType": "NONE",
            "Integration": {
                "Type": "AWS_PROXY",
                "IntegrationHttpMethod": "POST",
                "Uri": {
                    "Fn::Sub": "arn:aws:apigateway:${AWS::Region}:lambda:path/2015-03-31/functions/${LambdaForSimpleProxy.Arn}/invocations"
                }
            }
        }
    }
}
```

#### YAML

```yaml

ProxyResource:
  Type: 'AWS::ApiGateway::Resource'
  Properties:
    RestApiId: !Ref LambdaSimpleProxy
    ParentId: !GetAtt
      - LambdaSimpleProxy
      - RootResourceId
    PathPart: '{proxy+}'
ProxyResourceANY:
  Type: 'AWS::ApiGateway::Method'
  Properties:
    RestApiId: !Ref LambdaSimpleProxy
    ResourceId: !Ref ProxyResource
    HttpMethod: ANY
    AuthorizationType: NONE
    Integration:
      Type: AWS_PROXY
      IntegrationHttpMethod: POST
      Uri: !Sub >-
        arn:aws:apigateway:${AWS::Region}:lambda:path/2015-03-31/functions/${LambdaForSimpleProxy.Arn}/invocations

```

### Associated Request Validator

The following example creates a REST API, method, and request validator, and associates the request validator with the method. It also lets you specify how to convert the request payload.

#### JSON

```json

{
  "Parameters": {
    "contentHandling": {
      "Type": "String"
    },
    "operationName": {
      "Type": "String",
      "Default": "testoperationName"
    },
    "restApiName": {
      "Type": "String",
      "Default": "testrestApiName"
    },
    "validatorName": {
      "Type": "String",
      "Default": "testvalidatorName"
    },
    "validateRequestBody": {
      "Type": "String",
      "Default": "testvalidateRequestBody"
    },
    "validateRequestParameters": {
      "Type": "String",
      "Default": true
    }
  },
  "Resources": {
    "RestApi": {
      "Type": "AWS::ApiGateway::RestApi",
      "Properties": {
        "Name": {
          "Ref": "restApiName"
        }
      }
    },
    "Method": {
      "Type": "AWS::ApiGateway::Method",
      "Properties": {
        "HttpMethod": "POST",
        "ResourceId": {
          "Fn::GetAtt": [
            "RestApi",
            "RootResourceId"
          ]
        },
        "RestApiId": {
          "Ref": "RestApi"
        },
        "AuthorizationType": "NONE",
        "Integration": {
          "Type": "MOCK",
          "ContentHandling": {
            "Ref": "contentHandling"
          },
          "IntegrationResponses": [
            {
              "ContentHandling": {
                "Ref": "contentHandling"
              },
              "StatusCode": 400
            }
          ]
        },
        "RequestValidatorId": {
          "Ref": "RequestValidator"
        },
        "OperationName": {
          "Ref": "operationName"
        }
      }
    },
    "RequestValidator": {
      "Type": "AWS::ApiGateway::RequestValidator",
      "Properties": {
        "Name": {
          "Ref": "validatorName"
        },
        "RestApiId": {
          "Ref": "RestApi"
        },
        "ValidateRequestBody": {
          "Ref": "validateRequestBody"
        },
        "ValidateRequestParameters": {
          "Ref": "validateRequestParameters"
        }
      }
    }
  },
  "Outputs": {
    "RootResourceId": {
      "Value": {
        "Fn::GetAtt": [
          "RestApi",
          "RootResourceId"
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Parameters:
  contentHandling:
    Type: String
  operationName:
    Type: String
    Default: testoperationName
  restApiName:
    Type: String
    Default: testrestApiName
  validatorName:
    Type: String
    Default: testvalidatorName
  validateRequestBody:
    Type: String
    Default: testvalidateRequestBody
  validateRequestParameters:
    Type: String
    Default: true
Resources:
  RestApi:
    Type: AWS::ApiGateway::RestApi
    Properties:
      Name: !Ref restApiName
  Method:
    Type: AWS::ApiGateway::Method
    Properties:
      HttpMethod: POST
      ResourceId: !GetAtt RestApi.RootResourceId
      RestApiId: !Ref RestApi
      AuthorizationType: NONE
      Integration:
        Type: MOCK
        ContentHandling: !Ref contentHandling
        IntegrationResponses:
          - ContentHandling: !Ref contentHandling
            StatusCode: 400
      RequestValidatorId: !Ref RequestValidator
      OperationName: !Ref operationName
  RequestValidator:
    Type: AWS::ApiGateway::RequestValidator
    Properties:
      Name: !Ref validatorName
      RestApiId: !Ref RestApi
      ValidateRequestBody: !Ref validateRequestBody
      ValidateRequestParameters: !Ref validateRequestParameters
Outputs:
  RootResourceId:
    Value: !GetAtt RestApi.RootResourceId
```

## See also

- [method:put](../../../apigateway/latest/api/api-putmethod.md) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGateway::GatewayResponse

Integration

All content copied from https://docs.aws.amazon.com/.
