# MetricMathAnomalyDetector

Indicates the CloudWatch math expression that provides the time series the anomaly
detector uses as input. The designated math expression must return a single time
series.

## Contents

**MetricDataQueries**

An array of metric data query structures that enables you to create an anomaly
detector based on the result of a metric math expression. Each item in
`MetricDataQueries` gets a metric or performs a math expression. One item
in `MetricDataQueries` is the expression that provides the time series that
the anomaly detector uses as input. Designate the expression by setting
`ReturnData` to `true` for this object in the array. For all
other expressions and metrics, set `ReturnData` to `false`. The
designated expression must return a single time series.

Type: Array of [MetricDataQuery](api-metricdataquery.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/metricmathanomalydetector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/metricmathanomalydetector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/metricmathanomalydetector.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MetricDatum

MetricStat
