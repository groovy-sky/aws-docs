# DescribeScheduledInstanceAvailability

Finds available schedules that meet the specified criteria.

You can search for an available schedule no more than 3 months in advance. You must meet the minimum required duration of 1,200 hours per year. For example, the minimum daily schedule is 4 hours, the minimum weekly schedule is 24 hours, and the minimum monthly schedule is 100 hours.

After you find a schedule that meets your needs, call [PurchaseScheduledInstances](api-purchasescheduledinstances.md)
to purchase Scheduled Instances with that schedule.

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

**FirstSlotStartTimeRange**

The time period for the first schedule to start.

Type: [SlotDateTimeRangeRequest](api-slotdatetimerangerequest.md) object

Required: Yes

**MaxResults**

The maximum number of results to return in a single call.
This value can be between 5 and 300. The default value is 300.
To retrieve the remaining results, make another call with the returned
`NextToken` value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 300.

Required: No

**MaxSlotDurationInHours**

The maximum available duration, in hours. This value must be greater than `MinSlotDurationInHours`
and less than 1,720.

Type: Integer

Required: No

**MinSlotDurationInHours**

The minimum available duration, in hours. The minimum required duration is 1,200 hours per year. For example, the minimum daily schedule is 4 hours, the minimum weekly schedule is 24 hours, and the minimum monthly schedule is 100 hours.

Type: Integer

Required: No

**NextToken**

The token for the next set of results.

Type: String

Required: No

**Recurrence**

The schedule recurrence.

Type: [ScheduledInstanceRecurrenceRequest](api-scheduledinstancerecurrencerequest.md) object

Required: Yes

## Response Elements

The following elements are returned by the service.

**nextToken**

The token required to retrieve the next set of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**scheduledInstanceAvailabilitySet**

Information about the available Scheduled Instances.

Type: Array of [ScheduledInstanceAvailability](api-scheduledinstanceavailability.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describescheduledinstanceavailability.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describescheduledinstanceavailability.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeRouteTables

DescribeScheduledInstances
