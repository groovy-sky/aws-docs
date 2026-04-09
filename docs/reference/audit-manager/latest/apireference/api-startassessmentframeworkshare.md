# StartAssessmentFrameworkShare

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Creates a share request for a custom framework in AWS Audit Manager.

The share request specifies a recipient and notifies them that a custom framework is
available. Recipients have 120 days to accept or decline the request. If no action is
taken, the share request expires.

When you create a share request, Audit Manager stores a snapshot of your custom
framework in the US East (N. Virginia) AWS Region. Audit Manager also
stores a backup of the same snapshot in the US West (Oregon) AWS Region.

Audit Manager deletes the snapshot and the backup snapshot when one of the following
events occurs:

- The sender revokes the share request.

- The recipient declines the share request.

- The recipient encounters an error and doesn't successfully accept the share
request.

- The share request expires before the recipient responds to the request.

When a sender [resends a share request](../../../../services/audit-manager/latest/userguide/framework-sharing.md#framework-sharing-resend), the snapshot is replaced with an updated version that
corresponds with the latest version of the custom framework.

When a recipient accepts a share request, the snapshot is replicated into their AWS account under the AWS Region that was specified in the share
request.

###### Important

When you invoke the `StartAssessmentFrameworkShare` API, you are about to
share a custom framework with another AWS account. You may not share a
custom framework that is derived from a standard framework if the standard framework is
designated as not eligible for sharing by AWS, unless you have obtained
permission to do so from the owner of the standard framework. To learn more about which
standard frameworks are eligible for sharing, see [Framework sharing eligibility](../../../../services/audit-manager/latest/userguide/share-custom-framework-concepts-and-terminology.md#eligibility) in the _AWS Audit Manager User_
_Guide_.

## Request Syntax

```nohighlight

POST /assessmentFrameworks/frameworkId/shareRequests HTTP/1.1
Content-type: application/json

{
   "comment": "string",
   "destinationAccount": "string",
   "destinationRegion": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[frameworkId](#API_StartAssessmentFrameworkShare_RequestSyntax)**

The unique identifier for the custom framework to be shared.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[comment](#API_StartAssessmentFrameworkShare_RequestSyntax)**

An optional comment from the sender about the share request.

Type: String

Length Constraints: Maximum length of 500.

Pattern: `^[\w\W\s\S]*$`

Required: No

**[destinationAccount](#API_StartAssessmentFrameworkShare_RequestSyntax)**

The AWS account of the recipient.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: Yes

**[destinationRegion](#API_StartAssessmentFrameworkShare_RequestSyntax)**

The AWS Region of the recipient.

Type: String

Pattern: `^[a-z]{2}-[a-z]+-[0-9]{1}$`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "assessmentFrameworkShareRequest": {
      "comment": "string",
      "complianceType": "string",
      "creationTime": number,
      "customControlsCount": number,
      "destinationAccount": "string",
      "destinationRegion": "string",
      "expirationTime": number,
      "frameworkDescription": "string",
      "frameworkId": "string",
      "frameworkName": "string",
      "id": "string",
      "lastUpdated": number,
      "sourceAccount": "string",
      "standardControlsCount": number,
      "status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[assessmentFrameworkShareRequest](#API_StartAssessmentFrameworkShare_ResponseSyntax)**

The share request that's created by the `StartAssessmentFrameworkShare` API.

Type: [AssessmentFrameworkShareRequest](api-assessmentframeworksharerequest.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/startassessmentframeworkshare.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/startassessmentframeworkshare.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegisterOrganizationAdminAccount

TagResource

All content copied from https://docs.aws.amazon.com/.
