# GetIpamPrefixListResolverVersionEntries

Retrieves the CIDR entries for a specific version of an IPAM prefix list resolver. This shows the actual CIDRs that were selected and synchronized at a particular point in time.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPrefixListResolverId**

The ID of the IPAM prefix list resolver whose version entries you want to retrieve.

Type: String

Required: Yes

**IpamPrefixListResolverVersion**

The version number of the resolver for which to retrieve CIDR entries. If not specified, the latest version is used.

Type: Long

Required: Yes

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

**entrySet**

The CIDR entries for the specified resolver version.

Type: Array of [IpamPrefixListResolverVersionEntry](api-ipamprefixlistresolverversionentry.md) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetIpamPrefixListResolverVersionEntries)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetIpamPrefixListResolverVersionEntries)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getipamprefixlistresolverversionentries.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getipamprefixlistresolverversionentries.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getipamprefixlistresolverversionentries.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getipamprefixlistresolverversionentries.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getipamprefixlistresolverversionentries.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getipamprefixlistresolverversionentries.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetIpamPrefixListResolverVersionEntries)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getipamprefixlistresolverversionentries.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetIpamPrefixListResolverRules

GetIpamPrefixListResolverVersions
