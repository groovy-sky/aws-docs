# Purchase

Describes the result of the purchase.

## Contents

**currencyCode**

The currency in which the `UpfrontPrice` and `HourlyPrice`
amounts are specified. At this time, the only supported currency is
`USD`.

Type: String

Valid Values: `USD`

Required: No

**duration**

The duration of the reservation's term in seconds.

Type: Integer

Required: No

**HostIdSet.N**

The IDs of the Dedicated Hosts associated with the reservation.

Type: Array of strings

Required: No

**hostReservationId**

The ID of the reservation.

Type: String

Required: No

**hourlyPrice**

The hourly price of the reservation per hour.

Type: String

Required: No

**instanceFamily**

The instance family on the Dedicated Host that the reservation can be associated
with.

Type: String

Required: No

**paymentOption**

The payment option for the reservation.

Type: String

Valid Values: `AllUpfront | PartialUpfront | NoUpfront`

Required: No

**upfrontPrice**

The upfront price of the reservation.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/purchase.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/purchase.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/purchase.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PublicIpv4PoolRange

PurchaseRequest
