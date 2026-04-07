# CreateInterruptibleCapacityReservationAllocation

Creates an interruptible Capacity Reservation by specifying the number of unused instances you want to allocate from your source reservation. This helps you make unused capacity available for other workloads within your account while maintaining control to reclaim it.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId**

The ID of the source Capacity Reservation from which to create the interruptible Capacity Reservation. Your Capacity Reservation must be in active state with no end date set and have available capacity for allocation.

Type: String

Required: Yes

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.

Type: Boolean

Required: No

**InstanceCount**

The number of instances to allocate from your source reservation. You can only allocate available instances (also called unused capacity).

Type: Integer

Required: Yes

**TagSpecification.N**

The tags to apply to the interruptible Capacity Reservation during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**interruptionType**

The type of interruption applied to the interruptible reservation.

Type: String

Valid Values: `adhoc`

**requestId**

The ID of the request.

Type: String

**sourceCapacityReservationId**

The ID of the source Capacity Reservation from which the interruptible Capacity Reservation was created.

Type: String

**status**

The current status of the allocation request (creating, active, updating).

Type: String

Valid Values: `pending | active | updating | canceling | canceled | failed`

**targetInstanceCount**

The number of instances allocated to the interruptible reservation.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateInterruptibleCapacityReservationAllocation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateInterruptibleCapacityReservationAllocation)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createinterruptiblecapacityreservationallocation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createinterruptiblecapacityreservationallocation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createinterruptiblecapacityreservationallocation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createinterruptiblecapacityreservationallocation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createinterruptiblecapacityreservationallocation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createinterruptiblecapacityreservationallocation.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateInterruptibleCapacityReservationAllocation)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createinterruptiblecapacityreservationallocation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateInternetGateway

CreateIpam
