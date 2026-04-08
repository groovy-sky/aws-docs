# DescribeFlowLogs

Describes one or more flow logs.

To view the published flow log records, you must view the log destination. For example,
the CloudWatch Logs log group, the Amazon S3 bucket, or the Kinesis Data Firehose delivery stream.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `deliver-log-status` \- The status of the logs delivery ( `SUCCESS` \|
`FAILED`).

- `log-destination-type` \- The type of destination for the flow log
data ( `cloud-watch-logs` \| `s3` \|
`kinesis-data-firehose`).

- `flow-log-id` \- The ID of the flow log.

- `log-group-name` \- The name of the log group.

- `resource-id` \- The ID of the VPC, subnet, or network interface.

- `traffic-type` \- The type of traffic ( `ACCEPT` \|
`REJECT` \| `ALL`).

- `tag`:<key> - The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

- `tag-key` \- The key of a tag assigned to the resource. Use this filter to find all resources assigned a tag with a specific key, regardless of the tag value.

Type: Array of [Filter](api-filter.md) objects

Required: No

**FlowLogId.N**

One or more flow log IDs.

Constraint: Maximum of 1000 flow log IDs.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token to request the next page of items. Pagination continues from the end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**flowLogSet**

Information about the flow logs.

Type: Array of [FlowLog](api-flowlog.md) objects

**nextToken**

The token to request the next page of items. This value is `null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all of your flow logs.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeFlowLogs
&AUTHPARAMS
```

#### Sample Response

```

<DescribeFlowLogsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>3cb46f23-099e-4bf0-891c-EXAMPLE</requestId>
    <flowLogSet>
        <item>
            <logDestination>arn:aws:s3:::my-log-bucket/my-logs/</logDestination>
            <resourceId>vpc-001234183afc7cabc</resourceId>
            <logDestinationType>s3</logDestinationType>
            <creationTime>2020-02-04T11:46:13.831Z</creationTime>
            <trafficType>ALL</trafficType>
            <deliverLogsStatus>SUCCESS</deliverLogsStatus>
            <logFormat>${version} ${instance-id} ${interface-id} ${type} ${pkt-srcaddr} ${pkt-dstaddr} ${protocol} ${bytes} ${start} ${end} ${action}</logFormat>
            <flowLogStatus>ACTIVE</flowLogStatus>
            <flowLogId>fl-1234c5499532dbabc</flowLogId>
            <maxAggregationInterval>60</maxAggregationInterval>
            <tagSet>
                <item>
                    <key>Name</key>
                    <value>FlowsForVpcA</value>
                </item>
            </tagSet>
        </item>
        <item>
            <resourceId>vpc-1122e8183afc74455</resourceId>
            <logDestinationType>cloud-watch-logs</logDestinationType>
            <deliverLogsPermissionArn>arn:aws:iam::123456789101:role/flowlogsrole</deliverLogsPermissionArn>
            <creationTime>2019-07-24T13:11:42.383Z</creationTime>
            <trafficType>ALL</trafficType>
            <deliverLogsStatus>SUCCESS</deliverLogsStatus>
            <logFormat>${version} ${account-id} ${interface-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${packets} ${bytes} ${start} ${end} ${action} ${log-status}</logFormat>
            <flowLogStatus>ACTIVE</flowLogStatus>
            <logGroupName>FlowLogsForSubnetB</logGroupName>
            <flowLogId>fl-0abc1235983d13123</flowLogId>
            <maxAggregationInterval>600</maxAggregationInterval>
            <tagSet>
                <item>
                    <key>Name</key>
                    <value>FlowsForVpcB</value>
                </item>
            </tagSet>
        </item>
    </flowLogSet>
</DescribeFlowLogsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeflowlogs.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeflowlogs.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeFleets

DescribeFpgaImageAttribute
