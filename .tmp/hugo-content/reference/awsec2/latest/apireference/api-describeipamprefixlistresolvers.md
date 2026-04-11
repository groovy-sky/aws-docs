# DescribeIpamPrefixListResolvers

Describes one or more IPAM prefix list resolvers. Use this operation to view the configuration, status, and properties of your resolvers.

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

**IpamPrefixListResolverId.N**

The IDs of the IPAM prefix list resolvers to describe. If not specified, all resolvers in your account are described.

Type: Array of strings

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

**ipamPrefixListResolverSet**

Information about the IPAM prefix list resolvers.

Type: Array of [IpamPrefixListResolver](api-ipamprefixlistresolver.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeipamprefixlistresolvers.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeipamprefixlistresolvers.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeIpamPools

DescribeIpamPrefixListResolverTargets
