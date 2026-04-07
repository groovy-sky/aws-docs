# Models

Represents the collection of data models for an API. Supported only for WebSocket APIs. See [Create Models and Mapping Templates for Request and Response Mappings](../../../apigateway/latest/developerguide/models-mappings.md).

## URI

`/v2/apis/apiId/models`

## HTTP methods

### GET

**Operation ID:** `GetModels`

Gets the `Models` for an API.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``Models`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateModel`

Creates a `Model` for an API.

Path parametersNameTypeRequiredDescription`apiId`StringTrue

The API identifier.

ResponsesStatus codeResponse modelDescription`201``Model`

The request has succeeded and has resulted in the creation of a resource.

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`409``ConflictException`

The resource already exists.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

## Schemas

### Request bodies

```json

{
  "name": "string",
  "description": "string",
  "schema": "string",
  "contentType": "string"
}
```

### Response bodies

```json

{
  "items": [
    {
      "modelId": "string",
      "name": "string",
      "description": "string",
      "schema": "string",
      "contentType": "string"
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "modelId": "string",
  "name": "string",
  "description": "string",
  "schema": "string",
  "contentType": "string"
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "resourceType": "string"
}
```

```json

{
  "message": "string"
}
```

```json

{
  "message": "string",
  "limitType": "string"
}
```

## Properties

### BadRequestException

The request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### ConflictException

The requested operation would cause a conflict with the current state of a service resource associated with the request. Resolve the conflict before retrying this request. See the accompanying error message for details.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

### CreateModelInput

Represents the input parameters for a `CreateModel` request.

PropertyTypeRequiredDescription`contentType`

string

False

The content-type for the model, for example, "application/json".

`description`

string

False

The description of the model.

`name`

string

True

The name of the model. Must be alphanumeric.

`schema`

string

True

The schema for the model. For application/json models, this should be JSON schema draft 4 model.

### LimitExceededException

A limit has been exceeded. See the accompanying error message for details.

PropertyTypeRequiredDescription`limitType`

string

False

The limit type.

`message`

string

False

Describes the error encountered.

### Model

Represents a data model for an API. Supported only for WebSocket APIs. See [Create Models and Mapping Templates for Request and Response Mappings](../../../apigateway/latest/developerguide/models-mappings.md).

PropertyTypeRequiredDescription`contentType`

string

False

The content-type for the model, for example, "application/json".

`description`

string

False

The description of the model.

`modelId`

string

False

The model identifier.

`name`

string

True

The name of the model. Must be alphanumeric.

`schema`

string

False

The schema for the model. For application/json models, this should be JSON schema draft 4 model.

### Models

Represents a collection of data models. See [Create Models and Mapping Templates for Request and Response Mappings](../../../apigateway/latest/developerguide/models-mappings.md).

PropertyTypeRequiredDescription`items`

Array of type Model

False

The elements from this collection.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

### NotFoundException

The resource specified in the request was not found. See the `message` field for more information.

PropertyTypeRequiredDescription`message`

string

False

Describes the error encountered.

`resourceType`

string

False

The resource type.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### GetModels

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetModels)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetModels)

### CreateModel

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateModel)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateModel)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Model

ModelTemplate
