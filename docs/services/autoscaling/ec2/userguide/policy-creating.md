# Create a target tracking scaling policy

To create a target tracking scaling policy for your Auto Scaling group, use one of the
following methods.

Before you begin, confirm that your preferred metric is available at 1-minute
intervals (compared to the default 5-minute interval of Amazon EC2 metrics).

Console

###### To create a target tracking scaling policy for a new Auto Scaling group

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2), and choose **Auto Scaling Groups** from the navigation pane.

2. Choose **Create Auto Scaling group**.

3. In Steps 1, 2, and 3, choose the options as desired and
    proceed to **Step 4: Configure group size and**
**scaling policies**.

4. Under **Scaling**, specify the range that
    you want to scale between by updating the **Min**
**desired capacity** and **Max desired**
**capacity**. These two settings allow your Auto Scaling
    group to scale dynamically. For more information, see [Set scaling limits for your Auto Scaling group](asg-capacity-limits.md).

5. Under **Automatic scaling**, choose
    **Target tracking scaling**
**policy**.

6. To define a policy, do the following:
1. Specify a name for the policy.

2. For **Metric type**, choose a
       metric.

      If you chose **Application Load Balancer request count per**
      **target**, choose a target group in
       **Target group**.

3. Specify a **Target value** for
       the metric.

4. (Optional) For **Instance**
      **warmup**, update the instance warmup
       value as needed.

5. (Optional) Select **Disable scale in to**
      **create only a scale-out policy**. This
       allows you to create a separate scale-in policy of a
       different type if wanted.
7. Proceed to create the Auto Scaling group. Your scaling policy will
    be created after the Auto Scaling group has been created.

###### To create a target tracking scaling policy for an existing Auto Scaling group

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2), and choose **Auto Scaling Groups** from the navigation pane.

2. Select the check box next to your Auto Scaling group.

A split pane opens up in the bottom of the page.

3. Verify that the scaling limits are appropriately set. For
    example, if your group's desired capacity is already at its
    maximum, you need to specify a new maximum in order to scale
    out. For more information, see [Set scaling limits for your Auto Scaling group](asg-capacity-limits.md).

4. On the **Automatic scaling** tab, in
    **Dynamic scaling policies**, choose
    **Create dynamic scaling**
**policy**.

5. To define a policy, do the following:
1. For **Policy type**, keep the
       default of **Target tracking**
      **scaling**.

2. Specify a name for the policy.

3. For **Metric type**, choose a
       metric. You can choose only one metric type. To use
       more than one metric, create multiple
       policies.

      If you chose **Application Load Balancer request count per**
      **target**, choose a target group in
       **Target group**.

4. Specify a **Target value** for
       the metric.

5. (Optional) For **Instance**
      **warmup**, update the instance warmup
       value as needed.

6. (Optional) Select **Disable scale in to**
      **create only a scale-out policy**. This
       allows you to create a separate scale-in policy of a
       different type if wanted.
6. Choose **Create**.

AWS CLI

To create a target tracking scaling policy, you can use the
following example to help you get started. Replace each
`user input placeholder` with your own
information.

###### Note

For more examples, see [Example scaling policies for the AWS CLI](https://docs.aws.amazon.com/autoscaling/ec2/userguide/examples-scaling-policies.html).

###### To create a target tracking scaling policy (AWS CLI)

1. Use the following `cat` command to store a
    target value for your scaling policy and a predefined metric
    specification in a JSON file named `config.json`
    in your home directory. The following is an example target
    tracking configuration that keeps the average CPU
    utilization at 50 percent.

```nohighlight

$ cat ~/config.json
{
     "TargetValue": 50.0,
     "PredefinedMetricSpecification":
       {
         "PredefinedMetricType": "ASGAverageCPUUtilization"
       }
}
```

For more information, see [PredefinedMetricSpecification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredefinedMetricSpecification.html) in the
    _Amazon EC2 Auto Scaling API Reference_.

2. Use the [put-scaling-policy](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/autoscaling/put-scaling-policy.html) command, along with the
    `config.json` file that you created in the
    previous step, to create your scaling policy.

```nohighlight

aws autoscaling put-scaling-policy --policy-name cpu50-target-tracking-scaling-policy \
     --auto-scaling-group-name my-asg --policy-type TargetTrackingScaling \
     --target-tracking-configuration file://config.json
```

If successful, this command returns the ARNs and names of
    the two CloudWatch alarms created on your behalf.

```json

{
       "PolicyARN": "arn:aws:autoscaling:us-west-2:123456789012:scalingPolicy:228f02c2-c665-4bfd-aaac-8b04080bea3c:autoScalingGroupName/my-asg:policyName/cpu50-target-tracking-scaling-policy",
       "Alarms": [
           {
               "AlarmARN": "arn:aws:cloudwatch:us-west-2:123456789012:alarm:TargetTracking-my-asg-AlarmHigh-fc0e4183-23ac-497e-9992-691c9980c38e",
               "AlarmName": "TargetTracking-my-asg-AlarmHigh-fc0e4183-23ac-497e-9992-691c9980c38e"
           },
           {
               "AlarmARN": "arn:aws:cloudwatch:us-west-2:123456789012:alarm:TargetTracking-my-asg-AlarmLow-61a39305-ed0c-47af-bd9e-471a352ee1a2",
               "AlarmName": "TargetTracking-my-asg-AlarmLow-61a39305-ed0c-47af-bd9e-471a352ee1a2"
           }
       ]
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Target tracking scaling
policies

Create a policy using high-resolution metrics
