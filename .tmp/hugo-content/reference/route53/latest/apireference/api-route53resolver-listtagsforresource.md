# ListTagsForResource

Lists the tags that you associated with the specified resource.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "ResourceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_route53resolver_ListTagsForResource_RequestSyntax)**

The maximum number of tags that you want to return in the response to a `ListTagsForResource` request.
If you don't specify a value for `MaxResults`, Resolver returns up to 100 tags.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_route53resolver_ListTagsForResource_RequestSyntax)**

For the first `ListTagsForResource` request, omit this value.

If you have more than `MaxResults` tags, you can submit another `ListTagsForResource` request
to get the next group of tags for the resource. In the next request, specify the value of `NextToken` from the previous response.

Type: String

Required: No

**[ResourceArn](#API_route53resolver_ListTagsForResource_RequestSyntax)**

The Amazon Resource Name (ARN) for the resource that you want to list tags for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_route53resolver_ListTagsForResource_ResponseSyntax)**

If more than `MaxResults` tags match the specified criteria, you can submit another
`ListTagsForResource` request to get the next group of results. In the next request, specify the value of
`NextToken` from the previous response.

Type: String

**[Tags](#API_route53resolver_ListTagsForResource_ResponseSyntax)**

The tags that are associated with the resource that you specified in the `ListTagsForResource` request.

Type: Array of [Tag](api-route53resolver-tag.md) objects

Array Members: Maximum number of 200 items.

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

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### ListTagsForResource Example

This example illustrates one usage of ListTagsForResource.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 134
X-Amz-Target: Route53Resolver.ListTagsForResource
X-Amz-Date: 20181101T185222Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "MaxResults": 10,
    "ResourceArn": "arn:aws:route53resolver:us-east-2:123456789012:resolver-rule/rslvr-rr-5328a0899aexample"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 18:52:22 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 107
x-amzn-RequestId: bda80f7b-0f2c-41d1-9043-f36d3example
Connection: keep-alive

{
    "Tags": [
        {
            "Key": "account-id",
            "Value": "12345"
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/listtagsforresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/listtagsforresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListResolverRules

PutFirewallRuleGroupPolicy
