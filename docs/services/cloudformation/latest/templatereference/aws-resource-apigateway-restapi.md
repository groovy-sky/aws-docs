This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::RestApi

The `AWS::ApiGateway::RestApi` resource creates a REST API. For more information, see [restapi:create](../../../apigateway/latest/api/api-createrestapi.md) in the _Amazon API Gateway REST API Reference_.

###### Note

On January 1, 2016, the Swagger Specification was donated to the [OpenAPI initiative](https://www.openapis.org/), becoming the foundation of the OpenAPI Specification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::RestApi",
  "Properties" : {
      "ApiKeySourceType" : String,
      "BinaryMediaTypes" : [ String, ... ],
      "Body" : Json,
      "BodyS3Location" : S3Location,
      "CloneFrom" : String,
      "Description" : String,
      "DisableExecuteApiEndpoint" : Boolean,
      "EndpointAccessMode" : String,
      "EndpointConfiguration" : EndpointConfiguration,
      "FailOnWarnings" : Boolean,
      "MinimumCompressionSize" : Integer,
      "Mode" : String,
      "Name" : String,
      "Parameters" : {Key: Value, ...},
      "Policy" : Json,
      "SecurityPolicy" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::RestApi
Properties:
  ApiKeySourceType: String
  BinaryMediaTypes:
    - String
  Body: Json
  BodyS3Location:
    S3Location
  CloneFrom: String
  Description: String
  DisableExecuteApiEndpoint: Boolean
  EndpointAccessMode: String
  EndpointConfiguration:
    EndpointConfiguration
  FailOnWarnings: Boolean
  MinimumCompressionSize: Integer
  Mode: String
  Name: String
  Parameters:
    Key: Value
  Policy: Json
  SecurityPolicy: String
  Tags:
    - Tag

```

## Properties

`ApiKeySourceType`

The source of the API key for metering requests according to a usage plan. Valid values
are: `HEADER` to read the API key from the `X-API-Key` header of a
request. `AUTHORIZER` to read the API key from the `UsageIdentifierKey`
from a custom authorizer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BinaryMediaTypes`

The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Body`

An OpenAPI specification that defines a set of RESTful APIs in JSON format. For YAML templates, you can also provide the specification in YAML format.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BodyS3Location`

The Amazon Simple Storage Service (Amazon S3) location that points to an OpenAPI file, which defines a set of RESTful APIs in JSON or YAML format.

_Required_: No

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-restapi-s3location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloneFrom`

The ID of the RestApi that you want to clone from.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the RestApi.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableExecuteApiEndpoint`

Specifies whether clients can invoke your API by using the default `execute-api` endpoint.
By default, clients can invoke your API with the default
`https://{api_id}.execute-api.{region}.amazonaws.com` endpoint. To require that clients use a
custom domain name to invoke your API, disable the default endpoint

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointAccessMode`

The endpoint access mode for your RestApi.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointConfiguration`

A list of the endpoint types and IP address types of the API. Use this property when creating an API. When importing an existing API, specify the endpoint configuration types using the `Parameters` property.

_Required_: No

_Type_: [EndpointConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-restapi-endpointconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailOnWarnings`

A query parameter to indicate whether to rollback the API update ( `true`) or not ( `false`)
when a warning is encountered. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumCompressionSize`

A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

This property applies only when you use OpenAPI to define your REST API. The `Mode` determines how API Gateway handles resource updates.

Valid values are `overwrite` or `merge`.

For `overwrite`, the new API definition replaces the existing one. The existing
API identifier remains unchanged.

For `merge`, the new API definition is merged with the existing API.

If you don't specify this property, a default value is chosen. For REST APIs created before
March 29, 2021, the default is `overwrite`. For REST APIs created after March 29, 2021, the new API definition takes precedence, but any container types such as endpoint configurations and binary media types are merged with the existing API.

Use the default mode to define top-level `RestApi` properties in addition to using OpenAPI. Generally, it's preferred to use API Gateway's OpenAPI extensions to model these properties.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the RestApi. A name is required if the REST API is not based on an OpenAPI specification.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Custom header parameters as part of the request. For example, to exclude DocumentationParts from an imported API, set `ignore=documentation` as a `parameters` value, as in the AWS CLI command of `aws apigateway import-rest-api --parameters ignore=documentation --body 'file:///path/to/imported-api-body.json'`.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policy`

A policy document that contains the permissions for the `RestApi` resource. To set the ARN for the policy, use the `!Join` intrinsic function with `""` as delimiter and values of `"execute-api:/"` and `"*"`.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityPolicy`

The Transport Layer Security (TLS) version + cipher suite for this RestApi.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The key-value map of strings. The valid character set is \[a-zA-Z+-=.\_:/\]. The tag key can be up to 128 characters and must not start with `aws:`. The tag value can be up to 256 characters.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-restapi-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `RestApi` ID, such as `a1bcdef2gh`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RestApiId`

The string identifier of the associated RestApi.

`RootResourceId`

The root resource ID for a `RestApi` resource, such as `a0bc123d4e`.

## Examples

- [Based on OpenAPI specification](#aws-resource-apigateway-restapi--examples--Based_on_OpenAPI_specification)

- [With endpoint type](#aws-resource-apigateway-restapi--examples--With_endpoint_type)

- [With REGIONAL endpoint type](#aws-resource-apigateway-restapi--examples--With_REGIONAL_endpoint_type)

- [With ApiKeySourceType](#aws-resource-apigateway-restapi--examples--With_ApiKeySourceType)

### Based on OpenAPI specification

The following example creates an API Gateway RestApi resource based on an OpenAPI specification.

#### JSON

```json

{
    "MyRestApi": {
        "Type": "AWS::ApiGateway::RestApi",
        "Properties": {
            "Body": {
                "OpenAPI specification": null
            },
            "Description": "A test API",
            "Name": "MyRestAPI"
        }
    }
}
```

#### YAML

```yaml

MyRestApi:
  Type: 'AWS::ApiGateway::RestApi'
  Properties:
    Body:
      OpenAPI specification: null
    Description: A test API
    Name: MyRestAPI

```

### With endpoint type

The following example creates an API Gateway RestApi resource with an endpoint type.

#### JSON

```json

{
  "Parameters": {
    "apiName": {
      "Type": "String"
    },
    "type": {
      "Type": "String"
    }
  },
  "Resources": {
    "MyRestApi": {
      "Type": "AWS::ApiGateway::RestApi",
      "Properties": {
        "EndpointConfiguration": {
          "Types": [
            {
              "Ref": "type"
            }
          ]
        },
        "Name": {
          "Ref": "apiName"
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
  type:
    Type: String
Resources:
  MyRestApi:
    Type: AWS::ApiGateway::RestApi
    Properties:
      EndpointConfiguration:
        Types:
          - !Ref type
      Name: !Ref apiName
```

### With REGIONAL endpoint type

The following example imports an API Gateway RestApi resource with an endpoint type of REGIONAL.

#### JSON

```json

{
    "Resources": {
        "RestApi": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "Body": {
                    "swagger": 2,
                    "info": {
                        "version": "0.0.1",
                        "title": "test"
                    },
                    "basePath": "/pete",
                    "schemes": [
                        "https"
                    ],
                    "definitions": {
                        "Empty": {
                            "type": "object"
                        }
                    }
                },
                "Name": "myApi",
                "Parameters": {
                    "endpointConfigurationTypes": "REGIONAL"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources :
    RestApi :
        Type : AWS::ApiGateway::RestApi
        Properties :
            Body :
                swagger : 2.0
                info :
                    version : 0.0.1
                    title : test
                basePath : /pete
                schemes :
                    - https
                definitions:
                    Empty :
                        type : object
            Name : myApi
            Parameters:
                endpointConfigurationTypes: REGIONAL
```

### With ApiKeySourceType

The following example creates an API Gateway RestApi resource with ApiKeySourceType, BinaryMediaTypes and MinimumCompressionSize.

#### JSON

```json

{
    "Parameters": {
        "apiKeySourceType": {
            "Type": "String"
        },
        "apiName": {
            "Type": "String"
        },
        "binaryMediaType1": {
            "Type": "String"
        },
        "binaryMediaType2": {
            "Type": "String"
        },
        "minimumCompressionSize": {
            "Type": "String"
        }
    },
    "Resources": {
        "MyRestApi": {
            "Type": "AWS::ApiGateway::RestApi",
            "Properties": {
                "ApiKeySourceType": {
                    "Ref": "apiKeySourceType"
                },
                "BinaryMediaTypes": [
                    {
                        "Ref": "binaryMediaType1"
                    },
                    {
                        "Ref": "binaryMediaType2"
                    }
                ],
                "MinimumCompressionSize": {
                    "Ref": "minimumCompressionSize"
                },
                "Name": {
                    "Ref": "apiName"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Parameters:
    apiKeySourceType:
        Type: String
    apiName:
        Type: String
    binaryMediaType1:
        Type: String
    binaryMediaType2:
        Type: String
    minimumCompressionSize:
        Type: String
Resources:
    MyRestApi:
        Type: AWS::ApiGateway::RestApi
        Properties:
            ApiKeySourceType: !Ref apiKeySourceType
            BinaryMediaTypes:
                - !Ref binaryMediaType1
                - !Ref binaryMediaType2
            MinimumCompressionSize: !Ref minimumCompressionSize
            Name: !Ref apiName
```

## See also

- [restapi:create](../../../apigateway/latest/api/api-createrestapi.md) in the _Amazon API Gateway REST API Reference_

- [Tutorial: Building a private REST API](https://docs.aws.amazon.com/apigateway/latest/developerguide/private-api-tutorial.html) in the _API Gateway Developer Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::Resource

EndpointConfiguration
