---
title: "Process flow log records in CloudWatch Logs"
---

# Process flow log records in CloudWatch Logs

You can process flow log records as you would with any other log events
collected by CloudWatch Logs. For more information about monitoring log data and metric
filters, see [Creating metrics \
from log events using filter](../../../amazoncloudwatch/latest/logs/monitoringlogdata.md) in the _Amazon CloudWatch Logs User Guide_.

## Example: Create a CloudWatch metric filter and alarm for a flow log

In this example, you have a flow log for `eni-1a2b3c4d`. You want
to create an alarm that alerts you if there have been 10 or more rejected
attempts to connect to your instance over TCP port 22 (SSH) within a 1-hour time
period. First, you must create a metric filter that matches the pattern of the
traffic for which to create the alarm. Then, you can create an alarm for the
metric filter.

###### To create a metric filter for rejected SSH traffic and create an alarm for the filter

01. Open the CloudWatch console at
     [https://console.aws.amazon.com/cloudwatch/](https://console.aws.amazon.com/cloudwatch).

02. In the navigation pane, choose **Logs**,
     **Log groups**.

03. Select the check box for the log group, and then choose
     **Actions**, **Create metric**
    **filter**.

04. For **Filter pattern**, enter the following string.

    ```nohighlight

    [version, account, eni, source, destination, srcport, destport="22", protocol="6", packets, bytes, windowstart, windowend, action="REJECT", flowlogstatus]
    ```

05. For **Select log data to test**, select the log
     stream for your network interface. (Optional) To view the lines of log
     data that match the filter pattern, choose **Test**
    **pattern**.

06. When you're ready, choose **Next**.

07. Enter a filter name, metric namespace, and metric name. Set the metric
     value to 1. When you're done, choose **Next** and then
     choose **Create metric filter**.

08. In the navigation pane, choose **Alarms**,
     **All alarms**.

09. Choose **Create alarm**.

10. Select the metric name that you created and then choose
     **Select metric**.

11. Configure the alarm as follows, and then choose **Next**:

- For **Statistic**, choose **Sum**. This
ensure that you capture the total number of data points for the specified
time period.

- For **Period**, choose **1 hour**.

- For **Whenever TimeSinceLastActive is...**, choose **Greater/Equal**
and enter 10 for the threshold.

- For **Additional configuration**, **Datapoints**
**to alarm**, leave the default of 1.

12. Choose **Next**.

13. For **Notification**, select an existing SNS topic or choose
     **Create new topic** to create a new one. Choose **Next**.

14. Enter a name and description for the alarm and choose **Next**.

15. When you are done previewing the alarm, choose **Create**
    **alarm**.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Search flow log records

Publish to Amazon S3

All content copied from https://docs.aws.amazon.com/.
