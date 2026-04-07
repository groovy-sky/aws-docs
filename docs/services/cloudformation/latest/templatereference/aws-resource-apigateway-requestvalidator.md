This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::RequestValidator

The `AWS::ApiGateway::RequestValidator` resource sets up basic validation rules for incoming requests to your API. For more information, see [Enable Basic Request Validation for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html) in the _API Gateway Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::RequestValidator",
  "Properties" : {
      "Name" : String,
      "RestApiId" : String,
      "ValidateRequestBody" : Boolean,
      "ValidateRequestParameters" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::RequestValidator
Properties:
  Name: String
  RestApiId: String
  ValidateRequestBody: Boolean
  ValidateRequestParameters: Boolean

```

## Properties

`Name`

The name of this RequestValidator

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RestApiId`

The string identifier of the associated RestApi.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidateRequestBody`

A Boolean flag to indicate whether to validate a request body according to the configured Model schema.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidateRequestParameters`

A Boolean flag to indicate whether to validate request parameters ( `true`) or not ( `false`).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the request validator, such as `abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RequestValidatorId`

The ID for the request validator. For example: `abc123`.

## Examples

### Create request validator

The following example creates an API Gateway API with an associated request validator, based on the supplied parameters.

#### JSON

```json

{
    "Parameters": {
        "apiName": {
            "Type": "String"
        },
        "validatorName": {
            "Type": "String"
        },
        "validateRequestBody": {
            "Type": "String"
        },
        "validateRequestParameters": {
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
    }
}
```

#### YAML

```yaml

Parameters:
  apiName:
    Type: String
  validatorName:
    Type: String
  validateRequestBody:
    Type: String
  validateRequestParameters:
    Type: String
Resources:
  RestApi:
    Type: AWS::ApiGateway::RestApi
    Properties:
      Name: !Ref apiName
  RequestValidator:
    Type: AWS::ApiGateway::RequestValidator
    Properties:
      Name: !Ref validatorName
      RestApiId: !Ref RestApi
      ValidateRequestBody: !Ref validateRequestBody
      ValidateRequestParameters: !Ref validateRequestParameters
```

## See also

- [requestvalidator:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRequestValidator.html) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::Model

AWS::ApiGateway::Resource
