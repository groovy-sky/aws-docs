# CreateResolverEndpoint

Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound and outbound:

- An _inbound Resolver endpoint_ forwards DNS queries to the DNS service for a VPC
from your network.

- An _outbound Resolver endpoint_ forwards DNS queries from the DNS service for a VPC
to your network.

## Request Syntax

```nohighlight

{
   "CreatorRequestId": "string",
   "Direction": "string",
   "IpAddresses": [
      {
         "Ip": "string",
         "Ipv6": "string",
         "SubnetId": "string"
      }
   ],
   "Name": "string",
   "OutpostArn": "string",
   "PreferredInstanceType": "string",
   "Protocols": [ "string" ],
   "ResolverEndpointType": "string",
   "RniEnhancedMetricsEnabled": boolean,
   "SecurityGroupIds": [ "string" ],
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ],
   "TargetNameServerMetricsEnabled": boolean
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CreatorRequestId](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

A unique string that identifies the request and that allows failed requests to be retried
without the risk of running the operation twice. `CreatorRequestId` can be
any unique string, for example, a date/time stamp.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[Direction](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

Specify the applicable value:

- `INBOUND`: Resolver forwards DNS queries to the DNS service for a VPC from your network.

- `OUTBOUND`: Resolver forwards DNS queries from the DNS service for a VPC to your network.

- `INBOUND_DELEGATION`: Resolver delegates queries to Route 53 private hosted zones from your network.

Type: String

Valid Values: `INBOUND | OUTBOUND | INBOUND_DELEGATION`

Required: Yes

**[IpAddresses](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward
DNS queries to (for inbound endpoints). The subnet ID uniquely identifies a VPC.

###### Note

Even though the minimum is 1, Route 53 requires that you create at least two.

Type: Array of [IpAddressRequest](api-route53resolver-ipaddressrequest.md) objects

Array Members: Minimum number of 2 items. Maximum number of 20 items.

Required: Yes

**[Name](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**[OutpostArn](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

The Amazon Resource Name (ARN) of the Outpost. If you specify this, you must also specify a
value for the `PreferredInstanceType`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^arn:aws([a-z-]+)?:outposts:[a-z\d-]+:\d{12}:outpost/op-[a-f0-9]{17}$`

Required: No

**[PreferredInstanceType](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

The instance type. If you specify this, you must also specify a value for the `OutpostArn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**[Protocols](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

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

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 2 items.

Valid Values: `DoH | Do53 | DoH-FIPS`

Required: No

**[ResolverEndpointType](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

For the endpoint type you can choose either IPv4, IPv6, or dual-stack.
A dual-stack endpoint means that it will resolve via both IPv4 and IPv6. This
endpoint type is applied to all IP addresses.

Type: String

Valid Values: `IPV6 | IPV4 | DUALSTACK`

Required: No

**[RniEnhancedMetricsEnabled](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

Specifies whether RNI enhanced metrics are enabled for the Resolver endpoints.
When set to true, one-minute granular metrics are published in CloudWatch for each RNI associated with this endpoint.
When set to false, metrics are not published. Default is false.

###### Note

Standard CloudWatch pricing and charges are applied for using the Route 53 Resolver
endpoint RNI enhanced metrics. For more information, see [Detailed metrics](../../../../services/route53/latest/developerguide/monitoring-resolver-with-cloudwatch.md).

Type: Boolean

Required: No

**[SecurityGroupIds](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

The ID of one or more security groups that you want to use to control access to this VPC. The security group that you specify
must include one or more inbound rules (for inbound Resolver endpoints) or outbound rules (for outbound Resolver endpoints).
Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port
that you're using for DNS queries on your network.

Some security group rules will cause your connection to be tracked. For outbound resolver endpoint, it can potentially impact the
maximum queries per second from outbound endpoint to your target name server. For inbound resolver endpoint, it can bring down the overall maximum queries per second per IP address to as low as 1500.
To avoid connection tracking caused by security group, see
[Untracked connections](../../../../services/ec2/latest/userguide/security-group-connection-tracking.md#untracked-connectionsl).

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[Tags](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

A list of the tag keys and values that you want to associate with the endpoint.

Type: Array of [Tag](api-route53resolver-tag.md) objects

Array Members: Maximum number of 200 items.

Required: No

**[TargetNameServerMetricsEnabled](#API_route53resolver_CreateResolverEndpoint_RequestSyntax)**

Specifies whether target name server metrics are enabled for the outbound Resolver endpoints.
When set to true, one-minute granular metrics are published in CloudWatch for each target name server associated with this endpoint.
When set to false, metrics are not published. Default is false. This is not supported for inbound Resolver endpoints.

###### Note

Standard CloudWatch pricing and charges are applied for using the Route 53 Resolver
endpoint target name server metrics. For more information, see [Detailed metrics](../../../../services/route53/latest/developerguide/monitoring-resolver-with-cloudwatch.md).

Type: Boolean

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

**[ResolverEndpoint](#API_route53resolver_CreateResolverEndpoint_ResponseSyntax)**

Information about the `CreateResolverEndpoint` request, including the status of the request.

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

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## Examples

### CreateResolverEndpoint Example

This example illustrates one usage of CreateResolverEndpoint.

#### Sample Request

```

POST / HTTP/1.1
Host: route53resolver.us-east-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 283
X-Amz-Target: Route53Resolver.CreateResolverEndpoint
X-Amz-Date: 20181101T191344Z
User-Agent: aws-cli/1.16.45 Python/2.7.10 Darwin/16.7.0 botocore/1.12.35
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256
               Credential=AKIAJJ2SONIPEXAMPLE/20181101/us-east-2/route53resolver/aws4_request,
               SignedHeaders=content-type;host;x-amz-date;x-amz-target,
               Signature=[calculated-signature]

{
    "Direction": "OUTBOUND",
    "Name": "MyOutbound",
    "Tags": [
        {
            "Key": "LineOfBusiness",
            "Value": "Engineering"
        }
    ],
    "CreatorRequestId": "5678",
    "SecurityGroupIds": [
        "sg-071b99f42example"
    ],
    "IpAddresses": [
        {
            "SubnetId": "subnet-0bca4d363dexample"
        },
        {
            "SubnetId": "subnet-0bca4d363dexample"
        }
    ],
    "RniEnhancedMetricsEnabled": false,
    "TargetNameServerMetricsEnabled": true
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
        "Status": "CREATING",
        "StatusMessage": "[Trace id: 1-5bdb5068-e0bdc4d232b1a3fe9c344c10] Creating the Resolver Endpoint",
        "TargetNameServerMetricsEnabled": true
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/createresolverendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/createresolverendpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateOutpostResolver

CreateResolverQueryLogConfig
