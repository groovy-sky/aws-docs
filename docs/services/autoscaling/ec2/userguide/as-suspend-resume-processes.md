# Suspend and resume Amazon EC2 Auto Scaling processes

This topic describes how to suspend and then resume one or more of the processes for
your Auto Scaling group to temporarily disable certain operations.

Suspending processes can be useful when you need to investigate or troubleshoot an
issue without interference from scaling policies or scheduled actions. It also helps
prevent Amazon EC2 Auto Scaling from marking instances unhealthy and replacing them while you are
making changes to your Auto Scaling group.

###### Topics

- [Types of processes](#process-types)

- [Considerations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/suspend-resume-considerations.html)

- [Suspend processes](https://docs.aws.amazon.com/autoscaling/ec2/userguide/suspend-processes.html)

- [Resume processes](https://docs.aws.amazon.com/autoscaling/ec2/userguide/resume-processes.html)

- [How suspended processes affect other processes](https://docs.aws.amazon.com/autoscaling/ec2/userguide/understand-how-suspending-processes-affects-other-processes.html)

###### Note

In addition to suspensions that you initiate, Amazon EC2 Auto Scaling can also suspend processes
for Auto Scaling groups that repeatedly fail to launch instances. This is known as an
_administrative suspension_. An administrative suspension
most commonly applies to Auto Scaling groups that have been trying to launch instances for
over 24 hours but have not succeeded in launching any instances. You can resume
processes that were suspended by Amazon EC2 Auto Scaling for administrative reasons.

## Types of processes

The suspend-resume feature supports the following processes:

- `Launch` – Adds instances to the Auto Scaling group when the
group scales out, or when Amazon EC2 Auto Scaling chooses to launch instances for other
reasons, such as when it adds instances to a warm pool.

- `Terminate` – Removes instances from the Auto Scaling group when
the group scales in, or when Amazon EC2 Auto Scaling chooses to terminate instances for
other reasons, such as when an instance is terminated for exceeding its
maximum lifetime duration or failing a health check.

- `AddToLoadBalancer` – Adds instances to the attached
load balancer target group or Classic Load Balancer when they are launched. For more
information, see [Use Elastic Load Balancing to distribute incoming application traffic in your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html).

- `AlarmNotification` – Accepts notifications from CloudWatch
alarms that are associated with dynamic scaling policies. For more
information, see [Dynamic scaling for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html).

- `AZRebalance` – Balances the number of EC2 instances in
the group evenly across all of the specified Availability Zones when the
group becomes unbalanced, for example, when a previously unavailable
Availability Zone returns to a healthy state. For more information, see
[Rebalancing activities](https://docs.aws.amazon.com/autoscaling/ec2/userguide/auto-scaling-benefits.html#AutoScalingBehavior.InstanceUsage).

- `HealthCheck` – Checks the health of the instances and
marks an instance as unhealthy if Amazon EC2 or Elastic Load Balancing tells Amazon EC2 Auto Scaling that the
instance is unhealthy. This process can override the health status of an
instance that you set manually. For more information, see [Health checks for instances in an Auto Scaling group](ec2-auto-scaling-health-checks.md).

- `InstanceRefresh` – Terminates and replaces instances
using the instance refresh feature. For more information, see [Use an instance refresh to update instances in an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html).

- `ReplaceUnhealthy` – Terminates instances that are
marked as unhealthy and then creates new instances to replace them. For more
information, see [Health checks for instances in an Auto Scaling group](ec2-auto-scaling-health-checks.md).

- `ScheduledActions` – Performs the scheduled scaling
actions that you create or that are created for you when you create an
AWS Auto Scaling scaling plan and turn on predictive scaling. For more information,
see [Scheduled scaling for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scheduled-scaling.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Design for graceful instance termination

Considerations
