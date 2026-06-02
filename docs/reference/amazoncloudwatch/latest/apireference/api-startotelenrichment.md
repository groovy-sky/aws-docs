---
title: "StartOTelEnrichment"
---

# StartOTelEnrichment

Enables enrichment and PromQL access for CloudWatch vended metrics for [supported AWS resources](../../../../services/amazoncloudwatch/latest/monitoring/usingresourcetagsfortelemetry.md) in the account. Once enabled,
metrics that contain a resource identifier dimension (for example, EC2
`CPUUtilization` with an `InstanceId` dimension) are enriched
with resource ARN and resource tag labels and become queryable using PromQL.

Before calling this operation, you must enable resource tags on telemetry for your
account. For more information, see [Enable\
resource tags on telemetry](../../../../services/amazoncloudwatch/latest/monitoring/enableresourcetagsontelemetry.md).

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/monitoring-2010-08-01/StartOTelEnrichment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/monitoring-2010-08-01/StartOTelEnrichment)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartMetricStreams

StopMetricStreams

All content copied from https://docs.aws.amazon.com/.
