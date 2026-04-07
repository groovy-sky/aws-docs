# Auto Scaling group Availability Zone distribution

The following information describes Auto Scaling group Availability Zone strategies.

**Balanced best effort**

Auto Scaling maintains an equal number of instances across enabled Availability Zones. If launch attempts fail in an Availability Zone, Auto Scaling attempts to launch
instances in another healthy Availability Zone. This strategy is important for applications that need Availability Zone redundancy and are not impacted by imbalanced groups.

**Balanced only**

Auto Scaling maintains an equal number of instances across enabled Availability Zones. If
launch attempts fail in an Availability Zone, Auto Scaling will continue to attempt to launch
instances in the Availability Zone. This strategy is important to meet certain
requirements such as quorum-based workloads or if your Auto Scaling group can
tolerate the loss of an Availability Zone because you have sufficient capacity in the
remaining Availability Zones.

The Availability Zone distribution strategy selection is in the **Network** section of the AWS Management Console or
you can use the [create-auto-scaling-group](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/autoscaling/create-auto-scaling-group.html) or
[update-auto-scaling-group](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/autoscaling/update-auto-scaling-group.html) commands.

For more information, see [Create Auto Scaling groups using launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-auto-scaling-groups-launch-template.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Enable zonal shift using the AWS Management Console or the AWS CLI

Detach-attach instances
