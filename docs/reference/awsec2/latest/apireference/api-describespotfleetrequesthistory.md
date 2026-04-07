# DescribeSpotFleetRequestHistory

Describes the events for the specified Spot Fleet request during the specified
time.

Spot Fleet events are delayed by up to 30 seconds before they can be described. This
ensures that you can query by the last evaluated time and not miss a recorded event.
Spot Fleet events are available for 48 hours.

For more information, see [Monitor fleet events using Amazon\
EventBridge](../../../../services/ec2/latest/userguide/fleet-monitor.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**EventType**

The type of events to describe. By default, all events are described.

Type: String

Valid Values: `instanceChange | fleetRequestChange | error | information`

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

Required: No

**SpotFleetRequestId**

The ID of the Spot Fleet request.

Type: String

Required: Yes

**StartTime**

The starting date and time for the events, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).

Type: Timestamp

Required: Yes

## Response Elements

The following elements are returned by the service.

**historyRecordSet**

Information about the events in the history of the Spot Fleet request.

Type: Array of [HistoryRecord](api-historyrecord.md) objects

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

**spotFleetRequestId**

The ID of the Spot Fleet request.

Type: String

**startTime**

The starting date and time for the events, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the events for Spot Fleet request
sfr-123f8fc2-cb31-425e-abcd-example2710 from the specified start time.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeSpotFleetRequestHistory
&SpotFleetRequestId=sfr-123f8fc2-cb31-425e-abcd-example2710
&StartTime=2015-07-01T12:00:00Z
&AUTHPARAMS
```

#### Sample Response

```

<DescribeSpotFleetRequestHistoryResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>30be3aaf-afd2-408c-b62b-example</requestId>
    <lastEvaluatedTime>2015-07-01T13:29:40+0000</lastEvaluatedTime>
    <spotFleetRequestId>sfr-123f8fc2-cb31-425e-abcd-example2710</spotFleetRequestId>
    <startTime>2015-07-01T12:00:00Z</startTime>
    <historyRecordSet>
        <item>
            <eventInformation>
                <eventSubType>submitted</eventSubType>
            </eventInformation>
            <eventType>fleetRequestChange</eventType>
            <timestamp>2015-07-01T13:10:10.219Z</timestamp>
        </item>
        <item>
            <eventInformation>
                <eventSubType>active</eventSubType>
            </eventInformation>
            <eventType>fleetRequestChange</eventType>
            <timestamp>2015-07-01T13:10:11.624Z</timestamp>
        </item>
        <item>
            <eventInformation>
                <eventDescription>m3.medium, ami-1ecae776, Linux/UNIX (Amazon VPC); old price: 0.0153, new price: 0.0153</eventDescription>
                <eventSubType>price_update</eventSubType>
            </eventInformation>
            <eventType>fleetRequestChange</eventType>
            <timestamp>2015-07-01T13:10:13.365Z</timestamp>
        </item>
        <item>
            <eventInformation>
                <instanceId>i-1234567890abcdef0</instanceId>
                <eventSubType>launched</eventSubType>
            </eventInformation>
            <eventType>instanceChange</eventType>
            <timestamp>2015-07-01T13:19:53.795Z</timestamp>
        </item>
        <item>
            <eventInformation>
                <instanceId>i-1234567890abcdef1</instanceId>
                <eventSubType>launched</eventSubType>
            </eventInformation>
            <eventType>instanceChange</eventType>
            <timestamp>2015-07-01T13:20:39.777Z</timestamp>
        </item>
        <item>
            <eventInformation>
                <instanceId>i-1234567890abcdef2</instanceId>
                <eventSubType>launched</eventSubType>
            </eventInformation>
            <eventType>instanceChange</eventType>
            <timestamp>2015-07-01T13:20:57.773Z</timestamp>
        </item>
        <item>
            <eventInformation>
                <instanceId>i-1234567890abcdef3</instanceId>
                <eventSubType>launched</eventSubType>
            </eventInformation>
            <eventType>instanceChange</eventType>
            <timestamp>2015-07-01T13:22:05.696Z</timestamp>
        </item>
        <item>
            <eventInformation>
                <instanceId>i-1234567890abcdef4</instanceId>
                <eventSubType>launched</eventSubType>
            </eventInformation>
            <eventType>instanceChange</eventType>
            <timestamp>2015-07-01T13:23:58.927Z</timestamp>
        </item>
    </historyRecordSet>
</DescribeSpotFleetRequestHistoryResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeSpotFleetRequestHistory)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeSpotFleetRequestHistory)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describespotfleetrequesthistory.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describespotfleetrequesthistory.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describespotfleetrequesthistory.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describespotfleetrequesthistory.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describespotfleetrequesthistory.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describespotfleetrequesthistory.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeSpotFleetRequestHistory)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describespotfleetrequesthistory.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSpotFleetInstances

DescribeSpotFleetRequests
