# DescribeVpcClassicLink

###### Note

This action is deprecated.

Describes the ClassicLink status of the specified VPCs.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `is-classic-link-enabled` \- Whether the VPC is enabled for ClassicLink
( `true` \| `false`).

- `tag` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

Type: Array of [Filter](api-filter.md) objects

Required: No

**VpcId.N**

The VPCs for which you want to describe the ClassicLink status.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpcSet**

The ClassicLink status of the VPCs.

Type: Array of [VpcClassicLink](api-vpcclassiclink.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describevpcclassiclink.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevpcclassiclink.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVpcBlockPublicAccessOptions

DescribeVpcClassicLinkDnsSupport
