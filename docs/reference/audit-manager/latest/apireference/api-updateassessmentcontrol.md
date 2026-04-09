# UpdateAssessmentControl

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Updates a control within an assessment in AWS Audit Manager.

## Request Syntax

```nohighlight

PUT /assessments/assessmentId/controlSets/controlSetId/controls/controlId HTTP/1.1
Content-type: application/json

{
   "commentBody": "string",
   "controlStatus": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[assessmentId](#API_UpdateAssessmentControl_RequestSyntax)**

The unique identifier for the assessment.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

**[controlId](#API_UpdateAssessmentControl_RequestSyntax)**

The unique identifier for the control.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

**[controlSetId](#API_UpdateAssessmentControl_RequestSyntax)**

The unique identifier for the control set.

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[commentBody](#API_UpdateAssessmentControl_RequestSyntax)**

The comment body text for the control.

Type: String

Length Constraints: Maximum length of 500.

Pattern: `^[\w\W\s\S]*$`

Required: No

**[controlStatus](#API_UpdateAssessmentControl_RequestSyntax)**

The status of the control.

Type: String

Valid Values: `UNDER_REVIEW | REVIEWED | INACTIVE`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "control": {
      "assessmentReportEvidenceCount": number,
      "comments": [
         {
            "authorName": "string",
            "commentBody": "string",
            "postedDate": number
         }
      ],
      "description": "string",
      "evidenceCount": number,
      "evidenceSources": [ "string" ],
      "id": "string",
      "name": "string",
      "response": "string",
      "status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[control](#API_UpdateAssessmentControl_ResponseSyntax)**

The name of the updated control set that the `UpdateAssessmentControl` API
returned.

Type: [AssessmentControl](api-assessmentcontrol.md) object

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

**ValidationException**

The request has invalid or missing parameters.

**fields**

The fields that caused the error, if applicable.

**reason**

The reason the request failed validation.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/updateassessmentcontrol.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/updateassessmentcontrol.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAssessment

UpdateAssessmentControlSetStatus

All content copied from https://docs.aws.amazon.com/.
