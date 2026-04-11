# CreatePublicIpv4Pool

Creates a public IPv4 address pool. A public IPv4 pool is an EC2 IP address pool required for the public IPv4 CIDRs that you own and bring to AWS to manage with IPAM. IPv6 addresses you bring to AWS, however, use IPAM pools only. To monitor the status of pool creation, use [DescribePublicIpv4Pools](api-describepublicipv4pools.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**NetworkBorderGroup**

The Availability Zone (AZ) or Local Zone (LZ) network border group that the resource that the IP address is assigned to is in. Defaults to an AZ network border group. For more information on available Local Zones, see [Local Zone availability](../../../../services/ec2/latest/userguide/ec2-byoip.md#byoip-zone-avail) in the _Amazon EC2 User Guide_.

Type: String

Required: No

**TagSpecification.N**

The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**poolId**

The ID of the public IPv4 pool.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createpublicipv4pool.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createpublicipv4pool.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreatePlacementGroup

CreateReplaceRootVolumeTask
