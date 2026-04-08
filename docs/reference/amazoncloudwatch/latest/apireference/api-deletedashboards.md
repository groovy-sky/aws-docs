# DeleteDashboards

Deletes all dashboards that you specify. You can specify up to 100 dashboards to
delete. If there is an error during this call, no dashboards are deleted.

## Request Parameters

**DashboardNames**

The dashboards to be deleted. This parameter is required.

Type: Array of strings

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

This operation attempted to create a resource that already exists.

HTTP Status Code: 409

**InternalServiceError**

Request processing has failed due to some unknown error, exception, or
failure.

**Message**

HTTP Status Code: 500

**InvalidParameterValue**

The value of an input parameter is bad or out-of-range.

**message**

HTTP Status Code: 400

**ResourceNotFound**

The specified dashboard does not exist.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/deletedashboards.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/deletedashboards.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteAnomalyDetector

DeleteInsightRules
