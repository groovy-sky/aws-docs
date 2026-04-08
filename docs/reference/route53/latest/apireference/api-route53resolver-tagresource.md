# TagResource

Adds one or more tags to a specified resource.

## Request Syntax

```nohighlight

{
   "ResourceArn": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResourceArn](#API_route53resolver_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) for the resource that you want to add tags to. To get the ARN for a resource, use the applicable
`Get` or `List` command:

- [GetResolverEndpoint](api-route53resolver-getresolverendpoint.md)

- [GetResolverRule](api-route53resolver-getresolverrule.md)

- [GetResolverRuleAssociation](api-route53resolver-getresolverruleassociation.md)

- [ListResolverEndpoints](api-route53resolver-listresolverendpoints.md)

- [ListResolverRuleAssociations](api-route53resolver-listresolverruleassociations.md)

- [ListResolverRules](api-route53resolver-listresolverrules.md)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[Tags](#API_route53resolver_TagResource_RequestSyntax)**

The tags that you want to add to the specified resource.

Type: Array of [Tag](api-route53resolver-tag.md) objects

Array Members: Maximum number of 200 items.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

**InvalidTagException**

The specified tag is invalid.

HTTP Status Code: 400

**LimitExceededException**

The request caused one or more limits to be exceeded.

**ResourceType**

For a `LimitExceededException` error, the type of resource that exceeded the current limit.

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

### TagResource Example

This example illustrates one usage of TagResource.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 217
X-Amz-Target: Route53Resolver.TagResource
X-Amz-Date: 20181101T185222Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "ResourceArn": "arn:aws:route53resolver:us-east-2:123456789012:resolver-rule/rslvr-rr-5328a0899aexample",
    "Tags": [
        {
            "Key": "account-id",
            "Value": "12345"
        }
    ]
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 18:52:22 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 2
x-amzn-RequestId: bda80f7b-0f2c-41d1-9043-f36d3example
Connection: keep-alive

{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/tagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/tagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PutResolverRulePolicy

UntagResource
