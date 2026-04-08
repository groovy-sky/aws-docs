# Monitoring Amazon Aurora metrics with Amazon CloudWatch

Amazon CloudWatch is a metrics repository. The repository collects and processes raw data from
Amazon Aurora into readable,
near real-time metrics. For a complete list of Amazon Aurora metrics
sent to CloudWatch, see [Metrics reference for Amazon Aurora](../../../en-us/amazonrds/latest/aurorauserguide/metrics-reference.md).

To analyze and troubleshoot the performance of your databases at scale, use [CloudWatch Database Insights](user-databaseinsights.md).

###### Topics

- [Overview of Amazon Aurora and Amazon CloudWatch](#cw-metrics-overview)

- [Viewing DB cluster metrics in the CloudWatch console and AWS CLI](metrics-dimensions.md)

- [Exporting Performance Insights metrics to CloudWatch](pi-metrics-export-cw.md)

- [Creating CloudWatch alarms to monitor Amazon Aurora](creating-alarms.md)

## Overview of Amazon Aurora and Amazon CloudWatch

By default, Amazon Aurora automatically sends metric data to
CloudWatch in 1-minute periods. For example, the `CPUUtilization` metric records the percentage of CPU utilization for a DB instance over
time. Data points with a period of 60 seconds (1 minute) are available for 15 days. This means that you can access historical information and
see how your web application or service is performing.

You can now export Performance Insights metrics dashboards from Amazon RDS
to Amazon CloudWatch. You can export either the preconfigured or customized metrics dashboards as
a new dashboard or add them to an existing CloudWatch dashboard. The exported dashboard is available to
view in the CloudWatch console. For more information on how to export the Performance Insights metrics dashboards to
CloudWatch, see [Exporting Performance Insights metrics to CloudWatch](pi-metrics-export-cw.md).

As shown in the following diagram, you can set up alarms for your CloudWatch metrics. For example, you might create an alarm that signals when the
CPU utilization for an instance is over 70%. You can configure Amazon Simple Notification Service to email you when the threshold is passed.

![RDS metrics in AWS CloudWatch](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/rds-cloudwatch.png)

Amazon RDS publishes the following types of metrics to Amazon CloudWatch:

- Aurora metrics at both the cluster and instance level

For a table of these metrics, see [Amazon CloudWatch metrics for Amazon Aurora](aurora-auroramonitoring-metrics.md).

- Performance Insights metrics

For a table of these metrics, see [Amazon CloudWatch metrics for Amazon RDS Performance Insights](user-perfinsights-cloudwatch.md) and
[Performance Insights counter metrics](user-perfinsights-counters.md).

- Enhanced Monitoring metrics (published to Amazon CloudWatch Logs)

For a table of these metrics, see [OS metrics in Enhanced Monitoring](user-monitoring-available-os-metrics.md).

- Usage metrics for the Amazon RDS service quotas in your AWS account

For a table of these metrics, see [Amazon CloudWatch usage metrics for Amazon Aurora](aurora-auroramonitoring-metrics.md#rds-metrics-usage). For more information about
Amazon RDS quotas, see [Quotas and constraints for Amazon Aurora](chap-limits.md).

For more information about CloudWatch, see [What is Amazon CloudWatch?](../../../amazoncloudwatch/latest/developerguide/whatiscloudwatch.md) in the
_Amazon CloudWatch User Guide_. For more information about CloudWatch metrics retention, see [Metrics retention](../../../amazoncloudwatch/latest/developerguide/cloudwatch-concepts.md#metrics-retention).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Choosing the preconfigured dashboard

Viewing CloudWatch metrics

All content copied from https://docs.aws.amazon.com/.
