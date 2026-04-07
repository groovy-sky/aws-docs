# CreateSnapshot

Creates a snapshot of an EBS volume and stores it in Amazon S3. You can use snapshots for
backups, to make copies of EBS volumes, and to save data before shutting down an
instance.

The location of the source EBS volume determines where you can create the snapshot.

- If the source volume is in a Region, you must create the snapshot in the same
Region as the volume.

- If the source volume is in a Local Zone, you can create the snapshot in the same Local
Zone or in its parent AWS Region.

- If the source volume is on an Outpost, you can create the snapshot on the same
Outpost or in its parent AWS Region.

When a snapshot is created, any AWS Marketplace product codes that are associated with the
source volume are propagated to the snapshot.

You can take a snapshot of an attached volume that is in use. However, snapshots only
capture data that has been written to your Amazon EBS volume at the time the snapshot command is
issued; this might exclude any data that has been cached by any applications or the operating
system. If you can pause any file systems on the volume long enough to take a snapshot, your
snapshot should be complete. However, if you cannot pause all file writes to the volume, you
should unmount the volume from within the instance, issue the snapshot command, and then
remount the volume to ensure a consistent and complete snapshot. You may remount and use your
volume while the snapshot status is `pending`.

When you create a snapshot for an EBS volume that serves as a root device, we recommend
that you stop the instance before taking the snapshot.

Snapshots that are taken from encrypted volumes are automatically encrypted. Volumes that
are created from encrypted snapshots are also automatically encrypted. Your encrypted volumes
and any associated snapshots always remain protected. For more information, see [Amazon EBS encryption](../../../../services/ebs/latest/userguide/ebs-encryption.md)
in the _Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Description**

A description for the snapshot.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Location**

###### Note

Only supported for volumes in Local Zones. If the source volume is not in a Local Zone,
omit this parameter.

- To create a local snapshot in the same Local Zone as the source volume, specify
`local`.

- To create a regional snapshot in the parent Region of the Local Zone, specify
`regional` or omit this parameter.

Default value: `regional`

Type: String

Valid Values: `regional | local`

Required: No

**OutpostArn**

###### Note

Only supported for volumes on Outposts. If the source volume is not on an Outpost,
omit this parameter.

- To create the snapshot on the same Outpost as the source volume, specify the
ARN of that Outpost. The snapshot must be created on the same Outpost as the volume.

- To create the snapshot in the parent Region of the Outpost, omit this parameter.

For more information, see [Create local snapshots from volumes on an Outpost](../../../../services/ebs/latest/userguide/snapshots-outposts.md#create-snapshot) in the _Amazon EBS User Guide_.

Type: String

Required: No

**TagSpecification.N**

The tags to apply to the snapshot during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**VolumeId**

The ID of the Amazon EBS volume.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**availabilityZone**

The Availability Zone or Local Zone of the snapshot. For example, `us-west-1a`
(Availability Zone) or `us-west-2-lax-1a` (Local Zone).

Type: String

**completionDurationMinutes**

###### Note

Only for snapshot copies created with time-based snapshot copy operations.

The completion duration requested for the time-based snapshot copy operation.

Type: Integer

**completionTime**

The time stamp when the snapshot was completed.

Type: Timestamp

**dataEncryptionKeyId**

The data encryption key identifier for the snapshot. This value is a unique identifier
that corresponds to the data encryption key that was used to encrypt the original volume or
snapshot copy. Because data encryption keys are inherited by volumes created from snapshots,
and vice versa, if snapshots share the same data encryption key identifier, then they belong
to the same volume/snapshot lineage. This parameter is only returned by [DescribeSnapshots](api-describesnapshots.md).

Type: String

**description**

The description for the snapshot.

Type: String

**encrypted**

Indicates whether the snapshot is encrypted.

Type: Boolean

**fullSnapshotSizeInBytes**

The full size of the snapshot, in bytes.

###### Important

This is **not** the incremental size of the snapshot.
This is the full snapshot size and represents the size of all the blocks that were
written to the source volume at the time the snapshot was created.

Type: Long

**kmsKeyId**

The Amazon Resource Name (ARN) of the AWS KMS key that was used to protect the
volume encryption key for the parent volume.

Type: String

**outpostArn**

The ARN of the Outpost on which the snapshot is stored. For more information, see [Amazon EBS local snapshots on Outposts](../../../../services/ebs/latest/userguide/snapshots-outposts.md) in the
_Amazon EBS User Guide_.

Type: String

**ownerAlias**

The AWS owner alias, from an Amazon-maintained list ( `amazon`). This is not
the user-configured AWS account alias set using the IAM console.

Type: String

**ownerId**

The ID of the AWS account that owns the EBS snapshot.

Type: String

**progress**

The progress of the snapshot, as a percentage.

Type: String

**requestId**

The ID of the request.

Type: String

**restoreExpiryTime**

Only for archived snapshots that are temporarily restored. Indicates the date and
time when a temporarily restored snapshot will be automatically re-archived.

Type: Timestamp

**snapshotId**

The ID of the snapshot. Each snapshot receives a unique identifier when it is
created.

Type: String

**sseType**

Reserved for future use.

Type: String

Valid Values: `sse-ebs | sse-kms | none`

**startTime**

The time stamp when the snapshot was initiated.

Type: Timestamp

**status**

The snapshot state.

Type: String

Valid Values: `pending | completed | error | recoverable | recovering`

**statusMessage**

Encrypted Amazon EBS snapshots are copied asynchronously. If a snapshot copy operation fails
(for example, if the proper AWS KMS permissions are not obtained) this field displays error
state details to help you diagnose why the error occurred. This parameter is only returned by
[DescribeSnapshots](api-describesnapshots.md).

Type: String

**storageTier**

The storage tier in which the snapshot is stored. `standard` indicates
that the snapshot is stored in the standard snapshot storage tier and that it is ready
for use. `archive` indicates that the snapshot is currently archived and that
it must be restored before it can be used.

Type: String

Valid Values: `archive | standard`

**tagSet**

Any tags assigned to the snapshot.

Type: Array of [Tag](api-tag.md) objects

**transferType**

###### Note

Only for snapshot copies.

Indicates whether the snapshot copy was created with a standard or time-based
snapshot copy operation. Time-based snapshot copy operations complete within the
completion duration specified in the request. Standard snapshot copy operations
are completed on a best-effort basis.

- `standard` \- The snapshot copy was created with a standard
snapshot copy operation.

- `time-based` \- The snapshot copy was created with a time-based
snapshot copy operation.

Type: String

Valid Values: `time-based | standard`

**volumeId**

The ID of the volume that was used to create the snapshot. Snapshots created by a copy
snapshot operation have an arbitrary volume ID that you should not use for any purpose.

Type: String

**volumeSize**

The size of the volume, in GiB.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a snapshot of the volume with the ID
`vol-1234567890abcdef0`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateSnapshot
&VolumeId=vol-1234567890abcdef0
&Description=Daily+Backup
&AUTHPARAMS
```

#### Sample Response

```

<CreateSnapshotResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <snapshotId>snap-1234567890abcdef0</snapshotId>
  <volumeId>vol-1234567890abcdef0</volumeId>
  <status>pending</status>
  <startTime>YYYY-MM-DDTHH:MM:SS.000Z</startTime>
  <progress>60%</progress>
  <ownerId>111122223333</ownerId>
  <volumeSize>30</volumeSize>
  <description>Daily Backup</description>
</CreateSnapshotResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateSnapshot)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateSnapshot)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createsnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createsnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createsnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createsnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createsnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createsnapshot.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateSnapshot)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createsnapshot.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateSecurityGroup

CreateSnapshots
