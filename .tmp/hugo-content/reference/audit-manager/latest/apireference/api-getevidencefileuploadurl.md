# GetEvidenceFileUploadUrl

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Creates a presigned Amazon S3 URL that can be used to upload a file as manual
evidence. For instructions on how to use this operation, see [Upload a file from your browser](../../../../services/audit-manager/latest/userguide/upload-evidence.md#how-to-upload-manual-evidence-files) in the _AWS Audit Manager User_
_Guide_.

The following restrictions apply to this operation:

- Maximum size of an individual evidence file: 100 MB

- Number of daily manual evidence uploads per control: 100

- Supported file formats: See [Supported file types for manual evidence](../../../../services/audit-manager/latest/userguide/upload-evidence.md#supported-manual-evidence-files) in the _AWS Audit Manager User Guide_

For more information about Audit Manager service restrictions, see [Quotas and\
restrictions for AWS Audit Manager](../../../../services/audit-manager/latest/userguide/service-quotas.md).

## Request Syntax

```nohighlight

GET /evidenceFileUploadUrl?fileName=fileName HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[fileName](#API_GetEvidenceFileUploadUrl_RequestSyntax)**

The file that you want to upload. For a list of supported file formats, see [Supported file types for manual evidence](../../../../services/audit-manager/latest/userguide/upload-evidence.md#supported-manual-evidence-files) in the _AWS Audit Manager_
_User Guide_.

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `[^\/]*`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "evidenceFileName": "string",
   "uploadUrl": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[evidenceFileName](#API_GetEvidenceFileUploadUrl_ResponseSyntax)**

The name of the uploaded manual evidence file that the presigned URL was generated
for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

**[uploadUrl](#API_GetEvidenceFileUploadUrl_ResponseSyntax)**

The presigned URL that was generated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

Your account isn't registered with AWS Audit Manager. Check the delegated
administrator setup on the Audit Manager settings page, and try again.

HTTP Status Code: 403

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ThrottlingException**

The request was denied due to request throttling.

HTTP Status Code: 429

**ValidationException**

The request has invalid or missing parameters.

**fields**

The fields that caused the error, if applicable.

**reason**

The reason the request failed validation.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/getevidencefileuploadurl.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/getevidencefileuploadurl.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetEvidenceByEvidenceFolder

GetEvidenceFolder

All content copied from https://docs.aws.amazon.com/.
