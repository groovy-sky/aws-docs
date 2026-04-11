# MetricPoint

Indicates whether the network was healthy or degraded at a particular point. The value is aggregated from the `startDate` to the `endDate`. Currently only `five_minutes` is supported.

## Contents

**endDate**

The end date for the metric point. The ending time must be formatted as `yyyy-mm-ddThh:mm:ss`. For example, `2022-06-12T12:00:00.000Z`.

Type: Timestamp

Required: No

**startDate**

The start date for the metric point. The starting date for the metric point. The starting time must be formatted
as `yyyy-mm-ddThh:mm:ss`. For example, `2022-06-10T12:00:00.000Z`.

Type: Timestamp

Required: No

**status**

The status of the metric point.

Type: String

Required: No

**value**

Type: Float

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/metricpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/metricpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/metricpoint.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MetricDataResult

MetricValue
