# Resume processes

To resume a suspended process for an Auto Scaling group, use one of the following
methods:

Console

###### To resume a suspended process

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2), and choose **Auto Scaling Groups** from the navigation pane.

2. Select the check box next to the Auto Scaling group.

A split pane opens up in the bottom of the page.

3. On the **Details** tab, choose
    **Advanced configurations**,
    **Edit**.

4. For **Suspended processes**, remove the
    suspended process.

5. Choose **Update**.

AWS CLI

To resume a suspended process, use the following [resume-processes](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/autoscaling/resume-processes.html) command.

```nohighlight

aws autoscaling resume-processes --auto-scaling-group-name my-asg --scaling-processes HealthCheck
```

To resume all suspended processes, omit the
`--scaling-processes` option, as follows.

```nohighlight

aws autoscaling resume-processes --auto-scaling-group-name my-asg
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Suspend processes

How suspended processes affect other processes
