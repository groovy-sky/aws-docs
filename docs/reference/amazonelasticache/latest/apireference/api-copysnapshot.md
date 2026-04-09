# CopySnapshot

Makes a copy of an existing snapshot.

###### Note

This operation is valid for Valkey or Redis OSS only.

###### Important

Users or groups that have permissions to use the `CopySnapshot`
operation can create their own Amazon S3 buckets and copy snapshots to it. To
control access to your snapshots, use an IAM policy to control who has the ability
to use the `CopySnapshot` operation. For more information about using IAM
to control the use of ElastiCache operations, see [Exporting\
Snapshots](../../../../services/amazonelasticache/latest/dg/backups-exporting.md) and [Authentication & Access\
Control](../../../../services/amazonelasticache/latest/dg/iam.md).

You could receive the following error messages.

###### Error Messages

- **Error Message:** The S3 bucket %s is outside of
the region.

**Solution:** Create an Amazon S3 bucket in the
same region as your snapshot. For more information, see [Step 1: Create an Amazon S3 Bucket](../../../../services/amazonelasticache/latest/dg/backups-exporting.md#backups-exporting-create-s3-bucket) in the ElastiCache User
Guide.

- **Error Message:** The S3 bucket %s does not
exist.

**Solution:** Create an Amazon S3 bucket in the
same region as your snapshot. For more information, see [Step 1: Create an Amazon S3 Bucket](../../../../services/amazonelasticache/latest/dg/backups-exporting.md#backups-exporting-create-s3-bucket) in the ElastiCache User
Guide.

- **Error Message:** The S3 bucket %s is not owned
by the authenticated user.

**Solution:** Create an Amazon S3 bucket in the
same region as your snapshot. For more information, see [Step 1: Create an Amazon S3 Bucket](../../../../services/amazonelasticache/latest/dg/backups-exporting.md#backups-exporting-create-s3-bucket) in the ElastiCache User
Guide.

- **Error Message:** The authenticated user does
not have sufficient permissions to perform the desired activity.

**Solution:** Contact your system administrator
to get the needed permissions.

- **Error Message:** The S3 bucket %s already
contains an object with key %s.

**Solution:** Give the
`TargetSnapshotName` a new and unique value. If exporting a
snapshot, you could alternatively create a new Amazon S3 bucket and use this
same value for `TargetSnapshotName`.

- **Error Message:** ElastiCache has not been
granted READ permissions %s on the S3 Bucket.

**Solution:** Add List and Read permissions on
the bucket. For more information, see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket](../../../../services/amazonelasticache/latest/dg/backups-exporting.md#backups-exporting-grant-access) in the
ElastiCache User Guide.

- **Error Message:** ElastiCache has not been
granted WRITE permissions %s on the S3 Bucket.

**Solution:** Add Upload/Delete permissions on
the bucket. For more information, see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket](../../../../services/amazonelasticache/latest/dg/backups-exporting.md#backups-exporting-grant-access) in the
ElastiCache User Guide.

- **Error Message:** ElastiCache has not been
granted READ\_ACP permissions %s on the S3 Bucket.

**Solution:** Add View Permissions on the bucket.
For more information, see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket](../../../../services/amazonelasticache/latest/dg/backups-exporting.md#backups-exporting-grant-access) in the
ElastiCache User Guide.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**SourceSnapshotName**

The name of an existing snapshot from which to make a copy.

Type: String

Required: Yes

**TargetSnapshotName**

A name for the snapshot copy. ElastiCache does not permit overwriting a snapshot,
therefore this name must be unique within its context - ElastiCache or an Amazon S3
bucket if exporting. This value is stored as a lowercase string.

Type: String

Required: Yes

**KmsKeyId**

The ID of the KMS key used to encrypt the target snapshot.

Type: String

Required: No

**Tags.Tag.N**

A list of tags to be added to this resource. A tag is a key-value pair. A tag key must
be accompanied by a tag value, although null is accepted.

Type: Array of [Tag](api-tag.md) objects

Required: No

**TargetBucket**

The Amazon S3 bucket to which the snapshot is exported. This parameter is used only
when exporting a snapshot for external access.

When using this parameter to export a snapshot, be sure Amazon ElastiCache has the
needed permissions to this S3 bucket. For more information, see [Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket](../../../../services/amazonelasticache/latest/dg/backups-exporting.md#backups-exporting-grant-access) in the
_Amazon ElastiCache User Guide_.

For more information, see [Exporting a\
Snapshot](../../../../services/amazonelasticache/latest/dg/backups-exporting.md) in the _Amazon ElastiCache User Guide_.

Type: String

Required: No

## Response Elements

The following element is returned by the service.

**Snapshot**

Represents a copy of an entire Valkey or Redis OSS cluster as of the time when the snapshot was
taken.

Type: [Snapshot](api-snapshot.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterCombination**

Two or more incompatible parameters were specified.

**message**

Two or more parameters that must not be used together were used together.

HTTP Status Code: 400

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**InvalidSnapshotState**

The current state of the snapshot does not allow the requested operation to
occur.

HTTP Status Code: 400

**SnapshotAlreadyExistsFault**

You already have a snapshot with the given name.

HTTP Status Code: 400

**SnapshotNotFoundFault**

The requested snapshot name does not refer to an existing snapshot.

HTTP Status Code: 404

**SnapshotQuotaExceededFault**

The request cannot be processed because it would exceed the maximum number of
snapshots.

HTTP Status Code: 400

**TagQuotaPerResourceExceeded**

The request cannot be processed because it would cause the resource to have more than
the allowed number of tags. The maximum number of tags permitted on a resource is
50.

HTTP Status Code: 400

## Examples

### Snapshot copy

The following example makes a copy of the snapshot
`automatic.my-redis-primary-2016-04-27-03-15` named
`my-snapshot-copy`.

#### Sample Request

```

https://elasticache.us-west-2.amazonaws.com/
    ?Action=CopySnapshot
    &SourceSnapshotName=automatic.my-redis-primary-2016-04-27-03-15
    &TargetSnapshotName=my-snapshot-copy
    &SignatureVersion=4
    &SignatureMethod=HmacSHA256
    &Timestamp=20141201T220302Z
    &Version=2015-02-02
    &X-Amz-Algorithm=AWS4-HMAC-SHA256
    &X-Amz-Date=20141201T220302Z
    &X-Amz-SignedHeaders=Host
    &X-Amz-Expires=20141201T220302Z
    &X-Amz-Credential=<credential>
    &X-Amz-Signature=<signature>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/elasticache-2015-02-02/copysnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/copysnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyServerlessCacheSnapshot

CreateCacheCluster

All content copied from https://docs.aws.amazon.com/.
