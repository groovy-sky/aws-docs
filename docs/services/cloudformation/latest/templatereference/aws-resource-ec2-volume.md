This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Volume

Specifies an Amazon Elastic Block Store (Amazon EBS) volume. You can create an empty volume, a
volume from a snapshot, or a volume copy from an existing source volume.

###### Important

- When you use AWS CloudFormation to update an Amazon EBS volume that modifies
`Iops`, `Size`, or `VolumeType`, there is a cooldown
period before another operation can occur. This can cause your stack to report being in
`UPDATE_IN_PROGRESS` or `UPDATE_ROLLBACK_IN_PROGRESS` for long
periods of time. Some common scenarios when you might encounter a cooldown period for Amazon EBS
include:

- You successfully update an Amazon EBS volume and the update succeeds. When you
attempt another update within the cooldown window, that update will be subject to a
cooldown period.

- You successfully update an Amazon EBS volume and the update succeeds but another
change in your `update-stack` call fails. The rollback will be subject to
a cooldown period.

For more information, see [Requirements for EBS volume modifications](https://docs.aws.amazon.com/ebs/latest/userguide/modify-volume-requirements.html).

- Amazon EBS does not support sizing down an Amazon EBS volume. AWS CloudFormation
does not attempt to modify an Amazon EBS volume to a smaller size on rollback.

**DeletionPolicy attribute**

To control how AWS CloudFormation handles the volume when the stack is deleted,
set a deletion policy for your volume. You can choose to retain the volume, to delete the
volume, or to create a snapshot of the volume. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).

###### Note

