# StartMetricStreams

Starts the streaming of metrics for one or more of your metric streams.

## Request Parameters

**Names**

The array of the names of metric streams to start streaming.

This is an "all or nothing" operation. If you do not have permission to access all of
the metric streams that you list here, then none of the streams that you list in the
operation will start streaming.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceError**

Request processing has failed due to some unknown error, exception, or
failure.

**Message**

HTTP Status Code: 500

**InvalidParameterValue**

The value of an input parameter is bad or out-of-range.

**message**

HTTP Status Code: 400

**MissingParameter**

An input parameter that is required is missing.

**message**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/startmetricstreams.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/startmetricstreams.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SetAlarmState

StopMetricStreams
