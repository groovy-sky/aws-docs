This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment

The `AWS::Batch::ComputeEnvironment` resource defines your AWS Batch compute environment. You can define `MANAGED` or
`UNMANAGED` compute environments. `MANAGED` compute environments
can use Amazon EC2 or AWS Fargate resources. `UNMANAGED`
compute environments can only use EC2 resources. For more information, see [Compute\
Environments](../../../batch/latest/userguide/compute-environments.md) in the _AWS Batch User Guide_.

In a managed compute environment, AWS Batch manages the capacity and
instance types of the compute resources within the environment. This is based on the
compute resource specification that you define or the [launch template](../../../ec2/latest/userguide/ec2-launch-templates.md) that you
specify when you create the compute environment. You can choose either to use EC2 On-Demand
Instances and EC2 Spot Instances, or to use Fargate and Fargate Spot capacity in your
managed compute environment. You can optionally set a maximum price so that Spot Instances
only launch when the Spot Instance price is below a specified percentage of the On-Demand
price.

###### Note

Multi-node parallel jobs are not supported on Spot Instances.

In an unmanaged compute environment, you can manage your own EC2 compute resources and
have a lot of flexibility with how you configure your compute resources. For example, you
can use custom AMI. However, you need to verify that your AMI meets the Amazon ECS
container instance AMI specification. For more information, see [container instance\
AMIs](../../../amazonecs/latest/developerguide/container-instance-amis.md) in the _Amazon Elastic Container Service Developer_
_Guide_. After you have created your unmanaged compute environment, you can use
the [DescribeComputeEnvironments](../../../../reference/batch/latest/apireference/api-describecomputeenvironments.md) operation to find the Amazon ECS cluster that is
associated with it. Then, manually launch your container instances into that Amazon ECS
cluster. For more information, see [Launching an Amazon\
ECS container instance](../../../amazonecs/latest/developerguide/launch-container-instance.md) in the _Amazon Elastic Container Service_
_Developer Guide_.

###### Note

To create a compute environment that uses EKS resources, the caller must have
permissions to call `eks:DescribeCluster`.

###### Note

AWS Batch doesn't upgrade the AMIs in a compute environment after it's
created except under specific conditions. For example, it doesn't automatically update
the AMIs when a newer version of the Amazon ECS optimized AMI is available. Therefore,
you're responsible for the management of the guest operating system (including updates
and security patches) and any additional application software or utilities that you
install on the compute resources. There are two ways to use a new AMI for your AWS Batch jobs. The original method is to complete these steps:

1. Create a new compute environment with the new AMI.

2. Add the compute environment to an existing job queue.

3. Remove the earlier compute environment from your job queue.

4. Delete the earlier compute environment.

In April 2022, AWS Batch added enhanced support for updating compute
environments. For example, the `UpdateComputeEnvironent` API lets you use the
`ReplaceComputeEnvironment` property to dynamically update compute
environment parameters such as the launch template or instance type without replacement.
For more information, see [Updating compute\
environments](../../../batch/latest/userguide/updating-compute-environments.md) in the _AWS Batch User Guide_.

To use the enhanced updating of compute environments to update AMIs, follow these
rules:

