# DeleteAnomalyDetector

Deletes the specified anomaly detection model from your account. For more information
about how to delete an anomaly detection model, see [Deleting an anomaly detection model](../../../../services/amazoncloudwatch/latest/monitoring/create-anomaly-detection-alarm.md#Delete_Anomaly_Detection_Model) in the _CloudWatch User_
_Guide_.

## Request Parameters

**Dimensions**

_This parameter has been deprecated._

The metric dimensions associated with the anomaly detection model to delete.

Type: Array of [Dimension](api-dimension.md) objects

Array Members: Maximum number of 30 items.

Required: No

**MetricMathAnomalyDetector**

The metric math anomaly detector to be deleted.

When using `MetricMathAnomalyDetector`, you cannot include following
parameters in the same operation:

- `Dimensions`,

- `MetricName`

- `Namespace`

- `Stat`

- the `SingleMetricAnomalyDetector` parameters of
`DeleteAnomalyDetectorInput`

Instead, specify the metric math anomaly detector attributes as part of the
`MetricMathAnomalyDetector` property.

Type: [MetricMathAnomalyDetector](api-metricmathanomalydetector.md) object

Required: No

**MetricName**

_This parameter has been deprecated._

The metric name associated with the anomaly detection model to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**Namespace**

_This parameter has been deprecated._

The namespace associated with the anomaly detection model to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[^:].*`

Required: No

**SingleMetricAnomalyDetector**

A single metric anomaly detector to be deleted.

When using `SingleMetricAnomalyDetector`, you cannot include the following
parameters in the same operation:

- `Dimensions`,

- `MetricName`

- `Namespace`

- `Stat`

- the `MetricMathAnomalyDetector` parameters of
`DeleteAnomalyDetectorInput`

Instead, specify the single metric anomaly detector attributes as part of the
`SingleMetricAnomalyDetector` property.

Type: [SingleMetricAnomalyDetector](api-singlemetricanomalydetector.md) object

Required: No

**Stat**

_This parameter has been deprecated._

The statistic associated with the anomaly detection model to delete.

Type: String

Length Constraints: Maximum length of 50.

Pattern: `(SampleCount|Average|Sum|Minimum|Maximum|IQM|(p|tc|tm|ts|wm)(\d{1,2}(\.\d{0,10})?|100)|[ou]\d+(\.\d*)?)(_E|_L|_H)?|(TM|TC|TS|WM)\(((((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%)?:((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%|((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%:(((\d{1,2})(\.\d{0,10})?|100(\.0{0,10})?)%)?)\)|(TM|TC|TS|WM|PR)\(((\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)):((\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)))?|((\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)))?:(\d+(\.\d{0,10})?|(\d+(\.\d{0,10})?[Ee][+-]?\d+)))\)`

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceError**

Request processing has failed due to some unknown error, exception, or
failure.

**Message**

HTTP Status Code: 500

**InvalidParameterCombination**

Parameters were used together that cannot be used together.

**message**

HTTP Status Code: 400

**InvalidParameterValue**

The value of an input parameter is bad or out-of-range.

**message**

HTTP Status Code: 400

**MissingParameter**

An input parameter that is required is missing.

**message**

HTTP Status Code: 400

**ResourceNotFoundException**

The named resource does not exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/deleteanomalydetector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/deleteanomalydetector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAlarms

DeleteDashboards

All content copied from https://docs.aws.amazon.com/.
