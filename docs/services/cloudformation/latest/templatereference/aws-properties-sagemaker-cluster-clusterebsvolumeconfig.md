This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterEbsVolumeConfig

Defines the configuration for attaching an additional Amazon Elastic Block Store (EBS)
volume to each instance of the SageMaker HyperPod cluster instance group. To learn more, see [SageMaker HyperPod release notes: June 20, 2024](../../../sagemaker/latest/dg/sagemaker-hyperpod-release-notes.md#sagemaker-hyperpod-release-notes-20240620).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RootVolume" : Boolean,
  "VolumeKmsKeyId" : String,
  "VolumeSizeInGB" : Integer
}

```

### YAML

```yaml

  RootVolume: Boolean
  VolumeKmsKeyId: String
  VolumeSizeInGB: Integer

```

## Properties

`RootVolume`

Specifies whether the configuration is for the cluster's root or secondary Amazon EBS
volume. You can specify two `ClusterEbsVolumeConfig` fields to configure both
the root and secondary volumes. Set the value to `True` if you'd like to provide
your own customer managed AWS KMS key to encrypt the root volume. When
`True`:

- The configuration is applied to the root volume.

- You can't specify the `VolumeSizeInGB` field. The size of the root
volume is determined for you.

- You must specify a KMS key ID for `VolumeKmsKeyId` to encrypt the
root volume with your own KMS key instead of an AWS owned
KMS key.

Otherwise, by default, the value is `False`, and the following
applies:

- The configuration is applied to the secondary volume, while the root volume is
encrypted with an AWS owned key.

- You must specify the `VolumeSizeInGB` field.

- You can optionally specify the `VolumeKmsKeyId` to encrypt the
secondary volume with your own KMS key instead of an AWS owned
KMS key.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeKmsKeyId`

The ID of a KMS key to encrypt the Amazon EBS volume.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9:/_-]*$`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeSizeInGB`

The size in gigabytes (GB) of the additional EBS volume to be attached to the instances
in the SageMaker HyperPod cluster instance group. The additional EBS volume is attached to each
instance within the SageMaker HyperPod cluster instance group and mounted to
`/opt/sagemaker`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterCapacityRequirements

ClusterFsxLustreConfig

All content copied from https://docs.aws.amazon.com/.
