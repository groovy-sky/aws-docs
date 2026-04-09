# ListControls

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Returns a list of controls from AWS Audit Manager.

## Request Syntax

```nohighlight

GET /controls?controlCatalogId=controlCatalogId&controlType=controlType&maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[controlCatalogId](#API_ListControls_RequestSyntax)**

A filter that narrows the list of controls to a specific resource from the AWS
Control Catalog.

To use this parameter, specify the ARN of the Control Catalog resource. You can specify either
a control domain, a control objective, or a common control. For information about how to find the ARNs
for these resources, see [`ListDomains`](../../../controlcatalog/latest/apireference/api-listdomains.md), [`ListObjectives`](../../../controlcatalog/latest/apireference/api-listobjectives.md), and [`ListCommonControls`](../../../controlcatalog/latest/apireference/api-listcommoncontrols.md).

###### Note

You can only filter by one Control Catalog resource at a time.
Specifying multiple resource ARNs isn’t currently supported. If you want to filter by more
than one ARN, we recommend that you run the `ListControls` operation separately
for each ARN.

Alternatively, specify `UNCATEGORIZED` to list controls that aren't
mapped to a Control Catalog resource. For example, this operation might return a list of
custom controls that don't belong to any control domain or control objective.

Length Constraints: Minimum length of 13. Maximum length of 2048.

Pattern: `^arn:.*:controlcatalog:.*|UNCATEGORIZED`

**[controlType](#API_ListControls_RequestSyntax)**

A filter that narrows the list of controls to a specific type.

Valid Values: `Standard | Custom | Core`

Required: Yes

**[maxResults](#API_ListControls_RequestSyntax)**

The maximum number of results on a page or for an API request call.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[nextToken](#API_ListControls_RequestSyntax)**

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
   "controlMetadataList": [
      {
         "arn": "string",
         "controlSources": "string",
         "createdAt": number,
         "id": "string",
         "lastUpdatedAt": number,
         "name": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[controlMetadataList](#API_ListControls_ResponseSyntax)**

A list of metadata that the `ListControls` API returns for each
control.

Type: Array of [ControlMetadata](api-controlmetadata.md) objects

**[nextToken](#API_ListControls_ResponseSyntax)**

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

### Retrieving a list of custom controls

This shows a sample response that the `ListControls` operation might
return if you had three custom controls.

#### Sample Response

```json

{
    "controlMetadataList": [
        {
            "arn": "arn:aws:auditmanager:us-east-1:111122223333:control/a1b2c3d4-5678-90ab-cdef-example11111",
            "id": "a1b2c3d4-5678-90ab-cdef-example11111",
            "name": "Example custom control",
            "controlSources": "AWS API calls, AWS Config, Manual",
            "createdAt": "2022-03-28T18:08:24.446000-07:00",
            "lastUpdatedAt": "2022-03-28T18:08:24.446000-07:00"
        },
        {
            "arn": "arn:aws:auditmanager:us-east-1:111122223333:control/a1b2c3d4-5678-90ab-cdef-example22222",
            "id": "a1b2c3d4-5678-90ab-cdef-example22222",
            "name": "Another custom control",
            "controlSources": "AWS API calls, AWS Config, Manual",
            "createdAt": "2022-03-28T17:45:15.622000-07:00",
            "lastUpdatedAt": "2022-03-28T18:06:56.171000-07:00"
        },
        {
            "arn": "arn:aws:auditmanager:us-east-1:111122223333:control/a1b2c3d4-5678-90ab-cdef-example33333",
            "id": "a1b2c3d4-5678-90ab-cdef-example33333",
            "name": "Custom control - CloudTrail Volume Events",
            "controlSources": "AWS CloudTrail",
            "createdAt": "2021-03-08T11:31:12.468000-08:00",
            "lastUpdatedAt": "2022-01-10T14:13:37.979000-08:00"
        }
    ]
}
```

### Retrieving a list of core controls

This example shows how you can use the `ListControls` operation to
return a filtered list of core controls that fall under a specific control domain
( _Asset management_).

###### To retrieve a list of core controls

1. Use the [`ListDomains`](../../../controlcatalog/latest/apireference/api-listdomains.md) operation to see the control domains that
    you can use as filters. Find the control domain that you want to use as a
    filter ( _Asset management_), and take note of its
    ARN.

2. Use the `ListControls` operation and use the following attributes
    as filters:

- Use `ControlType` and specify `Core`
as the value.

- Use `ControlCatalogId` and specify the control
domain ARN that you got from step 1.

The sample request uses the following ARN as a filter:

- This ARN represents the Asset management control domain:
`arn:aws:controlcatalog:::domain/d4msesd9vvmzmmuvlv06m92uq`

The sample response shows the results that the `ListControls` operation
might return if two core controls matched the domain criteria of _Asset_
_management_.

#### Sample Request

```json

aws auditmanager list-controls --controlType Core --controlCatalogId "arn:aws:controlcatalog:::domain/d4msesd9vvmzmmuvlv06m92uq"

```

#### Sample Response

```json

{
   "controlMetadataList": [
    {
        "arn": "arn:aws:auditmanager:us-west-2::control/a1b2c3d4-5678-90ab-cdef-example11111",
        "controlSources": "Manual",
        "createdAt": 1.681511681787E9,
        "id": "a1b2c3d4-5678-90ab-cdef-example11111",
        "lastUpdatedAt": 1.681511681788E9,
        "name": "164.312(a)(2)(iii)"
    }, {
        "arn": "arn:aws:auditmanager:us-west-2::control/a1b2c3d4-5678-90ab-cdef-example22222",
        "controlSources": "AWS API calls, AWS Config, AWS Security Hub",
        "createdAt": 1.681511681787E9,
        "id": "a1b2c3d4-5678-90ab-cdef-example22222",
        "lastUpdatedAt": 1.681511681788E9,
        "name": "164.312(c)(1)"
   } ],
   "nextToken": "string"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/listcontrols.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/listcontrols.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListControlInsightsByControlDomain

ListKeywordsForDataSource

All content copied from https://docs.aws.amazon.com/.
