# DescribeInstanceImageMetadata

Describes the AMI that was used to launch an instance, even if the AMI is deprecated,
deregistered, made private (no longer public or shared with your account), or not
allowed.

If you specify instance IDs, the output includes information for only the specified
instances. If you specify filters, the output includes information for only those instances
that meet the filter criteria. If you do not specify instance IDs or filters, the output
includes information for all instances, which can affect performance.

If you specify an instance ID that is not valid, an instance that doesn't exist, or an
instance that you do not own, an error ( `InvalidInstanceID.NotFound`) is
returned.

Recently terminated instances might appear in the returned results. This interval is
usually less than one hour.

In the rare case where an Availability Zone is experiencing a service disruption and you
specify instance IDs that are in the affected Availability Zone, or do not specify any
instance IDs at all, the call fails. If you specify only instance IDs that are in an
unaffected Availability Zone, the call works normally.

###### Note

The order of the elements in the response, including those within nested structures,
might vary. Applications should not assume the elements appear in a particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `availability-zone` \- The name of the Availability Zone (for example,
`us-west-2a`) or Local Zone (for example, `us-west-2-lax-1b`) of
the instance.

- `instance-id` \- The ID of the instance.

- `image-allowed` \- A Boolean that indicates whether the image meets the
criteria specified for Allowed AMIs.

- `instance-state-name` \- The state of the instance ( `pending` \|
`running` \| `shutting-down` \| `terminated` \|
`stopping` \| `stopped`).

- `instance-type` \- The type of instance (for example,
`t3.micro`).

- `launch-time` \- The time when the instance was launched, in the ISO 8601
format in the UTC time zone (YYYY-MM-DDThh:mm:ss.sssZ), for example,
`2023-09-29T11:04:43.305Z`. You can use a wildcard ( `*`), for
example, `2023-09-29T*`, which matches an entire day.

- `owner-alias` \- The owner alias ( `amazon` \|
`aws-marketplace` \| `aws-backup-vault`). The valid aliases are
defined in an Amazon-maintained list. This is not the AWS account alias that can be set
using the IAM console. We recommend that you use the `Owner` request parameter
instead of this filter.

- `owner-id` \- The AWS account ID of the owner. We recommend that you use
the `Owner` request parameter instead of this filter.

- `tag:<key>` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

- `zone-id` \- The ID of the Availability Zone (for example,
`usw2-az2`) or Local Zone (for example, `usw2-lax1-az1`) of the
instance.

Type: Array of [Filter](api-filter.md) objects

Required: No

**InstanceId.N**

The instance IDs.

If you don't specify an instance ID or filters, the output includes information for all
instances.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Default: 1000

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**instanceImageMetadataSet**

Information about the instance and the AMI used to launch the instance.

Type: Array of [InstanceImageMetadata](api-instanceimagemetadata.md) objects

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeInstanceImageMetadata)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeInstanceImageMetadata)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeinstanceimagemetadata.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeinstanceimagemetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeinstanceimagemetadata.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeinstanceimagemetadata.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeinstanceimagemetadata.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeinstanceimagemetadata.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeInstanceImageMetadata)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeinstanceimagemetadata.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeInstanceEventWindows

DescribeInstances
