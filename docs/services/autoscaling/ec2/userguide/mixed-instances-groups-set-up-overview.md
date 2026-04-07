# Setup overview for creating a mixed instances group

This topic provides an overview and best practices for creating an Auto Scaling mixed instances
group.

###### Contents

- [Overview](#mixed-instances-groups-overview)

- [Instance type flexibility](#mixed-instances-group-instance-flexibility)

- [Availability Zone flexibility](#mixed-instances-group-az-flexibility)

- [Spot max price](#mixed-instances-group-spot-max-price)

- [Proactive capacity rebalancing](#use-capacity-rebalancing)

- [Scaling behavior](#mixed-instances-group-scaling-behavior)

- [Regional availability of instance types](#setup-overview-regional-availability-of-instance-types)

- [Related resources](#setup-overview-related-resources)

- [Limitations](#setup-overview-limitations)

## Overview

To create a mixed instances group, you have two options:

- [Attribute-based instance type selection](create-mixed-instances-group-attribute-based-instance-type-selection.md) – Define your
compute requirements to choose your instance types automatically based on
their specific instance attributes.

- [Manual instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-mixed-instances-group-manual-instance-type-selection.html) – Manually choose the
instance types that suit your workload.

Manual selection

The following steps describe how to create a mixed instances group by
manually choosing instance types:

1. Choose a launch template that has the parameters to launch an
    EC2 instance. Parameters in launch templates are optional, but
    Amazon EC2 Auto Scaling can't launch an instance if the amilong;
    (AMI) ID is missing from the launch template.

2. Choose the option to override the launch template.

3. Manually choose the instance types that suit your
    workload.

4. Specify the percentages of On-Demand Instances and Spot
    Instances to launch.

5. Choose allocation strategies that determine how Amazon EC2 Auto Scaling
    fulfills your On-Demand and Spot capacities from the possible
    instance types.

6. Choose the Availability Zones and VPC subnets to launch your
    instances in.

7. Specify the initial size of the group (the desired capacity)
    and the minimum and maximum size of the group.

Overrides are necessary to override the instance type declared in the
launch template and use multiple instances types that are embedded in
the Auto Scaling group's own resource definition. For more information about the
instance types that are available, see [Instance types](../../../ec2/latest/userguide/instance-types.md) in
the _Amazon EC2 User Guide_.

You can also configure the following optional parameters for each
instance type:

- `LaunchTemplateSpecification` – You can
assign a different launch template to an instance type as
needed. This option is currently not available from the console.
For more information, see [Use multiple launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html).

- `WeightedCapacity` – You decide how much the
instance counts toward the desired capacity relative to the rest
of the instances in your group. If you specify a
`WeightedCapacity` value for one instance type,
you must specify a `WeightedCapacity` value for all
of them. By default, each instance counts as one toward your
desired capacity. For more information, see [Configure an Auto Scaling group to use instance weights](ec2-auto-scaling-mixed-instances-groups-instance-weighting.md).

Attribute-based selection

To let Amazon EC2 Auto Scaling choose your instance types automatically based on
their specific instance attributes, use the following steps to create a
mixed instances group by specifying your compute requirements:

1. Choose a launch template that has the parameters to launch an
    EC2 instance. Parameters in launch templates are optional, but
    Amazon EC2 Auto Scaling can't launch an instance if the amilong;
    (AMI) ID is missing from the launch template.

2. Choose the option to override the launch template.

3. Specify instance attributes that match your compute
    requirements, such as vCPUs and memory requirements.

4. Specify the percentages of On-Demand Instances and Spot
    Instances to launch.

5. Choose allocation strategies that determine how Amazon EC2 Auto Scaling
    fulfills your On-Demand and Spot capacities from the possible
    instance types.

6. Choose the Availability Zones and VPC subnets to launch your
    instances in.

7. Specify the initial size of the group (the desired capacity)
    and the minimum and maximum size of the group.

Overrides are necessary to override the instance type declared in the
launch template and use a set of instance attributes that describe your
compute requirements. For supported attributes, see [InstanceRequirements](../../../../reference/autoscaling/ec2/apireference/api-instancerequirements.md) in the
_Amazon EC2 Auto Scaling API Reference_. Alternatively, you can use a
launch template that already has your instance attributes definition.

You can also configure the `LaunchTemplateSpecification`
parameter within the overrides structure to assign a different launch
template to a set of instance requirements as needed. This option is
currently not available from the console. For more information, see
[LaunchTemplateOverrides](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchTemplateOverrides.html) in the
_Amazon EC2 Auto Scaling API Reference_.

By default, you set the number of instances as the desired capacity of
your Auto Scaling group.

Alternatively, you can set the value for desired capacity to the
number of vCPUs or the amount of memory. To do so, use the
`DesiredCapacityType` property in the
`CreateAutoScalingGroup` API operation or the
**Desired capacity type** dropdown field in the
AWS Management Console. This is a useful alternative to [instance weights](ec2-auto-scaling-mixed-instances-groups-instance-weighting.md).

## Instance type flexibility

To enhance availability, deploy your application across multiple instance types.
It's a best practice to use multiple instance types to satisfy capacity
requirements. This way, Amazon EC2 Auto Scaling can launch another instance type if there is
insufficient instance capacity in your chosen Availability Zones.

If there is insufficient instance capacity with Spot
Instances, Amazon EC2 Auto Scaling keeps trying to launch from other Spot Instance pools. (The
pools it uses are determined by your choice of instance types and allocation
strategy.) Amazon EC2 Auto Scaling helps you leverage the cost savings of Spot Instances by
launching them instead of On-Demand Instances.

We recommend being flexible across at least 10 instance types for each workload.
When choosing your instance types, don't limit yourself to
the most popular new instance types. Choosing earlier generation instance types
tends to result in fewer Spot interruptions because they are less in demand from
On-Demand customers.

## Availability Zone flexibility

We strongly recommend that you span your Auto Scaling group across multiple Availability
Zones. With multiple Availability Zones, you can design applications that
automatically fail over between zones for greater resiliency.

As an added benefit, you can access a deeper Amazon EC2 capacity pool when compared to
groups in a single Availability Zone. Because capacity fluctuates independently for
each instance type in each Availability Zone, you can often get more compute
capacity with flexibility for both the instance type and the Availability Zone.

For more information about using multiple Availability Zones, see [Example: Distribute instances across Availability Zones](auto-scaling-benefits.md#arch-AutoScalingMultiAZ).

## Spot max price

When you create your Auto Scaling group using the AWS CLI or an SDK, you can specify the
`SpotMaxPrice` parameter. The `SpotMaxPrice` parameter
determines the maximum price that you're willing to pay for a Spot Instance hour.

When you specify the `WeightedCapacity` parameter in your overrides (or
`"DesiredCapacityType": "vcpu"` or `"DesiredCapacityType":
                    "memory-mib"` at the group level), the maximum price represents the
maximum unit price, not the maximum price for a whole instance.

We strongly recommend that you do not specify a maximum price. Your application
might not run if you do not receive any Spot Instances, such as when your maximum
price is too low. If you don't specify a maximum price, the default maximum price is
the On-Demand price. You pay only the Spot price for the Spot Instances that you
launch. You still receive the steep discounts provided by Spot Instances. These
discounts are possible because of the stable Spot pricing that's available with the
[Spot\
pricing model](https://aws.amazon.com/blogs/compute/new-amazon-ec2-spot-pricing). For more information, see [Pricing and\
savings](../../../ec2/latest/userguide/using-spot-instances.md#spot-pricing) in the _Amazon EC2 User Guide_.

## Proactive capacity rebalancing

If your use case allows, we recommend Capacity Rebalancing. Capacity Rebalancing
helps you maintain your workload availability by proactively replacing
Spot Instances at risk of interruption.

When Capacity Rebalancing is enabled, Amazon EC2 Auto Scaling attempts to proactively replace
Spot Instances that have received an EC2 instance rebalance recommendation. This provides an
opportunity to rebalance your workload to new Spot Instances that are not at an
elevated risk of interruption.

For more information, see [Capacity Rebalancing in Auto Scaling to replace at-risk Spot Instances](ec2-auto-scaling-capacity-rebalancing.md).

## Scaling behavior

When you create a mixed instances group, it uses On-Demand Instances by default.
To use Spot Instances, you must modify the percentage of the group to be launched as
On-Demand Instances. You can specify any number from 0 to 100 for the On-Demand
percentage.

Optionally, you can also designate a base number of On-Demand Instances to start
with. If you do so, Amazon EC2 Auto Scaling waits to launch Spot Instances until after it launches
the base capacity of On-Demand Instances when the group scales out. Anything beyond
the base capacity uses the On-Demand percentage to determine how many On-Demand
Instances and Spot Instances to launch.

Amazon EC2 Auto Scaling converts the percentage to the equivalent number of instances. If the
result creates a fractional number, it rounds up to the next integer in favor of
On-Demand Instances.

The following table demonstrates the behavior of the Auto Scaling group as it increases
and decreases in size.

Example: Scaling behaviorPurchase optionsGroup size and number of
running instances across purchase options**10****20****30****40**

**Example 1**: base of 10,
50/50% On-Demand/Spot

On-Demand Instances (base amount)10101010On-Demand Instances051015Spot Instances051015

**Example 2**: base of 0,
0/100% On-Demand/Spot

On-Demand Instances (base amount)0000On-Demand Instances0000Spot Instances10203040

**Example 3**: base of 0,
60/40% On-Demand/Spot

On-Demand Instances (base amount)0000On-Demand Instances6121824Spot Instances481216

**Example 4**: base of 0,
100/0% On-Demand/Spot

On-Demand Instances (base amount)0000On-Demand Instances10203040Spot Instances0000

**Example 5**: base of 12,
0/100% On-Demand/Spot

On-Demand Instances (base amount)10121212On-Demand Instances0000Spot Instances081828

When the size of the group _increases_, Amazon EC2 Auto Scaling attempts to
balance your capacity evenly across your specified Availability Zones. Then, it
launches instance types according to the specified allocation strategy.

When the size of the group _decreases_, Amazon EC2 Auto Scaling first
identifies which of the two types (Spot or On-Demand) should be terminated. Then, it
tries to terminate instances in a balanced way across your specified Availability
Zones. It also favors terminating instances in a way that aligns closer to your
allocation strategies. For information about termination policies, see [Configure termination policies for Amazon EC2 Auto Scaling](ec2-auto-scaling-termination-policies.md).

## Regional availability of instance types

The availability of EC2 instance types varies depending on your AWS Region. For
example, the newest generation instance types might not yet be available in a given
Region. Due to the variances in instance availability across Regions, you might
encounter issues when making programmatic requests if multiple instance types in
your overrides are not available in your Region. Using multiple instance types that
are not available in your Region might cause the request to fail entirely. To solve
the issue, retry the request with different instance types, making sure that each
instance type is available in the Region. To search for instance types offered by
location, use the [describe-instance-type-offerings](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/ec2/describe-instance-type-offerings.html) command. For more information, see
[Finding an Amazon EC2 instance\
type](../../../ec2/latest/userguide/instance-discovery.md) in the _Amazon EC2 User Guide_.

## Related resources

For more best practices for Spot Instances, see [Best practices for EC2 Spot](../../../ec2/latest/userguide/spot-best-practices.md)
in the _Amazon EC2 User Guide_.

## Limitations

After you add overrides to an Auto Scaling group using a [mixed instances\
policy](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_MixedInstancesPolicy.html), you can update the overrides with the
`UpdateAutoScalingGroup` API call but not delete them. To completely
remove the overrides, you must first switch the Auto Scaling group to use a launch template
or launch configuration instead of a mixed instances policy. Then, you can add a
mixed instances policy again without any overrides.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Use multiple instance types and purchase options

Allocation strategies for multiple instance types
