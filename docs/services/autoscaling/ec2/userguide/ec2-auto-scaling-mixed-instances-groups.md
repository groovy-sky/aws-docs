# Auto Scaling groups with multiple instance types and purchase options

You can launch and automatically scale a fleet of On-Demand
Instances and Spot Instances within a single Auto Scaling group. In addition to receiving discounts
for using Spot Instances, you can use Reserved Instances or a Savings Plans to receive
discounts on the regular On-Demand Instance pricing. These factors help you optimize your
cost savings for EC2 instances and get the desired scale and performance for your
application.

Spot Instances are spare capacity available at steep discounts
compared to the EC2 On-Demand price. Spot Instances are a cost-effective choice if you can
be flexible about when your applications run and if your applications can be interrupted.
They can be used for various fault-tolerant and flexible applications. Examples include
stateless web servers, API endpoints, big data and analytics applications, containerized
workloads, CI/CD pipelines, high performance and high throughput computing (HPC/HTC),
rendering workloads, and other flexible workloads.

For more information, see [Instance purchasing options](../../../ec2/latest/userguide/instance-purchasing-options.md) in the
_Amazon EC2 User Guide_.

###### Topics

- [Setup overview for creating a mixed instances group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/mixed-instances-groups-set-up-overview.html)

- [Allocation strategies for multiple instance types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/allocation-strategies.html)

- [Create mixed instances group using attribute-based instance type selection](create-mixed-instances-group-attribute-based-instance-type-selection.md)

- [Create a mixed instances group by manually choosing instance types](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-mixed-instances-group-manual-instance-type-selection.html)

- [Configure an Auto Scaling group to use instance weights](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.html)

- [Use multiple launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Create a group using the EC2 launch wizard

Setup overview for creating a mixed instances group
