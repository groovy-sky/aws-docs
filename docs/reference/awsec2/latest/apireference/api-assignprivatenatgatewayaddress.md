# AssignPrivateNatGatewayAddress

Assigns private IPv4 addresses to a private NAT gateway. For more information, see
[Work with NAT gateways](../../../../services/vpc/latest/userguide/nat-gateway-working-with.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

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

The private IPv4 addresses you want to assign to the private NAT gateway.

Type: Array of strings

Required: No

**PrivateIpAddressCount**

The number of private IP addresses to assign to the NAT gateway. You can't specify this parameter when also specifying private IP addresses.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 31.

Required: No

## Response Elements

The following elements are returned by the service.

**natGatewayAddressSet**

NAT gateway IP addresses.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssignPrivateNatGatewayAddress)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssignPrivateNatGatewayAddress)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/assignprivatenatgatewayaddress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/assignprivatenatgatewayaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/assignprivatenatgatewayaddress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/assignprivatenatgatewayaddress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/assignprivatenatgatewayaddress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/assignprivatenatgatewayaddress.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssignPrivateNatGatewayAddress)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/assignprivatenatgatewayaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssignPrivateIpAddresses

AssociateAddress