- Either do not set the [ServiceRole](../userguide/aws-resource-batch-computeenvironment.md#cfn-batch-computeenvironment-servicerole) property or set it to the **AWSServiceRoleForBatch** service-linked role.

- Set the [AllocationStrategy](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-allocationstrategy) property to `BEST_FIT_PROGRESSIVE` or
`SPOT_CAPACITY_OPTIMIZED`.

- Set the [ReplaceComputeEnvironment](../userguide/aws-resource-batch-computeenvironment.md#cfn-batch-computeenvironment-replacecomputeenvironment) property to `false`.

###### Important

Set the `ReplaceComputeEnvironment` property to `true`
if the compute environment uses the `BEST_FIT` allocation
strategy.

###### Note

If the `ReplaceComputeEnvironment` property is set to
`false`, you might receive an error message when you update the
CFN template for a compute environment. This issue occurs if the updated
`desiredvcpus` value is less than the current
`desiredvcpus` value. As a workaround, delete the
`desiredvcpus` value from the updated template or use the
`minvcpus` property to manage the number of vCPUs. For
information, see [Error message when you update the `DesiredvCpus`\
setting](../../../batch/latest/userguide/troubleshooting.md#error-desired-vcpus-update).

- Set the [UpdateToLatestImageVersion](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion) property to `true`. This
property is used when you update a compute environment. The [UpdateToLatestImageVersion](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion) property is ignored when you create a
compute environment.

- Either do not specify an image ID in [ImageId](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-imageid) or [ImageIdOverride](../userguide/aws-properties-batch-computeenvironment-ec2configurationobject.md#cfn-batch-computeenvironment-ec2configurationobject-imageidoverride) properties, or in the launch template identified by
the [Launch Template](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-launchtemplate) property. In that case AWS Batch will
select the latest Amazon ECS optimized AMI supported by AWS Batch at
the time the infrastructure update is initiated. Alternatively you can specify the
AMI ID in the `ImageId` or `ImageIdOverride` properties, or
the launch template identified by the `LaunchTemplate` properties.
Changing any of these properties will trigger an infrastructure update.

If these rules are followed, any update that triggers an infrastructure update will
cause the AMI ID to be re-selected. If the [Version](../userguide/aws-properties-batch-computeenvironment-launchtemplatespecification.md#cfn-batch-computeenvironment-launchtemplatespecification-version) property of the [LaunchTemplateSpecification](../userguide/aws-properties-batch-computeenvironment-launchtemplatespecification.md) is set to `$Latest` or
`$Default`, the latest or default version of the launch template will be
evaluated up at the time of the infrastructure update, even if the
`LaunchTemplateSpecification` was not updated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Batch::ComputeEnvironment",
  "Properties" : {
      "ComputeEnvironmentName" : String,
      "ComputeResources" : ComputeResources,
      "Context" : String,
      "EksConfiguration" : EksConfiguration,
      "ReplaceComputeEnvironment" : Boolean,
      "ServiceRole" : String,
      "State" : String,
      "Tags" : {Key: Value, ...},
      "Type" : String,
      "UnmanagedvCpus" : Integer,
      "UpdatePolicy" : UpdatePolicy
    }
}

```

### YAML

```yaml

Type: AWS::Batch::ComputeEnvironment
Properties:
  ComputeEnvironmentName: String
  ComputeResources:
    ComputeResources
  Context: String
  EksConfiguration:
    EksConfiguration
  ReplaceComputeEnvironment: Boolean
  ServiceRole: String
  State: String
  Tags:
    Key: Value
  Type: String
  UnmanagedvCpus: Integer
  UpdatePolicy:
    UpdatePolicy

```

## Properties

`ComputeEnvironmentName`

The name for your compute environment. It can be up to 128 characters long. It can contain uppercase and
lowercase letters, numbers, hyphens (-), and underscores (\_).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComputeResources`

The ComputeResources property type specifies details of the compute resources managed by
the compute environment. This parameter is required for managed compute environments. For
more information, see [Compute Environments](../../../batch/latest/userguide/compute-environments.md) in
the _AWS Batch User Guide_.

_Required_: No

_Type_: [ComputeResources](aws-properties-batch-computeenvironment-computeresources.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Context`

Reserved.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EksConfiguration`

The details for the Amazon EKS cluster that supports the compute environment.

###### Note

To create a compute environment that uses EKS resources, the caller must have
permissions to call `eks:DescribeCluster`.

_Required_: No

_Type_: [EksConfiguration](aws-properties-batch-computeenvironment-eksconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReplaceComputeEnvironment`

Specifies whether the compute environment is replaced if an update is made that requires
replacing the instances in the compute environment. The default value is `true`.
To enable more properties to be updated, set this property to `false`. When
changing the value of this property to `false`, do not change any other
properties at the same time. If other properties are changed at the same time, and the
change needs to be rolled back but it can't, it's possible for the stack to go into the
`UPDATE_ROLLBACK_FAILED` state. You can't update a stack that is in the
`UPDATE_ROLLBACK_FAILED` state. However, if you can continue to roll it back,
you can return the stack to its original settings and then try to update it again. For more
information, see [Continue rolling back an update](../userguide/using-cfn-updating-stacks-continueupdaterollback.md) in the _AWS CloudFormation User Guide_.

`ReplaceComputeEnvironment` is not applicable for Fargate compute environments. Fargate compute environments are always updated without interruption.

The properties that can't be changed without replacing the compute environment are in
the [`ComputeResources`](../userguide/aws-properties-batch-computeenvironment-computeresources.md) property type: [`AllocationStrategy`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-allocationstrategy), [`BidPercentage`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-bidpercentage), [`Ec2Configuration`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-ec2configuration), [`Ec2KeyPair`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-ec2keypair), [`Ec2KeyPair`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-ec2keypair), [`ImageId`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-imageid), [`InstanceRole`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-instancerole), [`InstanceTypes`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-instancetypes), [`LaunchTemplate`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-launchtemplate), [`MaxvCpus`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-maxvcpus), [`MinvCpus`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-minvcpus), [`PlacementGroup`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-placementgroup), [`SecurityGroupIds`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-securitygroupids), [`Subnets`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-subnets), [Tags](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-tags), [`Type`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-type), and [`UpdateToLatestImageVersion`](../userguide/aws-properties-batch-computeenvironment-computeresources.md#cfn-batch-computeenvironment-computeresources-updatetolatestimageversion).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceRole`

The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS
services on your behalf. For more information, see [AWS Batch service IAM role](../../../batch/latest/userguide/service-iam-role.md) in the
_AWS Batch User Guide_.

###### Important

If your account already created the AWS Batch service-linked role, that role is used by
default for your compute environment unless you specify a different role here. If the
AWS Batch service-linked role doesn't exist in your account, and no role is specified here,
the service attempts to create the AWS Batch service-linked role in your account.

If your specified role has a path other than `/`, then you must specify either
the full role ARN (recommended) or prefix the role name with the path. For example, if a
role with the name `bar` has a path of `/foo/`, specify
`/foo/bar` as the role name. For more information, see [Friendly\
names and paths](../../../iam/latest/userguide/reference-identifiers.md#identifiers-friendly-names) in the _IAM User Guide_.

###### Note

Depending on how you created your AWS Batch service role, its ARN might contain the
`service-role` path prefix. When you only specify the name of the service role,
AWS Batch assumes that your ARN doesn't use the `service-role` path prefix. Because
of this, we recommend that you specify the full ARN of your service role when you create
compute environments.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the compute environment. If the state is `ENABLED`, then the
compute environment accepts jobs from a queue and can scale out automatically based on
queues.

If the state is `ENABLED`, then the AWS Batch scheduler can attempt to place jobs
from an associated job queue on the compute resources within the environment. If the compute
environment is managed, then it can scale its instances out or in automatically, based on the
job queue demand.

If the state is `DISABLED`, then the AWS Batch scheduler doesn't attempt to place
jobs within the environment. Jobs in a `STARTING` or `RUNNING` state
continue to progress normally. Managed compute environments in the `DISABLED` state
don't scale out.

###### Note

Compute environments in a `DISABLED` state may continue to incur billing
charges. To prevent additional charges, turn off and then delete the compute environment.
For more information, see [State](../../../batch/latest/userguide/compute-environment-parameters.md#compute_environment_state) in the _AWS Batch User Guide_.

When an instance is idle, the instance scales down to the `minvCpus` value.
However, the instance size doesn't change. For example, consider a `c5.8xlarge`
instance with a `minvCpus` value of `4` and a `desiredvCpus`
value of `36`. This instance doesn't scale down to a `c5.large`
instance.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags applied to the compute environment.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of the compute environment: `MANAGED` or `UNMANAGED`. For
more information, see [Compute Environments](../../../batch/latest/userguide/compute-environments.md) in the _AWS Batch User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `MANAGED | UNMANAGED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UnmanagedvCpus`

The maximum number of vCPUs for an unmanaged compute environment. This parameter is only
used for fair-share scheduling to reserve vCPU capacity for new share identifiers. If this
parameter isn't provided for a fair-share job queue, no vCPU capacity is reserved.

###### Note

This parameter is only supported when the `type` parameter is set to
`UNMANAGED`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdatePolicy`

Specifies the infrastructure update policy for the compute environment. For more information
about infrastructure updates, see [Updating compute environments](../../../batch/latest/userguide/updating-compute-environments.md) in the
_AWS Batch User Guide_.

_Required_: No

_Type_: [UpdatePolicy](aws-properties-batch-computeenvironment-updatepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the compute environment ARN, such as `arn:aws:batch:us-east-1:555555555555:compute-environment/M4OnDemand`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ComputeEnvironmentArn`

Returns the compute environment ARN, such as `arn:aws:batch:us-east-1:111122223333:compute-environment/ComputeEnvironmentName`.

## Examples

### Managed Compute Environment

The following example creates a managed compute environment called
`C4OnDemand` that uses C4 On-Demand instances and a custom AMI.

#### JSON

```json

{
  "ComputeEnvironment": {
    "Type": "AWS::Batch::ComputeEnvironment",
    "Properties": {
      "Type": "MANAGED",
      "arn:aws:iam::111122223333:role/aws-service-role/batch.amazonaws.com/AWSServiceRoleForBatch",
      "ComputeEnvironmentName": "C4OnDemand",
      "ComputeResources": {
        "MaxvCpus": 128,
        "SecurityGroupIds": [
          "sg-abcd1234"
        ],
        "Type": "EC2",
        "Subnets": [
          "subnet-aaaaaaaa",
          "subnet-bbbbbbbb",
          "subnet-cccccccc"
        ],
        "MinvCpus": 0,
        "ImageId": "ami-a1b2c3d4",
        "InstanceRole": "ecsInstanceRole",
        "InstanceTypes": [
          "c4.large",
          "c4.xlarge",
          "c4.2xlarge",
          "c4.4xlarge",
          "c4.8xlarge"
        ],
        "Ec2KeyPair": "id_rsa",
        "Tags": {
          "Name": "Batch Instance - C4OnDemand"
        },
        "DesiredvCpus": 48
      },
      "State": "ENABLED"
    }
  }
}
```

#### YAML

```yaml

ComputeEnvironment:
  Type: 'AWS::Batch::ComputeEnvironment'
  Properties:
    Type: MANAGED
    ServiceRole:
     - arn:aws:iam::111122223333:role/aws-service-role/batch.amazonaws.com/AWSServiceRoleForBatch
    ComputeEnvironmentName: C4OnDemand
    ComputeResources:
      MaxvCpus: 128
      SecurityGroupIds:
        - sg-abcd1234
      Type: EC2
      Subnets:
        - subnet-aaaaaaaa
        - subnet-bbbbbbbb
        - subnet-cccccccc
      MinvCpus: 0
      ImageId: ami-a1b2c3d4
      InstanceRole: ecsInstanceRole
      InstanceTypes:
        - c4.large
        - c4.xlarge
        - c4.2xlarge
        - c4.4xlarge
        - c4.8xlarge
      Ec2KeyPair: id_rsa
      Tags:
        Name: Batch Instance - C4OnDemand
      DesiredvCpus: 48
    State: ENABLED

```

## See also

- [Compute Environments](../../../batch/latest/userguide/compute-environments.md) in the _AWS Batch User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Batch

ComputeResources

All content copied from https://docs.aws.amazon.com/.
