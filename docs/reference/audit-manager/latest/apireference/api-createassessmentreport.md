# CreateAssessmentReport

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Creates an assessment report for the specified assessment.

## Request Syntax

```nohighlight

POST /assessments/assessmentId/reports HTTP/1.1
Content-type: application/json

{
   "description": "string",
   "name": "string",
   "queryStatement": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[assessmentId](#API_CreateAssessmentReport_RequestSyntax)**

The identifier for the assessment.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[description](#API_CreateAssessmentReport_RequestSyntax)**

The description of the assessment report.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**[name](#API_CreateAssessmentReport_RequestSyntax)**

The name of the new assessment report.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[a-zA-Z0-9-_\.]+$`

Required: Yes

**[queryStatement](#API_CreateAssessmentReport_RequestSyntax)**

A SQL statement that represents an evidence finder query.

Provide this parameter when you want to generate an assessment report from the results
of an evidence finder search query. When you use this parameter, Audit Manager
generates a one-time report using only the evidence from the query output. This report does
not include any assessment evidence that was manually [added to a report using the console](https://docs.aws.amazon.com/audit-manager/latest/userguide/generate-assessment-report.html#generate-assessment-report-include-evidence), or [associated with a report using the API](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_BatchAssociateAssessmentReportEvidence.html).

To use this parameter, the [enablementStatus](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_EvidenceFinderEnablement.html#auditmanager-Type-EvidenceFinderEnablement-enablementStatus) of evidence finder must be `ENABLED`.

For examples and help resolving `queryStatement` validation exceptions, see
[Troubleshooting evidence finder issues](../../../../services/audit-manager/latest/userguide/evidence-finder-issues.md#querystatement-exceptions) in the
_AWS Audit Manager User Guide._

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10000.

Pattern: `(?s).*`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "assessmentReport": {
      "assessmentId": "string",
      "assessmentName": "string",
      "author": "string",
      "awsAccountId": "string",
      "creationTime": number,
      "description": "string",
      "id": "string",
      "name": "string",
      "status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[assessmentReport](#API_CreateAssessmentReport_ResponseSyntax)**

The new assessment report that the `CreateAssessmentReport` API returned.

Type: [AssessmentReport](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_AssessmentReport.html) object

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/auditmanager-2017-07-25/CreateAssessmentReport)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/CreateAssessmentReport)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateAssessmentFramework

CreateControl
