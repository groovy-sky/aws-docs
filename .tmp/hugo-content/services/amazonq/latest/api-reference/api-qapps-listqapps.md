# ListQApps

Lists the Amazon Q Apps owned by or associated with the user either because they created it
or because they used it from the library in the past. The user identity is extracted from the
credentials used to invoke this operation..

## Request Syntax

```nohighlight

GET /apps.list?limit=limit&nextToken=nextToken HTTP/1.1
instance-id: instanceId

```

## URI Request Parameters

The request uses the following URI parameters.

**[instanceId](#API_qapps_ListQApps_RequestSyntax)**

The unique identifier of the Amazon Q Business application environment instance.

Required: Yes

**[limit](#API_qapps_ListQApps_RequestSyntax)**

The maximum number of Q Apps to return in the response.

Valid Range: Minimum value of 1. Maximum value of 100.

**[nextToken](#API_qapps_ListQApps_RequestSyntax)**

The token to request the next page of results.

Length Constraints: Minimum length of 0. Maximum length of 300.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "apps": [
      {
         "appArn": "string",
         "appId": "string",
         "canEdit": boolean,
         "createdAt": "string",
         "description": "string",
         "isVerified": boolean,
         "status": "string",
         "title": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[apps](#API_qapps_ListQApps_ResponseSyntax)**

The list of Amazon Q Apps meeting the request criteria.

Type: Array of [UserAppItem](api-qapps-userappitem.md) objects

**[nextToken](#API_qapps_ListQApps_ResponseSyntax)**

The token to use to request the next page of results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The client is not authorized to perform the requested operation.

HTTP Status Code: 403

**InternalServerException**

An internal service error occurred while processing the request.

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

HTTP Status Code: 500

**ThrottlingException**

The requested operation could not be completed because too many requests were sent at
once. Wait a bit and try again later.

**quotaCode**

The code of the quota that was exceeded

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

**serviceCode**

The code for the service where the quota was exceeded

HTTP Status Code: 429

**UnauthorizedException**

The client is not authenticated or authorized to perform the requested operation.

HTTP Status Code: 401

**ValidationException**

The input failed to satisfy the constraints specified by the service.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/qapps-2023-11-27/listqapps.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qapps-2023-11-27/listqapps.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/listqapps.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qapps-2023-11-27/listqapps.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/listqapps.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qapps-2023-11-27/listqapps.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qapps-2023-11-27/listqapps.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qapps-2023-11-27/listqapps.md)

- [AWS SDK for Python](../../../goto/boto3/qapps-2023-11-27/listqapps.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/listqapps.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListLibraryItems

ListQAppSessionData

All content copied from https://docs.aws.amazon.com/.
