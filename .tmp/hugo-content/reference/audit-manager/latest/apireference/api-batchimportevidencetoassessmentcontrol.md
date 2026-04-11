# BatchImportEvidenceToAssessmentControl

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Adds one or more pieces of evidence to a control in an Audit Manager assessment.

You can import manual evidence from any S3 bucket by specifying the S3 URI of the
object. You can also upload a file from your browser, or enter plain text in response to a
risk assessment question.

The following restrictions apply to this action:

- `manualEvidence` can be only one of the following:
`evidenceFileName`, `s3ResourcePath`, or
`textResponse`

- Maximum size of an individual evidence file: 100 MB

- Number of daily manual evidence uploads per control: 100

- Supported file formats: See [Supported file types for manual evidence](../../../../services/audit-manager/latest/userguide/upload-evidence.md#supported-manual-evidence-files) in the _AWS Audit Manager User Guide_

For more information about Audit Manager service restrictions, see [Quotas and\
restrictions for AWS Audit Manager](../../../../services/audit-manager/latest/userguide/service-quotas.md).

## Request Syntax

```nohighlight

POST /assessments/assessmentId/controlSets/controlSetId/controls/controlId/evidence HTTP/1.1
Content-type: application/json

{
   "manualEvidence": [
      {
         "evidenceFileName": "string",
         "s3ResourcePath": "string",
         "textResponse": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[assessmentId](#API_BatchImportEvidenceToAssessmentControl_RequestSyntax)**

The identifier for the assessment.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

**[controlId](#API_BatchImportEvidenceToAssessmentControl_RequestSyntax)**

The identifier for the control.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

**[controlSetId](#API_BatchImportEvidenceToAssessmentControl_RequestSyntax)**

The identifier for the control set.

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[manualEvidence](#API_BatchImportEvidenceToAssessmentControl_RequestSyntax)**

The list of manual evidence objects.

Type: Array of [ManualEvidence](api-manualevidence.md) objects

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "errors": [
      {
         "errorCode": "string",
         "errorMessage": "string",
         "manualEvidence": {
            "evidenceFileName": "string",
            "s3ResourcePath": "string",
            "textResponse": "string"
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[errors](#API_BatchImportEvidenceToAssessmentControl_ResponseSyntax)**

A list of errors that the `BatchImportEvidenceToAssessmentControl` API
returned.

Type: Array of [BatchImportEvidenceToAssessmentControlError](api-batchimportevidencetoassessmentcontrolerror.md) objects

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

**ResourceNotFoundException**

The resource that's specified in the request can't be found.

**resourceId**

The unique identifier for the resource.

**resourceType**

The type of resource that's affected by the error.

HTTP Status Code: 404

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/batchimportevidencetoassessmentcontrol.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchDisassociateAssessmentReportEvidence

CreateAssessment

All content copied from https://docs.aws.amazon.com/.
