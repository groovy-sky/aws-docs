# DeleteResolverEndpoint

Deletes a Resolver endpoint. The effect of deleting a Resolver endpoint depends on whether it's an inbound or an outbound
Resolver endpoint:

- **Inbound**: DNS queries from your network are no longer routed
to the DNS service for the specified VPC.

- **Outbound**: DNS queries from a VPC are no longer routed to your network.

## Request Syntax

```nohighlight

{
   "ResolverEndpointId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResolverEndpointId](#API_route53resolver_DeleteResolverEndpoint_RequestSyntax)**

The ID of the Resolver endpoint that you want to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

{
   "ResolverEndpoint": {
      "Arn": "string",
      "CreationTime": "string",
      "CreatorRequestId": "string",
      "Direction": "string",
      "HostVPCId": "string",
      "Id": "string",
      "IpAddressCount": number,
      "ModificationTime": "string",
      "Name": "string",
      "OutpostArn": "string",
      "PreferredInstanceType": "string",
      "Protocols": [ "string" ],
      "ResolverEndpointType": "string",
      "RniEnhancedMetricsEnabled": boolean,
      "SecurityGroupIds": [ "string" ],
      "Status": "string",
      "StatusMessage": "string",
      "TargetNameServerMetricsEnabled": boolean
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ResolverEndpoint](#API_route53resolver_DeleteResolverEndpoint_ResponseSyntax)**

Information about the `DeleteResolverEndpoint` request, including the status of the request.

Type: [ResolverEndpoint](api-route53resolver-resolverendpoint.md) object

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

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### DeleteResolverEndpoint Example

This example illustrates one usage of DeleteResolverEndpoint.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 283
X-Amz-Target: Route53Resolver.DeleteResolverEndpoint
X-Amz-Date: 20181101T191344Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "ResolverEndpointId": "rslvr-out-fdc049932dexample"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 19:13:44 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 531
x-amzn-RequestId: 08afd081-9d67-4281-a277-b3880example
Connection: keep-alive

{
    "ResolverEndpoint": {
        "Arn": "arn:aws:route53resolver:us-east-2:123456789012:resolver-endpoint/rslvr-out-fdc049932dexample",
        "CreationTime": "2018-11-01T19:13:44.830Z",
        "CreatorRequestId": "5678",
        "Direction": "OUTBOUND",
        "HostVPCId": "vpc-0dd415a0edexample",
        "Id": "rslvr-out-fdc049932dexample",
        "IpAddressCount": 2,
        "ModificationTime": "2018-11-01T19:13:44.830Z",
        "Name": "MyOutbound",
        "Protocols": [
            "DoH"
        ],
        "ResolverEndpointType": "IPV4",
        "RniEnhancedMetricsEnabled": false,
        "SecurityGroupIds": [
            "sg-071b99f42example"
        ],
        "Status": "DELETING",
        "StatusMessage": "[Trace id: 1-5bdb5068-e0bdc4d232b1a3fe9c344c10] Successfully deleted Resolver Endpoint",
        "TargetNameServerMetricsEnabled": true
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/deleteresolverendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/deleteresolverendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteOutpostResolver

DeleteResolverQueryLogConfig
