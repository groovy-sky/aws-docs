# DeleteAssessmentReport

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Deletes an assessment report in AWS Audit Manager.

When you run the `DeleteAssessmentReport` operation, Audit Manager
attempts to delete the following data:

1. The specified assessment report that’s stored in your S3 bucket

2. The associated metadata that’s stored in Audit Manager

If Audit Manager can’t access the assessment report in your S3 bucket, the report
isn’t deleted. In this event, the `DeleteAssessmentReport` operation doesn’t
fail. Instead, it proceeds to delete the associated metadata only. You must then delete the
assessment report from the S3 bucket yourself.

This scenario happens when Audit Manager receives a `403 (Forbidden)` or
`404 (Not Found)` error from Amazon S3. To avoid this, make sure that
your S3 bucket is available, and that you configured the correct permissions for Audit Manager to delete resources in your S3 bucket. For an example permissions policy that
you can use, see [Assessment report destination permissions](../../../../services/audit-manager/latest/userguide/security-iam-id-based-policy-examples.md#full-administrator-access-assessment-report-destination) in the _AWS Audit Manager User Guide_. For information about the issues that could cause a `403
            (Forbidden)` or `404 (Not Found`) error from Amazon S3, see
[List of Error Codes](../../../../services/s3/latest/api/errorresponses.md#ErrorCodeList) in the _Amazon Simple Storage Service API_
_Reference_.

## Request Syntax

```nohighlight

DELETE /assessments/assessmentId/reports/assessmentReportId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[assessmentId](#API_DeleteAssessmentReport_RequestSyntax)**

The unique identifier for the assessment.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

**[assessmentReportId](#API_DeleteAssessmentReport_RequestSyntax)**

The unique identifier for the assessment report.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/deleteassessmentreport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/deleteassessmentreport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAssessmentFrameworkShare

DeleteControl

All content copied from https://docs.aws.amazon.com/.
