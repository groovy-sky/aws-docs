# UpdateAssessment

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Edits an Audit Manager assessment.

## Request Syntax

```nohighlight

PUT /assessments/assessmentId HTTP/1.1
Content-type: application/json

{
   "assessmentDescription": "string",
   "assessmentName": "string",
   "assessmentReportsDestination": {
      "destination": "string",
      "destinationType": "string"
   },
   "roles": [
      {
         "roleArn": "string",
         "roleType": "string"
      }
   ],
   "scope": {
      "awsAccounts": [
         {
            "emailAddress": "string",
            "id": "string",
            "name": "string"
         }
      ],
      "awsServices": [
         {
            "serviceName": "string"
         }
      ]
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[assessmentId](#API_UpdateAssessment_RequestSyntax)**

The unique identifier for the assessment.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[assessmentDescription](#API_UpdateAssessment_RequestSyntax)**

The description of the assessment.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**[assessmentName](#API_UpdateAssessment_RequestSyntax)**

The name of the assessment to be updated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**[assessmentReportsDestination](#API_UpdateAssessment_RequestSyntax)**

The assessment report storage destination for the assessment that's being updated.

Type: [AssessmentReportsDestination](api-assessmentreportsdestination.md) object

Required: No

**[roles](#API_UpdateAssessment_RequestSyntax)**

The list of roles for the assessment.

Type: Array of [Role](api-role.md) objects

Required: No

**[scope](#API_UpdateAssessment_RequestSyntax)**

The scope of the assessment.

Type: [Scope](api-scope.md) object

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "assessment": {
      "arn": "string",
      "awsAccount": {
         "emailAddress": "string",
         "id": "string",
         "name": "string"
      },
      "framework": {
         "arn": "string",
         "controlSets": [
            {
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
         ],
         "id": "string",
         "metadata": {
            "complianceType": "string",
            "description": "string",
            "logo": "string",
            "name": "string"
         }
      },
      "metadata": {
         "assessmentReportsDestination": {
            "destination": "string",
            "destinationType": "string"
         },
         "complianceType": "string",
         "creationTime": number,
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
         "lastUpdated": number,
         "name": "string",
         "roles": [
            {
               "roleArn": "string",
               "roleType": "string"
            }
         ],
         "scope": {
            "awsAccounts": [
               {
                  "emailAddress": "string",
                  "id": "string",
                  "name": "string"
               }
            ],
            "awsServices": [
               {
                  "serviceName": "string"
               }
            ]
         },
         "status": "string"
      },
      "tags": {
         "string" : "string"
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[assessment](#API_UpdateAssessment_ResponseSyntax)**

The response object for the `UpdateAssessment` API. This is the name of the
updated assessment.

Type: [Assessment](api-assessment.md) object

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

**ServiceQuotaExceededException**

You've reached your account quota for this resource type. To perform the requested
action, delete some existing resources or [request a quota increase](../../../../general/general/latest/gr/aws-service-limits.md) from
the Service Quotas console. For a list of Audit Manager service quotas, see [Quotas and\
restrictions for AWS Audit Manager](../../../../services/audit-manager/latest/userguide/service-quotas.md).

HTTP Status Code: 402

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/updateassessment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/updateassessment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UntagResource

UpdateAssessmentControl
