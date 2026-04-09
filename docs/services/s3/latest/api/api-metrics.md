# Metrics

A container specifying replication metrics-related settings enabling replication metrics and
events.

## Contents

**Status**

Specifies whether the replication metrics are enabled.

Type: String

Valid Values: `Enabled | Disabled`

Required: Yes

**EventThreshold**

A container specifying the time threshold for emitting the
`s3:Replication:OperationMissedThreshold` event.

Type: [ReplicationTimeValue](api-replicationtimevalue.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/metrics.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/metrics.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/metrics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataTableEncryptionConfiguration

MetricsAndOperator

All content copied from https://docs.aws.amazon.com/.
