This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterInstanceStorageConfig

Defines the configuration for attaching additional storage to the instances in the
SageMaker HyperPod cluster instance group. To learn more, see [SageMaker HyperPod release notes: June 20, 2024](../../../sagemaker/latest/dg/sagemaker-hyperpod-release-notes.md#sagemaker-hyperpod-release-notes-20240620).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EbsVolumeConfig" : ClusterEbsVolumeConfig,
  "FsxLustreConfig" : ClusterFsxLustreConfig,
  "FsxOpenZfsConfig" : ClusterFsxOpenZfsConfig
}

```

### YAML

```yaml

  EbsVolumeConfig:
    ClusterEbsVolumeConfig
  FsxLustreConfig:
    ClusterFsxLustreConfig
  FsxOpenZfsConfig:
    ClusterFsxOpenZfsConfig

```

## Properties

`EbsVolumeConfig`

Defines the configuration for attaching additional Amazon Elastic Block Store (EBS)
volumes to the instances in the SageMaker HyperPod cluster instance group. The additional EBS volume is
attached to each instance within the SageMaker HyperPod cluster instance group and mounted to
`/opt/sagemaker`.

_Required_: No

_Type_: [ClusterEbsVolumeConfig](aws-properties-sagemaker-cluster-clusterebsvolumeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FsxLustreConfig`

Defines the configuration for attaching an Amazon FSx for Lustre file system to the instances in the
SageMaker HyperPod cluster instance group.

_Required_: No

_Type_: [ClusterFsxLustreConfig](aws-properties-sagemaker-cluster-clusterfsxlustreconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FsxOpenZfsConfig`

Defines the configuration for attaching an Amazon FSx for OpenZFS file system to the
instances in the SageMaker HyperPod cluster instance group.

_Required_: No

_Type_: [ClusterFsxOpenZfsConfig](aws-properties-sagemaker-cluster-clusterfsxopenzfsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterInstanceGroup

ClusterKubernetesConfig

All content copied from https://docs.aws.amazon.com/.
