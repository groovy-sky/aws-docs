This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance Ebs

Specifies a block device for an EBS volume.

`Ebs` is a property of the [BlockDeviceMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-blockdevicemapping.html) property type.

###### Important

After the instance is running, you can modify only the
`DeleteOnTermination` parameters for the attached volumes without
interrupting the instance. Modifying any other parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeleteOnTermination" : Boolean,
  "Encrypted" : Boolean,
  "Iops" : Integer,
  "KmsKeyId" : String,
  "SnapshotId" : String,
  "VolumeSize" : Integer,
  "VolumeType" : String
}

```

### YAML

```yaml

  DeleteOnTermination: Boolean
  Encrypted: Boolean
  Iops: Integer
  KmsKeyId: String
  SnapshotId: String
  VolumeSize: Integer
  VolumeType: String

```

## Properties

`DeleteOnTermination`

Indicates whether the EBS volume is deleted on instance termination. For more
information, see [Preserving Amazon EBS volumes on instance termination](../../../ec2/latest/userguide/terminating-instances.md#preserving-volumes-on-termination) in the _Amazon_
_EC2 User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Encrypted`

Indicates whether the volume should be encrypted. The effect of setting the encryption
state to `true` depends on the volume origin (new or from a snapshot), starting
encryption state, ownership, and whether encryption by default is enabled. For more
information, see [Encryption by\
default](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default) in the _Amazon Elastic Compute Cloud User_
_Guide_.

Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS
encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Iops`

The number of I/O operations per second (IOPS). For `gp3`, `io1`,
and `io2` volumes, this represents the number of IOPS that are provisioned for
the volume. For `gp2` volumes, this represents the baseline performance of the
volume and the rate at which the volume accumulates I/O credits for bursting.

The following are the supported values for each volume type:

- `gp3`: 3,000-16,000 IOPS

- `io1`: 100-64,000 IOPS

- `io2`: 100-64,000 IOPS

For `io1` and `io2` volumes, we guarantee 64,000 IOPS only for
[Instances built on\
the Nitro System](../../../ec2/latest/userguide/instance-types.md#ec2-nitro-instances). Other instance families guarantee performance up to 32,000
IOPS.

This parameter is required for `io1` and `io2` volumes. The
default for `gp3` volumes is 3,000 IOPS. This parameter is not supported for
`gp2`, `st1`, `sc1`, or `standard`
volumes.

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: Conditional

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`KmsKeyId`

The identifier of the AWS KMS key to use for Amazon EBS encryption. If
`KmsKeyId` is specified, the encrypted state must be `true`. If
the encrypted state is `true` but you do not specify `KmsKeyId`, your
KMS key for EBS is used.

You can specify the KMS key using any of the following:

- Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.

- Key alias. For example, alias/ExampleAlias.

- Key ARN. For example,
arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.

- Alias ARN. For example,
arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias.

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SnapshotId`

The ID of the snapshot.

If you specify both `SnapshotId` and `VolumeSize`,
`VolumeSize` must be equal or greater than the size of the snapshot.

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: Conditional

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VolumeSize`

The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size.
If you specify a snapshot, the default is the snapshot size. You can specify a volume size
that is equal to or larger than the snapshot size.

The following are the supported volumes sizes for each volume type:

- `gp2` and `gp3`:1-16,384

- `io1` and `io2`: 4-16,384

- `st1` and `sc1`: 125-16,384

- `standard`: 1-1,024

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: Conditional

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`VolumeType`

The volume type. For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the
_Amazon EC2 User Guide_. If the volume type is `io1` or
`io2`, you must specify the IOPS that the volume supports.

###### Important

After the instance is running, modifying this parameter results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

_Required_: No

_Type_: String

_Allowed values_: `standard | io1 | io2 | gp2 | sc1 | st1 | gp3`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Examples

### Create an EBS volume from a snapshot

This example creates a 50GB io1 EBS volume from a snapshot, and configures it to
support 1000 IOPS and to persist after terminating the instance to which it is
attached.

#### JSON

```json

{
    "DeviceName": "/dev/sdc",
    "Ebs": {
        "SnapshotId": "snap-xxxxxxxx",
        "VolumeSize": "50",
        "VolumeType": "io1",
        "Iops": "1000",
        "DeleteOnTermination": "false"
    }
}
```

#### YAML

```yaml

BlockDeviceMappings:
  - DeviceName: /dev/sdc
    Ebs:
      SnapshotId: snap-xxxxxxxx
      VolumeSize: 50
      VolumeType: io1
      Iops: 1000
      DeleteOnTermination: false
```

## See also

- [CreateVolume](../../../../reference/awsec2/latest/apireference/api-createvolume.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreditSpecification

ElasticGpuSpecification
