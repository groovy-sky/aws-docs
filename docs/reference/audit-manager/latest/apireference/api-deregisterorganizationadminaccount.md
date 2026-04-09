# DeregisterOrganizationAdminAccount

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Removes the specified AWS account as a delegated administrator for
AWS Audit Manager.

When you remove a delegated administrator from your Audit Manager settings, you
continue to have access to the evidence that you previously collected under that account.
This is also the case when you deregister a delegated administrator from AWS Organizations. However, Audit Manager stops collecting and attaching evidence to
that delegated administrator account moving forward.

###### Important

Keep in mind the following cleanup task if you use evidence finder:

Before you use your management account to remove a delegated administrator, make sure
that the current delegated administrator account signs in to Audit Manager and
disables evidence finder first. Disabling evidence finder automatically deletes the
event data store that was created in their account when they enabled evidence finder. If
this task isn’t completed, the event data store remains in their account. In this case,
we recommend that the original delegated administrator goes to CloudTrail Lake
and manually [deletes the\
event data store](../../../../services/awscloudtrail/latest/userguide/query-eds-disable-termination.md).

This cleanup task is necessary to ensure that you don't end up with multiple event
data stores. Audit Manager ignores an unused event data store after you remove or
change a delegated administrator account. However, the unused event data store continues
to incur storage costs from CloudTrail Lake if you don't delete it.

When you deregister a delegated administrator account for Audit Manager, the data
for that account isn’t deleted. If you want to delete resource data for a delegated
administrator account, you must perform that task separately before you deregister the
account. Either, you can do this in the Audit Manager console. Or, you can use one of
the delete API operations that are provided by Audit Manager.

To delete your Audit Manager resource data, see the following instructions:

- [DeleteAssessment](api-deleteassessment.md) (see also: [Deleting an\
assessment](../../../../services/audit-manager/latest/userguide/delete-assessment.md) in the _AWS Audit Manager User_
_Guide_)

- [DeleteAssessmentFramework](api-deleteassessmentframework.md) (see also: [Deleting a\
custom framework](../../../../services/audit-manager/latest/userguide/delete-custom-framework.md) in the _AWS Audit Manager User_
_Guide_)

- [DeleteAssessmentFrameworkShare](api-deleteassessmentframeworkshare.md) (see also: [Deleting a share request](../../../../services/audit-manager/latest/userguide/deleting-shared-framework-requests.md) in the _AWS Audit Manager User_
_Guide_)

- [DeleteAssessmentReport](api-deleteassessmentreport.md) (see also: [Deleting an assessment report](../../../../services/audit-manager/latest/userguide/generate-assessment-report.md#delete-assessment-report-steps) in the _AWS Audit Manager User_
_Guide_)

- [DeleteControl](api-deletecontrol.md) (see also: [Deleting a custom\
control](../../../../services/audit-manager/latest/userguide/delete-controls.md) in the _AWS Audit Manager User_
_Guide_)

At this time, Audit Manager doesn't provide an option to delete evidence for a
specific delegated administrator. Instead, when your management account deregisters Audit Manager, we perform a cleanup for the current delegated administrator account at the
time of deregistration.

## Request Syntax

```nohighlight

POST /account/deregisterOrganizationAdminAccount HTTP/1.1
Content-type: application/json

{
   "adminAccountId": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[adminAccountId](#API_DeregisterOrganizationAdminAccount_RequestSyntax)**

The identifier for the administrator account.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: No

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/deregisterorganizationadminaccount.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeregisterAccount

DisassociateAssessmentReportEvidenceFolder

All content copied from https://docs.aws.amazon.com/.
