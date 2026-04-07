# ApiMappings

Represents a collection of API mappings.

## URI

`/v2/domainnames/domainName/apimappings`

## HTTP methods

### GET

**Operation ID:** `GetApiMappings`

Gets API mappings.

Path parametersNameTypeRequiredDescription`domainName`StringTrue

The domain name.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The next page of elements from this collection. Not valid for the last element of the collection.

`maxResults`StringFalse

The maximum number of elements to be returned for this resource.

ResponsesStatus codeResponse modelDescription`200``ApiMappings`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

### POST

**Operation ID:** `CreateApiMapping`

Creates an API mapping.

Path parametersNameTypeRequiredDescription`domainName`StringTrue

The domain name.

ResponsesStatus codeResponse modelDescription`201``ApiMapping`

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
  "apiMappingKey": "string",
  "apiId": "string",
  "stage": "string"
}
```

### Response bodies

```json

{
  "items": [
    {
      "apiMappingId": "string",
      "apiMappingKey": "string",
      "apiId": "string",
      "stage": "string"
    }
  ],
  "nextToken": "string"
}
```

```json

{
  "apiMappingId": "string",
  "apiMappingKey": "string",
  "apiId": "string",
  "stage": "string"
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

### ApiMapping

Represents an API mapping.

PropertyTypeRequiredDescription`apiId`

string

True

The API identifier.

`apiMappingId`

string

False

The API mapping identifier.

`apiMappingKey`

string

False

The API mapping key.

`stage`

string

True

The API stage.

### ApiMappings

Represents a collection of `ApiMappings` resources.

PropertyTypeRequiredDescription`items`

Array of type ApiMapping

False

The elements from this collection.

`nextToken`

string

False

The next page of elements from this collection. Not valid for the last element of the collection.

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

### CreateApiMappingInput

Represents the input parameters for a `CreateApiMapping` request.

PropertyTypeRequiredDescription`apiId`

string

True

The API identifier.

`apiMappingKey`

string

False

The API mapping key.

`stage`

string

True

The API stage.

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

### GetApiMappings

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/GetApiMappings)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/GetApiMappings)

### CreateApiMapping

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/apigatewayv2-2018-11-29/CreateApiMapping)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/apigatewayv2-2018-11-29/CreateApiMapping)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApiMapping

Apis
