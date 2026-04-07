# GetIpamPrefixListResolverVersions

Retrieves version information for an IPAM prefix list resolver.

Each version is a snapshot of what CIDRs matched your rules at that moment in time. The version number increments every time the CIDR list
changes due to infrastructure changes.

**Version example:**

**Initial State (Version 1)**

Production environment:

- vpc-prod-web (10.1.0.0/16) - tagged env=prod

- vpc-prod-db (10.2.0.0/16) - tagged env=prod

Resolver rule: Include all VPCs tagged env=prod

**Version 1 CIDRs:** 10.1.0.0/16, 10.2.0.0/16

**Infrastructure Change (Version 2)**

New VPC added:

- vpc-prod-api (10.3.0.0/16) - tagged env=prod

IPAM automatically detects the change and creates a new version.

**Version 2 CIDRs:** 10.1.0.0/16, 10.2.0.0/16, 10.3.0.0/16

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters to limit the results.

Type: Array of [Filter](api-filter.md) objects

Required: No

**IpamPrefixListResolverId**

The ID of the IPAM prefix list resolver whose versions you want to retrieve.

Type: String

Required: Yes

**IpamPrefixListResolverVersion.N**

Specific version numbers to retrieve. If not specified, all versions are returned.

Type: Array of longs

Required: No

**MaxResults**

The maximum number of items to return for this request. To get the next page of items, make another request with the token returned in the output. For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**ipamPrefixListResolverVersionSet**

Information about the IPAM prefix list resolver versions.

Type: Array of [IpamPrefixListResolverVersion](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamPrefixListResolverVersion.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetIpamPrefixListResolverVersions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetIpamPrefixListResolverVersionEntries

GetIpamResourceCidrs
