# PriceSchedule

Describes the price for a Reserved Instance.

## Contents

**active**

The current price schedule, as determined by the term remaining for the Reserved Instance
in the listing.

A specific price schedule is always in effect, but only one price schedule can be active
at any time. Take, for example, a Reserved Instance listing that has five months remaining in
its term. When you specify price schedules for five months and two months, this means that
schedule 1, covering the first three months of the remaining term, will be active during
months 5, 4, and 3. Then schedule 2, covering the last two months of the term, will be active
for months 2 and 1.

Type: Boolean

Required: No

**currencyCode**

The currency for transacting the Reserved Instance resale. At this time, the only
supported currency is `USD`.

Type: String

Valid Values: `USD`

Required: No

**price**

The fixed price for the term.

Type: Double

Required: No

**term**

The number of months remaining in the reservation. For example, 2 is the second to the
last month before the capacity reservation expires.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/priceschedule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/priceschedule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/priceschedule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PrefixListId

PriceScheduleSpecification
