# PurchaseScheduledInstances

###### Note

You can no longer purchase Scheduled Instances.

Purchases the Scheduled Instances with the specified schedule.

Scheduled Instances enable you to purchase Amazon EC2 compute capacity by the hour for a one-year term.
Before you can purchase a Scheduled Instance, you must call [DescribeScheduledInstanceAvailability](api-describescheduledinstanceavailability.md)
to check for available schedules and obtain a purchase token. After you purchase a Scheduled Instance,
you must call [RunScheduledInstances](api-runscheduledinstances.md) during each scheduled time period.

After you purchase a Scheduled Instance, you can't cancel, modify, or resell your purchase.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that ensures the idempotency of the request.
For more information, see [Ensuring Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PurchaseRequest.N**

The purchase requests.

Type: Array of [PurchaseRequest](api-purchaserequest.md) objects

Array Members: Minimum number of 1 item.

Required: Yes

## Response Elements

The following elements are returned by the service.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/purchasescheduledinstances.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/purchasescheduledinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PurchaseReservedInstancesOffering

RebootInstances
