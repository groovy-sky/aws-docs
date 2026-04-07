# ModifyCapacityReservationFleet

Modifies a Capacity Reservation Fleet.

When you modify the total target capacity of a Capacity Reservation Fleet, the Fleet
automatically creates new Capacity Reservations, or modifies or cancels existing
Capacity Reservations in the Fleet to meet the new total target capacity. When you
modify the end date for the Fleet, the end dates for all of the individual Capacity
Reservations in the Fleet are updated accordingly.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationFleetId**

The ID of the Capacity Reservation Fleet to modify.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EndDate**

The date and time at which the Capacity Reservation Fleet expires. When the Capacity
Reservation Fleet expires, its state changes to `expired` and all of the
Capacity Reservations in the Fleet expire.

The Capacity Reservation Fleet expires within an hour after the specified time. For
example, if you specify `5/31/2019`, `13:30:55`, the Capacity
Reservation Fleet is guaranteed to expire between `13:30:55` and
`14:30:55` on `5/31/2019`.

You can't specify **EndDate** and **RemoveEndDate** in the same request.

Type: Timestamp

Required: No

**RemoveEndDate**

Indicates whether to remove the end date from the Capacity Reservation Fleet. If you
remove the end date, the Capacity Reservation Fleet does not expire and it remains
active until you explicitly cancel it using the **CancelCapacityReservationFleet** action.

You can't specify **RemoveEndDate** and **EndDate** in the same request.

Type: Boolean

Required: No

**TotalTargetCapacity**

The total number of capacity units to be reserved by the Capacity Reservation Fleet.
This value, together with the instance type weights that you assign to each instance
type used by the Fleet determine the number of instances for which the Fleet reserves
capacity. Both values are based on units that make sense for your workload. For more
information, see [Total target\
capacity](../../../../services/ec2/latest/userguide/crfleet-concepts.md#target-capacity) in the _Amazon EC2 User Guide_.

Type: Integer

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyCapacityReservationFleet)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyCapacityReservationFleet)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifycapacityreservationfleet.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifycapacityreservationfleet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifycapacityreservationfleet.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifycapacityreservationfleet.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifycapacityreservationfleet.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifycapacityreservationfleet.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyCapacityReservationFleet)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifycapacityreservationfleet.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyCapacityReservation

ModifyClientVpnEndpoint
