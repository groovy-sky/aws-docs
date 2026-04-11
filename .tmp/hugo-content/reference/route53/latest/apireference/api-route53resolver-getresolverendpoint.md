# GetResolverEndpoint

Gets information about a specified Resolver endpoint, such as whether it's an inbound or an outbound Resolver endpoint, and the
current status of the endpoint.

## Request Syntax

```nohighlight

{
   "ResolverEndpointId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResolverEndpointId](#API_route53resolver_GetResolverEndpoint_RequestSyntax)**

The ID of the Resolver endpoint that you want to get information about.

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

**[ResolverEndpoint](#API_route53resolver_GetResolverEndpoint_ResponseSyntax)**

Information about the Resolver endpoint that you specified in a `GetResolverEndpoint` request.

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

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### GetResolverEndpoint Example

This example illustrates one usage of GetResolverEndpoint.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 52
X-Amz-Target: Route53Resolver.GetResolverEndpoint
X-Amz-Date: 20181101T184736Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{"ResolverEndpointId": "rslvr-in-60b9fd8fdbexample"}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 18:47:36 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 490
x-amzn-RequestId: 98d13818-b790-4aa8-8ed8-87a1eexample
Connection: keep-alive

{
    "ResolverEndpoint": {
        "Arn": "arn:aws:route53resolver:us-east-2:123456789012:resolver-endpoint/rslvr-in-60b9fd8fdbexample",
        "CreationTime": "2018-11-01T18:44:50.372Z",
        "CreatorRequestId": "1234",
        "Direction": "INBOUND",
        "HostVPCId": "vpc-03cf94c75cexample",
        "Id": "rslvr-in-60b9fd8fdbexample",
        "IpAddressCount": 2,
        "ModificationTime": "2018-11-01T18:44:50.372Z",
        "Name": "MyInbound",
        "Protocols": [
           "DoH"
        ],
        "ResolverEndpointType": "IPV4",
        "RniEnhancedMetricsEnabled": false,
        "SecurityGroupIds": [
            "sg-020a3554aexample"
        ],
        "Status": "OPERATIONAL",
        "StatusMessage": "This Resolver Endpoint is operational."
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/getresolverendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/getresolverendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetResolverDnssecConfig

GetResolverQueryLogConfig
