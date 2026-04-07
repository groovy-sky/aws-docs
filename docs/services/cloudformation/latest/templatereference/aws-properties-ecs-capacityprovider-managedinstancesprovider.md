This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::CapacityProvider ManagedInstancesProvider

The configuration for a Amazon ECS Managed Instances provider. Amazon ECS uses this
configuration to automatically launch, manage, and terminate Amazon EC2 instances on
your behalf. Managed instances provide access to the full range of Amazon EC2 instance
types and features while offloading infrastructure management to AWS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InfrastructureOptimization" : InfrastructureOptimization,
  "InfrastructureRoleArn" : String,
  "InstanceLaunchTemplate" : InstanceLaunchTemplate,
  "PropagateTags" : String
}

```

### YAML

```yaml

  InfrastructureOptimization:
    InfrastructureOptimization
  InfrastructureRoleArn: String
  InstanceLaunchTemplate:
    InstanceLaunchTemplate
  PropagateTags: String

```

## Properties

`InfrastructureOptimization`

Defines how Amazon ECS Managed Instances optimizes the infrastastructure in your capacity provider.
Configure it to turn on or off the infrastructure optimization in your capacity provider, and to control the idle or underutilized EC2 instances optimization delay.

_Required_: No

_Type_: [InfrastructureOptimization](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-capacityprovider-infrastructureoptimization.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InfrastructureRoleArn`

The Amazon Resource Name (ARN) of the infrastructure role that Amazon ECS assumes to
manage instances. This role must include permissions for Amazon EC2 instance lifecycle
management, networking, and any additional AWS services required for your
workloads.

For more information, see [Amazon ECS\
infrastructure IAM role](../../../amazonecs/latest/developerguide/infrastructure-iam-role.md) in the _Amazon ECS Developer_
_Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceLaunchTemplate`

The launch template that defines how Amazon ECS launches Amazon ECS Managed Instances.
This includes the instance profile for your tasks, network and storage configuration,
and instance requirements that determine which Amazon EC2 instance types can be
used.

For more information, see [Store instance launch\
parameters in Amazon EC2 launch templates](../../../ec2/latest/userguide/ec2-launch-templates.md) in the _Amazon EC2 User_
_Guide_.

_Required_: Yes

_Type_: [InstanceLaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-capacityprovider-instancelaunchtemplate.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropagateTags`

Determines whether tags from the capacity provider are automatically applied to Amazon
ECS Managed Instances. This helps with cost allocation and resource management by
ensuring consistent tagging across your infrastructure.

_Required_: No

_Type_: String

_Allowed values_: `CAPACITY_PROVIDER | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ManagedInstancesNetworkConfiguration

ManagedInstancesStorageConfiguration
