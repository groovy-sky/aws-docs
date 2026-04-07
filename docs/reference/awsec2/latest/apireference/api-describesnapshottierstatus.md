# DescribeSnapshotTierStatus

Describes the storage tier status of one or more Amazon EBS snapshots.

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

- `snapshot-id` \- The snapshot ID.

- `volume-id` \- The ID of the volume the snapshot is for.

- `last-tiering-operation` \- The state of the last archive or restore action. ( `archival-in-progress` \| `archival-completed` \|
`archival-failed` \| `permanent-restore-in-progress` \| `permanent-restore-completed` \| `permanent-restore-failed` \|
`temporary-restore-in-progress` \| `temporary-restore-completed` \| `temporary-restore-failed`)

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token returned from a previous paginated request.
Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items.
This value is `null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**snapshotTierStatusSet**

Information about the snapshot's storage tier.

Type: Array of [SnapshotTierStatus](api-snapshottierstatus.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeSnapshotTierStatus)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeSnapshotTierStatus)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describesnapshottierstatus.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describesnapshottierstatus.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describesnapshottierstatus.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describesnapshottierstatus.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describesnapshottierstatus.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describesnapshottierstatus.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeSnapshotTierStatus)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describesnapshottierstatus.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSnapshots

DescribeSpotDatafeedSubscription
