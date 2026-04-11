# ListAssessmentFrameworks

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Returns a list of the frameworks that are available in the Audit Manager framework
library.

## Request Syntax

```nohighlight

GET /assessmentFrameworks?frameworkType=frameworkType&maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[frameworkType](#API_ListAssessmentFrameworks_RequestSyntax)**

The type of framework, such as a standard framework or a custom framework.

Valid Values: `Standard | Custom`

Required: Yes

**[maxResults](#API_ListAssessmentFrameworks_RequestSyntax)**

Represents the maximum number of results on a page or for an API request call.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[nextToken](#API_ListAssessmentFrameworks_RequestSyntax)**

The pagination token that's used to fetch the next set of results.

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `^[A-Za-z0-9+\/=]*$`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "frameworkMetadataList": [
      {
         "arn": "string",
         "complianceType": "string",
         "controlsCount": number,
         "controlSetsCount": number,
         "createdAt": number,
         "description": "string",
         "id": "string",
         "lastUpdatedAt": number,
         "logo": "string",
         "name": "string",
         "type": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[frameworkMetadataList](#API_ListAssessmentFrameworks_ResponseSyntax)**

A list of metadata that the `ListAssessmentFrameworks` API returns for each
framework.

Type: Array of [AssessmentFrameworkMetadata](api-assessmentframeworkmetadata.md) objects

**[nextToken](#API_ListAssessmentFrameworks_ResponseSyntax)**

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

### Retrieving a list of custom frameworks

This shows a sample response that the `ListAssessmentFrameworks`
operation might return if you had two custom frameworks in your AWS account.

#### Sample Response

```json

{
    "frameworkMetadataList": [
        {
            "arn": "arn:aws:auditmanager:us-east-1:111122223333:assessmentFramework/a1b2c3d4-5678-90ab-cdef-example11111",
            "id": "a1b2c3d4-5678-90ab-cdef-example11111",
            "type": "Custom",
            "name": "My custom framework",
            "description": "My custom framework for internal audits",
            "complianceType": "Internal",
            "controlsCount": 3,
            "controlSetsCount": 1,
            "createdAt": "2022-03-29T14:57:31.634000-07:00",
            "lastUpdatedAt": "2022-03-29T14:57:31.635000-07:00"
        },
        {
            "arn": "arn:aws:auditmanager:us-east-1:111122223333:assessmentFramework/a1b2c3d4-5678-90ab-cdef-example22222",
            "id": "a1b2c3d4-5678-90ab-cdef-example22222",
            "type": "Custom",
            "name": "My custom HIPAA framework",
            "description": "My custom framework for HIPAA audits",
            "complianceType": "HIPAA",
            "controlsCount": 4,
            "controlSetsCount": 2,
            "createdAt": "2021-03-02T16:35:24.177000-08:00",
            "lastUpdatedAt": "2022-01-10T14:32:18.855000-08:00"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/listassessmentframeworks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/listassessmentframeworks.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAssessmentControlInsightsByControlDomain

ListAssessmentFrameworkShareRequests

All content copied from https://docs.aws.amazon.com/.
