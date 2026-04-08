# MetricCharacteristics

This object includes parameters that you can use to provide information to CloudWatch to help it build more accurate anomaly detection models.

## Contents

**PeriodicSpikes**

Set this parameter to `true` if values for this metric consistently include
spikes that should not be considered to be anomalies. With this set to
`true`, CloudWatch will expect to see spikes that occurred
consistently during the model training period, and won't flag future similar spikes as
anomalies.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/metriccharacteristics.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/metriccharacteristics.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/metriccharacteristics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricAlarm

MetricDataQuery

All content copied from https://docs.aws.amazon.com/.
