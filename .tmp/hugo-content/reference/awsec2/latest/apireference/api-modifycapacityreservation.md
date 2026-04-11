# ModifyCapacityReservation

Modifies a Capacity Reservation's capacity, instance eligibility, and the conditions
under which it is to be released. You can't modify a Capacity Reservation's instance
type, EBS optimization, platform, instance store settings, Availability Zone, or
tenancy. If you need to modify any of these attributes, we recommend that you cancel the
Capacity Reservation, and then create a new one with the required attributes. For more
information, see [Modify an active\
Capacity Reservation](../../../../services/ec2/latest/userguide/capacity-reservations-modify.md).

The allowed modifications depend on the state of the Capacity Reservation:

- `assessing` or `scheduled` state - You can modify the
tags only.

- `pending` state - You can't modify the Capacity Reservation in any
way.

- `active` state but still within the commitment duration - You can't
decrease the instance count or set an end date that is within the commitment
duration. All other modifications are allowed.

- `active` state with no commitment duration or elapsed commitment
duration - All modifications are allowed.

- `expired`, `cancelled`, `unsupported`, or
`failed` state - You can't modify the Capacity Reservation in any
way.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Accept**

Reserved. Capacity Reservations you have created are accepted by default.

Type: Boolean

Required: No

**AdditionalInfo**

Reserved for future use.

Type: String

Required: No

**CapacityReservationId**

The ID of the Capacity Reservation.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EndDate**

The date and time at which the Capacity Reservation expires. When a Capacity
Reservation expires, the reserved capacity is released and you can no longer launch
instances into it. The Capacity Reservation's state changes to `expired` when
it reaches its end date and time.

The Capacity Reservation is cancelled within an hour from the specified time. For
example, if you specify 5/31/2019, 13:30:55, the Capacity Reservation is guaranteed to
end between 13:30:55 and 14:30:55 on 5/31/2019.

You must provide an `EndDate` value if `EndDateType` is
`limited`. Omit `EndDate` if `EndDateType` is
`unlimited`.

Type: Timestamp

Required: No

**EndDateType**

Indicates the way in which the Capacity Reservation ends. A Capacity Reservation can
have one of the following end types:

- `unlimited` \- The Capacity Reservation remains active until you
explicitly cancel it. Do not provide an `EndDate` value if
`EndDateType` is `unlimited`.

- `limited` \- The Capacity Reservation expires automatically at a
specified date and time. You must provide an `EndDate` value if
`EndDateType` is `limited`.

Type: String

Valid Values: `unlimited | limited`

Required: No

**InstanceCount**

The number of instances for which to reserve capacity. The number of instances can't
be increased or decreased by more than `1000` in a single request.

Type: Integer

Required: No

**InstanceMatchCriteria**

The matching criteria (instance eligibility) that you want to use in the modified
Capacity Reservation. If you change the instance eligibility of an existing Capacity
Reservation from `targeted` to `open`, any running instances that
match the attributes of the Capacity Reservation, have the
`CapacityReservationPreference` set to `open`, and are not yet
running in the Capacity Reservation, will automatically use the modified Capacity
Reservation.

To modify the instance eligibility, the Capacity Reservation must be completely idle
(zero usage).

Type: String

Valid Values: `open | targeted`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifycapacityreservation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifycapacityreservation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyAvailabilityZoneGroup

ModifyCapacityReservationFleet
