# AllocateAddress

Allocates an Elastic IP address to your AWS account. After you allocate the Elastic IP address you can associate
it with an instance or network interface. After you release an Elastic IP address, it is released to the IP address
pool and can be allocated to a different AWS account.

You can allocate an Elastic IP address from one of the following address pools:

- Amazon's pool of IPv4 addresses

- Public IPv4 address range that you own and bring to your AWS account using
[Bring Your Own IP Addresses (BYOIP)](../../../../services/ec2/latest/userguide/ec2-byoip.md)

- An IPv4 IPAM pool with an Amazon-provided or BYOIP public IPv4 address range

- IPv4 addresses from your on-premises network made available for use with an Outpost
using a [customer-owned IP address\
pool](../../../../services/outposts/latest/userguide/routing.md#ip-addressing) (CoIP pool)

For more information, see [Elastic IP Addresses](../../../../services/ec2/latest/userguide/elastic-ip-addresses-eip.md) in the _Amazon EC2 User Guide_.

If you release an Elastic IP address, you might be able to recover it. You cannot recover
an Elastic IP address that you released after it is allocated to another AWS account. To attempt to recover an Elastic IP address that you released, specify
it in this operation.

You can allocate a carrier IP address which is a public IP address from a telecommunication carrier,
to a network interface which resides in a subnet in a Wavelength Zone (for example an EC2 instance).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Address**

The Elastic IP address to recover or an IPv4 address from an address pool.

Type: String

Required: No

**CustomerOwnedIpv4Pool**

The ID of a customer-owned address pool. Use this parameter to let Amazon EC2
select an address from the address pool. Alternatively, specify a specific
address from the address pool.

Type: String

Required: No

**Domain**

The network ( `vpc`).

Type: String

Valid Values: `vpc | standard`

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPoolId**

The ID of an IPAM pool which has an Amazon-provided or BYOIP public IPv4 CIDR provisioned to it. For more information, see [Allocate sequential Elastic IP addresses from an IPAM pool](../../../../services/vpc/latest/ipam/tutorials-eip-pool.md) in the _Amazon VPC IPAM User Guide_.

Type: String

Required: No

**NetworkBorderGroup**

A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS
advertises IP addresses. Use this parameter to limit the IP address to this location. IP
addresses cannot move between network border groups.

Use [DescribeAvailabilityZones](api-describeavailabilityzones.md) to view the network border groups.

Type: String

Required: No

**PublicIpv4Pool**

The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.
To specify a specific address from the address pool, use the `Address` parameter instead.

Type: String

Required: No

**TagSpecification.N**

The tags to assign to the Elastic IP address.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**allocationId**

The ID that represents the allocation of the Elastic IP address.

Type: String

**carrierIp**

The carrier IP address. Available only for network interfaces that reside
in a subnet in a Wavelength Zone.

Type: String

**customerOwnedIp**

The customer-owned IP address.

Type: String

**customerOwnedIpv4Pool**

The ID of the customer-owned address pool.

Type: String

**domain**

The network ( `vpc`).

Type: String

Valid Values: `vpc | standard`

**networkBorderGroup**

The set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises
IP addresses.

Type: String

**publicIp**

The Amazon-owned IP address. Not available when using an address pool that you own.

Type: String

**publicIpv4Pool**

The ID of an address pool that you own.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example for Allocation

This example request allocates an Elastic IP address.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AllocateAddress
&Domain=vpc
&AUTHPARAMS
```

#### Sample Response

```

<AllocateAddressResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <publicIp>198.51.100.1</publicIp>
   <domain>vpc</domain>
   <allocationId>eipalloc-5723d13e</allocationId>
</AllocateAddressResponse>
```

### Example for Recovery

This example request shows how to recover an Elastic IP address that you previously released.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AllocateAddress
&Domain=vpc
&Address=203.0.113.3
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/allocateaddress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/allocateaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AdvertiseByoipCidr

AllocateHosts
