# CreateQApp

Creates a new Amazon Q App based on the provided definition. The Q App definition specifies
the cards and flow of the Q App. This operation also calculates the dependencies between the
cards by inspecting the references in the prompts.

## Request Syntax

```nohighlight

POST /apps.create HTTP/1.1
instance-id: instanceId
Content-type: application/json

{
   "appDefinition": {
      "cards": [
         { ... }
      ],
      "initialPrompt": "string"
   },
   "description": "string",
   "tags": {
      "string" : "string"
   },
   "title": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[instanceId](#API_qapps_CreateQApp_RequestSyntax)**

The unique identifier of the Amazon Q Business application environment instance.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[appDefinition](#API_qapps_CreateQApp_RequestSyntax)**

The definition of the new Q App, specifying the cards and flow.

Type: [AppDefinitionInput](https://docs.aws.amazon.com/amazonq/latest/api-reference/API_qapps_AppDefinitionInput.html) object

Required: Yes

**[description](#API_qapps_CreateQApp_RequestSyntax)**

The description of the new Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

Required: No

**[tags](#API_qapps_CreateQApp_RequestSyntax)**

Optional tags to associate with the new Q App.

Type: String to string map

Required: No

**[title](#API_qapps_CreateQApp_RequestSyntax)**

The title of the new Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "appArn": "string",
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

**[appArn](#API_qapps_CreateQApp_ResponseSyntax)**

The Amazon Resource Name (ARN) of the new Q App.

Type: String

**[appId](#API_qapps_CreateQApp_ResponseSyntax)**

The unique identifier of the new Q App.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

**[appVersion](#API_qapps_CreateQApp_ResponseSyntax)**

The version of the new Q App.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2147483647.

**[createdAt](#API_qapps_CreateQApp_ResponseSyntax)**

The date and time the Q App was created.

Type: Timestamp

**[createdBy](#API_qapps_CreateQApp_ResponseSyntax)**

The user who created the Q App.

Type: String

**[description](#API_qapps_CreateQApp_ResponseSyntax)**

The description of the new Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 500.

**[initialPrompt](#API_qapps_CreateQApp_ResponseSyntax)**

The initial prompt displayed when the Q App is started.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

**[requiredCapabilities](#API_qapps_CreateQApp_ResponseSyntax)**

The capabilities required to run the Q App, such as file upload or third-party
integrations.

Type: Array of strings

Valid Values: `FileUpload | CreatorMode | RetrievalMode | PluginMode`

**[status](#API_qapps_CreateQApp_ResponseSyntax)**

The status of the new Q App, such as "Created".

Type: String

Valid Values: `PUBLISHED | DRAFT | DELETED`

**[title](#API_qapps_CreateQApp_ResponseSyntax)**

The title of the new Q App.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Pattern: `[^{}\\"<>]+`

**[updatedAt](#API_qapps_CreateQApp_ResponseSyntax)**

The date and time the Q App was last updated.

Type: Timestamp

**[updatedBy](#API_qapps_CreateQApp_ResponseSyntax)**

The user who last updated the Q App.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/amazonq/latest/api-reference/CommonErrors.html).

**AccessDeniedException**

The client is not authorized to perform the requested operation.

HTTP Status Code: 403

**ConflictException**

The requested operation could not be completed due to a conflict with the current state of
the resource.

**resourceId**

The unique identifier of the resource

**resourceType**

The type of the resource

HTTP Status Code: 409

**ContentTooLargeException**

The requested operation could not be completed because the content exceeds the maximum
allowed size.

**resourceId**

The unique identifier of the resource

**resourceType**

The type of the resource

HTTP Status Code: 413

**InternalServerException**

An internal service error occurred while processing the request.

**retryAfterSeconds**

The number of seconds to wait before retrying the operation

HTTP Status Code: 500

**ServiceQuotaExceededException**

The requested operation could not be completed because it would exceed the service's quota
or limit.

**quotaCode**

The code of the quota that was exceeded

**resourceId**

The unique identifier of the resource

**resourceType**

The type of the resource

**serviceCode**

The code for the service where the quota was exceeded

HTTP Status Code: 402

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/qapps-2023-11-27/CreateQApp)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/qapps-2023-11-27/CreateQApp)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qapps-2023-11-27/CreateQApp)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/qapps-2023-11-27/CreateQApp)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qapps-2023-11-27/CreateQApp)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/qapps-2023-11-27/CreateQApp)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/qapps-2023-11-27/CreateQApp)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/qapps-2023-11-27/CreateQApp)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/qapps-2023-11-27/CreateQApp)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qapps-2023-11-27/CreateQApp)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreatePresignedUrl

DeleteLibraryItem
