# GetInsights

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Gets the latest analytics data for all your current active assessments.

## Request Syntax

```

GET /insights HTTP/1.1

```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "insights": {
      "activeAssessmentsCount": number,
      "assessmentControlsCountByNoncompliantEvidence": number,
      "compliantEvidenceCount": number,
      "inconclusiveEvidenceCount": number,
      "lastUpdated": number,
      "noncompliantEvidenceCount": number,
      "totalAssessmentControlsCount": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[insights](#API_GetInsights_ResponseSyntax)**

The analytics data that the `GetInsights` API returned.

Type: [Insights](api-insights.md) object

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/getinsights.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/getinsights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetEvidenceFoldersByAssessmentControl

GetInsightsByAssessment

All content copied from https://docs.aws.amazon.com/.
