# Metrics

A container that specifies replication metrics-related settings.

## Contents

**Status**

Specifies whether replication metrics are enabled.

Type: String

Valid Values: `Enabled | Disabled`

Required: Yes

**EventThreshold**

A container that specifies the time threshold for emitting the
`s3:Replication:OperationMissedThreshold` event.

###### Note

This is not supported by Amazon S3 on Outposts buckets.

Type: [ReplicationTimeValue](api-control-replicationtimevalue.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/metrics.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/metrics.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/metrics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MatchObjectSize

MultiRegionAccessPointPolicyDocument

All content copied from https://docs.aws.amazon.com/.
