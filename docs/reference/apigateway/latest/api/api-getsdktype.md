# GetSdkType

Gets an SDK type.

## Request Syntax

```nohighlight

GET /sdktypes/sdktype_id HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[sdktype\_id](#API_GetSdkType_RequestSyntax)**

The identifier of the queried SdkType instance.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

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
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[configurationProperties](#API_GetSdkType_ResponseSyntax)**

A list of configuration properties of an SdkType.

Type: Array of [SdkConfigurationProperty](api-sdkconfigurationproperty.md) objects

**[description](#API_GetSdkType_ResponseSyntax)**

The description of an SdkType.

Type: String

**[friendlyName](#API_GetSdkType_ResponseSyntax)**

The user-friendly name of an SdkType instance.

Type: String

**[id](#API_GetSdkType_ResponseSyntax)**

The identifier of an SdkType instance.

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

- [AWS Command Line Interface V2](../../../goto/cli2/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for Python](../../../goto/boto3/apigateway-2015-07-09/getsdktype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/apigateway-2015-07-09/getsdktype.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetSdk

GetSdkTypes

All content copied from https://docs.aws.amazon.com/.
