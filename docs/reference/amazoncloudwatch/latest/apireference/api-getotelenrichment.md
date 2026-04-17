---
title: "GetOTelEnrichment"
---

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/monitoring-2010-08-01/GetOTelEnrichment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/monitoring-2010-08-01/GetOTelEnrichment)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetMetricWidgetImage

ListAlarmMuteRules

All content copied from https://docs.aws.amazon.com/.
