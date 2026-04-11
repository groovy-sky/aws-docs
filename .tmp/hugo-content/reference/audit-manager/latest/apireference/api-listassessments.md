# ListAssessments

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Returns a list of current and past assessments from AWS Audit Manager.

## Request Syntax

```nohighlight

GET /assessments?maxResults=maxResults&nextToken=nextToken&status=status HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_ListAssessments_RequestSyntax)**

Represents the maximum number of results on a page or for an API request call.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[nextToken](#API_ListAssessments_RequestSyntax)**

The pagination token that's used to fetch the next set of results.

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `^[A-Za-z0-9+\/=]*$`

**[status](#API_ListAssessments_RequestSyntax)**

The current status of the assessment.

Valid Values: `ACTIVE | INACTIVE`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "assessmentMetadata": [
      {
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
         "id": "string",
         "lastUpdated": number,
         "name": "string",
         "roles": [
            {
               "roleArn": "string",
               "roleType": "string"
            }
         ],
         "status": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[assessmentMetadata](#API_ListAssessments_ResponseSyntax)**

The metadata that the `ListAssessments` API returns for each
assessment.

Type: Array of [AssessmentMetadataItem](api-assessmentmetadataitem.md) objects

**[nextToken](#API_ListAssessments_ResponseSyntax)**

The pagination token that's used to fetch the next set of results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `^[A-Za-z0-9+\/=]*$`

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

**ValidationException**

The request has invalid or missing parameters.

**fields**

The fields that caused the error, if applicable.

**reason**

The reason the request failed validation.

HTTP Status Code: 400

## Examples

### Retrieving a list of assessments

This shows a sample response that the `ListAssessments` operation might
return if you had three active assessments.

#### Sample Response

```json

{
    "assessmentMetadata": [
        {
            "name": "my GxP assessment",
            "id": "a1b2c3d4-5678-90ab-cdef-example11111",
            "complianceType": "GxP 21 CFR Part 11",
            "status": "ACTIVE",
            "roles": [
                {
                    "roleType": "PROCESS_OWNER",
                    "roleArn": "arn:aws:iam::111122223333:role/Administrator"
                }
            ],
            "delegations": [],
            "creationTime": "2022-12-20T16:38:09.215000-08:00",
            "lastUpdated": "2022-12-20T16:38:09.397000-08:00"
        },
        {
            "name": "My control tower assessment",
            "id": "a1b2c3d4-5678-90ab-cdef-example22222",
            "complianceType": "AWS Control Tower Guardrails",
            "status": "ACTIVE",
            "roles": [
                {
                    "roleType": "PROCESS_OWNER",
                    "roleArn": "arn:aws:iam::111122223333:role/Administrator"
                }
            ],
            "delegations": [],
            "creationTime": "2022-02-23T10:25:41.083000-08:00",
            "lastUpdated": "2022-02-23T17:06:54.132000-08:00"
        },
        {
            "name": "My PCI DSS assessment",
            "id": "a1b2c3d4-5678-90ab-cdef-example33333",
            "complianceType": "PCI DSS",
            "status": "ACTIVE",
            "roles": [
                {
                    "roleType": "PROCESS_OWNER",
                    "roleArn": "arn:aws:iam::111122223333:role/Administrator"
                }
            ],
            "delegations": [],
            "creationTime": "2022-02-23T09:34:42.850000-08:00",
            "lastUpdated": "2022-02-23T09:34:42.982000-08:00"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/listassessments.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/listassessments.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAssessmentReports

ListControlDomainInsights

All content copied from https://docs.aws.amazon.com/.
