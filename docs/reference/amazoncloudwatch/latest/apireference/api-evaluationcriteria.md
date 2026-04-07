# EvaluationCriteria

The evaluation criteria for an alarm. This is a union type that currently
supports `PromQLCriteria`.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**PromQLCriteria**

The PromQL criteria for the alarm evaluation.

Type: [AlarmPromQLCriteria](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_AlarmPromQLCriteria.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/monitoring-2010-08-01/EvaluationCriteria)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/monitoring-2010-08-01/EvaluationCriteria)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/monitoring-2010-08-01/EvaluationCriteria)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EntityMetricData

InsightRule
