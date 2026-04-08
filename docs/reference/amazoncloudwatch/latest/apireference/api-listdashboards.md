# ListDashboards

Returns a list of the dashboards for your account. If you include
`DashboardNamePrefix`, only those dashboards with names starting with the
prefix are listed. Otherwise, all dashboards in your account are listed.

`ListDashboards` returns up to 1000 results on one page. If there are
more than 1000 dashboards, you can call `ListDashboards` again and include
the value you received for `NextToken` in the first call, to receive the next
1000 results.

## Request Parameters

**DashboardNamePrefix**

If you specify this parameter, only the dashboards with names starting with the
specified string are listed. The maximum length is 255, and valid characters are A-Z,
a-z, 0-9, ".", "-", and "\_".

Type: String

Required: No

**NextToken**

The token returned by a previous call to indicate that there is more data
available.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**DashboardEntries**

The list of matching dashboards.

Type: Array of [DashboardEntry](api-dashboardentry.md) objects

**NextToken**

The token that marks the start of the next batch of returned results.

Type: String

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/listdashboards.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/listdashboards.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListAlarmMuteRules

ListManagedInsightRules
