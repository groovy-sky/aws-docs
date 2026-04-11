# AllocateIpamPoolCidr

Allocate a CIDR from an IPAM pool. The Region you use should be the IPAM pool locale. The locale is the AWS Region where this IPAM pool is available for allocations.

In IPAM, an allocation is a CIDR assignment from an IPAM pool to another IPAM pool or to a resource. For more information, see [Allocate CIDRs](../../../../services/vpc/latest/ipam/allocate-cidrs-ipam.md) in the _Amazon VPC IPAM User Guide_.

###### Note

This action creates an allocation with strong consistency. The returned CIDR will not overlap with any other allocations from the same pool.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllowedCidr.N**

Include a particular CIDR range that can be returned by the pool. Allowed CIDRs are only allowed if using netmask length for allocation.

Type: Array of strings

Required: No

**Cidr**

The CIDR you would like to allocate from the IPAM pool. Note the following:

- If there is no DefaultNetmaskLength allocation rule set on the pool, you must specify either the NetmaskLength or the CIDR.

- If the DefaultNetmaskLength allocation rule is set on the pool, you can specify either the NetmaskLength or the CIDR and the DefaultNetmaskLength allocation rule will be ignored.

Possible values: Any available IPv4 or IPv6 CIDR.

Type: String

Required: No

**ClientToken**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

A description for the allocation.

Type: String

Required: No

**DisallowedCidr.N**

Exclude a particular CIDR range from being returned by the pool. Disallowed CIDRs are only allowed if using netmask length for allocation.

Type: Array of strings

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPoolId**

The ID of the IPAM pool from which you would like to allocate a CIDR.

Type: String

Required: Yes

**NetmaskLength**

The netmask length of the CIDR you would like to allocate from the IPAM pool. Note the following:

- If there is no DefaultNetmaskLength allocation rule set on the pool, you must specify either the NetmaskLength or the CIDR.

- If the DefaultNetmaskLength allocation rule is set on the pool, you can specify either the NetmaskLength or the CIDR and the DefaultNetmaskLength allocation rule will be ignored.

Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.

Type: Integer

Required: No

**PreviewNextCidr**

A preview of the next available CIDR in a pool.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**ipamPoolAllocation**

Information about the allocation created.

Type: [IpamPoolAllocation](api-ipampoolallocation.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/allocateipampoolcidr.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/allocateipampoolcidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AllocateHosts

ApplySecurityGroupsToClientVpnTargetNetwork
