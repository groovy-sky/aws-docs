# UpdateResolverEndpoint

Updates the name, or endpoint type for an inbound or an outbound Resolver endpoint.
You can only update between IPV4 and DUALSTACK, IPV6 endpoint type can't be updated to other type.

## Request Syntax

```nohighlight

{
   "Name": "string",
   "Protocols": [ "string" ],
   "ResolverEndpointId": "string",
   "ResolverEndpointType": "string",
   "RniEnhancedMetricsEnabled": boolean,
   "TargetNameServerMetricsEnabled": boolean,
   "UpdateIpAddresses": [
      {
         "IpId": "string",
         "Ipv6": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Name](#API_route53resolver_UpdateResolverEndpoint_RequestSyntax)**

The name of the Resolver endpoint that you want to update.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**[Protocols](#API_route53resolver_UpdateResolverEndpoint_RequestSyntax)**

The protocols you want to use for the endpoint. DoH-FIPS is applicable for default inbound endpoints only.

For a default inbound endpoint you can apply the protocols as follows:

- Do53 and DoH in combination.

- Do53 and DoH-FIPS in combination.

- Do53 alone.

- DoH alone.

- DoH-FIPS alone.

- None, which is treated as Do53.

For a delegation inbound endpoint you can use Do53 only.

For an outbound endpoint you can apply the protocols as follows:

- Do53 and DoH in combination.

- Do53 alone.

- DoH alone.

- None, which is treated as Do53.

###### Important

You can't change the protocol of an inbound endpoint directly from only Do53 to only DoH, or DoH-FIPS.
This is to prevent a sudden disruption to incoming traffic that
relies on Do53. To change the protocol from Do53 to DoH, or DoH-FIPS, you must
first enable both Do53 and DoH, or Do53 and DoH-FIPS, to make sure that all incoming traffic
has transferred to using the DoH protocol, or DoH-FIPS, and then remove the
Do53.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 2 items.

Valid Values: `DoH | Do53 | DoH-FIPS`

Required: No

**[ResolverEndpointId](#API_route53resolver_UpdateResolverEndpoint_RequestSyntax)**

The ID of the Resolver endpoint that you want to update.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[ResolverEndpointType](#API_route53resolver_UpdateResolverEndpoint_RequestSyntax)**

Specifies the endpoint type for what type of IP address the endpoint uses to forward DNS queries.

Updating to `IPV6` type isn't currently supported.

Type: String

Valid Values: `IPV6 | IPV4 | DUALSTACK`

Required: No

**[RniEnhancedMetricsEnabled](#API_route53resolver_UpdateResolverEndpoint_RequestSyntax)**

Updates whether RNI enhanced metrics are enabled for the Resolver endpoints.
When set to true, one-minute granular metrics are published in CloudWatch for each RNI associated with this endpoint.
When set to false, metrics are not published.

###### Note

Standard CloudWatch pricing and charges are applied for using the Route 53 Resolver
endpoint RNI enhanced metrics. For more information, see [Detailed metrics](../../../../services/route53/latest/developerguide/monitoring-resolver-with-cloudwatch.md).

Type: Boolean

Required: No

**[TargetNameServerMetricsEnabled](#API_route53resolver_UpdateResolverEndpoint_RequestSyntax)**

Updates whether target name server metrics are enabled for the outbound Resolver endpoints.
When set to true, one-minute granular metrics are published in CloudWatch for each target name server associated with this endpoint.
When set to false, metrics are not published. This setting is not supported for inbound Resolver endpoints.

###### Note

Standard CloudWatch pricing and charges are applied for using the Route 53 Resolver
endpoint target name server metrics. For more information, see [Detailed metrics](../../../../services/route53/latest/developerguide/monitoring-resolver-with-cloudwatch.md).

Type: Boolean

Required: No

**[UpdateIpAddresses](#API_route53resolver_UpdateResolverEndpoint_RequestSyntax)**

Specifies the IPv6 address when you update the Resolver endpoint from IPv4 to dual-stack.
If you don't specify an IPv6 address, one will be automatically chosen from your subnet.

Type: Array of [UpdateIpAddress](api-route53resolver-updateipaddress.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

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

**[ResolverEndpoint](#API_route53resolver_UpdateResolverEndpoint_ResponseSyntax)**

The response to an `UpdateResolverEndpoint` request.

Type: [ResolverEndpoint](api-route53resolver-resolverendpoint.md) object

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

### UpdateResolverEndpoint Example

This example illustrates one usage of UpdateResolverEndpoint.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 2
X-Amz-Target: Route53Resolver.UpdateResolverEndpoint
X-Amz-Date: 20181101T192600Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "Name":"MyInbound",
    "ResolverEndpointId": "rslvr-in-60b9fd8fdbexample"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Date: Thu, 01 Nov 2018 18:52:22 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 479
x-amzn-RequestId: bda80f7b-0f2c-41d1-9043-f36d3example
Connection: keep-alive

{
    "ResolverEndpoint":{
        "Arn":"arn:aws:route53resolver:us-east-2:0123456789012:resolver-endpoint/rslvr-in-60b9fd8fdbexample",
        "CreationTime":"2018-11-01T18:44:50.372Z",
        "CreatorRequestId":"1234",
        "Direction":"INBOUND",
        "HostVPCId":"vpc-03cf94c75cexample",
        "Id":"rslvr-in-60b9fd8fdbexample",
        "IpAddressCount":3,
        "ModificationTime":"2018-11-01T18:44:50.372Z",
        "Name":"MyInbound",
        "Protocols": [
            "DoH"
        ],
        "ResolverEndpointType": "IPV4",
        "RniEnhancedMetricsEnabled": false,
        "SecurityGroupIds":[
            "sg-020a3554aexample"
        ],
        "Status":"UPDATING",
        "StatusMessage":"Updating the Resolver Endpoint"
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/updateresolverendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/updateresolverendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateResolverDnssecConfig

UpdateResolverRule
