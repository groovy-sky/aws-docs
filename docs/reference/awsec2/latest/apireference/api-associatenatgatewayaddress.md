# AssociateNatGatewayAddress

Associates Elastic IP addresses (EIPs) and private IPv4 addresses with a public NAT gateway. For more information,
see [Work with NAT gateways](../../../../services/vpc/latest/userguide/nat-gateway-working-with.md) in the _Amazon VPC User Guide_.

By default, you can associate up to 2 Elastic IP addresses per public NAT gateway. You can increase the limit by requesting a quota adjustment.
For more information, see [Elastic IP address quotas](../../../../services/vpc/latest/userguide/amazon-vpc-limits.md#vpc-limits-eips) in the _Amazon VPC User Guide_.

###### Important

When you associate an EIP or secondary EIPs with a public NAT gateway, the network border group of the EIPs
must match the network border group of the Availability Zone (AZ) that the public NAT gateway is in. If it's not the same,
the EIP will fail to associate. You can see the network border group for the subnet's AZ by viewing the details of the subnet.
Similarly, you can view the network border group of an EIP by viewing the details of the EIP address. For more information
about network border groups and EIPs, see [Allocate an Elastic IP address](../../../../services/vpc/latest/userguide/workwitheips.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllocationId.N**

The allocation IDs of EIPs that you want to associate with your NAT gateway.

Type: Array of strings

Required: Yes

**AvailabilityZone**

For regional NAT gateways only: The Availability Zone where you want to associate an Elastic IP address (EIP). The regional NAT gateway uses a separate EIP in each AZ to handle outbound NAT traffic from that AZ.

A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

Type: String

Required: No

**AvailabilityZoneId**

For regional NAT gateways only: The ID of the Availability Zone where you want to associate an Elastic IP address (EIP). The regional NAT gateway uses a separate EIP in each AZ to handle outbound NAT traffic from that AZ. Use this instead of AvailabilityZone for consistent identification of AZs across AWS Regions.

A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**NatGatewayId**

The ID of the NAT gateway.

Type: String

Required: Yes

**PrivateIpAddress.N**

The private IPv4 addresses that you want to assign to the NAT gateway.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**natGatewayAddressSet**

The IP addresses.

Type: Array of [NatGatewayAddress](api-natgatewayaddress.md) objects

**natGatewayId**

The ID of the NAT gateway.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/associatenatgatewayaddress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associatenatgatewayaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateIpamResourceDiscovery

AssociateRouteServer
