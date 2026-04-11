# CapacityReservationBillingRequest

Information about a request to assign billing of the unused capacity of a Capacity
Reservation.

## Contents

**capacityReservationId**

The ID of the Capacity Reservation.

Type: String

Required: No

**capacityReservationInfo**

Information about the Capacity Reservation.

Type: [CapacityReservationInfo](api-capacityreservationinfo.md) object

Required: No

**lastUpdateTime**

The date and time, in UTC time format, at which the request was initiated.

Type: Timestamp

Required: No

**requestedBy**

The ID of the AWS account that initiated the request.

Type: String

Required: No

**status**

The status of the request. For more information, see [View billing assignment\
requests for a shared Amazon EC2 Capacity Reservation](../../../../services/ec2/latest/userguide/view-billing-transfers.md).

Type: String

Valid Values: `pending | accepted | rejected | cancelled | revoked | expired`

Required: No

**statusMessage**

Information about the status.

Type: String

Required: No

**unusedReservationBillingOwnerId**

The ID of the AWS account to which the request was sent.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9]{12}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityreservationbillingrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityreservationbillingrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityreservationbillingrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityReservation

CapacityReservationCommitmentInfo
