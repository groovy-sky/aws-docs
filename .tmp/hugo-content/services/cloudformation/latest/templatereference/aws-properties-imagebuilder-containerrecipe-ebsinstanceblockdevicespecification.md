This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ContainerRecipe EbsInstanceBlockDeviceSpecification

Amazon EBS-specific block device mapping specifications.

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
  "Throughput" : Integer,
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
  Throughput: Integer
  VolumeSize: Integer
  VolumeType: String

```

## Properties

`DeleteOnTermination`

Use to configure delete on termination of the associated device.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Encrypted`

Use to configure device encryption.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Iops`

Use to configure device IOPS.

_Required_: No

_Type_: Integer

_Minimum_: `100`

_Maximum_: `64000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The Amazon Resource Name (ARN) that uniquely identifies the KMS key to use when encrypting the device.
This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](../../../kms/latest/developerguide/concepts.md#key-id-key-ARN)
in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotId`

The snapshot that defines the device contents.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Throughput`

**For GP3 volumes only** – The throughput in MiB/s
that the volume supports.

_Required_: No

_Type_: Integer

_Minimum_: `125`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeSize`

Use to override the device's volume size.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `16000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeType`

Use to override the device's volume type.

_Required_: No

_Type_: String

_Allowed values_: `standard | io1 | io2 | gp2 | gp3 | sc1 | st1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentParameter

InstanceBlockDeviceMapping

All content copied from https://docs.aws.amazon.com/.
