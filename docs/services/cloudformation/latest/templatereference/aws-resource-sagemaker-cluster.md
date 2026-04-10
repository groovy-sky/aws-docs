This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster

Creates an Amazon SageMaker HyperPod cluster. SageMaker HyperPod is a capability of SageMaker for creating and
managing persistent clusters for developing large machine learning models, such as large
language models (LLMs) and diffusion models. To learn more, see [Amazon SageMaker HyperPod](../../../sagemaker/latest/dg/sagemaker-hyperpod.md) in the
_Amazon SageMaker Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Cluster",
  "Properties" : {
      "AutoScaling" : ClusterAutoScalingConfig,
      "ClusterName" : String,
      "ClusterRole" : String,
      "InstanceGroups" : [ ClusterInstanceGroup, ... ],
      "NodeProvisioningMode" : String,
      "NodeRecovery" : String,
      "Orchestrator" : Orchestrator,
      "RestrictedInstanceGroups" : [ ClusterRestrictedInstanceGroup, ... ],
      "Tags" : [ Tag, ... ],
      "TieredStorageConfig" : TieredStorageConfig,
      "VpcConfig" : VpcConfig
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Cluster
Properties:
  AutoScaling:
    ClusterAutoScalingConfig
  ClusterName: String
  ClusterRole: String
  InstanceGroups:
    - ClusterInstanceGroup
  NodeProvisioningMode: String
  NodeRecovery: String
  Orchestrator:
    Orchestrator
  RestrictedInstanceGroups:
    - ClusterRestrictedInstanceGroup
  Tags:
    - Tag
  TieredStorageConfig:
    TieredStorageConfig
  VpcConfig:
    VpcConfig

```

## Properties

`AutoScaling`

The autoscaling configuration for the cluster. Enables automatic scaling of cluster
nodes based on workload demand using a Karpenter-based system.

_Required_: No

_Type_: [ClusterAutoScalingConfig](aws-properties-sagemaker-cluster-clusterautoscalingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterName`

The name of the SageMaker HyperPod cluster.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterRole`

The Amazon Resource Name (ARN) of the IAM role that HyperPod assumes to perform cluster
autoscaling operations. This role must have permissions for
`sagemaker:BatchAddClusterNodes` and
`sagemaker:BatchDeleteClusterNodes`. This is only required when autoscaling
is enabled and when HyperPod is performing autoscaling operations.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceGroups`

The instance groups of the SageMaker HyperPod cluster. To delete an instance group, remove it from the
array.

_Required_: No

_Type_: Array of [ClusterInstanceGroup](aws-properties-sagemaker-cluster-clusterinstancegroup.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeProvisioningMode`

The mode for provisioning nodes in the cluster. You can specify the following
modes:

- **Continuous**: Scaling behavior that enables 1)
concurrent operation execution within instance groups, 2) continuous retry mechanisms
for failed operations, 3) enhanced customer visibility into cluster events through
detailed event streams, 4) partial provisioning capabilities. Your clusters and
instance groups remain `InService` while scaling. This mode is only
supported for EKS orchestrated clusters.

_Required_: No

_Type_: String

_Allowed values_: `Continuous`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeRecovery`

Specifies whether to enable or disable the automatic node recovery feature of SageMaker HyperPod. Available
values are `Automatic` for enabling and `None` for disabling.

_Required_: No

_Type_: String

_Allowed values_: `Automatic | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Orchestrator`

The orchestrator type for the SageMaker HyperPod cluster. Currently, `'eks'` is the only available
option.

_Required_: No

_Type_: [Orchestrator](aws-properties-sagemaker-cluster-orchestrator.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestrictedInstanceGroups`

The specialized instance groups for training models like Amazon Nova to be created in the
SageMaker HyperPod cluster.

_Required_: No

_Type_: Array of [ClusterRestrictedInstanceGroup](aws-properties-sagemaker-cluster-clusterrestrictedinstancegroup.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A tag object that consists of a key and an optional value, used to manage metadata
for SageMaker AWS resources.

You can add tags to notebook instances, training jobs, hyperparameter tuning jobs,
batch transform jobs, models, labeling jobs, work teams, endpoint configurations, and
endpoints. For more information on adding tags to SageMaker resources, see [AddTags](../../../../reference/sagemaker/latest/apireference/api-addtags.md).

For more information on adding metadata to your AWS resources with
tagging, see [Tagging AWS resources](../../../../general/latest/gr/aws-tagging.md). For advice on best practices for
managing AWS resources with tagging, see [Tagging\
Best Practices: Implement an Effective AWS Resource Tagging\
Strategy](https://d1.awsstatic.com/whitepapers/aws-tagging-best-practices.pdf).

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-cluster-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TieredStorageConfig`

The configuration for managed tier checkpointing on the HyperPod cluster. When
enabled, this feature uses a multi-tier storage approach for storing model checkpoints,
providing faster checkpoint operations and improved fault tolerance across cluster
nodes.

_Required_: No

_Type_: [TieredStorageConfig](aws-properties-sagemaker-cluster-tieredstorageconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources
have access to. You can control access to and from your resources by configuring a VPC.
For more information, see [Give SageMaker Access to\
Resources in your Amazon VPC](../../../sagemaker/latest/dg/infrastructure-give-access.md).

_Required_: No

_Type_: [VpcConfig](aws-properties-sagemaker-cluster-vpcconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ClusterArn`

The Amazon Resource Name (ARN) of the SageMaker HyperPod cluster.

`ClusterStatus`

The status of the SageMaker HyperPod cluster.

`CreationTime`

The time when the SageMaker HyperPod cluster is created.

`FailureMessage`

The failure message of the SageMaker HyperPod cluster.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AlarmDetails

All content copied from https://docs.aws.amazon.com/.
