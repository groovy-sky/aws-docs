# ListControlDomainInsights

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Lists the latest analytics data for control domains across all of your active
assessments.

Audit Manager supports the control domains that are provided by AWS
Control Catalog. For information about how to find a list of available control domains, see
[`ListDomains`](../../../controlcatalog/latest/apireference/api-listdomains.md) in the AWS Control
Catalog API Reference.

###### Note

A control domain is listed only if at least one of the controls within that domain
collected evidence on the `lastUpdated` date of
`controlDomainInsights`. If this condition isn’t met, no data is listed
for that control domain.

## Request Syntax

```nohighlight

GET /insights/control-domains?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_ListControlDomainInsights_RequestSyntax)**

Represents the maximum number of results on a page or for an API request call.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[nextToken](#API_ListControlDomainInsights_RequestSyntax)**

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
   "controlDomainInsights": [
      {
         "controlsCountByNoncompliantEvidence": number,
         "evidenceInsights": {
            "compliantEvidenceCount": number,
            "inconclusiveEvidenceCount": number,
            "noncompliantEvidenceCount": number
         },
         "id": "string",
         "lastUpdated": number,
         "name": "string",
         "totalControlsCount": number
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[controlDomainInsights](#API_ListControlDomainInsights_ResponseSyntax)**

The control domain analytics data that the `ListControlDomainInsights` API
returned.

Type: Array of [ControlDomainInsights](api-controldomaininsights.md) objects

**[nextToken](#API_ListControlDomainInsights_ResponseSyntax)**

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/listcontroldomaininsights.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/listcontroldomaininsights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAssessments

ListControlDomainInsightsByAssessment

All content copied from https://docs.aws.amazon.com/.
