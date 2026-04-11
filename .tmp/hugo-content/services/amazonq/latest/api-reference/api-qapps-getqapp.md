# GetQApp

Retrieves the full details of an Q App, including its definition specifying the cards and
flow.

## Request Syntax

```nohighlight

GET /apps.get?appId=appId&appVersion=appVersion HTTP/1.1
instance-id: instanceId

```

## URI Request Parameters

The request uses the following URI parameters.

**[appId](#API_qapps_GetQApp_RequestSyntax)**

The unique identifier of the Q App to retrieve.

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**[appVersion](#API_qapps_GetQApp_RequestSyntax)**

The version of the Q App.

Valid Range: Minimum value of 0. Maximum value of 2147483647.

**[instanceId](#API_qapps_GetQApp_RequestSyntax)**

The unique identifier of the Amazon Q Business application environment instance.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "appArn": "string",
   "appDefinition": {
      "appDefinitionVersion": "string",
      "canEdit": boolean,
      "cards": [
         { ... }
      ]
   },
   "appId": "string",
   "appVersion": number,
   "createdAt": "string",
   "createdBy": "string",
   "description": "string",
   "initialPrompt": "string",
   "requiredCapabilities": [ "string" ],
   "status": "string",
   "title": "string",
   "updatedAt": "string",
   "updatedBy": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[appArn](#API_qapps_GetQApp_ResponseSyntax)**

The Amazon Resource Name (ARN) of the Q App.

Type: String

**[appDefinition](#API_qapps_GetQApp_ResponseSyntax)**

The full definition of the Q App, specifying the cards and flow.

Type: [AppDefinition](api-qapps-appdefinition.md) object

**[appId](#API_qapps_GetQApp_ResponseSyntax)**

The unique identifier of the Q App.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

**[appVersion](#API_qapps_GetQApp_ResponseSyntax)**

The version of the Q App.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2147483647.

**[createdAt](#API_qapps_GetQApp_ResponseSyntax)**

The date and time the Q App was created.

Type: Timestamp

**[createdBy](#API_qapps_GetQApp_ResponseSyntax)**

The user who created the Q App.

Type: String

**[description](#API_qapps_GetQApp_ResponseSyntax)**

The description of the Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

**[initialPrompt](#API_qapps_GetQApp_ResponseSyntax)**

The initial prompt displayed when the Q App is started.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

**[requiredCapabilities](#API_qapps_GetQApp_ResponseSyntax)**

The capabilities required to run the Q App, such as file upload or third-party
integrations.

Type: Array of strings

Valid Values: `FileUpload | CreatorMode | RetrievalMode | PluginMode`

**[status](#API_qapps_GetQApp_ResponseSyntax)**

The status of the Q App.

Type: String

Valid Values: `PUBLISHED | DRAFT | DELETED`

**[title](#API_qapps_GetQApp_ResponseSyntax)**

The title of the Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

**[updatedAt](#API_qapps_GetQApp_ResponseSyntax)**

The date and time the Q App was last updated.

Type: Timestamp

**[updatedBy](#API_qapps_GetQApp_ResponseSyntax)**

The user who last updated the Q App.

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

- [AWS Command Line Interface V2](../../../goto/cli2/qapps-2023-11-27/getqapp.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qapps-2023-11-27/getqapp.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/getqapp.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qapps-2023-11-27/getqapp.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/getqapp.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qapps-2023-11-27/getqapp.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qapps-2023-11-27/getqapp.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qapps-2023-11-27/getqapp.md)

- [AWS SDK for Python](../../../goto/boto3/qapps-2023-11-27/getqapp.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/getqapp.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetLibraryItem

GetQAppSession

All content copied from https://docs.aws.amazon.com/.
