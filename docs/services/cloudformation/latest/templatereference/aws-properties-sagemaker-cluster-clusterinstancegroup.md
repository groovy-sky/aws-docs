This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterInstanceGroup

The configuration information of the instance group within the HyperPod cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityRequirements" : ClusterCapacityRequirements,
  "CurrentCount" : Integer,
  "ExecutionRole" : String,
  "ImageId" : String,
  "InstanceCount" : Integer,
  "InstanceGroupName" : String,
  "InstanceStorageConfigs" : [ ClusterInstanceStorageConfig, ... ],
  "InstanceType" : String,
  "KubernetesConfig" : ClusterKubernetesConfig,
  "LifeCycleConfig" : ClusterLifeCycleConfig,
  "MinInstanceCount" : Integer,
  "OnStartDeepHealthChecks" : [ String, ... ],
  "OverrideVpcConfig" : VpcConfig,
  "ScheduledUpdateConfig" : ScheduledUpdateConfig,
  "SlurmConfig" : ClusterSlurmConfig,
  "ThreadsPerCore" : Integer,
  "TrainingPlanArn" : String
}

```

### YAML

```yaml

  CapacityRequirements:
    ClusterCapacityRequirements
  CurrentCount: Integer
  ExecutionRole: String
  ImageId: String
  InstanceCount: Integer
  InstanceGroupName: String
  InstanceStorageConfigs:
    - ClusterInstanceStorageConfig
  InstanceType: String
  KubernetesConfig:
    ClusterKubernetesConfig
  LifeCycleConfig:
    ClusterLifeCycleConfig
  MinInstanceCount: Integer
  OnStartDeepHealthChecks:
    - String
  OverrideVpcConfig:
    VpcConfig
  ScheduledUpdateConfig:
    ScheduledUpdateConfig
  SlurmConfig:
    ClusterSlurmConfig
  ThreadsPerCore: Integer
  TrainingPlanArn: String

```

## Properties

`CapacityRequirements`

The capacity requirements for the instance group, specifying on-demand and spot instance configurations.

_Required_: No

_Type_: [ClusterCapacityRequirements](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-clustercapacityrequirements.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CurrentCount`

The number of instances that are currently in the instance group of a SageMaker HyperPod
cluster.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

The execution role for the instance group to assume.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageId`

The ID of the Amazon Machine Image (AMI) to use for the instances in the group.

_Required_: No

_Type_: String

_Pattern_: `^ami-[0-9a-fA-F]{8,17}|default$`

_Minimum_: `7`

_Maximum_: `21`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceCount`

The number of instances in an instance group of the SageMaker HyperPod cluster.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceGroupName`

The name of the instance group of a SageMaker HyperPod cluster.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceStorageConfigs`

The configurations of additional storage specified to the instance group where the
instance (node) is launched.

_Required_: No

_Type_: Array of [ClusterInstanceStorageConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-clusterinstancestorageconfig.html)

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The instance type of the instance group of a SageMaker HyperPod cluster.

_Required_: Yes

_Type_: String

