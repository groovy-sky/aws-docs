# CreateSubnetCidrReservation

Creates a subnet CIDR reservation. For more information, see [Subnet CIDR reservations](../../../../services/vpc/latest/userguide/subnet-cidr-reservation.md)
in the _Amazon VPC User Guide_ and [Manage prefixes \
for your network interfaces](../../../../services/ec2/latest/userguide/work-with-prefixes.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Cidr**

The IPv4 or IPV6 CIDR range to reserve.

Type: String

Required: Yes

**Description**

The description to assign to the subnet CIDR reservation.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ReservationType**

The type of reservation. The reservation type determines how the reserved IP addresses are
assigned to resources.

- `prefix` \- AWS assigns the reserved IP addresses to
network interfaces.

- `explicit` \- You assign the reserved IP addresses to network interfaces.

Type: String

Valid Values: `prefix | explicit`

Required: Yes

**SubnetId**

The ID of the subnet.

Type: String

Required: Yes

**TagSpecification.N**

The tags to assign to the subnet CIDR reservation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**subnetCidrReservation**

Information about the created subnet CIDR reservation.

Type: [SubnetCidrReservation](api-subnetcidrreservation.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateSubnetCidrReservation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateSubnetCidrReservation)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createsubnetcidrreservation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createsubnetcidrreservation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createsubnetcidrreservation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createsubnetcidrreservation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createsubnetcidrreservation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createsubnetcidrreservation.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateSubnetCidrReservation)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createsubnetcidrreservation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateSubnet

CreateTags
