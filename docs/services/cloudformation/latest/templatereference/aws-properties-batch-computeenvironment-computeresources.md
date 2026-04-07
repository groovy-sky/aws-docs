This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment ComputeResources

Details about the compute resources managed by the compute environment. This parameter is
required for managed compute environments. For more information, see [Compute Environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html)
in the _AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "BidPercentage" : Integer,
  "DesiredvCpus" : Integer,
  "Ec2Configuration" : [ Ec2ConfigurationObject, ... ],
  "Ec2KeyPair" : String,
  "ImageId" : String,
  "InstanceRole" : String,
  "InstanceTypes" : [ String, ... ],
  "LaunchTemplate" : LaunchTemplateSpecification,
  "MaxvCpus" : Integer,
  "MinvCpus" : Integer,
  "PlacementGroup" : String,
  "ScalingPolicy" : ComputeScalingPolicy,
  "SecurityGroupIds" : [ String, ... ],
  "SpotIamFleetRole" : String,
  "Subnets" : [ String, ... ],
  "Tags" : {Key: Value, ...},
  "Type" : String,
  "UpdateToLatestImageVersion" : Boolean
}

```

### YAML

```yaml

  AllocationStrategy: String
  BidPercentage: Integer
  DesiredvCpus: Integer
  Ec2Configuration:
    - Ec2ConfigurationObject
  Ec2KeyPair: String
  ImageId: String
  InstanceRole: String
  InstanceTypes:
    - String
  LaunchTemplate:
    LaunchTemplateSpecification
  MaxvCpus: Integer
  MinvCpus: Integer
  PlacementGroup: String
  ScalingPolicy:
    ComputeScalingPolicy
  SecurityGroupIds:
    - String
  SpotIamFleetRole: String
  Subnets:
    - String
  Tags:
    Key: Value
  Type: String
  UpdateToLatestImageVersion: Boolean

```

## Properties

`AllocationStrategy`

The allocation strategy to use for the compute resource if not enough instances of the
best fitting instance type can be allocated. This might be because of availability of the
instance type in the Region or [Amazon EC2 service\
limits](../../../ec2/latest/userguide/ec2-resource-limits.md). For more information, see [Allocation strategies](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) in
the _AWS Batch User Guide_.

When updating a compute environment, changing the allocation strategy requires an
infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the _AWS Batch User Guide_. `BEST_FIT` is not supported
when updating a compute environment.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources, and
shouldn't be specified.

BEST\_FIT (default)

AWS Batch selects an instance type that best fits the needs of the
jobs with a preference for the lowest-cost instance type. If additional instances
of the selected instance type aren't available, AWS Batch waits for
the additional instances to be available. If there aren't enough instances
available, or if the user is reaching [Amazon EC2 service\
limits](../../../ec2/latest/userguide/ec2-resource-limits.md) then additional jobs aren't run until the currently running jobs
have completed. This allocation strategy keeps costs lower but can limit scaling.
If you are using Spot Fleets with `BEST_FIT` then the Spot Fleet IAM
role must be specified.

BEST\_FIT\_PROGRESSIVE

AWS Batch will select additional instance types that are large
enough to meet the requirements of the jobs in the queue, with a preference for
instance types with a lower cost per unit vCPU. If additional instances of the
previously selected instance types aren't available, AWS Batch will
select new instance types.

SPOT\_CAPACITY\_OPTIMIZED

AWS Batch will select one or more instance types that are large
enough to meet the requirements of the jobs in the queue, with a preference for
instance types that are less likely to be interrupted. This allocation strategy is
only available for Spot Instance compute resources.

SPOT\_PRICE\_CAPACITY\_OPTIMIZED

The price and capacity optimized allocation strategy looks at both price and
capacity to select the Spot Instance pools that are the least likely to be
interrupted and have the lowest possible price. This allocation strategy is only
available for Spot Instance compute resources.

###### Note

We recommend that you use `SPOT_PRICE_CAPACITY_OPTIMIZED` rather
than `SPOT_CAPACITY_OPTIMIZED` in most instances.

With `BEST_FIT_PROGRESSIVE`, `SPOT_CAPACITY_OPTIMIZED`, and
`SPOT_PRICE_CAPACITY_OPTIMIZED` allocation strategies using On-Demand or Spot
Instances, and the `BEST_FIT` strategy using Spot Instances, AWS Batch might need to go above `maxvCpus` to meet your capacity
requirements. In this event, AWS Batch never exceeds `maxvCpus`
by more than a single instance.

_Required_: No

_Type_: String

_Allowed values_: `BEST_FIT_PROGRESSIVE | SPOT_CAPACITY_OPTIMIZED | SPOT_PRICE_CAPACITY_OPTIMIZED`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`BidPercentage`

The maximum percentage that a Spot Instance price can be when compared with the On-Demand
price for that instance type before instances are launched. For example, if your maximum
percentage is 20%, the Spot price must be less than 20% of the current On-Demand price for that
Amazon EC2 instance. You always pay the lowest (market) price and never more than your maximum
percentage. For most use cases, we recommend leaving this field empty.

When updating a compute environment, changing the bid percentage requires an infrastructure
update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: Integer

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DesiredvCpus`

