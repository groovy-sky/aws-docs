# DescribeAlarmContributors

Returns the information of the current alarm contributors that are in `ALARM` state. This operation returns details about the individual time series that contribute to the alarm's state.

## Request Parameters

**AlarmName**

The name of the alarm for which to retrieve contributor information.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**NextToken**

The token returned by a previous call to indicate that there is more data available.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**AlarmContributors**

A list of alarm contributors that provide details about the individual time series contributing to the alarm's state.

Type: Array of [AlarmContributor](api-alarmcontributor.md) objects

**NextToken**

The token that marks the start of the next batch of returned results.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidNextToken**

The next token specified is invalid.

**message**

HTTP Status Code: 400

**ResourceNotFoundException**

The named resource does not exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/describealarmcontributors.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/describealarmcontributors.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteMetricStream

DescribeAlarmHistory

All content copied from https://docs.aws.amazon.com/.
