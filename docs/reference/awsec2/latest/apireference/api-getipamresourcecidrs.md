# GetIpamResourceCidrs

Returns resource CIDRs managed by IPAM in a given scope. If an IPAM is associated with more than one resource discovery, the resource CIDRs across all of the resource discoveries is returned. A resource discovery is an IPAM component that enables IPAM to manage and monitor resources that belong to the owning account.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters for the request. For more information about filtering, see [Filtering CLI output](https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html).

Type: Array of [Filter](api-filter.md) objects

Required: No

**IpamPoolId**

The ID of the IPAM pool that the resource is in.

Type: String

Required: No

**IpamScopeId**

The ID of the scope that the resource is in.

Type: String

Required: Yes

**MaxResults**

The maximum number of results to return in the request.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**ResourceId**

The ID of the resource.

Type: String

Required: No

**ResourceOwner**

The ID of the AWS account that owns the resource.

Type: String

Required: No

**ResourceTag**

The resource tag.

Type: [RequestIpamResourceTag](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestIpamResourceTag.html) object

Required: No

**ResourceType**

The resource type.

Type: String

Valid Values: `vpc | subnet | eip | public-ipv4-pool | ipv6-pool | eni | anycast-ip-list`

Required: No

## Response Elements

The following elements are returned by the service.

**ipamResourceCidrSet**

The resource CIDRs.

Type: Array of [IpamResourceCidr](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamResourceCidr.html) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetIpamResourceCidrs)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetIpamResourceCidrs)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetIpamPrefixListResolverVersions

GetLaunchTemplateData
