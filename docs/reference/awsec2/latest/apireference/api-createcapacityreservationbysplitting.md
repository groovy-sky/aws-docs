# CreateCapacityReservationBySplitting

Create a new Capacity Reservation by splitting the capacity of the source Capacity
Reservation. The new Capacity Reservation will have the same attributes as the source
Capacity Reservation except for tags. The source Capacity Reservation must be
`active` and owned by your AWS account.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensure Idempotency](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceCount**

The number of instances to split from the source Capacity Reservation.

Type: Integer

Required: Yes

**SourceCapacityReservationId**

The ID of the Capacity Reservation from which you want to split the capacity.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the new Capacity Reservation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**destinationCapacityReservation**

Information about the destination Capacity Reservation.

Type: [CapacityReservation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CapacityReservation.html) object

**instanceCount**

The number of instances in the new Capacity Reservation. The number of instances in
the source Capacity Reservation was reduced by this amount.

Type: Integer

**requestId**

The ID of the request.

Type: String

**sourceCapacityReservation**

Information about the source Capacity Reservation.

Type: [CapacityReservation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CapacityReservation.html) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateCapacityReservationBySplitting)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateCapacityReservationBySplitting)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateCapacityReservation

CreateCapacityReservationFleet
