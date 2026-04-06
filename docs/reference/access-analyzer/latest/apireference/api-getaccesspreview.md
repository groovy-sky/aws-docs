# GetAccessPreview

Retrieves information about an access preview for the specified analyzer.

## Request Syntax

```nohighlight

GET /access-preview/accessPreviewId?analyzerArn=analyzerArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[accessPreviewId](#API_GetAccessPreview_RequestSyntax)**

The unique ID for the access preview.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**[analyzerArn](#API_GetAccessPreview_RequestSyntax)**

The [ARN of\
the analyzer](../../../../services/iam/latest/userguide/access-analyzer-getting-started.md#permission-resources) used to generate the access preview.

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "accessPreview": {
      "analyzerArn": "string",
      "configurations": {
         "string" : { ... }
      },
      "createdAt": "string",
      "id": "string",
      "status": "string",
      "statusReason": {
         "code": "string"
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[accessPreview](#API_GetAccessPreview_ResponseSyntax)**

An object that contains information about the access preview.

Type: [AccessPreview](api-accesspreview.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 403

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/accessanalyzer-2019-11-01/GetAccessPreview)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/accessanalyzer-2019-11-01/GetAccessPreview)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/getaccesspreview.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/getaccesspreview.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/getaccesspreview.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/getaccesspreview.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/getaccesspreview.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/getaccesspreview.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/accessanalyzer-2019-11-01/GetAccessPreview)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/getaccesspreview.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GenerateFindingRecommendation

GetAnalyzedResource
