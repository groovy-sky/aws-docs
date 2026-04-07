# Analyzing metrics with the Performance Insights dashboard

###### Important

AWS has announced the end-of-life date for Performance Insights: June 30, 2026. After this date, Amazon RDS will no longer support the Performance Insights console experience,
flexible retention periods (1-24 months), and their associated pricing. The Performance Insights API will continue to exist with no pricing changes. Costs for the
Performance Insights API will appear in your AWS bill with the cost of CloudWatch Database Insights.

We recommend that you upgrade any DB clusters
using the paid tier of Performance Insights to the Advanced mode of Database Insights before June 30, 2026.
For information about upgrading to the Advanced mode of Database Insights, see
[Turning on the Advanced mode of Database Insights for Amazon Aurora](user-databaseinsights-turningonadvanced.md).

If you take no action, DB clusters using Performance Insights
will default to using the Standard mode of Database Insights. With Standard mode of Database Insights, you might lose access to performance data history beyond 7 days and might not be able to use execution plans
and on-demand analysis features in the Amazon RDS console. After June 30, 2026 only the Advanced mode of Database Insights will support execution plans and on-demand analysis.

With CloudWatch Database Insights, you can monitor database load for your fleet of databases and analyze and troubleshoot performance at scale.
For more information about Database Insights, see [Monitoring Amazon Aurora databases with CloudWatch Database Insights](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_DatabaseInsights.html).
For pricing information, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

The Performance Insights dashboard contains database performance information to help you analyze and troubleshoot
performance issues. On the main dashboard page, you can view information about the database load. You can "slice" DB
load by dimensions such as wait events or SQL.

###### Performance Insights dashboard

- [Overview of the Performance Insights dashboard](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.UsingDashboard.Components.html)

- [Accessing the Performance Insights dashboard](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.UsingDashboard.Opening.html)

- [Analyzing DB load by wait events](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.html)

- [Analyzing database performance for a period of time](user-perfinsights-usingdashboard-analyzeperformancetimeperiod.md)

- [Analyzing queries with the Top SQL tab in Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using tag-based access control for Performance Insights

Overview of the dashboard
