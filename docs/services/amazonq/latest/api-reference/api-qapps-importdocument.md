# ImportDocument

Uploads a file that can then be used either as a default in a `FileUploadCard`
from Q App definition or as a file that is used inside a single Q App run. The purpose of
the document is determined by a scope parameter that indicates whether it is at the app
definition level or at the app session level.

## Request Syntax

```nohighlight

POST /apps.importDocument HTTP/1.1
instance-id: instanceId
Content-type: application/json

{
   "appId": "string",
   "cardId": "string",
   "fileContentsBase64": "string",
   "fileName": "string",
   "scope": "string",
   "sessionId": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[instanceId](#API_qapps_ImportDocument_RequestSyntax)**

The unique identifier of the Amazon Q Business application environment instance.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[appId](#API_qapps_ImportDocument_RequestSyntax)**

The unique identifier of the Q App the file is associated with.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**[cardId](#API_qapps_ImportDocument_RequestSyntax)**

The unique identifier of the card the file is associated with.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: Yes

**[fileContentsBase64](#API_qapps_ImportDocument_RequestSyntax)**

The base64-encoded contents of the file to upload.

Type: String

Required: Yes

**[fileName](#API_qapps_ImportDocument_RequestSyntax)**

The name of the file being uploaded.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 100.

Required: Yes

**[scope](#API_qapps_ImportDocument_RequestSyntax)**

Whether the file is associated with a Q App definition or a specific Q App
session.

Type: String

Valid Values: `APPLICATION | SESSION`

Required: Yes

**[sessionId](#API_qapps_ImportDocument_RequestSyntax)**

The unique identifier of the Q App session the file is associated with, if
applicable.

Type: String

Pattern: `[\da-f]{8}-[\da-f]{4}-[45][\da-f]{3}-[89ABab][\da-f]{3}-[\da-f]{12}`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "fileId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[fileId](#API_qapps_ImportDocument_ResponseSyntax)**

The unique identifier assigned to the uploaded file.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The client is not authorized to perform the requested operation.

HTTP Status Code: 403

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

**ResourceNotFoundException**

The requested resource could not be found.

**resourceId**

The unique identifier of the resource

**resourceType**

The type of the resource

HTTP Status Code: 404

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

- [AWS Command Line Interface V2](../../../goto/cli2/qapps-2023-11-27/importdocument.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/qapps-2023-11-27/importdocument.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/importdocument.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/qapps-2023-11-27/importdocument.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/importdocument.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/qapps-2023-11-27/importdocument.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/qapps-2023-11-27/importdocument.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/qapps-2023-11-27/importdocument.md)

- [AWS SDK for Python](../../../goto/boto3/qapps-2023-11-27/importdocument.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/importdocument.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetQAppSessionMetadata

ListCategories

All content copied from https://docs.aws.amazon.com/.
