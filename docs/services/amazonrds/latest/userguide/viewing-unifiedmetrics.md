# Viewing combined metrics with the Performance Insights dashboard

###### Important

AWS has announced the end-of-life date for Performance Insights: June 30, 2026. After this date, Amazon RDS will no longer support the Performance Insights console experience,
flexible retention periods (1-24 months), and their associated pricing. The Performance Insights API will continue to exist with no pricing changes. Costs for the
Performance Insights API will appear in your AWS bill with the cost of CloudWatch Database Insights.

We recommend that you upgrade any DB instances
using the paid tier of Performance Insights to the Advanced mode of Database Insights before June 30, 2026.
For information about upgrading to the Advanced mode of Database Insights, see
[Turning on the Advanced mode of Database Insights for Amazon RDS](user-databaseinsights-turningonadvanced.md).

If you take no action, DB instances using Performance Insights
will default to using the Standard mode of Database Insights. With Standard mode of Database Insights, you might lose access to performance data history beyond 7 days and might not be able to use execution plans
and on-demand analysis features in the Amazon RDS console. After June 30, 2026 only the Advanced mode of Database Insights will support execution plans and on-demand analysis.

With CloudWatch Database Insights, you can monitor database load for your fleet of databases and analyze and troubleshoot performance at scale.
For more information about Database Insights, see [Monitoring Amazon RDS databases with CloudWatch Database Insights](user-databaseinsights.md).
For pricing information, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

Amazon RDS provides a consolidated view of Performance Insights and CloudWatch metrics for your DB instance in the
Performance Insights dashboard. You can use the preconfigured dashboard or create a custom dashboard. The
preconfigured dashboard provides the most commonly used metrics to help diagnose performance
issues for a database engine. Alternatively, you can create a custom dashboard with the
metrics for a database engine that meet your analysis requirements. Then, use this dashboard
for all the DB instances of that database engine type in your AWS account.

You can choose the monitoring view in the **Monitoring** tab or
**Performance Insights** in the navigation pane.

Performance Insights must be turned on for your DB instance to view the combined metrics in the Performance Insights dashboard.
For more information about turning on Performance Insights, see [Turning Performance Insights on and off for Amazon RDS](user-perfinsights-enabling.md).

In the following sections, you can learn to display Performance Insights and CloudWatch metrics.

###### Topics

- [Choosing the new monitoring view from the Monitoring tab](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Viewing_Unifiedmetrics.MonitoringTab.html)

- [Choosing the new monitoring view from the Performance Insights page](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Viewing_Unifiedmetrics.PInavigationPane.html)

- [Creating a custom dashboard with Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Viewing_Unifiedmetrics.PIcustomizeMetricslist.html)

- [Choosing the preconfigured dashboard with Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Viewing_Unifiedmetrics.PI-preconfigured-dashboard.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Viewing metrics in the Amazon RDS console

Choosing the new monitoring view from the Monitoring tab
