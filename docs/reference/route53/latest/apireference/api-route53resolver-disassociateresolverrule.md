# DisassociateResolverRule

Removes the association between a specified Resolver rule and a specified VPC.

###### Important

If you disassociate a Resolver rule from a VPC, Resolver stops forwarding DNS queries for the
domain name that you specified in the Resolver rule.

## Request Syntax

```nohighlight

{
   "ResolverRuleId": "string",
   "VPCId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResolverRuleId](#API_route53resolver_DisassociateResolverRule_RequestSyntax)**

The ID of the Resolver rule that you want to disassociate from the specified VPC.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[VPCId](#API_route53resolver_DisassociateResolverRule_RequestSyntax)**

The ID of the VPC that you want to disassociate the Resolver rule from.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

{
   "ResolverRuleAssociation": {
      "Id": "string",
      "Name": "string",
      "ResolverRuleId": "string",
      "Status": "string",
      "StatusMessage": "string",
      "VPCId": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ResolverRuleAssociation](#API_route53resolver_DisassociateResolverRule_ResponseSyntax)**

Information about the `DisassociateResolverRule` request, including the status of the request.

Type: [ResolverRuleAssociation](api-route53resolver-resolverruleassociation.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

We encountered an unknown error. Try again in a few minutes.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters in this request are not valid.

**FieldName**

For an `InvalidParameterException` error, the name of the parameter that's invalid.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### DisassociateResolverRule Example

This example illustrates one usage of DisassociateResolverRule.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 82
X-Amz-Target: Route53Resolver.DisassociateResolverRule
X-Amz-Date: 20181101T192826Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "ResolverRuleId": "rslvr-rr-5328a0899aexample"
    "VPCId": "vpc-03cf94c75cexample",
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 19:28:27 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 258
x-amzn-RequestId: 9827eb40-87e6-44d6-bee7-a1116example
Connection: keep-alive

{
    "ResolverRuleAssociation":{
        "Id":"rslvr-rrassoc-97242eaf88example",
        "Name":"forward example.com for gamma VPC",
        "ResolverRuleId":"rslvr-rr-5328a0899example",
        "Status":"DELETING",
        "StatusMessage":"[Trace id: 1-5bdb53db-75f3852c8384ad30fexample], Deleting the association.",
        "VPCId":"vpc-03cf94c75cexample"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/disassociateresolverrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/disassociateresolverrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateResolverQueryLogConfig

GetFirewallConfig
