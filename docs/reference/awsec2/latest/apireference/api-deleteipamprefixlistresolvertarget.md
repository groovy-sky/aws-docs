# DeleteIpamPrefixListResolverTarget

Deletes an IPAM prefix list resolver target. This removes the association between the resolver and the managed prefix list, stopping automatic CIDR synchronization.

For more information about IPAM prefix list resolver, see [Automate prefix list updates with IPAM](https://docs.aws.amazon.com/vpc/latest/ipam/automate-prefix-list-updates.html) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPrefixListResolverTargetId**

The ID of the IPAM prefix list resolver target to delete.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipamPrefixListResolverTarget**

Information about the IPAM prefix list resolver target that was deleted.

Type: [IpamPrefixListResolverTarget](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamPrefixListResolverTarget.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DeleteIpamPrefixListResolverTarget)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteIpamPrefixListResolver

DeleteIpamResourceDiscovery
