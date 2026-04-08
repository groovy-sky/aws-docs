# ListMetricStreams

Returns a list of metric streams in this account.

## Request Parameters

**MaxResults**

The maximum number of results to return in one operation.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 500.

Required: No

**NextToken**

Include this value, if it was returned by the previous call, to get the next set of
metric streams.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**Entries**

The array of metric stream information.

Type: Array of [MetricStreamEntry](api-metricstreamentry.md) objects

**NextToken**

The token that marks the start of the next batch of returned results. You can use
this token in a subsequent operation to get the next batch of results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceError**

Request processing has failed due to some unknown error, exception, or
failure.

**Message**

HTTP Status Code: 500

**InvalidNextToken**

The next token specified is invalid.

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/listmetricstreams.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/listmetricstreams.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListMetrics

ListTagsForResource