If you set a deletion policy that creates a snapshot, all tags on the volume are
included in the snapshot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::Volume",
  "Properties" : {
      "AutoEnableIO" : Boolean,
      "AvailabilityZone" : String,
      "AvailabilityZoneId" : String,
      "Encrypted" : Boolean,
      "Iops" : Integer,
      "KmsKeyId" : String,
      "MultiAttachEnabled" : Boolean,
      "OutpostArn" : String,
      "Size" : Integer,
      "SnapshotId" : String,
      "SourceVolumeId" : String,
      "Tags" : [ Tag, ... ],
      "Throughput" : Integer,
      "VolumeInitializationRate" : Integer,
      "VolumeType" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::Volume
Properties:
  AutoEnableIO: Boolean
  AvailabilityZone: String
  AvailabilityZoneId: String
  Encrypted: Boolean
  Iops: Integer
  KmsKeyId: String
  MultiAttachEnabled: Boolean
  OutpostArn: String
  Size: Integer
  SnapshotId: String
  SourceVolumeId: String
  Tags:
    - Tag
  Throughput: Integer
  VolumeInitializationRate: Integer
  VolumeType: String

```

## Properties

`AutoEnableIO`

Indicates whether the volume is auto-enabled for I/O operations. By default, Amazon EBS
disables I/O to the volume from attached EC2 instances when it determines that a volume's data is
potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that
the volume be made available immediately if it's impaired, you can configure the volume to
automatically enable I/O.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The ID of the Availability Zone in which to create the volume. For example, `us-east-1a`.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified, but not both.

If you are creating a volume copy, omit this parameter. The volume copy is created in the same Availability Zone
as the source volume.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`AvailabilityZoneId`

The ID of the Availability Zone in which to create the volume. For example, `use1-az1`.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified, but not both.

If you are creating a volume copy, omit this parameter. The volume copy is created in the same Availability Zone
as the source volume.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`Encrypted`

Indicates whether the volume should be encrypted. The effect of setting the encryption state to `true`
depends on the volume origin (new, from a snapshot, or from an existing volume), starting encryption state, ownership,
and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the
_Amazon EBS User Guide_.

If you are creating a volume copy, omit this parameter. The volume is automatically encrypted with the same KMS key
as the source volume. You can't copy unencrypted volumes.

_Required_: No

_Type_: Boolean

_Update requires_: Updates are not supported.

`Iops`

The number of I/O operations per second (IOPS) to provision for the volume.
Required for `io1` and `io2` volumes. Optional for `gp3`
volumes. Omit for all other volume types.

Valid ranges:

- gp3: `3,000 `( _default_) ` - 80,000` IOPS

- io1: `100 - 64,000` IOPS

- io2: `100 - 256,000` IOPS

###### Note

[Instances built on the Nitro System](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html) can support up to 256,000 IOPS. Other instances can support up to 32,000
IOPS.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The identifier of the AWS KMS key to use for Amazon EBS encryption. If `KmsKeyId`
is specified, the encrypted state must be `true`.

If you omit this property and your account is enabled for encryption by default, or **Encrypted**
is set to `true`, then the volume is encrypted using the default key specified for your account. If your account
does not have a default key, then the volume is encrypted using the AWS managed key.

Alternatively, if you want to specify a different key, you can specify one of the following:

- Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.

- Key alias. Specify the alias for the key, prefixed with `alias/`. For example, for a key with the alias
`my_cmk`, use `alias/my_cmk`. Or to specify the AWS managed key, use
`alias/aws/ebs`.

- Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.

- Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.

If you are creating a volume copy, omit this parameter. The volume is automatically encrypted with the same KMS key as the
source volume. You can't copy unencrypted volumes.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`MultiAttachEnabled`

Indicates whether Amazon EBS Multi-Attach is enabled.

AWS CloudFormation does not currently support updating a single-attach volume to
be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or
updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled
volume.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutpostArn`

The Amazon Resource Name (ARN) of the Outpost on which to create the volume.

If you intend to use a volume with an instance running on an outpost, then you must
create the volume on the same outpost as the instance. You can't use a volume created
in an AWS Region with an instance on an AWS outpost, or the other way around.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Size`

The size of the volume, in GiBs.

- Required for new empty volumes.

- Optional for volumes created from snapshots and volume copies. In this case, the size defaults to the size of the
snapshot or source volume. You can optionally specify a size that is equal to or larger than the size of the source
snapshot or volume.

Supported volume sizes:

- gp2: `1 - 16,384` GiB

- gp3: `1 - 65,536` GiB

- io1: `4 - 16,384` GiB

- io2: `4 - 65,536` GiB

- st1 and sc1: `125 - 16,384` GiB

- standard: `1 - 1024` GiB

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotId`

The snapshot from which to create the volume. Only specify to create a volume from a snapshot.
To create a new empty volume, omit this parameter and specify a value for `Size` instead. To create
a volume copy, omit this parameter and specify `SourceVolumeId` instead.

_Required_: Conditional

_Type_: String

_Update requires_: Updates are not supported.

`SourceVolumeId`

The ID of the source EBS volume to copy. When specified, the volume
is created as an exact copy of the specified volume. Only specify to
create a volume copy. To create a new empty volume or to create a volume
from a snapshot, omit this parameter,

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to apply to the volume during creation.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-volume-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Throughput`

The throughput to provision for a volume, with a maximum of 2,000 MiB/s.

This parameter is valid only for `gp3` volumes. The default value is 125.

Valid Range: Minimum value of 125. Maximum value of 2000.

The maximum ratio of throughput to IOPS is 0.25 MiB/s per IOPS. For example, a volume
with 3,000 IOPS can have a maximum throughput of 750 MiB/s (3,000 x 0.25).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeInitializationRate`

Specifies the Amazon EBS Provisioned Rate for Volume Initialization (volume initialization rate), in MiB/s, at which to download
the snapshot blocks from Amazon S3 to the volume. This is also known as _volume_
_initialization_. Specifying a volume initialization rate ensures that the volume is
initialized at a predictable and consistent rate after creation.

This parameter is supported only for volumes created from snapshots. Omit this parameter
if:

- You want to create the volume using fast snapshot restore. You must specify a snapshot
that is enabled for fast snapshot restore. In this case, the volume is fully initialized at
creation.

###### Note

If you specify a snapshot that is enabled for fast snapshot restore and a volume initialization rate,
the volume will be initialized at the specified rate instead of fast snapshot restore.

- You want to create a volume that is initialized at the default rate.

For more information, see [Initialize Amazon EBS volumes](https://docs.aws.amazon.com/ebs/latest/userguide/initalize-volume.html) in the _Amazon EC2 User Guide_.

Valid range: 100 - 300 MiB/s

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeType`

The volume type. This parameter can be one of the following values:

- General Purpose SSD: `gp2` \| `gp3`

- Provisioned IOPS SSD: `io1` \| `io2`

- Throughput Optimized HDD: `st1`

- Cold HDD: `sc1`

- Magnetic: `standard`

###### Important

Throughput Optimized HDD ( `st1`) and Cold HDD ( `sc1`) volumes can't be used as boot volumes.

For more information, see [Amazon EBS volume types](../../../ebs/latest/userguide/ebs-volume-types.md) in the
_Amazon EBS User Guide_.

Default: `gp2`

_Required_: No

_Type_: String

_Allowed values_: `standard | io1 | io2 | gp2 | sc1 | st1 | gp3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:
`vol-5cb85026`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`VolumeId`

The ID of the volume.

## Examples

- [Encrypted Amazon EBS volume with DeletionPolicy](#aws-resource-ec2-volume--examples--Encrypted_Amazon_EBS_volume_with_DeletionPolicy)

- [Provisioned IOPS SSD io1 volume](#aws-resource-ec2-volume--examples--Provisioned_IOPS_SSD_io1_volume)

- [Volume from snapshot with encryption and initialization rate](#aws-resource-ec2-volume--examples--Volume_from_snapshot_with_encryption_and_initialization_rate)

### Encrypted Amazon EBS volume with DeletionPolicy

The following example creates an encrypted `gp2` volume with a
DeletionPolicy attribute that creates a snapshot of the volume when the stack is
deleted.

#### JSON

```json

"NewVolume" : {
   "Type" : "AWS::EC2::Volume",
   "Properties" : {
      "Size" : "100",
      "Encrypted" : "true",
      "AvailabilityZone" : { "Fn::GetAtt" : [ "Ec2Instance", "AvailabilityZone" ] },
      "Tags" : [ {
         "Key" : "MyTag",
         "Value" : "TagValue"
      } ]
   },
   "DeletionPolicy" : "Snapshot"
}
```

#### YAML

```yaml

NewVolume:
  Type: AWS::EC2::Volume
  Properties:
    Size: 100
    Encrypted: true
    AvailabilityZone: !GetAtt Ec2Instance.AvailabilityZone
    Tags:
      - Key: MyTag
        Value: TagValue
  DeletionPolicy: Snapshot
```

### Provisioned IOPS SSD io1 volume

The following example creates a 100 GiB `io1` with `100`
provisioned IOPS.

#### JSON

```json

"NewVolume" : {
   "Type" : "AWS::EC2::Volume",
   "Properties" : {
      "Size" : "100",
      "VolumeType" : "io1",
      "Iops" : "100",
      "AvailabilityZone" : { "Fn::GetAtt" : [ "EC2Instance", "AvailabilityZone" ] }
   }
}
```

#### YAML

```yaml

NewVolume:
  Type: AWS::EC2::Volume
  Properties:
    Size: 100
    VolumeType: io1
    Iops: 100
    AvailabilityZone: !GetAtt Ec2Instance.AvailabilityZone
```

### Volume from snapshot with encryption and initialization rate

The following example creates an encrypted `gp3` volume from a
snapshot. The volume specifies a volume initialization rate for predictable
initialization performance.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Create an encrypted gp3 volume from a snapshot",
  "Parameters": {
    "SnapshotId": {
      "Type": "String",
      "Description": "The ID of the snapshot to create the volume from"
    },
    "KmsKeyId": {
      "Type": "String",
      "Description": "The ARN of the KMS key for volume encryption"
    }
  },
  "Resources": {
    "SnapshotVolume": {
      "Type": "AWS::EC2::Volume",
      "Properties": {
        "AvailabilityZone": {
          "Fn::Select": [
            "0",
            {
              "Fn::GetAZs": ""
            }
          ]
        },
        "SnapshotId": {
          "Ref": "SnapshotId"
        },
        "VolumeType": "gp3",
        "Size": 100,
        "Iops": 3000,
        "Throughput": 125,
        "Encrypted": true,
        "KmsKeyId": {
          "Ref": "KmsKeyId"
        },
        "VolumeInitializationRate": 100,
        "Tags": [
          {
            "Key": "Name",
            "Value": "SnapshotRestoredVolume"
          }
        ]
      },
      "DeletionPolicy": "Snapshot"
    }
  },
  "Outputs": {
    "VolumeId": {
      "Value": {
        "Ref": "SnapshotVolume"
      },
      "Description": "The ID of the created volume"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Create an encrypted gp3 volume from a snapshot
Parameters:
  SnapshotId:
    Type: String
    Description: The ID of the snapshot to create the volume from
  KmsKeyId:
    Type: String
    Description: The ARN of the KMS key for volume encryption
Resources:
  SnapshotVolume:
    Type: AWS::EC2::Volume
    Properties:
      AvailabilityZone: !Select
        - 0
        - !GetAZs ''
      SnapshotId: !Ref SnapshotId
      VolumeType: gp3
      Size: 100
      Iops: 3000
      Throughput: 125
      Encrypted: true
      KmsKeyId: !Ref KmsKeyId
      VolumeInitializationRate: 100
      Tags:
        - Key: Name
          Value: SnapshotRestoredVolume
    DeletionPolicy: Snapshot
Outputs:
  VolumeId:
    Value: !Ref SnapshotVolume
    Description: The ID of the created volume
```

## See also

- [CreateVolume](../../../../reference/awsec2/latest/apireference/api-createvolume.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
