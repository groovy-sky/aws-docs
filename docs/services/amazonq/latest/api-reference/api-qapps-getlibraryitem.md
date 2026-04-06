# GetLibraryItem

Retrieves details about a library item for an Amazon Q App, including its metadata,
categories, ratings, and usage statistics.

## Request Syntax

```nohighlight

GET /catalog.getItem?appId=appId&libraryItemId=libraryItemId HTTP/1.1
instance-id: instanceId

```

## URI Request Parameters

The request uses the following URI parameters.

**[appId](#API_qapps_GetLibraryItem_RequestSyntax)**

The unique identifier of the Amazon Q App associated with the library item.

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

**[instanceId](#API_qapps_GetLibraryItem_RequestSyntax)**

The unique identifier of the Amazon Q Business application environment instance.

Required: Yes

**[libraryItemId](#API_qapps_GetLibraryItem_RequestSyntax)**

The unique identifier of the library item to retrieve.

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "appId": "string",
   "appVersion": number,
   "categories": [
      {
         "appCount": number,
         "color": "string",
         "id": "string",
         "title": "string"
      }
   ],
   "createdAt": "string",
   "createdBy": "string",
   "isRatedByUser": boolean,
   "isVerified": boolean,
   "libraryItemId": "string",
   "ratingCount": number,
   "status": "string",
   "updatedAt": "string",
   "updatedBy": "string",
   "userCount": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[appId](#API_qapps_GetLibraryItem_ResponseSyntax)**

The unique identifier of the Q App associated with the library item.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

**[appVersion](#API_qapps_GetLibraryItem_ResponseSyntax)**

The version of the Q App associated with the library item.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2147483647.

**[categories](#API_qapps_GetLibraryItem_ResponseSyntax)**

The categories associated with the library item for discovery.

Type: Array of [Category](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_Category.html) objects

Array Members: Minimum number of 0 items. Maximum number of 3 items.

**[createdAt](#API_qapps_GetLibraryItem_ResponseSyntax)**

The date and time the library item was created.

Type: Timestamp

**[createdBy](#API_qapps_GetLibraryItem_ResponseSyntax)**

The user who created the library item.

Type: String

**[isRatedByUser](#API_qapps_GetLibraryItem_ResponseSyntax)**

Whether the current user has rated the library item.

Type: Boolean

**[isVerified](#API_qapps_GetLibraryItem_ResponseSyntax)**

Indicates whether the library item has been verified.

Type: Boolean

**[libraryItemId](#API_qapps_GetLibraryItem_ResponseSyntax)**

The unique identifier of the library item.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

**[ratingCount](#API_qapps_GetLibraryItem_ResponseSyntax)**

The number of ratings the library item has received from users.

Type: Integer

**[status](#API_qapps_GetLibraryItem_ResponseSyntax)**

The status of the library item, such as "Published".

Type: String

**[updatedAt](#API_qapps_GetLibraryItem_ResponseSyntax)**

The date and time the library item was last updated.

Type: Timestamp

**[updatedBy](#API_qapps_GetLibraryItem_ResponseSyntax)**

The user who last updated the library item.

Type: String

**[userCount](#API_qapps_GetLibraryItem_ResponseSyntax)**

The number of users who have associated the Q App with their account.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/amazonq/latest/api-reference/CommonErrors.html).

**AccessDeniedException**

The client is not authorized to perform the requested operation.

HTTP Status Code: 403

**InternalServerException**

An internal service error occurred while processing the request.

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

HTTP Status Code: 500

**ResourceNotFoundException**

The requested resource could not be found.

**resourceId**

The unique identifier of the resource

**resourceType**

The type of the resource

HTTP Status Code: 404

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qapps-2023-11-27/GetLibraryItem)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/GetLibraryItem)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExportQAppSessionData

GetQApp
