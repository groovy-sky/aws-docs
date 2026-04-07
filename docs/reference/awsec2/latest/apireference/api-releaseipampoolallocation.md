# ReleaseIpamPoolAllocation

Release an allocation within an IPAM pool. The Region you use should be the IPAM pool locale. The locale is the AWS Region where this IPAM pool is available for allocations. You can only use this action to release manual allocations. To remove an allocation for a resource without deleting the resource, set its monitored state to false using [ModifyIpamResourceCidr](api-modifyipamresourcecidr.md). For more information, see [Release an allocation](https://docs.aws.amazon.com/vpc/latest/ipam/release-alloc-ipam.html) in the _Amazon VPC IPAM User Guide_.

###### Note

All EC2 API actions follow an [eventual consistency](https://docs.aws.amazon.com/ec2/latest/devguide/eventual-consistency.html) model.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Cidr**

The CIDR of the allocation you want to release.

Type: String

Required: Yes

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPoolAllocationId**

The ID of the allocation.

Type: String

Required: Yes

**IpamPoolId**

The ID of the IPAM pool which contains the allocation you want to release.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**success**

Indicates if the release was successful.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ReleaseIpamPoolAllocation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ReleaseIpamPoolAllocation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReleaseHosts

ReplaceIamInstanceProfileAssociation
