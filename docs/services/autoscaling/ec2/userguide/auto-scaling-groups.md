# Auto Scaling groups

###### Note

If you are new to Auto Scaling groups, work through the steps in the [Create your first Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-your-first-auto-scaling-group.html)
tutorial to get started and see how an Auto Scaling group responds when an instance in the group
terminates.

An _Auto Scaling group_ contains a collection of EC2 instances that are treated
as a logical grouping for the purposes of automatic scaling and management. An Auto Scaling group
also lets you use Amazon EC2 Auto Scaling features such as health check replacements and scaling policies.
Both maintaining the number of instances in an Auto Scaling group and automatic scaling are the core
functionality of the Amazon EC2 Auto Scaling service.

The size of an Auto Scaling group depends on the number of instances that you set as the desired
capacity. You can adjust its size to meet demand, either manually or by using automatic
scaling.

An Auto Scaling group starts by launching enough instances to meet its desired capacity. It
maintains this number of instances by performing periodic health checks on the instances in
the group. The Auto Scaling group continues to maintain a fixed number of instances even if an
instance becomes unhealthy. If an instance becomes unhealthy, the group terminates the
unhealthy instance and launches another instance to replace it. For more information, see
[Health checks for instances in an Auto Scaling group](ec2-auto-scaling-health-checks.md).

You can use scaling policies to increase or decrease the number of instances in your group
dynamically to meet changing conditions. When the scaling policy is in effect, the Auto Scaling
group adjusts the desired capacity of the group, between the minimum and maximum capacity
values that you specify, and launches or terminates the instances as needed. You can also
scale on a schedule. For more information, see [Choose your scaling method](https://docs.aws.amazon.com/autoscaling/ec2/userguide/scaling-overview.html).

When creating an Auto Scaling group, you can choose whether to launch On-Demand Instances, Spot
Instances, or both. You can specify multiple purchase options for your Auto Scaling group only when
you use a launch template. For more information, see [Auto Scaling groups with multiple instance types and purchase options](ec2-auto-scaling-mixed-instances-groups.md).

Spot Instances provide you with access to unused EC2 capacity at steep discounts relative
to On-Demand prices. For more information, see [Amazon EC2 Spot Instances](https://aws.amazon.com/ec2/spot/pricing). There are key
differences between Spot Instances and On-Demand Instances:

- The price for Spot Instances varies based on demand

- Amazon EC2 can terminate an individual Spot Instance as the availability of, or price
for, Spot Instances changes

When a Spot Instance is terminated, the Auto Scaling group attempts to launch a replacement
instance to maintain the desired capacity for the group.

When instances are launched, if you specified multiple Availability Zones, the desired
capacity is distributed across these Availability Zones. If a scaling action occurs,
Amazon EC2 Auto Scaling automatically maintains balance across all of the Availability Zones that you
specify.

###### Contents

- [Create Auto Scaling groups using launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-auto-scaling-groups-launch-template.html)

- [Create Auto Scaling groups using launch configurations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-auto-scaling-groups-launch-configuration.html)

- [Launch instances synchronously](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-instances-synchronously.html)

- [Update an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/update-auto-scaling-group.html)

- [Tag Auto Scaling groups and instances](ec2-auto-scaling-tagging.md)

- [Instance maintenance policies](ec2-auto-scaling-instance-maintenance-policy.md)

- [Amazon EC2 Auto Scaling lifecycle hooks](lifecycle-hooks.md)

- [Decrease latency for applications with long boot times using warm pools](ec2-auto-scaling-warm-pools.md)

- [Auto Scaling group zonal shift](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.html)

- [Auto Scaling group Availability Zone distribution](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-availability-zone-balanced.html)

- [Detach or attach instances from your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-detach-attach-instances.html)

- [Temporarily remove instances from your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html)

- [Delete your Auto Scaling infrastructure](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-process-shutdown.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Change a launch configuration

Create Auto Scaling groups using launch templates
