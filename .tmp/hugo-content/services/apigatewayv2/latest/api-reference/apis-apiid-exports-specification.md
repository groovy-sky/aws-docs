# ExportedAPI

Represents an exported definition of an API in a particular output format, for example, YAML. The API is serialized to the requested specification, for example, OpenAPI 3.0.

## URI

`/v2/apis/apiId/exports/specification`

## HTTP methods

### GET

**Operation ID:** `ExportApi`

Exports a definition of an API in a particular output format and specification.

Path parametersNameTypeRequiredDescription`specification`StringTrue

The version of the API specification to use. `OAS30`, for OpenAPI 3.0, is the only supported value.

`apiId`StringTrue

The API identifier.

Query parametersNameTypeRequiredDescription`includeExtensions`StringFalse

Specifies whether to include
[API Gateway extensions](../../../apigateway/latest/developerguide/api-gateway-swagger-extensions.md) in the exported API definition. API Gateway extensions are included by default.

`stageName`StringFalse

The name of the API stage to export. If you don't specify this property, a representation of the latest API configuration is exported.

`exportVersion`StringFalse

The version of the API Gateway export algorithm. API Gateway uses the latest version by default. Currently, the only supported version is `1.0`.

`outputType`StringTrue

The output type of the exported definition file. Valid values are `JSON` and `YAML`.

ResponsesStatus codeResponse modelDescription`200``ExportedApi`

Success

`400``BadRequestException`

One of the parameters in the request is invalid.

`404``NotFoundException`

The resource specified in the request was not found.

`429``LimitExceededException`

The client is sending more than the allowed number of requests per unit of time.

## Schemas

### Response bodies

```json

"string"
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

### ExportApi

- [AWS Command Line Interface V2](../../../goto/cli2/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for Python](../../../goto/boto3/apigatewayv2-2018-11-29/exportapi.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigatewayv2-2018-11-29/exportapi.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainNames

Integration

All content copied from https://docs.aws.amazon.com/.
