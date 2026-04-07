# DisassociateNatGatewayAddress

Disassociates secondary Elastic IP addresses (EIPs) from a public NAT gateway.
You cannot disassociate your primary EIP. For more information, see [Edit secondary IP address associations](../../../../services/vpc/latest/userguide/nat-gateway-working-with.md#nat-gateway-edit-secondary) in the _Amazon VPC User Guide_.

While disassociating is in progress, you cannot associate/disassociate additional EIPs while the connections are being drained. You are, however, allowed to delete the NAT gateway.

An EIP is released only at the end of MaxDrainDurationSeconds. It stays
associated and supports the existing connections but does not support any new connections
(new connections are distributed across the remaining associated EIPs). As the existing
connections drain out, the EIPs (and the corresponding private IP addresses mapped to them)
are released.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationId.N**

The association IDs of EIPs that have been associated with the NAT gateway.

Type: Array of strings

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxDrainDurationSeconds**

The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 4000.

Required: No

**NatGatewayId**

The ID of the NAT gateway.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**natGatewayAddressSet**

Information about the NAT gateway IP addresses.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisassociateNatGatewayAddress)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisassociateNatGatewayAddress)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociatenatgatewayaddress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociatenatgatewayaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociatenatgatewayaddress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociatenatgatewayaddress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociatenatgatewayaddress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociatenatgatewayaddress.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisassociateNatGatewayAddress)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociatenatgatewayaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateIpamResourceDiscovery

DisassociateRouteServer
