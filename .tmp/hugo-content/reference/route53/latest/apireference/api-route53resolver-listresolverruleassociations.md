# ListResolverRuleAssociations

Lists the associations that were created between Resolver rules and VPCs using the current AWS account.

## Request Syntax

```nohighlight

{
   "Filters": [
      {
         "Name": "string",
         "Values": [ "string" ]
      }
   ],
   "MaxResults": number,
   "NextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Filters](#API_route53resolver_ListResolverRuleAssociations_RequestSyntax)**

An optional specification to return a subset of Resolver rules, such as Resolver rules that are associated with the same VPC ID.

###### Note

If you submit a second or subsequent `ListResolverRuleAssociations` request and specify the `NextToken` parameter,
you must use the same values for `Filters`, if any, as in the previous request.

Type: Array of [Filter](api-route53resolver-filter.md) objects

Required: No

**[MaxResults](#API_route53resolver_ListResolverRuleAssociations_RequestSyntax)**

The maximum number of rule associations that you want to return in the response to a `ListResolverRuleAssociations` request.
If you don't specify a value for `MaxResults`, Resolver returns up to 100 rule associations.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_route53resolver_ListResolverRuleAssociations_RequestSyntax)**

For the first `ListResolverRuleAssociation` request, omit this value.

If you have more than `MaxResults` rule associations, you can submit another `ListResolverRuleAssociation` request
to get the next group of rule associations. In the next request, specify the value of `NextToken` from the previous response.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "ResolverRuleAssociations": [
      {
         "Id": "string",
         "Name": "string",
         "ResolverRuleId": "string",
         "Status": "string",
         "StatusMessage": "string",
         "VPCId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[MaxResults](#API_route53resolver_ListResolverRuleAssociations_ResponseSyntax)**

The value that you specified for `MaxResults` in the request.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

**[NextToken](#API_route53resolver_ListResolverRuleAssociations_ResponseSyntax)**

If more than `MaxResults` rule associations match the specified criteria, you can submit another
`ListResolverRuleAssociation` request to get the next group of results. In the next request, specify the value of
`NextToken` from the previous response.

Type: String

**[ResolverRuleAssociations](#API_route53resolver_ListResolverRuleAssociations_ResponseSyntax)**

The associations that were created between Resolver rules and VPCs using the current AWS account, and that match the
specified filters, if any.

Type: Array of [ResolverRuleAssociation](api-route53resolver-resolverruleassociation.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

We encountered an unknown error. Try again in a few minutes.

HTTP Status Code: 400

**InvalidNextTokenException**

The value that you specified for `NextToken` in a `List` request isn't valid.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters in this request are not valid.

**FieldName**

For an `InvalidParameterException` error, the name of the parameter that's invalid.

HTTP Status Code: 400

**InvalidRequestException**

The request is invalid.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### ListResolverRuleAssociations Example

This example illustrates one usage of ListResolverRuleAssociations.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 2
X-Amz-Target: Route53Resolver.ListResolverRuleAssociations
X-Amz-Date: 20181101T193011Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 19:30:11 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 277
x-amzn-RequestId: 5ad8f14f-64c0-4cb0-a4f9-e5c48example
Connection: keep-alive

{
    "MaxResults": 30,
    "ResolverRuleAssociations": [
        {
            "Id": "rslvr-rrassoc-97242eaf88example",
            "Name":"forward example.com for gamma VPC",
            "ResolverRuleId": "rslvr-rr-5328a0899aexample",
            "Status": "COMPLETE",
            "StatusMessage": "[Trace id: 1-5bdb53db-75f3852c8384ad30fexample] Creating the association.",
            "VPCId": "vpc-03cf94c75eexample"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/listresolverruleassociations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/listresolverruleassociations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListResolverQueryLogConfigs

ListResolverRules
