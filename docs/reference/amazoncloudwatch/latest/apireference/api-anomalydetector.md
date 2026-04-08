# AnomalyDetector

An anomaly detection model associated with a particular CloudWatch metric, statistic,
or metric math expression. You can use the model to display a band of expected, normal
values when the metric is graphed.

If you have enabled unified cross-account observability, and this account is a
monitoring account, the metric can be in the same account or a source account.

## Contents

**Configuration**

The configuration specifies details about how the anomaly detection model is to be
trained, including time ranges to exclude from use for training the model, and the time
zone to use for the metric.

Type: [AnomalyDetectorConfiguration](api-anomalydetectorconfiguration.md) object

Required: No

**Dimensions**

_This member has been deprecated._

The metric dimensions associated with the anomaly detection model.

Type: Array of [Dimension](api-dimension.md) objects

Array Members: Maximum number of 30 items.

Required: No

**MetricCharacteristics**

This object includes parameters that you can use to provide information about your
metric to CloudWatch to help it build more accurate anomaly detection models.
Currently, it includes the `PeriodicSpikes` parameter.

Type: [MetricCharacteristics](api-metriccharacteristics.md) object

Required: No

**MetricMathAnomalyDetector**

The CloudWatch metric math expression for this anomaly detector.

Type: [MetricMathAnomalyDetector](api-metricmathanomalydetector.md) object

Required: No

**MetricName**

_This member has been deprecated._

The name of the metric associated with the anomaly detection model.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**Namespace**

_This member has been deprecated._

The namespace of the metric associated with the anomaly detection model.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[^:].*`

Required: No

**SingleMetricAnomalyDetector**

The CloudWatch metric and statistic for this anomaly detector.

Type: [SingleMetricAnomalyDetector](api-singlemetricanomalydetector.md) object

Required: No

**Stat**

_This member has been deprecated._

The statistic associated with the anomaly detection model.

Type: String

Length Constraints: Maximum length of 50.

Pattern: `(SampleCount|Average|Sum|Minimum|Maximum|IQM|(p|tc|tm|ts|wm)(\d{1,2}(\.\d{0,10})?|100)|[ou]\d+(\.\d*)?)(_E|_L|_H)?|(TM|TC|TS|WM)\(((((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%)?:((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%|((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%:(((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%)?)\)|(TM|TC|TS|WM|PR)\(((\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)):((\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)))?|((\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)))?:(\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)))\)`

Required: No

**StateValue**

The current status of the anomaly detector's training.

Type: String

Valid Values: `PENDING_TRAINING | TRAINED_INSUFFICIENT_DATA | TRAINED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/anomalydetector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/anomalydetector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/anomalydetector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlarmPromQLCriteria

AnomalyDetectorConfiguration

All content copied from https://docs.aws.amazon.com/.
