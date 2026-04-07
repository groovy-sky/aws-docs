# DeleteIpamPool

Delete an IPAM pool.

###### Note

You cannot delete an IPAM pool if there are allocations in it or CIDRs provisioned to it. To release
allocations, see [ReleaseIpamPoolAllocation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ReleaseIpamPoolAllocation.html). To deprovision pool
CIDRs, see [DeprovisionIpamPoolCidr](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeprovisionIpamPoolCidr.html).

For more information, see [Delete a pool](https://docs.aws.amazon.com/vpc/latest/ipam/delete-pool-ipam.html) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Cascade**

Enables you to quickly delete an IPAM pool and all resources within that pool, including
provisioned CIDRs, allocations, and other pools.

###### Important

You can only use this option to delete pools in the private scope or pools in the public scope with a source resource. A source resource is a resource used to provision CIDRs to a resource planning pool.

Type: Boolean

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPoolId**

The ID of the pool to delete.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipamPool**

Information about the results of the deletion.

Type: [IpamPool](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamPool.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteIpamPool)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DeleteIpamPool)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteIpamPolicy

DeleteIpamPrefixListResolver
