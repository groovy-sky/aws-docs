# Creating CloudWatch alarms to monitor Amazon RDS

You can create a CloudWatch alarm that sends an Amazon SNS message when the alarm changes state. An alarm watches a single metric over a time period that you specify. The alarm can also perform one or
more actions based on the value of the metric relative to a given threshold over a number of time periods. The action is a notification sent to an Amazon SNS topic or Amazon EC2 Auto Scaling policy.

Alarms invoke actions for sustained state changes only. CloudWatch alarms don't invoke actions simply because they are in a particular state. The state must have changed and have been
maintained for a specified number of time periods.

You can use the **DB\_PERF\_INSIGHTS**
metric math function in the CloudWatch console to query Amazon RDS for Performance Insights counter metrics. The
**DB\_PERF\_INSIGHTS** function also includes the DBLoad metric at
sub-minute intervals. You can set CloudWatch alarms on these metrics.

For more details on how to create an alarm, see
[Create an alarm on Performance Insights counter metrics from an AWS database](../../../amazoncloudwatch/latest/monitoring/cloudwatch-alarm-database-performance-insights.md).

###### To set an alarm using the AWS CLI

- Call [`put-metric-alarm`](../../../cli/latest/reference/cloudwatch/put-metric-alarm.md). For more information, see _[AWS CLI Command Reference](../../../cli/latest/reference.md)_.

###### To set an alarm using the CloudWatch API

- Call [`PutMetricAlarm`](../../../../reference/amazoncloudwatch/latest/apireference/api-putmetricalarm.md). For more information, see _[Amazon CloudWatch API Reference](../../../../reference/amazoncloudwatch/latest/apireference.md)_

For more information about setting up Amazon SNS topics and creating alarms, see [Using Amazon CloudWatch\
alarms](../../../amazoncloudwatch/latest/monitoring/alarmthatsendsemail.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Viewing a Performance Insights metric widget in CloudWatch

Tutorial: Creating a CloudWatch alarm for DB cluster replica lag

All content copied from https://docs.aws.amazon.com/.
