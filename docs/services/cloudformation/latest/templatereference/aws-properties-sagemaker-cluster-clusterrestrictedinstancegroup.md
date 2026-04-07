This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterRestrictedInstanceGroup

Specialized instance groups for training models like Amazon Nova in the SageMaker HyperPod cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CurrentCount" : Integer,
  "EnvironmentConfig" : EnvironmentConfig,
  "ExecutionRole" : String,
  "InstanceCount" : Integer,
  "InstanceGroupName" : String,
  "InstanceStorageConfigs" : [ ClusterInstanceStorageConfig, ... ],
  "InstanceType" : String,
  "OnStartDeepHealthChecks" : [ String, ... ],
  "OverrideVpcConfig" : VpcConfig,
  "ThreadsPerCore" : Integer,
  "TrainingPlanArn" : String
}

```

### YAML

```yaml

  CurrentCount: Integer
  EnvironmentConfig:
    EnvironmentConfig
  ExecutionRole: String
  InstanceCount: Integer
  InstanceGroupName: String
  InstanceStorageConfigs:
    - ClusterInstanceStorageConfig
  InstanceType: String
  OnStartDeepHealthChecks:
    - String
  OverrideVpcConfig:
    VpcConfig
  ThreadsPerCore: Integer
  TrainingPlanArn: String

```

## Properties

`CurrentCount`

The current number of instances in the restricted instance group.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentConfig`

Environment configuration for the restricted instance group, including FSx Lustre settings.

_Required_: Yes

_Type_: [EnvironmentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-environmentconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

The Amazon Resource Name (ARN) of the IAM execution role for the restricted instance group.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceCount`

The number of instances in the restricted instance group.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceGroupName`

The name of the restricted instance group.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceStorageConfigs`

Storage configurations for instances in the restricted instance group.

_Required_: No

_Type_: Array of [ClusterInstanceStorageConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-clusterinstancestorageconfig.html)

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceType`

The EC2 instance type for the restricted instance group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnStartDeepHealthChecks`

Deep health checks to run when instances start in the restricted instance group.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverrideVpcConfig`

VPC configuration override for the restricted instance group.

_Required_: No

_Type_: [VpcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-cluster-vpcconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThreadsPerCore`

The number of threads per CPU core for instances in the restricted instance group.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrainingPlanArn`

The Amazon Resource Name (ARN) of the training plan for the restricted instance group.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:training-plan/.*$`

_Minimum_: `50`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClusterOrchestratorSlurmConfig

ClusterSlurmConfig
