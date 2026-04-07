# GetResolverRuleAssociation

Gets information about an association between a specified Resolver rule and a VPC. You associate a Resolver rule and a VPC using
[AssociateResolverRule](api-route53resolver-associateresolverrule.md).

## Request Syntax

```nohighlight

{
   "ResolverRuleAssociationId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResolverRuleAssociationId](#API_route53resolver_GetResolverRuleAssociation_RequestSyntax)**

The ID of the Resolver rule association that you want to get information about.

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

**[ResolverRuleAssociation](#API_route53resolver_GetResolverRuleAssociation_ResponseSyntax)**

Information about the Resolver rule association that you specified in a `GetResolverRuleAssociation` request.

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

### GetResolverRuleAssociation Example

This example illustrates one usage of GetResolverRuleAssociation.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 48
X-Amz-Target: Route53Resolver.GetResolverRuleAssociation
X-Amz-Date: 20181101T192652Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "ResolverRuleAssociationId": "rslvr-rrassoc-97242eaf88example"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 19:26:52 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 314
x-amzn-RequestId: 1d7c8aad-1c63-413d-88a9-bc7b0example
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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/GetResolverRuleAssociation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/GetResolverRuleAssociation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetResolverRule

GetResolverRulePolicy
