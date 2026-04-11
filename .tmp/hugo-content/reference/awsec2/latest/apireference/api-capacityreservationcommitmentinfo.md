# CapacityReservationCommitmentInfo

Information about your commitment for a future-dated Capacity Reservation.

## Contents

**commitmentEndDate**

The date and time at which the commitment duration expires, in the ISO8601 format in
the UTC time zone ( `YYYY-MM-DDThh:mm:ss.sssZ`). You can't decrease the
instance count or cancel the Capacity Reservation before this date and time.

Type: Timestamp

Required: No

**committedInstanceCount**

The instance capacity that you committed to when you requested the future-dated
Capacity Reservation.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityreservationcommitmentinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityreservationcommitmentinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityreservationcommitmentinfo.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityReservationBillingRequest

CapacityReservationFleet
