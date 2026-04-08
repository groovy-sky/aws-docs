# GetOTelEnrichment

Returns the current status of vended metric enrichment for the account, including
whether CloudWatch vended metrics are enriched with resource ARN and resource tag
labels and queryable using PromQL. For the list of supported resources, see
[Supported AWS infrastructure metrics](../../../../services/amazoncloudwatch/latest/monitoring/usingresourcetagsfortelemetry.md).

## Response Elements

The following element is returned by the service.

**Status**

The status of OTel enrichment for the account. Valid values are
`Running` (enrichment is enabled) and `Stopped`
(enrichment is disabled).

Type: String

Valid Values: `Running | Stopped`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/getotelenrichment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/getotelenrichment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetMetricWidgetImage

ListAlarmMuteRules

All content copied from https://docs.aws.amazon.com/.
