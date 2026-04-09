# Destination

Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket
and S3 Replication Time Control (S3 RTC).

## Contents

**Bucket**

The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.

Type: String

Required: Yes

**AccessControlTranslation**

Specify this only in a cross-account scenario (where source and destination bucket owners are not
the same), and you want to change replica ownership to the AWS account that owns the destination
bucket. If this is not specified in the replication configuration, the replicas are owned by same
AWS account that owns the source object.

Type: [AccessControlTranslation](api-accesscontroltranslation.md) data type

Required: No

**Account**

Destination bucket owner account ID. In a cross-account scenario, if you direct Amazon S3 to change
replica ownership to the AWS account that owns the destination bucket by specifying the
`AccessControlTranslation` property, this is the account ID of the destination bucket
owner. For more information, see [Replication Additional Configuration: Changing\
the Replica Owner](../dev/replication-change-owner.md) in the _Amazon S3 User Guide_.

Type: String

Required: No

**EncryptionConfiguration**

A container that provides information about encryption. If `SourceSelectionCriteria` is
specified, you must specify this element.

Type: [EncryptionConfiguration](api-encryptionconfiguration.md) data type

Required: No

**Metrics**

A container specifying replication metrics-related settings enabling replication metrics and
events.

Type: [Metrics](api-metrics.md) data type

Required: No

**ReplicationTime**

A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all
objects and operations on objects must be replicated. Must be specified together with a
`Metrics` block.

Type: [ReplicationTime](api-replicationtime.md) data type

Required: No

**StorageClass**

The storage class to use when replicating objects, such as S3 Standard or reduced redundancy. By
default, Amazon S3 uses the storage class of the source object to create the object replica.

For valid values, see the `StorageClass` element of the [PUT Bucket replication](restbucketputreplication.md) action in the
_Amazon S3 API Reference_.

`FSX_OPENZFS` is not an accepted value when replicating objects.

Type: String

Valid Values: `STANDARD | REDUCED_REDUNDANCY | STANDARD_IA | ONEZONE_IA | INTELLIGENT_TIERING | GLACIER | DEEP_ARCHIVE | OUTPOSTS | GLACIER_IR | SNOW | EXPRESS_ONEZONE | FSX_OPENZFS | FSX_ONTAP`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/destination.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/destination.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/destination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteMarkerReplication

DestinationResult

All content copied from https://docs.aws.amazon.com/.
