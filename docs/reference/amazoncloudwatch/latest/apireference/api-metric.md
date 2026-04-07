# Metric

Represents a specific metric.

## Contents

**Dimensions**

The dimensions for the metric.

Type: Array of [Dimension](api-dimension.md) objects

Array Members: Maximum number of 30 items.

Required: No

**MetricName**

The name of the metric. This is a required field.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**Namespace**

The namespace of the metric.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[^:].*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/monitoring-2010-08-01/Metric)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/monitoring-2010-08-01/Metric)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/monitoring-2010-08-01/Metric)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MessageData

MetricAlarm
