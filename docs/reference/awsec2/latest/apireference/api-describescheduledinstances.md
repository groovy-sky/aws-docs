# DescribeScheduledInstances

Describes the specified Scheduled Instances or all your Scheduled Instances.

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

- `availability-zone` \- The Availability Zone (for example, `us-west-2a`).

- `instance-type` \- The instance type (for example, `c4.large`).

- `platform` \- The platform ( `Linux/UNIX` or `Windows`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return in a single call.
This value can be between 5 and 300. The default value is 100.
To retrieve the remaining results, make another call with the returned
`NextToken` value.

Type: Integer

Required: No

**NextToken**

The token for the next set of results.

Type: String

Required: No

**ScheduledInstanceId.N**

The Scheduled Instance IDs.

Type: Array of strings

Required: No

**SlotStartTimeRange**

The time period for the first schedule to start.

Type: [SlotStartTimeRangeRequest](api-slotstarttimerangerequest.md) object

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token required to retrieve the next set of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**scheduledInstanceSet**

Information about the Scheduled Instances.

Type: Array of [ScheduledInstance](api-scheduledinstance.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeScheduledInstances)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeScheduledInstances)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describescheduledinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describescheduledinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describescheduledinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describescheduledinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describescheduledinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describescheduledinstances.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeScheduledInstances)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describescheduledinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeScheduledInstanceAvailability

DescribeSecondaryInterfaces
