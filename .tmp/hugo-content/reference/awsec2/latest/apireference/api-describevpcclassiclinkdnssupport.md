# DescribeVpcClassicLinkDnsSupport

###### Note

This action is deprecated.

Describes the ClassicLink DNS support status of one or more VPCs. If enabled, the DNS
hostname of a linked EC2-Classic instance resolves to its private IP address when
addressed from an instance in the VPC to which it's linked. Similarly, the DNS hostname
of an instance in a VPC resolves to its private IP address when addressed from a linked
EC2-Classic instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 255.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**VpcIds.N**

The IDs of the VPCs.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there are no more items to return.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**requestId**

The ID of the request.

Type: String

**vpcs**

Information about the ClassicLink DNS support status of the VPCs.

Type: Array of [ClassicLinkDnsSupport](api-classiclinkdnssupport.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevpcclassiclinkdnssupport.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVpcClassicLink

DescribeVpcEncryptionControls
