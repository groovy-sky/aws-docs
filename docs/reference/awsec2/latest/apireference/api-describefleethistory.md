# DescribeFleetHistory

Describes the events for the specified EC2 Fleet during the specified time.

EC2 Fleet events are delayed by up to 30 seconds before they can be described. This ensures
that you can query by the last evaluated time and not miss a recorded event. EC2 Fleet events
are available for 48 hours.

For more information, see [Monitor fleet events using Amazon EventBridge](../../../../services/ec2/latest/userguide/fleet-monitor.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EventType**

The type of events to describe. By default, all events are described.

Type: String

Valid Values: `instance-change | fleet-change | service-error`

Required: No

**FleetId**

The ID of the EC2 Fleet.

Type: String

Required: Yes

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

**StartTime**

The start date and time for the events, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).

Type: Timestamp

Required: Yes

## Response Elements

The following elements are returned by the service.

**fleetId**

The ID of the EC Fleet.

Type: String

**historyRecordSet**

Information about the events in the history of the EC2 Fleet.

Type: Array of [HistoryRecordEntry](api-historyrecordentry.md) objects

**lastEvaluatedTime**

The last date and time for the events, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).
All records up to this time were retrieved.

If `nextToken` indicates that there are more items, this value is not
present.

Type: Timestamp

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**startTime**

The start date and time for the events, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeFleetHistory)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeFleetHistory)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describefleethistory.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describefleethistory.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describefleethistory.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describefleethistory.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describefleethistory.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describefleethistory.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeFleetHistory)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describefleethistory.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeFastSnapshotRestores

DescribeFleetInstances
