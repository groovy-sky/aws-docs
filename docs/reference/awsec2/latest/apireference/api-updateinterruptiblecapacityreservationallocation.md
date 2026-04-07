# UpdateInterruptibleCapacityReservationAllocation

Modifies the number of instances allocated to an interruptible reservation, allowing you to add more capacity or reclaim capacity to your source Capacity Reservation.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId**

The ID of the source Capacity Reservation containing the interruptible allocation to modify.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.

Type: Boolean

Required: No

**TargetInstanceCount**

The new number of instances to allocate. Enter a higher number to add more capacity to share, or a lower number to reclaim capacity to your source Capacity Reservation.

Type: Integer

Required: Yes

## Response Elements

The following elements are returned by the service.

**instanceCount**

The current number of instances allocated to the interruptible reservation.

Type: Integer

**interruptibleCapacityReservationId**

The ID of the interruptible Capacity Reservation that was modified.

Type: String

**interruptionType**

The interruption type for the interruptible reservation.

Type: String

Valid Values: `adhoc`

**requestId**

The ID of the request.

Type: String

**sourceCapacityReservationId**

The ID of the source Capacity Reservation to which capacity was reclaimed or from which capacity was allocated.

Type: String

**status**

The current status of the allocation (updating during reclamation, active when complete).

Type: String

Valid Values: `pending | active | updating | canceling | canceled | failed`

**targetInstanceCount**

The requested number of instances for the interruptible Capacity Reservation.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/UpdateInterruptibleCapacityReservationAllocation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateCapacityManagerOrganizationsAccess

UpdateSecurityGroupRuleDescriptionsEgress
