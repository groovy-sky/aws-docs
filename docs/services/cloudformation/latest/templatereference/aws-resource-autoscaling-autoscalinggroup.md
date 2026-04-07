This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup

The `AWS::AutoScaling::AutoScalingGroup` resource defines an Amazon EC2 Auto
Scaling group, which is a collection of Amazon EC2 instances that are treated as a logical
grouping for the purposes of automatic scaling and management.

For more information about Amazon EC2 Auto Scaling, see the [Amazon EC2 Auto Scaling\
User Guide](../../../autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.md).

###### Note

Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group
using either a [launch\
template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) or a launch configuration. We strongly recommend that you do not use
launch configurations. For more information, see [Launch configurations](../../../autoscaling/ec2/userguide/launch-configurations.md)
in the _Amazon EC2 Auto Scaling User Guide_.

For help migrating from launch configurations to launch templates, see [Migrate AWS CloudFormation stacks from launch configurations to launch\
templates](../../../autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.md) in the _Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AutoScaling::AutoScalingGroup",
  "Properties" : {
      "AutoScalingGroupName" : String,
      "AvailabilityZoneDistribution" : AvailabilityZoneDistribution,
      "AvailabilityZoneIds" : [ String, ... ],
      "AvailabilityZoneImpairmentPolicy" : AvailabilityZoneImpairmentPolicy,
      "AvailabilityZones" : [ String, ... ],
      "CapacityRebalance" : Boolean,
      "CapacityReservationSpecification" : CapacityReservationSpecification,
      "Context" : String,
      "Cooldown" : String,
      "DefaultInstanceWarmup" : Integer,
      "DeletionProtection" : String,
      "DesiredCapacity" : String,
      "DesiredCapacityType" : String,
      "HealthCheckGracePeriod" : Integer,
      "HealthCheckType" : String,
      "InstanceId" : String,
      "InstanceLifecyclePolicy" : InstanceLifecyclePolicy,
      "InstanceMaintenancePolicy" : InstanceMaintenancePolicy,
      "LaunchConfigurationName" : String,
      "LaunchTemplate" : LaunchTemplateSpecification,
      "LifecycleHookSpecificationList" : [ LifecycleHookSpecification, ... ],
      "LoadBalancerNames" : [ String, ... ],
      "MaxInstanceLifetime" : Integer,
      "MaxSize" : String,
      "MetricsCollection" : [ MetricsCollection, ... ],
      "MinSize" : String,
      "MixedInstancesPolicy" : MixedInstancesPolicy,
      "NewInstancesProtectedFromScaleIn" : Boolean,
      "NotificationConfigurations" : [ NotificationConfiguration, ... ],
      "PlacementGroup" : String,
      "ServiceLinkedRoleARN" : String,
      "SkipZonalShiftValidation" : Boolean,
      "Tags" : [ TagProperty, ... ],
      "TargetGroupARNs" : [ String, ... ],
      "TerminationPolicies" : [ String, ... ],
      "TrafficSources" : [ TrafficSourceIdentifier, ... ],
      "VPCZoneIdentifier" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AutoScaling::AutoScalingGroup
Properties:
  AutoScalingGroupName: String
  AvailabilityZoneDistribution:
    AvailabilityZoneDistribution
  AvailabilityZoneIds:
    - String
  AvailabilityZoneImpairmentPolicy:
    AvailabilityZoneImpairmentPolicy
  AvailabilityZones:
    - String
  CapacityRebalance: Boolean
  CapacityReservationSpecification:
    CapacityReservationSpecification
  Context: String
  Cooldown: String
  DefaultInstanceWarmup: Integer
  DeletionProtection: String
  DesiredCapacity: String
  DesiredCapacityType: String
  HealthCheckGracePeriod: Integer
  HealthCheckType: String
  InstanceId: String
  InstanceLifecyclePolicy:
    InstanceLifecyclePolicy
  InstanceMaintenancePolicy:
    InstanceMaintenancePolicy
  LaunchConfigurationName: String
  LaunchTemplate:
    LaunchTemplateSpecification
  LifecycleHookSpecificationList:
    - LifecycleHookSpecification
  LoadBalancerNames:
    - String
  MaxInstanceLifetime: Integer
  MaxSize: String
  MetricsCollection:
    - MetricsCollection
  MinSize: String
  MixedInstancesPolicy:
    MixedInstancesPolicy
  NewInstancesProtectedFromScaleIn: Boolean
  NotificationConfigurations:
    - NotificationConfiguration
  PlacementGroup: String
  ServiceLinkedRoleARN: String
  SkipZonalShiftValidation: Boolean
  Tags:
    - TagProperty
  TargetGroupARNs:
    - String
  TerminationPolicies:
    - String
  TrafficSources:
    - TrafficSourceIdentifier
  VPCZoneIdentifier:
    - String

```

## Properties

`AutoScalingGroupName`

The name of the Auto Scaling group. This name must be unique per Region per account.

The name can contain any ASCII character 33 to 126 including most punctuation
characters, digits, and upper and lowercased letters.

###### Note

You cannot use a colon (:) in the name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZoneDistribution`

The EC2 instance capacity distribution across Availability Zones for the Auto Scaling group.

_Required_: No

_Type_: [AvailabilityZoneDistribution](aws-properties-autoscaling-autoscalinggroup-availabilityzonedistribution.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZoneIds`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZoneImpairmentPolicy`

The Availability Zone impairment policy for the Auto Scaling group.

_Required_: No

_Type_: [AvailabilityZoneImpairmentPolicy](aws-properties-autoscaling-autoscalinggroup-availabilityzoneimpairmentpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZones`

A list of Availability Zones where instances in the Auto Scaling group can be created. Used
for launching into the default VPC subnet in each Availability Zone when not using the
`VPCZoneIdentifier` property, or for attaching a network interface when
an existing network interface ID is specified in a launch template.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityRebalance`

Indicates whether Capacity Rebalancing is enabled. Otherwise, Capacity Rebalancing is
disabled. When you turn on Capacity Rebalancing, Amazon EC2 Auto Scaling attempts to launch a Spot
Instance whenever Amazon EC2 notifies that a Spot Instance is at an elevated risk of
interruption. After launching a new instance, it then terminates an old instance. For
more information, see [Use Capacity\
Rebalancing to handle Amazon EC2 Spot Interruptions](../../../autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.md) in the in the _Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityReservationSpecification`

The capacity reservation specification for the Auto Scaling group.

_Required_: No

_Type_: [CapacityReservationSpecification](aws-properties-autoscaling-autoscalinggroup-capacityreservationspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Context`

Reserved.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cooldown`

_Only needed if you use simple scaling policies._

The amount of time, in seconds, between one scaling activity ending and another one
starting due to simple scaling policies. For more information, see [Scaling\
cooldowns for Amazon EC2 Auto Scaling](../../../autoscaling/ec2/userguide/ec2-auto-scaling-scaling-cooldowns.md) in the _Amazon EC2 Auto Scaling User Guide_.

Default: `300` seconds

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultInstanceWarmup`

The amount of time, in seconds, until a new instance is considered to have finished
initializing and resource consumption to become stable after it enters the
`InService` state.

During an instance refresh, Amazon EC2 Auto Scaling waits for the warm-up period after it replaces an
instance before it moves on to replacing the next instance. Amazon EC2 Auto Scaling also waits for the
warm-up period before aggregating the metrics for new instances with existing instances
in the Amazon CloudWatch metrics that are used for scaling, resulting in more reliable usage
data. For more information, see [Set\
the default instance warmup for an Auto Scaling group](../../../autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.md) in the
_Amazon EC2 Auto Scaling User Guide_.

###### Important

To manage various warm-up settings at the group level, we recommend that you set
the default instance warmup, _even if it is set to 0 seconds_. To
remove a value that you previously set, include the property but specify
`-1` for the value. However, we strongly recommend keeping the
default instance warmup enabled by specifying a value of `0` or other
nominal value.

Default: None

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtection`

The deletion protection setting for the Auto Scaling group.

_Required_: No

_Type_: String

_Allowed values_: `none | prevent-force-deletion | prevent-all-deletion`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DesiredCapacity`

The desired capacity is the initial capacity of the Auto Scaling group at the time of its
creation and the capacity it attempts to maintain. It can scale beyond this capacity if you
configure automatic scaling.

The number must be greater than or equal to the minimum size of the group and less than or
equal to the maximum size of the group. If you do not specify a desired capacity when creating
the stack, the default is the minimum size of the group.

CloudFormation marks the Auto Scaling group as successful (by setting its status to
CREATE\_COMPLETE) when the desired capacity is reached. However, if a maximum Spot price is set
in the launch template or launch configuration that you specified, then desired capacity is
not used as a criteria for success. Whether your request is fulfilled depends on Spot Instance
capacity and your maximum price.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DesiredCapacityType`

The unit of measurement for the value specified for desired capacity. Amazon EC2 Auto Scaling
supports `DesiredCapacityType` for attribute-based instance type selection
only. For more information, see [Create a mixed instances group using attribute-based instance type\
selection](../../../autoscaling/ec2/userguide/create-mixed-instances-group-attribute-based-instance-type-selection.md) in the _Amazon EC2 Auto Scaling User Guide_.

By default, Amazon EC2 Auto Scaling specifies `units`, which translates into number of
instances.

Valid values: `units` \| `vcpu` \| `memory-mib`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckGracePeriod`

The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status
of an EC2 instance that has come into service and marking it unhealthy due to a failed
health check. This is useful if your instances do not immediately pass their health
checks after they enter the `InService` state. For more information, see
[Set the health check\
grace period for an Auto Scaling group](../../../autoscaling/ec2/userguide/health-check-grace-period.md) in the
_Amazon EC2 Auto Scaling User Guide_.

Default: `0` seconds

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckType`

A comma-separated value string of one or more health check types.

The valid values are `EC2`, `EBS`, `ELB`, and
`VPC_LATTICE`. `EC2` is the default health check and cannot be
disabled. For more information, see [Health checks\
for instances in an Auto Scaling group](../../../autoscaling/ec2/userguide/ec2-auto-scaling-health-checks.md) in the
_Amazon EC2 Auto Scaling User Guide_.

Only specify `EC2` if you must clear a value that was previously
set.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceId`

The ID of the instance used to base the launch configuration on. For more information, see
[Create an Auto Scaling group using an EC2 instance](../../../autoscaling/ec2/userguide/create-asg-from-instance.md) in the _Amazon EC2 Auto_
_Scaling User Guide_.

If you specify `LaunchTemplate`, `MixedInstancesPolicy`, or
`LaunchConfigurationName`, don't specify `InstanceId`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceLifecyclePolicy`

The instance lifecycle policy for the Auto Scaling group.

_Required_: No

_Type_: [InstanceLifecyclePolicy](aws-properties-autoscaling-autoscalinggroup-instancelifecyclepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceMaintenancePolicy`

An instance maintenance policy. For more information, see [Set instance maintenance policy](../../../autoscaling/ec2/userguide/ec2-auto-scaling-instance-maintenance-policy.md) in the
_Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: [InstanceMaintenancePolicy](aws-properties-autoscaling-autoscalinggroup-instancemaintenancepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchConfigurationName`

The name of the launch configuration to use to launch instances.

Required only if you don't specify `LaunchTemplate`,
`MixedInstancesPolicy`, or `InstanceId`.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LaunchTemplate`

Information used to specify the launch template and version to use to launch instances.
You can alternatively associate a launch template to the Auto Scaling group by specifying a
`MixedInstancesPolicy`. For more information about creating launch templates, see
[Create a launch template for an Auto Scaling group](../../../autoscaling/ec2/userguide/create-launch-template.md) in the _Amazon EC2 Auto_
_Scaling User Guide_.

If you omit this property, you must specify `MixedInstancesPolicy`,
`LaunchConfigurationName`, or `InstanceId`.

_Required_: No

_Type_: [LaunchTemplateSpecification](aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LifecycleHookSpecificationList`

One or more lifecycle hooks to add to the Auto Scaling group before instances are
launched.

_Required_: No

_Type_: Array of [LifecycleHookSpecification](aws-properties-autoscaling-autoscalinggroup-lifecyclehookspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadBalancerNames`

A list of Classic Load Balancers associated with this Auto Scaling group. For Application Load Balancers, Network Load Balancers, and Gateway Load Balancers,
specify the `TargetGroupARNs` property instead.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxInstanceLifetime`

The maximum amount of time, in seconds, that an instance can be in service. The
default is null. If specified, the value must be either 0 or a number equal to or
greater than 86,400 seconds (1 day). For more information, see [Replace Auto Scaling instances based on maximum instance lifetime](../../../autoscaling/ec2/userguide/asg-max-instance-lifetime.md) in the
_Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSize`

The maximum size of the group.

###### Note

With a mixed instances policy that uses instance weighting, Amazon EC2 Auto Scaling may need to
go above `MaxSize` to meet your capacity requirements. In this event,
Amazon EC2 Auto Scaling will never go above `MaxSize` by more than your largest instance
weight (weights that define how many units each instance contributes to the desired
capacity of the group).

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsCollection`

Enables the monitoring of group metrics of an Auto Scaling group. By default, these
metrics are disabled.

_Required_: No

_Type_: [Array](aws-properties-autoscaling-autoscalinggroup-metricscollection.md) of [MetricsCollection](aws-properties-autoscaling-autoscalinggroup-metricscollection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinSize`

The minimum size of the group.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MixedInstancesPolicy`

An embedded object that specifies a mixed instances policy.

The policy includes properties that not only define the distribution of On-Demand
Instances and Spot Instances, the maximum price to pay for Spot Instances (optional), and how
the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities, but
also the properties that specify the instance configuration information—the launch template
and instance types. The policy can also include a weight for each instance type and different
launch templates for individual instance types.

For more information, see [Auto Scaling\
groups with multiple instance types and purchase options](../../../autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.md) in the _Amazon EC2_
_Auto Scaling User Guide_.

_Required_: No

_Type_: [MixedInstancesPolicy](aws-properties-autoscaling-autoscalinggroup-mixedinstancespolicy.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`NewInstancesProtectedFromScaleIn`

Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling
when scaling in. For more information about preventing instances from terminating on
scale in, see [Use\
instance scale-in protection](../../../autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.md) in the
_Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationConfigurations`

Configures an Auto Scaling group to send notifications when specified events take
place.

_Required_: No

_Type_: Array of [NotificationConfiguration](aws-properties-autoscaling-autoscalinggroup-notificationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PlacementGroup`

The name of the placement group into which to launch your instances. For more
information, see [Placement groups](../../../ec2/latest/userguide/placement-groups.md) in the
_Amazon EC2 User Guide_.

###### Note

A _cluster_ placement group is a logical grouping of instances
within a single Availability Zone. You cannot specify multiple Availability Zones
and a cluster placement group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceLinkedRoleARN`

The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to
call other AWS service on your behalf. By default, Amazon EC2 Auto Scaling uses a service-linked role
named `AWSServiceRoleForAutoScaling`, which it creates if it does not exist.
For more information, see [Service-linked\
roles](../../../autoscaling/ec2/userguide/autoscaling-service-linked-role.md) in the _Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkipZonalShiftValidation`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags. You can tag your Auto Scaling group and propagate the tags to the Amazon EC2
instances it launches. Tags are not propagated to Amazon EBS volumes. To add tags to Amazon EBS
volumes, specify the tags in a launch template but use caution. If the launch template
specifies an instance tag with a key that is also specified for the Auto Scaling group, Amazon EC2 Auto Scaling
overrides the value of that instance tag with the value specified by the Auto Scaling group. For
more information, see [Tag Auto Scaling groups and\
instances](../../../autoscaling/ec2/userguide/ec2-auto-scaling-tagging.md) in the _Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Array of [TagProperty](aws-properties-autoscaling-autoscalinggroup-tagproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroupARNs`

The Amazon Resource Names (ARN) of the Elastic Load Balancing target groups to associate with the Auto Scaling
group. Instances are registered as targets with the target groups. The target groups
receive incoming traffic and route requests to one or more registered targets. For more
information, see [Use Elastic Load Balancing to\
distribute traffic across the instances in your Auto Scaling group](../../../autoscaling/ec2/userguide/autoscaling-load-balancer.md) in the
_Amazon EC2 Auto Scaling User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminationPolicies`

A policy or a list of policies that are used to select the instance to terminate.
These policies are executed in the order that you list them. For more information, see
[Configure\
termination policies for Amazon EC2 Auto Scaling](../../../autoscaling/ec2/userguide/ec2-auto-scaling-termination-policies.md) in the
_Amazon EC2 Auto Scaling User Guide_.

Valid values: `Default` \| `AllocationStrategy` \|
`ClosestToNextInstanceHour` \| `NewestInstance` \|
`OldestInstance` \| `OldestLaunchConfiguration` \|
`OldestLaunchTemplate` \|
`arn:aws:lambda:region:account-id:function:my-function:my-alias`

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficSources`

The traffic sources associated with this Auto Scaling group.

_Required_: No

_Type_: Array of [TrafficSourceIdentifier](aws-properties-autoscaling-autoscalinggroup-trafficsourceidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VPCZoneIdentifier`

A list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling
group can be created.

If this resource specifies public subnets and is also in a VPC that is defined in the same
stack template, you must use the [DependsOn\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html).

###### Note

When you update `VPCZoneIdentifier`, this retains the same Auto Scaling group
and replaces old instances with new ones, according to the specified subnets. You can
optionally specify how CloudFormation handles these updates by using an [UpdatePolicy\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html).

Required to launch instances into a nondefault VPC. If you specify
`VPCZoneIdentifier` with `AvailabilityZones`, the subnets that you
specify for this property must reside in those Availability Zones.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:
`mystack-myasgroup-NT5EUXTNTXXD`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutoScalingGroupARN`

The Amazon Resource Name (ARN) of the Auto Scaling group.

## Remarks

When you update the launch template or launch configuration for an Auto Scaling group,
this update action does not deploy any change across the running Amazon EC2 instances in the
Auto Scaling group. All new instances will get the updated configuration, but existing
instances continue to run with the configuration that they were originally launched with.
This works the same way as any other Auto Scaling group.

You can add an [UpdatePolicy\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html) to your stack to perform rolling updates (or replace the group) when a
change has been made to the group. You can find a sample update policy for rolling updates
in [Configure Amazon\
EC2 Auto Scaling resources](../userguide/quickref-ec2-auto-scaling.md). Alternatively, you can force a rolling update on your
instances at any time after updating the stack by starting an instance refresh. For more
information, see [Replace Auto Scaling instances\
based on an instance refresh](../../../autoscaling/ec2/userguide/asg-instance-refresh.md) in the _Amazon EC2 Auto Scaling User_
_Guide_.

You can use a [CreationPolicy\
attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-creationpolicy.html) with an Auto Scaling group to prevent its status from reaching create
complete until CloudFormation receives a specified number of success signals. For more
information, see [Use a\
CreationPolicy to Wait for On-Instance Configurations](https://aws.amazon.com/blogs/devops/use-a-creationpolicy-to-wait-for-on-instance-configurations) on the AWS
DevOps Blog. For an example template, see [Configure Amazon\
EC2 Auto Scaling resources](../userguide/quickref-ec2-auto-scaling.md).

Note that Amazon EC2 Auto Scaling provides scaling activities to help you monitor the
progress of your Auto Scaling group and to assist in troubleshooting any configuration
issues when launching Amazon EC2 instances. For more information, see [Verify a scaling activity for an Auto Scaling group](../../../autoscaling/ec2/userguide/as-verify-scaling-activity.md) in the _Amazon EC2_
_Auto Scaling User Guide_.

## Examples

The following examples create or make changes to an Auto Scaling group.

- [An Auto Scaling group and a launch template with a parameters section](#aws-resource-autoscaling-autoscalinggroup--examples--An_Auto_Scaling_group_and_a_launch_template_with_a_parameters_section)

- [Auto Scaling group with CloudWatch monitoring enabled and custom tags](#aws-resource-autoscaling-autoscalinggroup--examples--Auto_Scaling_group_with_CloudWatch_monitoring_enabled_and_custom_tags)

### An Auto Scaling group and a launch template with a parameters section

The following example shows an Auto Scaling group. You specify values for the
`MaxSize` and `MinSize` properties.

It also shows an [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) resource that contains the instance configuration
information for the group, which uses the `LaunchTemplate` property to specify
the launch template. The [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) intrinsic function gets the ID of the `AWS::EC2::LaunchTemplate`
resource `myLaunchTemplate`. The [GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html) function gets the latest version number (for example, `1`) of
the launch template for the `Version` property.

This example references [parameters](../userguide/parameters-section-structure.md) to specify the `ImageId` and `InstanceType`
properties for the launch template and the `VPCZoneIdentifier` property for the
group. Parameters are variables that you can specify when you create or update the stack.
By default, the `ImageId` property of the launch template references the latest
Amazon Linux 2 AMI from the AWS Systems Manager Parameter Store. For more
information, see [AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) in the _AWS Systems Manager User Guide_ and the blog post [Query for the latest Amazon Linux AMI IDs using AWS Systems Manager\
Parameter Store](https://aws.amazon.com/blogs/compute/query-for-the-latest-amazon-linux-ami-ids-using-aws-systems-manager-parameter-store) on the AWS Compute Blog.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "LatestAmiId": {
            "Description": "Region specific image from the Parameter Store",
            "Type": "AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>",
            "Default": "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
        },
        "InstanceType": {
            "Description": "Amazon EC2 instance type for the instances",
            "Type": "String",
            "AllowedValues": [
                "t3.micro",
                "t3.small",
                "t3.medium"
            ],
            "Default": "t3.micro"
        },
        "Subnets": {
            "Type": "List<AWS::EC2::Subnet::Id>",
            "Description": "A list of subnets for the Auto Scaling group"
        }
    },
    "Resources": {
        "myLaunchTemplate": {
            "Type": "AWS::EC2::LaunchTemplate",
            "Properties": {
                "LaunchTemplateName": { "Fn::Sub": "${AWS::StackName}-launch-template" },
                "LaunchTemplateData": {
                    "ImageId": { "Ref": "LatestAmiId" },
                    "InstanceType": { "Ref": "InstanceType" }
                }
            }
        },
        "myASG": {
            "Type": "AWS::AutoScaling::AutoScalingGroup",
            "Properties": {
                "LaunchTemplate": {
                    "LaunchTemplateId": { "Ref": "myLaunchTemplate" },
                    "Version": { "Fn::GetAtt": [ "myLaunchTemplate", "LatestVersionNumber" ] }
                },
                "MaxSize": "1",
                "MinSize": "1",
                "VPCZoneIdentifier": { "Ref": "Subnets" }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Parameters:
  LatestAmiId:
    Description: Region specific image from the Parameter Store
    Type: 'AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>'
    Default: '/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2'
  InstanceType:
    Description: Amazon EC2 instance type for the instances
    Type: String
    AllowedValues:
      - t3.micro
      - t3.small
      - t3.medium
    Default: t3.micro
  Subnets:
    Type: 'List<AWS::EC2::Subnet::Id>'
    Description: A list of subnets for the Auto Scaling group
Resources:
  myLaunchTemplate:
    Type: AWS::EC2::LaunchTemplate
    Properties:
      LaunchTemplateName: !Sub ${AWS::StackName}-launch-template
      LaunchTemplateData:
        ImageId: !Ref LatestAmiId
        InstanceType: !Ref InstanceType
  myASG:
    Type: AWS::AutoScaling::AutoScalingGroup
    Properties:
      LaunchTemplate:
        LaunchTemplateId: !Ref myLaunchTemplate
        Version: !GetAtt myLaunchTemplate.LatestVersionNumber
      MaxSize: '1'
      MinSize: '1'
      VPCZoneIdentifier: !Ref Subnets
```

### Auto Scaling group with CloudWatch monitoring enabled and custom tags

The following snippet shows an Auto Scaling group with CloudWatch monitoring enabled
and custom tags. The `LaunchTemplate` property references an [AWS::EC2::LaunchTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) resource with the logical name
`myLaunchTemplate` that is defined elsewhere in your template.

You specify the CloudWatch metrics to monitor using the `MetricsCollection`
property. If you keep the metrics as they are, only `GroupMinSize` and
`GroupMaxSize` metrics are enabled.

You specify the tag keys and tag key values for the `Tags` property. If you
keep the provided tags, the first tag, `Environment` = `Production`,
is assigned to the Auto Scaling group and to any EC2 instances launched as part of the
Auto Scaling group. The second tag, `Purpose` = `WebServerGroup`, is
assigned only to the Auto Scaling group itself.

You also specify values for the `MaxSize`, `MinSize`, and
`VPCZoneIdentifier` properties.

#### JSON

```json

{
    "Resources": {
        "myASG": {
            "Type": "AWS::AutoScaling::AutoScalingGroup",
            "Properties": {
                "LaunchTemplate": {
                    "LaunchTemplateId": { "Ref": "myLaunchTemplate" },
                    "Version": { "Fn::GetAtt": [ "myLaunchTemplate", "LatestVersionNumber" ] }
                },
                "MaxSize": "1",
                "MinSize": "1",
                "VPCZoneIdentifier": [
                    "subnetIdAz1",
                    "subnetIdAz2",
                    "subnetIdAz3"
                ],
                "MetricsCollection": [
                    {
                        "Granularity": "1Minute",
                        "Metrics": [
                            "GroupMinSize",
                            "GroupMaxSize"
                        ]
                    }
                ],
                "Tags": [
                    {
                        "Key": "Environment",
                        "Value": "Production",
                        "PropagateAtLaunch": "true"
                    },
                    {
                        "Key": "Purpose",
                        "Value": "WebServerGroup",
                        "PropagateAtLaunch": "false"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  myASG:
    Type: AWS::AutoScaling::AutoScalingGroup
    Properties:
      LaunchTemplate:
        LaunchTemplateId: !Ref myLaunchTemplate
        Version: !GetAtt myLaunchTemplate.LatestVersionNumber
      MaxSize: '1'
      MinSize: '1'
      VPCZoneIdentifier:
        - subnetIdAz1
        - subnetIdAz2
        - subnetIdAz3
      MetricsCollection:
        - Granularity: 1Minute
          Metrics:
            - GroupMinSize
            - GroupMaxSize
      Tags:
        - Key: Environment
          Value: Production
          PropagateAtLaunch: true
        - Key: Purpose
          Value: WebServerGroup
          PropagateAtLaunch: false
```

## See also

- You can find additional useful snippets in the following sections of the _AWS CloudFormation User Guide_:

- For examples of Auto Scaling groups, see [Configure\
Amazon EC2 Auto Scaling resources](../userguide/quickref-ec2-auto-scaling.md).

- For examples of launch templates, see [Create\
launch templates](../userguide/quickref-ec2-launch-templates.md).

- [Migrate AWS CloudFormation stacks from launch configurations to\
launch templates](../../../autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.md) in the _Amazon EC2 Auto Scaling User_
_Guide_

- [CreateAutoScalingGroup](../../../../reference/autoscaling/ec2/apireference/api-createautoscalinggroup.md) and [UpdateAutoScalingGroup](../../../../reference/autoscaling/ec2/apireference/api-updateautoscalinggroup.md) in the _Amazon EC2 Auto Scaling API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon EC2 Auto Scaling

AcceleratorCountRequest
