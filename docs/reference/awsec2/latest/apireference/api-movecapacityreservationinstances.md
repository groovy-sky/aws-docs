# MoveCapacityReservationInstances

Move available capacity from a source Capacity Reservation to a destination Capacity
Reservation. The source Capacity Reservation and the destination Capacity Reservation
must be `active`, owned by your AWS account, and share the following:

- Instance type

- Platform

- Availability Zone

- Tenancy

- Placement group

- Capacity Reservation end time - `At specific time` or
`Manually`.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensure Idempotency](run-instance-idempotency.md).

Type: String

Required: No

**DestinationCapacityReservationId**

The ID of the Capacity Reservation that you want to move capacity into.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceCount**

The number of instances that you want to move from the source Capacity Reservation.

Type: Integer

Required: Yes

**SourceCapacityReservationId**

The ID of the Capacity Reservation from which you want to move capacity.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**destinationCapacityReservation**

Information about the destination Capacity Reservation.

Type: [CapacityReservation](api-capacityreservation.md) object

**instanceCount**

The number of instances that were moved from the source Capacity Reservation to the
destination Capacity Reservation.

Type: Integer

**requestId**

The ID of the request.

Type: String

**sourceCapacityReservation**

Information about the source Capacity Reservation.

Type: [CapacityReservation](api-capacityreservation.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/movecapacityreservationinstances.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/movecapacityreservationinstances.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MoveByoipCidrToIpam

ProvisionByoipCidr
