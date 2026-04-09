# UpdateAssessmentControlSetStatus

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Updates the status of a control set in an Audit Manager assessment.

## Request Syntax

```nohighlight

PUT /assessments/assessmentId/controlSets/controlSetId/status HTTP/1.1
Content-type: application/json

{
   "comment": "string",
   "status": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[assessmentId](#API_UpdateAssessmentControlSetStatus_RequestSyntax)**

The unique identifier for the assessment.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

**[controlSetId](#API_UpdateAssessmentControlSetStatus_RequestSyntax)**

The unique identifier for the control set.

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[comment](#API_UpdateAssessmentControlSetStatus_RequestSyntax)**

The comment that's related to the status update.

Type: String

Length Constraints: Maximum length of 350.

Pattern: `^[\w\W\s\S]*$`

Required: Yes

**[status](#API_UpdateAssessmentControlSetStatus_RequestSyntax)**

The status of the control set that's being updated.

Type: String

Valid Values: `ACTIVE | UNDER_REVIEW | REVIEWED`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "controlSet": {
      "controls": [
         {
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
      ],
      "delegations": [
         {
            "assessmentId": "string",
            "assessmentName": "string",
            "comment": "string",
            "controlSetId": "string",
            "createdBy": "string",
            "creationTime": number,
            "id": "string",
            "lastUpdated": number,
            "roleArn": "string",
            "roleType": "string",
            "status": "string"
         }
      ],
      "description": "string",
      "id": "string",
      "manualEvidenceCount": number,
      "roles": [
         {
            "roleArn": "string",
            "roleType": "string"
         }
      ],
      "status": "string",
      "systemEvidenceCount": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[controlSet](#API_UpdateAssessmentControlSetStatus_ResponseSyntax)**

The name of the updated control set that the
`UpdateAssessmentControlSetStatus` API returned.

Type: [AssessmentControlSet](api-assessmentcontrolset.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/updateassessmentcontrolsetstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAssessmentControl

UpdateAssessmentFramework

All content copied from https://docs.aws.amazon.com/.
