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

One or more filters for the request. For more information about filtering, see [Filtering CLI output](../../../../services/cli/latest/userguide/cli-usage-filter.md).

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

Type: [RequestIpamResourceTag](api-requestipamresourcetag.md) object

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

Type: Array of [IpamResourceCidr](api-ipamresourcecidr.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getipamresourcecidrs.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getipamresourcecidrs.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetIpamPrefixListResolverVersions

GetLaunchTemplateData
