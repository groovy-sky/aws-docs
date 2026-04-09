# GetSdkTypes

Gets SDK types

## Request Syntax

```nohighlight

GET /sdktypes?limit=limit&position=position HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[limit](#API_GetSdkTypes_RequestSyntax)**

The maximum number of returned results per page. The default value is 25 and the maximum value is 500.

**[position](#API_GetSdkTypes_RequestSyntax)**

The current pagination position in the paged result set.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "item": [
      {
         "configurationProperties": [
            {
               "defaultValue": "string",
               "description": "string",
               "friendlyName": "string",
               "name": "string",
               "required": boolean
            }
         ],
         "description": "string",
         "friendlyName": "string",
         "id": "string"
      }
   ],
   "position": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[item](#API_GetSdkTypes_ResponseSyntax)**

The current page of elements from this collection.

Type: Array of [SdkType](api-sdktype.md) objects

**[position](#API_GetSdkTypes_ResponseSyntax)**

The current pagination position in the paged result set.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](../../../../services/apigateway/latest/api/commonerrors.md).

**BadRequestException**

The submitted request is not valid, for example, the input is incomplete or incorrect. See the accompanying error message for details.

HTTP Status Code: 400

**NotFoundException**

The requested resource is not found. Make sure that the request URI is correct.

HTTP Status Code: 404

**TooManyRequestsException**

The request has reached its throttling limit. Retry after the specified time period.

HTTP Status Code: 429

**UnauthorizedException**

The request is denied because the caller has insufficient permissions.

HTTP Status Code: 401

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getsdktypes.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/apigateway-2015-07-09/getsdktypes.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetSdkType

GetStage

All content copied from https://docs.aws.amazon.com/.
