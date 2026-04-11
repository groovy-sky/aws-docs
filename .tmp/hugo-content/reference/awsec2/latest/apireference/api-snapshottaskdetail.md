# SnapshotTaskDetail

Details about the import snapshot task.

## Contents

**description**

The description of the disk image being imported.

Type: String

Required: No

**diskImageSize**

The size of the disk in the snapshot, in GiB.

Type: Double

Required: No

**encrypted**

Indicates whether the snapshot is encrypted.

Type: Boolean

Required: No

**format**

The format of the disk image from which the snapshot is created.

Type: String

Required: No

**kmsKeyId**

The identifier for the KMS key that was used to create the encrypted snapshot.

Type: String

Required: No

**progress**

The percentage of completion for the import snapshot task.

Type: String

Required: No

**snapshotId**

The snapshot ID of the disk being imported.

Type: String

Required: No

**status**

A brief status for the import snapshot task.

Type: String

Required: No

**statusMessage**

A detailed status message for the import snapshot task.

Type: String

Required: No

**url**

The URL of the disk image from which the snapshot is created.

Type: String

Required: No

**userBucket**

The Amazon S3 bucket for the disk image.

Type: [UserBucketDetails](api-userbucketdetails.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/snapshottaskdetail.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/snapshottaskdetail.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/snapshottaskdetail.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SnapshotRecycleBinInfo

SnapshotTierStatus
