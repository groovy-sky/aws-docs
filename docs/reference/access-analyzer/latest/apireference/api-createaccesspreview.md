# CreateAccessPreview

Creates an access preview that allows you to preview IAM Access Analyzer findings for your
resource before deploying resource permissions.

## Request Syntax

```nohighlight

PUT /access-preview HTTP/1.1
Content-type: application/json

{
   "analyzerArn": "string",
   "clientToken": "string",
   "configurations": {
      "string" : { ... }
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[analyzerArn](#API_CreateAccessPreview_RequestSyntax)**

The [ARN of\
the account analyzer](../../../../services/iam/latest/userguide/access-analyzer-getting-started.md#permission-resources) used to generate the access preview. You can only create an
access preview for analyzers with an `Account` type and `Active`
status.

Type: String

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

**[clientToken](#API_CreateAccessPreview_RequestSyntax)**

A client token.

Type: String

Required: No

**[configurations](#API_CreateAccessPreview_RequestSyntax)**

Access control configuration for your resource that is used to generate the access
preview. The access preview includes findings for external access allowed to the resource
with the proposed access control configuration. The configuration must contain exactly one
element.

Type: String to [Configuration](api-configuration.md) object map

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "id": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[id](#API_CreateAccessPreview_ResponseSyntax)**

The unique ID for the access preview.

Type: String

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 403

**ConflictException**

A conflict exception error.

**resourceId**

The ID of the resource.

**resourceType**

The resource type.

HTTP Status Code: 409

**InternalServerException**

Internal server error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource could not be found.

**resourceId**

The ID of the resource.

**resourceType**

The type of the resource.

HTTP Status Code: 404

**ServiceQuotaExceededException**

Service quote met error.

**resourceId**

The resource ID.

**resourceType**

The resource type.

HTTP Status Code: 402

**ThrottlingException**

Throttling limit exceeded error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 429

**ValidationException**

Validation exception error.

**fieldList**

A list of fields that didn't validate.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for Python](../../../../services/goto/boto3/accessanalyzer-2019-11-01/createaccesspreview.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/createaccesspreview.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CheckNoPublicAccess

CreateAnalyzer
