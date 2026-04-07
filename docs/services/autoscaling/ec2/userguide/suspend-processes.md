# Suspend processes

To suspend a process for an Auto Scaling group, use one of the following methods:

Console

###### To suspend a process

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2), and choose **Auto Scaling Groups** from the navigation pane.

2. Select the check box next to the Auto Scaling group.

A split pane opens up in the bottom of the page.

3. On the **Details** tab, choose
    **Advanced configurations**,
    **Edit**.

4. For **Suspended processes**, choose the
    process to suspend.

5. Choose **Update**.

AWS CLI

Use the following [suspend-processes](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/autoscaling/suspend-processes.html) command to suspend individual
processes.

```nohighlight

aws autoscaling suspend-processes --auto-scaling-group-name my-asg --scaling-processes HealthCheck ReplaceUnhealthy
```

To suspend all processes, omit the `--scaling-processes`
option, as follows.

```nohighlight

aws autoscaling suspend-processes --auto-scaling-group-name my-asg
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Considerations

Resume processes
