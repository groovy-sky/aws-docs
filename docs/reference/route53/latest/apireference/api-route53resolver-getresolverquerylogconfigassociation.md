# GetResolverQueryLogConfigAssociation

Gets information about a specified association between a Resolver query logging configuration and an Amazon VPC. When you associate a VPC
with a query logging configuration, Resolver logs DNS queries that originate in that VPC.

## Request Syntax

```nohighlight

{
   "ResolverQueryLogConfigAssociationId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResolverQueryLogConfigAssociationId](#API_route53resolver_GetResolverQueryLogConfigAssociation_RequestSyntax)**

The ID of the Resolver query logging configuration association that you want to get information about.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

{
   "ResolverQueryLogConfigAssociation": {
      "CreationTime": "string",
      "Error": "string",
      "ErrorMessage": "string",
      "Id": "string",
      "ResolverQueryLogConfigId": "string",
      "ResourceId": "string",
      "Status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ResolverQueryLogConfigAssociation](#API_route53resolver_GetResolverQueryLogConfigAssociation_ResponseSyntax)**

Information about the Resolver query logging configuration association that you specified in a `GetQueryLogConfigAssociation` request.

Type: [ResolverQueryLogConfigAssociation](api-route53resolver-resolverquerylogconfigassociation.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified Resolver operation.

This error can also be thrown when a customer has reached the 5120 character limit for a
resource policy for CloudWatch Logs.

HTTP Status Code: 400

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

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### GetResolverQueryLogConfigAssociation Example

This example illustrates one usage of GetResolverQueryLogConfigAssociation.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 48
X-Amz-Target: Route53Resolver.GetResolverQueryLogConfigAssociation
X-Amz-Date: 20200731T192652Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "ResolverQueryLogConfigAssociationId": "rqlca-ff85e1ffexample"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 19:26:52 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 464
x-amzn-RequestId: 1d7c8aad-1c63-413d-88a9-bc7b0example
Connection: keep-alive

{
    "ResolverQueryLogConfigAssociation":{
        "CreationTime": "2020-07-31T18:52:23.30Z",
        "Error": "",
        "ErrorMessage": "",
        "Id": "rqlca-ff85e1ffexample",
        "ResolverQueryLogConfigId": "rqlc-8ca61fe7cexample",
        "Resource": "vpc-03cf94c75cexample",
        "Status": "CREATED"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/getresolverquerylogconfigassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetResolverQueryLogConfig

GetResolverQueryLogConfigPolicy
