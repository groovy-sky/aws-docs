# Create an alarm from a metric on a graph

You can graph a metric and then create an alarm from the metric on the graph, which has
the benefit of populating many of the alarm fields for you.

###### To create an alarm from a metric on a graph

01. Open the CloudWatch console at
     [https://console.aws.amazon.com/cloudwatch/](https://console.aws.amazon.com/cloudwatch).

02. In the navigation pane, choose **Metrics**, **All**
    **metrics**.

03. Select a metric namespace (for example, **EC2**) and then a metric
     dimension (for example, **Per-Instance Metrics**).

04. The **All metrics** tab displays all metrics for that dimension in
     that namespace. To graph a metric, select the check box next to the metric.

05. To create an alarm for the metric, choose the **Graphed metrics**
     tab. For **Actions**, choose the alarm icon.

    ![Create an alarm from a graphed metric](https://docs.aws.amazon.com/images/AmazonCloudWatch/latest/monitoring/images/metric_graph_alarm.png)

06. Under **Conditions**, choose **Static** or
     **Anomaly detection** to specify whether to use a static threshold or
     anomaly detection model for the alarm.

    Depending on your choice, enter the rest of the data for the alarm
     conditions.

07. Choose **Additional configuration**. For **Datapoints to**
    **alarm**, specify how many evaluation periods (data points) must be in the
     `ALARM` state to trigger the alarm. If the two values here match, you
     create an alarm that goes to `ALARM` state if that many consecutive periods
     are breaching.

    To create an M out of N alarm, specify a lower number for the first value than you
     specify for the second value. For more information, see [Alarm evaluation](alarm-evaluation.md).

08. For **Missing data treatment**, choose how to have the alarm behave
     when some data points are missing. For more information, see [Configuring how CloudWatch alarms treat missing data](alarms-and-missing-data.md).

09. Choose **Next**.

10. Under **Notification**, select an SNS topic to notify when the
     alarm is in `ALARM` state, `OK` state, or
     `INSUFFICIENT_DATA` state.

    To have the alarm send multiple notifications for the same alarm state or for
     different alarm states, choose **Add notification**.

    To have the alarm not send notifications, choose **Remove**.

11. To have the alarm perform Auto Scaling or EC2 actions, choose the appropriate button
     and choose the alarm state and action to perform.

12. When finished, choose **Next**.

13. Enter a name and description for the alarm. The name must contain only ASCII
     characters. Then choose **Next**.

14. Under **Preview and create**, confirm that the information and
     conditions are what you want, then choose **Create alarm**.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Modify the y-axis for a graph

Using anomaly detection

All content copied from https://docs.aws.amazon.com/.
