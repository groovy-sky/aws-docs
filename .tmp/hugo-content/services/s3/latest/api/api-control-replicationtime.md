# ReplicationTime

A container that specifies S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC
is enabled and the time when all objects and operations on objects must be
replicated.

###### Note

This is not supported by Amazon S3 on Outposts buckets.

## Contents

**Status**

Specifies whether S3 Replication Time Control (S3 RTC) is enabled.

Type: String

Valid Values: `Enabled | Disabled`

Required: Yes

**Time**

A container that specifies the time by which replication should be complete for all
objects and operations on objects.

Type: [ReplicationTimeValue](api-control-replicationtimevalue.md) data type

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/replicationtime.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/replicationtime.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/replicationtime.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationRuleFilter

ReplicationTimeValue

All content copied from https://docs.aws.amazon.com/.
