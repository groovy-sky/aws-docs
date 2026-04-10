This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket ReplicationDestination

A container for information about the replication destination and its configurations including
enabling the S3 Replication Time Control (S3 RTC).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessControlTranslation" : AccessControlTranslation,
  "Account" : String,
  "Bucket" : String,
  "EncryptionConfiguration" : EncryptionConfiguration,
  "Metrics" : Metrics,
  "ReplicationTime" : ReplicationTime,
  "StorageClass" : String
}

```

### YAML

```yaml

  AccessControlTranslation:
    AccessControlTranslation
  Account: String
  Bucket: String
  EncryptionConfiguration:
    EncryptionConfiguration
  Metrics:
    Metrics
  ReplicationTime:
    ReplicationTime
  StorageClass: String

```

## Properties

`AccessControlTranslation`

Specify this only in a cross-account scenario (where source and destination bucket owners are not
the same), and you want to change replica ownership to the AWS account that owns the destination
bucket. If this is not specified in the replication configuration, the replicas are owned by same
AWS account that owns the source object.

_Required_: No

_Type_: [AccessControlTranslation](aws-properties-s3-bucket-accesscontroltranslation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Account`

Destination bucket owner account ID. In a cross-account scenario, if you direct Amazon S3
to change replica ownership to the AWS account that owns the destination
bucket by specifying the `AccessControlTranslation` property, this is the account
ID of the destination bucket owner. For more information, see [Cross-Region Replication Additional\
Configuration: Change Replica Owner](../../../s3/latest/dev/crr-change-owner.md) in the _Amazon S3 User_
_Guide_.

If you specify the `AccessControlTranslation` property, the
`Account` property is required.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bucket`

The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfiguration`

Specifies encryption-related information.

_Required_: No

_Type_: [EncryptionConfiguration](aws-properties-s3-bucket-encryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metrics`

A container specifying replication metrics-related settings enabling replication metrics and
events.

_Required_: No

_Type_: [Metrics](aws-properties-s3-bucket-metrics.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationTime`

A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all
objects and operations on objects must be replicated. Must be specified together with a
`Metrics` block.

_Required_: No

_Type_: [ReplicationTime](aws-properties-s3-bucket-replicationtime.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageClass`

The storage class to use when replicating objects, such as S3 Standard or reduced redundancy. By
default, Amazon S3 uses the storage class of the source object to create the object replica.

For valid values, see the `StorageClass` element of the [PUT Bucket replication](../../../s3/latest/api/restbucketputreplication.md) action in the
_Amazon S3 API Reference_.

`FSX_OPENZFS` is not an accepted value when replicating objects.

_Required_: No

_Type_: String

_Allowed values_: `DEEP_ARCHIVE | GLACIER | GLACIER_IR | INTELLIGENT_TIERING | ONEZONE_IA | REDUCED_REDUNDANCY | STANDARD | STANDARD_IA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationConfiguration

ReplicationRule

All content copied from https://docs.aws.amazon.com/.
