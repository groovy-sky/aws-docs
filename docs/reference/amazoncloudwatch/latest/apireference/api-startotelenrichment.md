# StartOTelEnrichment

Enables enrichment and PromQL access for CloudWatch vended metrics for
[supported AWS resources](../../../../services/amazoncloudwatch/latest/monitoring/usingresourcetagsfortelemetry.md) in the account. Once enabled, metrics that
contain a resource identifier dimension (for example, EC2
`CPUUtilization` with an `InstanceId` dimension) are enriched
with resource ARN and resource tag labels and become queryable using
PromQL.

Before calling this operation, you must enable resource tags on telemetry for
your account. For more information, see [Enable resource tags on telemetry](../../../../services/amazoncloudwatch/latest/monitoring/enableresourcetagsontelemetry.md).

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/startotelenrichment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/startotelenrichment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartMetricStreams

StopMetricStreams

All content copied from https://docs.aws.amazon.com/.
