# CapacityBlockExtension

Describes a Capacity Block extension. With an extension, you can extend the duration
of time for an existing Capacity Block.

## Contents

**availabilityZone**

The Availability Zone of the Capacity Block extension.

Type: String

Required: No

**availabilityZoneId**

The Availability Zone ID of the Capacity Block extension.

Type: String

Required: No

**capacityBlockExtensionDurationHours**

The duration of the Capacity Block extension in hours.

Type: Integer

Required: No

**capacityBlockExtensionEndDate**

The end date of the Capacity Block extension.

Type: Timestamp

Required: No

**capacityBlockExtensionOfferingId**

The ID of the Capacity Block extension offering.

Type: String

Required: No

**capacityBlockExtensionPurchaseDate**

The date when the Capacity Block extension was purchased.

Type: Timestamp

Required: No

**capacityBlockExtensionStartDate**

The start date of the Capacity Block extension.

Type: Timestamp

Required: No

**capacityBlockExtensionStatus**

The status of the Capacity Block extension. A Capacity Block extension can have one of
the following statuses:

- `payment-pending` \- The Capacity Block extension payment is
processing. If your payment can't be processed within 12 hours, the Capacity
Block extension is failed.

- `payment-failed` \- Payment for the Capacity Block extension request
was not successful.

- `payment-succeeded` \- Payment for the Capacity Block extension
request was successful. You receive an invoice that reflects the one-time
upfront payment. In the invoice, you can associate the paid amount with the
Capacity Block reservation ID.

Type: String

Valid Values: `payment-pending | payment-failed | payment-succeeded`

Required: No

**capacityReservationId**

The reservation ID of the Capacity Block extension.

Type: String

Required: No

**currencyCode**

The currency of the payment for the Capacity Block extension.

Type: String

Required: No

**instanceCount**

The number of instances in the Capacity Block extension.

Type: Integer

Required: No

**instanceType**

The instance type of the Capacity Block extension.

Type: String

Required: No

**upfrontFee**

The total price to be paid up front.

Type: String

Required: No

**zoneType**

The type of zone where the Capacity Block extension is located.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacityblockextension.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacityblockextension.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacityblockextension.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityBlock

CapacityBlockExtensionOffering
