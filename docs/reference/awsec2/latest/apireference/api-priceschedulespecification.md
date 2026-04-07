# PriceScheduleSpecification

Describes the price for a Reserved Instance.

## Contents

**CurrencyCode**

The currency for transacting the Reserved Instance resale. At this time, the only
supported currency is `USD`.

Type: String

Valid Values: `USD`

Required: No

**Price**

The fixed price for the term.

Type: Double

Required: No

**Term**

The number of months remaining in the reservation. For example, 2 is the second to the
last month before the capacity reservation expires.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/PriceScheduleSpecification)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/PriceScheduleSpecification)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/PriceScheduleSpecification)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PriceSchedule

PricingDetail