The desired number of vCPUS in the compute environment. AWS Batch modifies this value between
the minimum and maximum values based on job queue demand.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

###### Note

AWS Batch doesn't support changing the desired number of vCPUs of an existing compute
environment. Don't specify this parameter for compute environments using Amazon EKS clusters.

###### Note

When you update the `desiredvCpus` setting, the value must be between the
`minvCpus` and `maxvCpus` values.

Additionally, the updated `desiredvCpus` value must be greater than or equal to
the current `desiredvCpus` value. For more information, see [Troubleshooting\
AWS Batch](https://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html#error-desired-vcpus-update) in the _AWS Batch User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2Configuration`

Provides information used to select Amazon Machine Images (AMIs) for Amazon EC2 instances in the
compute environment. If `Ec2Configuration` isn't specified, the default is
`ECS_AL2023` for EC2 (ECS) compute environments and `EKS_AL2023` for EKS compute environments.

When updating a compute environment, changing this setting requires an infrastructure update
of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_. To remove the Amazon EC2 configuration and any custom AMI ID
specified in `imageIdOverride`, set this value to an empty string.

One or two values can be provided.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: Array of [Ec2ConfigurationObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-computeenvironment-ec2configurationobject.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Ec2KeyPair`

The Amazon EC2 key pair that's used for instances launched in the compute environment. You can
use this key pair to log in to your instances with SSH. To remove the Amazon EC2 key pair, set this
value to an empty string.

When updating a compute environment, changing the Amazon EC2 key pair requires an infrastructure
update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ImageId`

The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
This parameter is overridden by the `imageIdOverride` member of the
`Ec2Configuration` structure. To remove the custom AMI ID and use the default AMI ID,
set this value to an empty string.

When updating a compute environment, changing the AMI ID requires an infrastructure update
of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

###### Note

The AMI that you choose for a compute environment must match the architecture of the instance types that
you intend to use for that compute environment. For example, if your compute environment uses A1 instance types,
the compute resource AMI that you choose must support ARM instances. Amazon ECS vends both x86 and ARM versions of the
Amazon ECS-optimized Amazon Linux 2023 AMI. For more information, see [Amazon ECS-optimized\
Amazon Linux 2023 AMI](../../../amazonecs/latest/developerguide/ecs-optimized-ami.md#ecs-optimized-ami-linux-variants.html)
in the _Amazon Elastic Container Service Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InstanceRole`

The Amazon ECS instance profile applied to Amazon EC2 instances in a compute environment.
Required for Amazon EC2 instances. You can specify the short name or full Amazon Resource Name (ARN) of an instance
profile. For example, `
                            ecsInstanceRole
                        ` or
`arn:aws:iam::<aws_account_id>:instance-profile/ecsInstanceRole`.
For more information, see [Amazon ECS instance role](https://docs.aws.amazon.com/batch/latest/userguide/instance_IAM_role.html) in the _AWS Batch User Guide_.

When updating a compute environment, changing this setting requires an infrastructure update
of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`InstanceTypes`

The instances types that can be launched. You can specify instance families to launch any
instance type within those families (for example, `c5` or `p3`), or you can
specify specific sizes within a family (such as `c5.8xlarge`).

AWS Batch can select the instance type for you if you choose one of the following:

- `optimal` to select instance types (from the `c4`, `m4`,
`r4`, `c5`, `m5`, and `r5`
instance families) that match the demand of your job queues.

- `default_x86_64` to choose x86 based instance types (from the `m6i`,
`c6i`, `r6i`, and `c7i` instance families) that matches the resource demands of the job queue.

- `default_arm64` to choose x86 based instance types (from the `m6g`,
`c6g`, `r6g`, and `c7g` instance families) that matches the resource demands of the job queue.

###### Note

Starting on 11/01/2025 the behavior of `optimal` is going to be changed to match
`default_x86_64`. During the change your instance families could be updated to a
newer generation. You do not need to perform any actions for the upgrade to happen. For more
information about change, see [Optimal instance type configuration to receive automatic instance family\
updates](https://docs.aws.amazon.com/batch/latest/userguide/optimal-default-instance-troubleshooting.html).

###### Note

Instance family availability varies by AWS Region. For example, some AWS Regions may not have any fourth generation instance families but have fifth and
sixth generation instance families.

When using `default_x86_64` or `default_arm64` instance bundles,
AWS Batch selects instance families based on a balance of cost-effectiveness and performance.
While newer generation instances often provide better price-performance, AWS Batch may choose an
earlier generation instance family if it provides the optimal combination of availability, cost,
and performance for your workload. For example, in an AWS Region where both c6i
and c7i instances are available, AWS Batch might select c6i instances if they offer better
cost-effectiveness for your specific job requirements. For more information on AWS Batch instance
types and AWS Region availability, see [Instance\
type compute table](https://docs.aws.amazon.com/batch/latest/userguide/instance-type-compute-table.html) in the _AWS Batch User Guide_.

AWS Batch periodically updates your instances in default bundles to newer,
more cost-effective options. Updates happen automatically without requiring any
action from you. Your workloads continue running during updates with no interruption

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

###### Note

When you create a compute environment, the instance types that you select for the compute environment must
share the same architecture. For example, you can't mix x86 and ARM instances in the same compute
environment.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LaunchTemplate`

The launch template to use for your compute resources. Any other compute resource
parameters that you specify in a [CreateComputeEnvironment](https://docs.aws.amazon.com/batch/latest/APIReference/API_CreateComputeEnvironment.html) API operation override the same parameters in the
launch template. You must specify either the launch template ID or launch template name in
the request, but not both. For more information, see [Launch Template Support](https://docs.aws.amazon.com/batch/latest/userguide/launch-templates.html) in
the _AWS Batch User Guide_. Removing the launch template from a compute environment will not remove the
AMI specified in the launch template. In order to update the AMI specified in a launch
template, the `updateToLatestImageVersion` parameter must be set to
`true`.

When updating a compute environment, changing the launch template requires an
infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the _AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs running on Fargate resources, and shouldn't
be specified.

_Required_: No

_Type_: [LaunchTemplateSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-computeenvironment-launchtemplatespecification.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`MaxvCpus`

The maximum number of Amazon EC2 vCPUs that an environment can reach.

###### Note

With `BEST_FIT_PROGRESSIVE`, `SPOT_CAPACITY_OPTIMIZED` and
`SPOT_PRICE_CAPACITY_OPTIMIZED` (recommended) strategies using On-Demand or Spot
Instances, and the `BEST_FIT` strategy using Spot Instances, AWS Batch might need to
exceed `maxvCpus` to meet your capacity requirements. In this event, AWS Batch never
exceeds `maxvCpus` by more than a single instance.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinvCpus`

The minimum number of vCPUs that an environment should maintain (even if the compute environment
is `DISABLED`).

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlacementGroup`

The Amazon EC2 placement group to associate with your compute resources. If you intend to submit
multi-node parallel jobs to your compute environment, you should consider creating a cluster
placement group and associate it with your compute resources. This keeps your multi-node parallel
job on a logical grouping of instances within a single Availability Zone with high network flow
potential. For more information, see [Placement groups](../../../ec2/latest/userguide/placement-groups.md) in the
_Amazon EC2 User Guide for Linux Instances_.

When updating a compute environment, changing the placement group requires an infrastructure
update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ScalingPolicy`

The scaling policy configuration for the compute environment.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: [ComputeScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-batch-computeenvironment-computescalingpolicy.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SecurityGroupIds`

The Amazon EC2 security groups that are associated with instances launched in the compute
environment. This parameter is required for Fargate compute resources, where it can contain up
to 5 security groups. For Fargate compute resources, providing an empty list is handled as if
this parameter wasn't specified and no change is made. For Amazon EC2 compute resources, providing an
empty list removes the security groups from the compute resource.

When updating a compute environment, changing the Amazon EC2 security groups requires an
infrastructure update of the compute environment. For more information, see [Updating compute\
environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the _AWS Batch User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`SpotIamFleetRole`

The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a `SPOT` compute
environment. This role is required if the allocation strategy set to `BEST_FIT` or if
the allocation strategy isn't specified. For more information, see [Amazon EC2 spot fleet role](https://docs.aws.amazon.com/batch/latest/userguide/spot_fleet_IAM_role.html) in the
_AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

###### Important

To tag your Spot Instances on creation, the Spot Fleet IAM role specified here must use
the newer **AmazonEC2SpotFleetTaggingRole** managed policy. The
previously recommended **AmazonEC2SpotFleetRole** managed policy
doesn't have the required permissions to tag Spot Instances. For more information, see [Spot instances\
not tagged on creation](https://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html#spot-instance-no-tag) in the _AWS Batch User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subnets`

The VPC subnets where the compute resources are launched. Fargate compute resources can
contain up to 16 subnets. For Fargate compute resources, providing an empty list will be
handled as if this parameter wasn't specified and no change is made. For Amazon EC2 compute resources,
providing an empty list removes the VPC subnets from the compute resource. For more information,
see [VPCs and\
subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the _Amazon VPC User Guide_.

When updating a compute environment, changing the VPC subnets requires an infrastructure
update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

###### Note

AWS Batch on Amazon EC2 and AWS Batch on Amazon EKS support Local Zones. For more information, see [Local\
Zones](../../../ec2/latest/userguide/using-regions-availability-zones.md#concepts-local-zones) in the _Amazon EC2 User Guide for Linux Instances_, [Amazon EKS and AWS Local\
Zones](https://docs.aws.amazon.com/eks/latest/userguide/local-zones.html) in the _Amazon EKS User Guide_ and [Amazon ECS\
clusters in Local Zones, Wavelength Zones, and AWS Outposts](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-regions-zones.html#clusters-local-zones) in the _Amazon ECS_
_Developer Guide_.

AWS Batch on Fargate doesn't currently support Local Zones.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

Key-value pair tags to be applied to Amazon EC2 resources that are launched in the compute
environment. For AWS Batch, these take the form of `"String1": "String2"`, where
`String1` is the tag key and `String2` is the tag value (for example,
`{ "Name": "Batch Instance - C4OnDemand" }`). This is helpful for recognizing your
Batch instances in the Amazon EC2 console. These tags aren't seen when using the AWS Batch `ListTagsForResource` API operation.

When updating a compute environment, changing this setting requires an infrastructure update
of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources. Don't specify it.

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Type`

The type of compute environment: `EC2`, `SPOT`,
`FARGATE`, or `FARGATE_SPOT`. For more information, see [Compute\
environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the _AWS Batch User Guide_.

If you choose `SPOT`, you must also specify an Amazon EC2 Spot Fleet role
with the `spotIamFleetRole` parameter. For more information, see [Amazon EC2\
spot fleet role](https://docs.aws.amazon.com/batch/latest/userguide/spot_fleet_IAM_role.html) in the _AWS Batch User Guide_.

When updating compute environment, changing the type of a compute environment requires
an infrastructure update of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the _AWS Batch User Guide_.

When updating the type of a compute environment, changing between `EC2` and
`SPOT` or between `FARGATE` and `FARGATE_SPOT` will
initiate an infrastructure update, but if you switch between `EC2` and
`FARGATE`, CloudFormation will create a new compute environment.

_Required_: Yes

_Type_: String

_Allowed values_: `EC2 | SPOT | FARGATE | FARGATE_SPOT`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UpdateToLatestImageVersion`

Specifies whether the AMI ID is updated to the latest one that's supported by AWS Batch when
the compute environment has an infrastructure update. The default value is
`false`.

###### Note

An AMI ID can either be specified in the `imageId` or
`imageIdOverride` parameters or be determined by the launch template that's
specified in the `launchTemplate` parameter. If an AMI ID is specified any of these
ways, this parameter is ignored. For more information about to update AMI IDs during an
infrastructure update, see [Updating\
the AMI ID](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html#updating-compute-environments-ami) in the _AWS Batch User Guide_.

When updating a compute environment, changing this setting requires an infrastructure update
of the compute environment. For more information, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Compute Environments](https://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) in the _AWS Batch User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Batch::ComputeEnvironment

ComputeScalingPolicy
