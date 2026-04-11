# UpdateAssessmentFrameworkShare

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Updates a share request for a custom framework in AWS Audit Manager.

## Request Syntax

```nohighlight

PUT /assessmentFrameworkShareRequests/requestId HTTP/1.1
Content-type: application/json

{
   "action": "string",
   "requestType": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[requestId](#API_UpdateAssessmentFrameworkShare_RequestSyntax)**

The unique identifier for the share request.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[action](#API_UpdateAssessmentFrameworkShare_RequestSyntax)**

Specifies the update action for the share request.

Type: String

Valid Values: `ACCEPT | DECLINE | REVOKE`

Required: Yes

**[requestType](#API_UpdateAssessmentFrameworkShare_RequestSyntax)**

Specifies whether the share request is a sent request or a received request.

Type: String

Valid Values: `SENT | RECEIVED`

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

**[assessmentFrameworkShareRequest](#API_UpdateAssessmentFrameworkShare_ResponseSyntax)**

The updated share request that's returned by the
`UpdateAssessmentFrameworkShare` operation.

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

**ServiceQuotaExceededException**

You've reached your account quota for this resource type. To perform the requested
action, delete some existing resources or [request a quota increase](../../../../general/latest/gr/aws-service-limits.md) from
the Service Quotas console. For a list of Audit Manager service quotas, see [Quotas and\
restrictions for AWS Audit Manager](../../../../services/audit-manager/latest/userguide/service-quotas.md).

HTTP Status Code: 402

**ValidationException**

The request has invalid or missing parameters.

**fields**

The fields that caused the error, if applicable.

**reason**

The reason the request failed validation.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/updateassessmentframeworkshare.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAssessmentFramework

UpdateAssessmentStatus

All content copied from https://docs.aws.amazon.com/.
