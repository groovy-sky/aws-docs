# Considerations for suspending processes

Consider the following before suspending processes:

- Suspending `AlarmNotification` allows you to temporarily stop
the group's target tracking, step, and simple scaling policies without
deleting the scaling policies or their associated CloudWatch alarms. To
temporarily stop individual scaling policies instead, see [Disable a scaling policy for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enable-disable-scaling-policy.html).

- You might choose to suspend the `HealthCheck` and
`ReplaceUnhealthy` processes to reboot instances without
Amazon EC2 Auto Scaling terminating the instances based on its health checks. However, if
you need Amazon EC2 Auto Scaling to continue performing health checks on the remaining
instances, use the standby feature instead. For more information, see [Temporarily remove instances from your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enter-exit-standby.html).

- If you suspend the `Launch` and `Terminate`
processes, or `AZRebalance`, and then you make changes to your
Auto Scaling group, for example, by detaching instances or changing the Availability
Zones that are specified, your group can become unbalanced between
Availability Zones. If that happens, after you resume the suspended
processes, Amazon EC2 Auto Scaling gradually redistributes instances evenly between the
Availability Zones.

- If you suspend the `Terminate` process, you can still force
instances to be terminated by using the [delete-auto-scaling-group](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/autoscaling/delete-auto-scaling-group.html) command with the force delete
option.

- Suspending the `Terminate` process applies only to instances
that are currently in the `InService` state. It does not prevent
the termination of instances in other states, such as `Pending`,
or instances that fail to resume properly from standby.

- The `RemoveFromLoadBalancerLowPriority` process can be ignored
when it is present in calls to describe Auto Scaling groups using the AWS CLI or SDKs.
This process is deprecated and is retained only for backward compatibility.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Suspend-resume processes

Suspend processes
