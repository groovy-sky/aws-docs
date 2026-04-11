# ProvisionByoipCidr

Provisions an IPv4 or IPv6 address range for use with your AWS resources through bring your own IP
addresses (BYOIP) and creates a corresponding address pool. After the address range is
provisioned, it is ready to be advertised.

AWS verifies that you own the address range and are authorized to advertise it.
You must ensure that the address range is registered to you and that you created an
RPKI ROA to authorize Amazon ASNs 16509 and 14618 to advertise the address range.
For more information, see [Bring your own IP addresses (BYOIP)](../../../../services/ec2/latest/userguide/ec2-byoip.md) in the _Amazon EC2 User Guide_.

Provisioning an address range is an asynchronous operation, so the call returns immediately,
but the address range is not ready to use until its status changes from `pending-provision`
to `provisioned`. For more information, see [Onboard your address range](../../../../services/ec2/latest/userguide/byoip-onboard.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Cidr**

The public IPv4 or IPv6 address range, in CIDR notation. The most specific IPv4 prefix that you can
specify is /24. The most specific IPv6 address range that you can bring is /48 for CIDRs that are publicly advertisable and /56 for CIDRs that are not publicly advertisable. The address range cannot overlap with another address range that you've
brought to this or another Region.

Type: String

Required: Yes

**CidrAuthorizationContext**

A signed document that proves that you are authorized to bring the specified IP address
range to Amazon using BYOIP.

Type: [CidrAuthorizationContext](api-cidrauthorizationcontext.md) object

Required: No

**Description**

A description for the address range and the address pool.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MultiRegion**

Reserved.

Type: Boolean

Required: No

**NetworkBorderGroup**

If you have [Local Zones](../../../../services/local-zones/latest/ug/how-local-zones-work.md) enabled, you can choose a network border group for Local Zones when you provision and advertise a BYOIPv4 CIDR. Choose the network border group carefully as the EIP and the AWS resource it is associated with must reside in the same network border group.

You can provision BYOIP address ranges to and advertise them in the following Local Zone network border groups:

- us-east-1-dfw-2

- us-west-2-lax-1

- us-west-2-phx-2

###### Note

You cannot provision or advertise BYOIPv6 address ranges in Local Zones at this time.

Type: String

Required: No

**PoolTagSpecification.N**

The tags to apply to the address pool.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**PubliclyAdvertisable**

(IPv6 only) Indicate whether the address range will be publicly advertised to the
internet.

Default: true

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**byoipCidr**

Information about the address range.

Type: [ByoipCidr](api-byoipcidr.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/provisionbyoipcidr.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/provisionbyoipcidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MoveCapacityReservationInstances

ProvisionIpamByoasn
