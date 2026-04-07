# GetIpamDiscoveredResourceCidrs

Returns the resource CIDRs that are monitored as part of a resource discovery. A discovered resource is a resource CIDR monitored under a resource discovery. The following resources can be discovered: VPCs, Public IPv4 pools, VPC subnets, and Elastic IP addresses.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

Filters.

Type: Array of [Filter](api-filter.md) objects

Required: No

**IpamResourceDiscoveryId**

A resource discovery ID.

Type: String

Required: Yes

**MaxResults**

The maximum number of discovered resource CIDRs to return in one page of results.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

Specify the pagination token from a previous request to retrieve the next page of results.

Type: String

Required: No

**ResourceRegion**

A resource Region.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipamDiscoveredResourceCidrSet**

Discovered resource CIDRs.

Type: Array of [IpamDiscoveredResourceCidr](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamDiscoveredResourceCidr.html) objects

**nextToken**

Specify the pagination token from a previous request to retrieve the next page of results.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetIpamDiscoveredResourceCidrs)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetIpamDiscoveredPublicAddresses

GetIpamPolicyAllocationRules
