# Create an Auto Scaling group using a launch template

When you create an Auto Scaling group, you must specify the necessary information to configure the
Amazon EC2 instances, the Availability Zones and VPC subnets for the instances, the desired
capacity, and the minimum and maximum capacity limits.

To configure Amazon EC2 instances that are launched by your Auto Scaling group, you can specify a
launch template or a launch configuration. The following procedure demonstrates how to
create an Auto Scaling group using a launch template.

###### Prerequisites

- You must have created a launch template. For more information, see [Create a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html).

###### To create an Auto Scaling group using a launch template (console)

01. Open the Amazon EC2 console at
     [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2), and choose **Auto Scaling Groups** from the navigation pane.

02. On the navigation bar at the top of the screen, choose the same AWS Region that
     you used when you created the launch template.

03. Choose **Create an Auto Scaling group**.

04. On the **Choose launch template or configuration** page, do the
     following:
    1. For **Auto Scaling group name**, enter a name for your Auto Scaling
        group.

    2. For **Launch template**, choose an existing launch
        template.

    3. For **Launch template version**, choose whether the Auto Scaling
        group uses the default, the latest, or a specific version of the launch
        template when scaling out.

    4. Verify that your launch template supports all of the options that you are
        planning to use, and then choose **Next**.
05. On the **Choose instance launch options**
     page, if you're not using multiple instance types, you can skip the
     **Instance type requirements** section to use the EC2 instance
     type that is specified in the launch template.

    To use multiple instance types, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html).

06. Under **Network**, for **VPC**, choose a VPC.
     The Auto Scaling group must be created in the same VPC as the security group you specified
     in your launch template.

07. For **Availability Zones and**
    **subnets**, choose one or more subnets in the specified VPC.
     Use subnets in multiple Availability Zones for high availability. For more
     information, see [Considerations when choosing VPC subnets](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-in-vpc.html#as-vpc-considerations).

08. For **Availability Zone distribution**, select a distribution strategy. For more information, see
     [Auto Scaling group Availability Zone distribution](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-availability-zone-balanced.html).

09. If you created a launch template with an instance type specified, then you can
     continue to the next step to create an Auto Scaling group that uses the instance type in the
     launch template.

    Alternatively, you can choose the **Override launch template**
     option if no instance type is specified in your launch template or if you want to
     use multiple instance types for auto scaling. For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html).

10. Choose **Next** to continue to the next step.

    Or, you can accept the rest of the defaults, and choose **Skip to**
    **review**.

11. (Optional) On the **Integrate with other services** page, configure
     the following options, and then choose **Next**:
    1. For **Load balancing**, choose whether to attach your Auto Scaling group to a load balancer. For more information, see
        [Elastic Load Balancing](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html).

    2. For **VPC Lattice integration options**, choose whether to use VPC Lattice. For more information, see
        [Manage traffic flow with a VPC Lattice target group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-vpc-lattice.html).

    3. For **Amazon Application Recovery Controller (ARC) zonal shift**, select the checkbox to enable zonal shift. For more information, see
        [Auto Scaling group zonal shift](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.html).
       1. If you enable zonal shift, for **Health check behavior**, select Ignore unhealthy or Replace unhealthy. For more information, see
           [How zonal shift works for Auto Scaling groups](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.html#asg-zonal-shift-how-it-works).
    4. Under **Health checks**, for **Additional**
       **health check types**, select **Turn on Amazon EBS health**
       **checks**. For more information, see [Monitor Auto Scaling instances with impaired Amazon EBS volumes using health checks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/monitor-and-replace-instances-with-impaired-ebs-volumes.html).

    5. For **Health check grace period**, enter the
        amount of time, in seconds. This amount of time is how long Amazon EC2 Auto Scaling needs
        to wait before checking the health status of an instance after it enters the
        `InService` state. For more information, see [Set the health check grace period for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html).
12. (Optional) On the **Configure group size and scaling**
     page, configure the following options, and then choose
     **Next**:
    1. Under **Group size**, for **Desired**
       **capacity**, enter the initial number of instances to launch.

    2. Under **Scaling**, **Scaling**
       **limits**, if your new value for **Desired**
       **capacity** is greater than **Min desired**
       **capacity** and **Max desired capacity**, the
        **Max desired capacity** is automatically increased to
        the new desired capacity value. You can change these limits as needed. For
        more information, see [Set scaling limits for your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-capacity-limits.html).

    3. For **Automatic scaling**, choose whether you want to
        create a target tracking scaling policy. You can also create this policy
        after your create your Auto Scaling group.

       If you choose **Target tracking scaling policy**, follow
        the directions in [Create a target tracking scaling policy](https://docs.aws.amazon.com/autoscaling/ec2/userguide/policy_creating.html) to create the policy.

    4. Under **Instance maintenance policy**, choose whether you
        want to create an instance maintenance policy. You can also create this
        policy after your create your Auto Scaling group. Follow the directions in [Set an instance maintenance policy](https://docs.aws.amazon.com/autoscaling/ec2/userguide/set-instance-maintenance-policy.html) to create the
        policy.

    5. Under **Additional capacity settings**, **Capacity Reservation preference**, choose whether you want to use a Capacity Reservation preference. For more information, see
        [Reserve capacity in specific Availability Zones with Capacity Reservations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/use-ec2-capacity-reservations.html).

    6. Under **Additional settings**, **Instance scale-in protection**, choose whether to
        enable instance scale-in protection. For more information, see [Use instance scale-in protection to control instance termination](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html).

    7. For **Monitoring**, choose whether to enable CloudWatch group
        metrics collection. These metrics provide measurements that can be
        indicators of a potential issue, such as number of terminating instances or
        number of pending instances. For more information, see [Monitor CloudWatch metrics for your Auto Scaling groups and instances](ec2-auto-scaling-cloudwatch-monitoring.md).

    8. For **default instance warmup**, select this
        option and choose the warmup time for your application. If you are creating
        an Auto Scaling group that has a scaling policy, the default instance warmup feature
        improves the Amazon CloudWatch metrics used for dynamic scaling. For more
        information, see [Set the default instance warmup for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.html).
13. (Optional) On the **Add notifications** page,
     configure the notification, and then choose **Next**. For more
     information, see [Amazon SNS notification options for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-sns-notifications.html).

14. (Optional) On the **Add tags** page, choose **Add tag**, provide a tag key and
     value for each tag, and then choose **Next**. For more information,
     see [Tag Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-tagging.html).

15. On the **Review** page, choose **Create Auto Scaling**
    **group**.

###### To create an Auto Scaling group using the command line

You can use one of the following commands:

- [create-auto-scaling-group](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/autoscaling/create-auto-scaling-group.html) (AWS CLI)

- [New-ASAutoScalingGroup](https://docs.aws.amazon.com/powershell/latest/reference/items/New-ASAutoScalingGroup.html) (AWS Tools for Windows PowerShell)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Create Auto Scaling groups using launch templates

Create a group using the EC2 launch wizard