_Allowed values_: `ml.p4d.24xlarge | ml.p4de.24xlarge | ml.p5.48xlarge | ml.p5.4xlarge | ml.p6e-gb200.36xlarge | ml.trn1.32xlarge | ml.trn1n.32xlarge | ml.g5.xlarge | ml.g5.2xlarge | ml.g5.4xlarge | ml.g5.8xlarge | ml.g5.12xlarge | ml.g5.16xlarge | ml.g5.24xlarge | ml.g5.48xlarge | ml.c5.large | ml.c5.xlarge | ml.c5.2xlarge | ml.c5.4xlarge | ml.c5.9xlarge | ml.c5.12xlarge | ml.c5.18xlarge | ml.c5.24xlarge | ml.c5n.large | ml.c5n.2xlarge | ml.c5n.4xlarge | ml.c5n.9xlarge | ml.c5n.18xlarge | ml.m5.large | ml.m5.xlarge | ml.m5.2xlarge | ml.m5.4xlarge | ml.m5.8xlarge | ml.m5.12xlarge | ml.m5.16xlarge | ml.m5.24xlarge | ml.t3.medium | ml.t3.large | ml.t3.xlarge | ml.t3.2xlarge | ml.g6.xlarge | ml.g6.2xlarge | ml.g6.4xlarge | ml.g6.8xlarge | ml.g6.16xlarge | ml.g6.12xlarge | ml.g6.24xlarge | ml.g6.48xlarge | ml.gr6.4xlarge | ml.gr6.8xlarge | ml.g6e.xlarge | ml.g6e.2xlarge | ml.g6e.4xlarge | ml.g6e.8xlarge | ml.g6e.16xlarge | ml.g6e.12xlarge | ml.g6e.24xlarge | ml.g6e.48xlarge | ml.p5e.48xlarge | ml.p5en.48xlarge | ml.p6-b200.48xlarge | ml.trn2.3xlarge | ml.trn2.48xlarge | ml.c6i.large | ml.c6i.xlarge | ml.c6i.2xlarge | ml.c6i.4xlarge | ml.c6i.8xlarge | ml.c6i.12xlarge | ml.c6i.16xlarge | ml.c6i.24xlarge | ml.c6i.32xlarge | ml.m6i.large | ml.m6i.xlarge | ml.m6i.2xlarge | ml.m6i.4xlarge | ml.m6i.8xlarge | ml.m6i.12xlarge | ml.m6i.16xlarge | ml.m6i.24xlarge | ml.m6i.32xlarge | ml.r6i.large | ml.r6i.xlarge | ml.r6i.2xlarge | ml.r6i.4xlarge | ml.r6i.8xlarge | ml.r6i.12xlarge | ml.r6i.16xlarge | ml.r6i.24xlarge | ml.r6i.32xlarge | ml.i3en.large | ml.i3en.xlarge | ml.i3en.2xlarge | ml.i3en.3xlarge | ml.i3en.6xlarge | ml.i3en.12xlarge | ml.i3en.24xlarge | ml.m7i.large | ml.m7i.xlarge | ml.m7i.2xlarge | ml.m7i.4xlarge | ml.m7i.8xlarge | ml.m7i.12xlarge | ml.m7i.16xlarge | ml.m7i.24xlarge | ml.m7i.48xlarge | ml.r7i.large | ml.r7i.xlarge | ml.r7i.2xlarge | ml.r7i.4xlarge | ml.r7i.8xlarge | ml.r7i.12xlarge | ml.r7i.16xlarge | ml.r7i.24xlarge | ml.r7i.48xlarge | ml.p6-b300.48xlarge`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KubernetesConfig`

The Kubernetes configuration for the instance group, including labels and taints.

_Required_: No

_Type_: [ClusterKubernetesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-clusterkubernetesconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifeCycleConfig`

The lifecycle configuration for a SageMaker HyperPod cluster.

_Required_: Yes

_Type_: [ClusterLifeCycleConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-clusterlifecycleconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinInstanceCount`

The minimum number of instances to maintain in the instance group.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnStartDeepHealthChecks`

A flag indicating whether deep health checks should be performed when the HyperPod cluster instance group is
created or updated. Deep health checks are comprehensive, invasive tests that validate the health of the underlying
hardware and infrastructure components.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideVpcConfig`

The customized Amazon VPC configuration at the instance group level that overrides the default Amazon VPC configuration of the SageMaker HyperPod cluster.

_Required_: No

_Type_: [VpcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-vpcconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScheduledUpdateConfig`

Configuration for scheduled updates to the instance group.

_Required_: No

_Type_: [ScheduledUpdateConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-scheduledupdateconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlurmConfig`

The Slurm workload manager configuration for the instance group.

_Required_: No

_Type_: [ClusterSlurmConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-clusterslurmconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreadsPerCore`

The number of threads per CPU core you specified under
`CreateCluster`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrainingPlanArn`

The Amazon Resource Name (ARN) of the training plan associated with the instance group.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:training-plan/.*$`

_Minimum_: `50`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClusterFsxOpenZfsConfig

ClusterInstanceStorageConfig
