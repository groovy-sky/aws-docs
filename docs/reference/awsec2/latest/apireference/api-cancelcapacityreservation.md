# CancelCapacityReservation

Cancels the specified Capacity Reservation, releases the reserved capacity, and
changes the Capacity Reservation's state to `cancelled`.

You can cancel a Capacity Reservation that is in the following states:

- `assessing`

- `active` and there is no commitment duration or the commitment
duration has elapsed. You can't cancel a future-dated Capacity Reservation
during the commitment duration.

###### Note

You can't modify or cancel a Capacity Block. For more information, see [Capacity Blocks for ML](../../../../services/ec2/latest/userguide/ec2-capacity-blocks.md).

If a future-dated Capacity Reservation enters the `delayed` state, the
commitment duration is waived, and you can cancel it as soon as it enters the
`active` state.

Instances running in the reserved capacity continue running until you stop them.
Stopped instances that target the Capacity Reservation can no longer launch. Modify
these instances to either target a different Capacity Reservation, launch On-Demand
Instance capacity, or run in any open Capacity Reservation that has matching attributes
and sufficient capacity.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId**

The ID of the Capacity Reservation to be cancelled.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CancelCapacityReservation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CancelCapacityReservation)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cancelcapacityreservation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/cancelcapacityreservation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cancelcapacityreservation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/cancelcapacityreservation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/cancelcapacityreservation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/cancelcapacityreservation.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CancelCapacityReservation)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cancelcapacityreservation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelBundleTask

CancelCapacityReservationFleets
