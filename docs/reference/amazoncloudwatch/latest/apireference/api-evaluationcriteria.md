# EvaluationCriteria

The evaluation criteria for an alarm. This is a union type that currently
supports `PromQLCriteria`.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**PromQLCriteria**

The PromQL criteria for the alarm evaluation.

Type: [AlarmPromQLCriteria](api-alarmpromqlcriteria.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/evaluationcriteria.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/evaluationcriteria.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/evaluationcriteria.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EntityMetricData

InsightRule
