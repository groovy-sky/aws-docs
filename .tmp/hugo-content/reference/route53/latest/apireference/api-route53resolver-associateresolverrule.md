# AssociateResolverRule

Associates a Resolver rule with a VPC. When you associate a rule with a VPC, Resolver forwards all DNS queries
for the domain name that is specified in the rule and that originate in the VPC. The queries are forwarded to the
IP addresses for the DNS resolvers that are specified in the rule. For more information about rules, see
[CreateResolverRule](api-route53resolver-createresolverrule.md).

## Request Syntax

```nohighlight

{
   "Name": "string",
   "ResolverRuleId": "string",
   "VPCId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Name](#API_route53resolver_AssociateResolverRule_RequestSyntax)**

A name for the association that you're creating between a Resolver rule and a VPC.

The name can be up to 64 characters long and can contain letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (\_), and spaces. The name cannot consist of only numbers.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**[ResolverRuleId](#API_route53resolver_AssociateResolverRule_RequestSyntax)**

The ID of the Resolver rule that you want to associate with the VPC. To list the existing Resolver rules, use
[ListResolverRules](api-route53resolver-listresolverrules.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[VPCId](#API_route53resolver_AssociateResolverRule_RequestSyntax)**

The ID of the VPC that you want to associate the Resolver rule with.

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

**[ResolverRuleAssociation](#API_route53resolver_AssociateResolverRule_ResponseSyntax)**

Information about the `AssociateResolverRule` request, including the status of the request.

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

**InvalidRequestException**

The request is invalid.

HTTP Status Code: 400

**LimitExceededException**

The request caused one or more limits to be exceeded.

**ResourceType**

For a `LimitExceededException` error, the type of resource that exceeded the current limit.

HTTP Status Code: 400

**ResourceExistsException**

The resource that you tried to create already exists.

**ResourceType**

For a `ResourceExistsException` error, the type of resource that the error applies to.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ResourceUnavailableException**

The specified resource isn't available.

**ResourceType**

For a `ResourceUnavailableException` error, the type of resource that isn't available.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### AssociateResolverRule Example

This example illustrates one usage of AssociateResolverRule.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 82
X-Amz-Target: Route53Resolver.AssociateResolverRule
X-Amz-Date: 20181101T192826Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
   "VPCId": "vpc-03cf94c75cexample",
   "ResolverRuleId": "rslvr-rr-5328a0899aexample"
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
        "Status":"CREATING",
        "StatusMessage":"[Trace id: 1-5bdb53db-75f3852c8384ad30fexample], Creating the association.",
        "VPCId":"vpc-03cf94c75cexample"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/associateresolverrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/associateresolverrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateResolverQueryLogConfig

CreateFirewallDomainList
