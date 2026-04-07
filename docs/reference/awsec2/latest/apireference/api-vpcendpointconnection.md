# VpcEndpointConnection

Describes a VPC endpoint connection to a service.

## Contents

**creationTimestamp**

The date and time that the VPC endpoint was created.

Type: Timestamp

Required: No

**DnsEntrySet.N**

The DNS entries for the VPC endpoint.

Type: Array of [DnsEntry](api-dnsentry.md) objects

Required: No

**GatewayLoadBalancerArnSet.N**

The Amazon Resource Names (ARNs) of the Gateway Load Balancers for the service.

Type: Array of strings

Required: No

**ipAddressType**

The IP address type for the endpoint.

Type: String

Valid Values: `ipv4 | dualstack | ipv6`

Required: No

**NetworkLoadBalancerArnSet.N**

The Amazon Resource Names (ARNs) of the network load balancers for the service.

Type: Array of strings

Required: No

**serviceId**

The ID of the service to which the endpoint is connected.

Type: String

Required: No

**TagSet.N**

The tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcEndpointConnectionId**

The ID of the VPC endpoint connection.

Type: String

Required: No

**vpcEndpointId**

The ID of the VPC endpoint.

Type: String

Required: No

**vpcEndpointOwner**

The ID of the AWS account that owns the VPC endpoint.

Type: String

Required: No

**vpcEndpointRegion**

The Region of the endpoint.

Type: String

Required: No

**vpcEndpointState**

The state of the VPC endpoint.

Type: String

Valid Values: `PendingAcceptance | Pending | Available | Deleting | Deleted | Rejected | Failed | Expired | Partial`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpcendpointconnection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpcendpointconnection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpcendpointconnection.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpcEndpointAssociation

VpcIpv6CidrBlockAssociation
