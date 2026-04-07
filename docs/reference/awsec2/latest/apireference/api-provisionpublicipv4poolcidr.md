# ProvisionPublicIpv4PoolCidr

Provision a CIDR to a public IPv4 pool.

For more information about IPAM, see [What is IPAM?](../../../../services/vpc/latest/ipam/what-is-it-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPoolId**

The ID of the IPAM pool you would like to use to allocate this CIDR.

Type: String

Required: Yes

**NetmaskLength**

The netmask length of the CIDR you would like to allocate to the public IPv4 pool. The least specific netmask length you can define is 24.

Type: Integer

Required: Yes

**NetworkBorderGroup**

The Availability Zone (AZ) or Local Zone (LZ) network border group that the resource that the IP address is assigned to is in. Defaults to an AZ network border group. For more information on available Local Zones, see [Local Zone availability](../../../../services/ec2/latest/userguide/ec2-byoip.md#byoip-zone-avail) in the _Amazon EC2 User Guide_.

Type: String

Required: No

**PoolId**

The ID of the public IPv4 pool you would like to use for this CIDR.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**poolAddressRange**

Information about the address range of the public IPv4 pool.

Type: [PublicIpv4PoolRange](api-publicipv4poolrange.md) object

**poolId**

The ID of the pool that you want to provision the CIDR to.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ProvisionPublicIpv4PoolCidr)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ProvisionPublicIpv4PoolCidr)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/provisionpublicipv4poolcidr.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/provisionpublicipv4poolcidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/provisionpublicipv4poolcidr.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/provisionpublicipv4poolcidr.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/provisionpublicipv4poolcidr.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/provisionpublicipv4poolcidr.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ProvisionPublicIpv4PoolCidr)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/provisionpublicipv4poolcidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ProvisionIpamPoolCidr

PurchaseCapacityBlock
