# DisassociateCapacityReservationBillingOwner

Cancels a pending request to assign billing of the unused capacity of a Capacity
Reservation to a consumer account, or revokes a request that has already been accepted.
For more information, see [Billing assignment for shared\
Amazon EC2 Capacity Reservations](../../../../services/ec2/latest/userguide/assign-billing.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityReservationId**

The ID of the Capacity Reservation.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**UnusedReservationBillingOwnerId**

The ID of the consumer account to which the request was sent.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9]{12}`

Required: Yes

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociatecapacityreservationbillingowner.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateAddress

DisassociateClientVpnTargetNetwork
