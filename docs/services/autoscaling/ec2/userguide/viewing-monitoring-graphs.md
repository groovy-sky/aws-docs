# View monitoring graphs in the Amazon EC2 Auto Scaling console

In the Amazon EC2 Auto Scaling section of the Amazon EC2 console, you can monitor minute-by-minute
progress of individual Auto Scaling groups using CloudWatch metrics.

You can monitor the following types of metrics:

- Auto Scaling metrics – Auto Scaling metrics are turned
on only when you enable them. For more information, see [Enable Auto Scaling group metrics (console)](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html#as-enable-group-metrics).
When Auto Scaling metrics are enabled, the monitoring graphs show data published at
one-minute granularity for Auto Scaling metrics.

- EC2 metrics – The Amazon EC2 instance metrics
are always enabled. When detailed monitoring is enabled, the monitoring graphs
show data published at one-minute granularity for instance metrics. For more
information, see [Configure monitoring for Auto Scaling instances](enable-as-instance-metrics.md).

###### To view monitoring graphs using the Amazon EC2 Auto Scaling console

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2), and choose **Auto Scaling Groups** from the navigation pane.

2. Select the check box next to the Auto Scaling group that you want to view metrics for.

A split pane opens up in the bottom part of the **Auto Scaling**
**groups** page.

3. Choose the **Monitoring** tab.

Amazon EC2 Auto Scaling displays monitoring graphs for **Auto Scaling**
    metrics.

4. To view monitoring graphs of the aggregated instance metrics for the group,
    choose **EC2**.

**Graph actions**

- Hover on a data point to view a data pop-up for a specific time in UTC.

- To enlarge a graph, choose **Enlarge** from the menu tool
(the three vertical dots) in the upper right of the graph. Alternatively, choose
the maximize icon at the top of the graph.

- Adjust the time period for data displayed in the graph by selecting one of the
predefined time period values. If the graph is enlarged, you can choose
**Custom** to define your own time period.

- Choose **Refresh** from the menu tool to update the data in a
graph.

- Drag your cursor over the graph data to select a specific range. You can then
choose **Apply time range** in the menu tool.

- Choose **View logs** from the menu tool to view associated
log streams (if any) in the CloudWatch console.

- To view a graph in CloudWatch, choose **View in metrics** from the
menu tool. This takes you to the CloudWatch page for that graph. There, you can view
more information or access historical information to gain a better understanding
of how your Auto Scaling group changed over an extended period.

## Graph metrics for your Auto Scaling groups

After you create an Auto Scaling group, you can open the Amazon EC2 Auto Scaling console and view the
monitoring graphs for the group on the **Monitoring**
tab.

In the **Auto Scaling** section, the graph metrics include the following
metrics. These metrics provide measurements that can be indicators of a potential
issue, such as number of terminating instances or number of pending instances. You
can find definitions for these metrics in [Amazon CloudWatch metrics for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html).

Display nameCloudWatch metric name**Minimum Group Size**`GroupMinSize`**Maximum Group Size**`GroupMaxSize`**Desired Capacity**`GroupDesiredCapacity`**In Service Instances**`GroupInServiceInstances`**Pending Instances**`GroupPendingInstances`**Standby Instances**`GroupStandbyInstances`**Terminating Instances**`GroupTerminatingInstances`**Total Instances**`GroupTotalInstances`

In the **EC2** section, you can find the following graph metrics
based on key performance metrics for your Amazon EC2 instances. These EC2 metrics are an
aggregate of metrics for all instances in the group. You can find definitions for
these metrics in [List the\
available CloudWatch metrics for your instances](../../../ec2/latest/userguide/viewing-metrics-with-cloudwatch.md) in the
_Amazon EC2 User Guide_.

Display nameCloudWatch metric name**CPU Utilization**`CPUUtilization`**Disk Reads**`DiskReadBytes`**Disk Read Operations**`DiskReadOps`**Disk Writes**`DiskWriteBytes`**Disk Write Operations**`DiskWriteOps`**Network In**`NetworkIn`**Network Out**`NetworkOut`**Status Check Failed (Any)**`StatusCheckFailed`**Status Check Failed (Instance)**`StatusCheckFailed_Instance`**Status Check Failed (System)**`StatusCheckFailed_System`

In addition, some metrics are available for specific use cases in the
**Auto Scaling** graph metrics.

The following metrics are useful if your group uses weights that define how many
units each instance contributes to the desired capacity of the group. You can find
definitions for these metrics in [Amazon CloudWatch metrics for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html).

Display nameCloudWatch metric name**In Service Capacity Units**`GroupInServiceCapacity`**Pending Capacity Units**`GroupPendingCapacity`**Standby Capacity Units**`GroupStandbyCapacity`**Terminating Capacity Units**`GroupTerminatingCapacity`**Total Capacity Units**`GroupTotalCapacity`

The following metrics are useful if your group uses the [warm pool](ec2-auto-scaling-warm-pools.md) feature. You can find
definitions for these metrics in [Amazon CloudWatch metrics for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html).

Display nameCloudWatch metric name**Warm Pool Minimum Size**`WarmPoolMinSize`**Warm Pool Desired Capacity**`WarmPoolDesiredCapacity`**Warm Pool Pending Capacity Units**`WarmPoolPendingCapacity`**Warm Pool Terminating Capacity Units**`WarmPoolTerminatingCapacity`**Warm Pool Warmed Capacity Units**`WarmPoolWarmedCapacity`**Warm Pool Total Capacity Units**
**Launched**`WarmPoolTotalCapacity`**Group and Warm Pool Desired Capacity**`GroupAndWarmPoolDesiredCapacity`**Group and Warm Pool Total Capacity Units**
**Launched**`GroupAndWarmPoolTotalCapacity`

### Related resources

- To monitor per-instance metrics, see [Graph metrics for your instances](../../../ec2/latest/userguide/graphs-in-the-aws-management-console.md) in the
_Amazon EC2 User Guide_.

- CloudWatch dashboards are customizable home pages in the CloudWatch console. You
can use these pages to monitor your resources in a single view, even
including resources that are spread across different Regions. You can
use CloudWatch dashboards to create customized views of the metrics and alarms
for your AWS resources. For more information, see the
[Amazon CloudWatch User Guide](../../../amazoncloudwatch/latest/monitoring.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Monitor CloudWatch metrics

CloudWatch metrics for Amazon EC2 Auto Scaling
