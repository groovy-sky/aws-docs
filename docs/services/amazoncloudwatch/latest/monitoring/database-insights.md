# CloudWatch Database Insights

Use CloudWatch Database Insights to monitor and troubleshoot Amazon Aurora MySQL, Amazon Aurora PostgreSQL, Amazon Aurora PostgreSQL Limitless, Amazon RDS for SQL Server, RDS for MySQL, RDS for PostgreSQL, RDS for Oracle, and RDS for MariaDB databases at scale.

With Database Insights, you can monitor your database fleet with pre-built, opinionated dashboards. To help you analyze the performance of your fleet, the Database Insights dashboards display curated metrics and visualizations, and you
can customize these dashboards.
By presenting metrics in a single dashboard for all databases in your fleet, Database Insights allows you to monitor your databases simultaneously.

For example, you can use Database Insights to find a database that is performing poorly within a fleet of hundreds of database instances. You can then choose that instance and use Database Insights to troubleshoot issues.

For information about engine, AWS Region, and instance class support, see
[Aurora DB engine, Region, and instance class support for Database Insights](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_DatabaseInsights.Engines.html)
and [Amazon RDS DB engine, Region, and instance class support for Database Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DatabaseInsights.Engines.html).

Database Insights supports monitoring workloads across multiple accounts and regions.
To learn more about the cross-account cross-region monitoring feauture of Database Insights see
[Set up cross-account cross-region monitoring for CloudWatch Database Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Database-Insights-Cross-Account-Cross-Region.html)

To get started with Database Insights, see the following topics.

###### Topics

- [Get started with CloudWatch Database Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Database-Insights-Get-Started.html)

- [Viewing the Fleet Health Dashboard for CloudWatch Database Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Database-Insights-Fleet-Health-Dashboard.html)

- [Viewing the Database Instance Dashboard for CloudWatch Database Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Database-Insights-Database-Instance-Dashboard.html)

- [Troubleshooting for CloudWatch Database Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Database-Insights-Troubleshooting.html)

## Modes for Database Insights

Database Insights has an Advanced mode and a Standard mode. Standard mode is the default for Database Insights, and you can turn on the Advanced mode for your database.

The following table shows which features CloudWatch supports for the Advanced mode and Standard mode of Database Insights.

FeatureStandard modeAdvanced modeAnalyze the top contributors to DB Load by dimensionSupportedSupportedQuery, graph, and set alarms on database metrics with up to 7 days of retentionSupportedSupportedDefine fine‐grained access control policies to restrict access to potentially sensitive dimensions such as SQL textSupportedSupported

Analyze operating system processes happening in your databases with detailed metrics per running process

[Amazon RDS\
Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights_Counters.html) is required for this feature to work.

Not supportedSupportedDefine and save fleet‐wide monitoring views to assess database health at scaleNot supportedSupportedAnalyze SQL locks with 15 months of retention and a guided UXNot supportedSupported only for Aurora PostgreSQLAnalyze SQL execution plans with 15 months of retention and guided UXNot supportedSupported only for Aurora PostgreSQL, RDS for Oracle, and RDS for SQL ServerVisualize per‐query statisticsNot supportedSupportedAnalyze slow SQL queries

Export of database logs to CloudWatch Logs is required for
this feature to work.

Not supportedSupportedView calling services with CloudWatch Application SignalsNot supportedSupportedView a consolidated dashboard for all database telemetry, including metrics,
logs, events, and applications

Export of database logs to CloudWatch Logs is required to
view database logs in the Database Insights console.

Not supportedSupportedImport Performance Insights counter metrics into CloudWatch automaticallyNot supportedSupportedView Amazon RDS events in CloudWatchNot supportedSupportedAnalyze database performance for a time period of your choice with on‐demand analysisNot supportedSupported

###### Note

Database Insights feature availability differs in different AWS Regions, because not all Advanced Mode features are available in all Regions.

## Data retention

The Advanced mode of Database Insights retains 15 months of metrics collected by Performance Insights.

If Performance Insights is enabled for the Standard mode, Amazon RDS retains 7 days of Performance Insights counter metrics.

For information about counter metrics for Performance Insights, see [Performance Insights counter metrics](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights_Counters.html).

For information about the retention period for CloudWatch metrics collected by Database Insights, see the following topics.

- [Amazon CloudWatch metrics for Amazon Aurora](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.AuroraMonitoring.Metrics.html) in the _Amazon Aurora User Guide_

- [Amazon CloudWatch metrics for Amazon Relational Database Service](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-metrics.html) in the _Amazon RDS User Guide_

- [Amazon CloudWatch metrics for Amazon RDS Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.Cloudwatch.html) in the _Amazon Aurora User Guide_

- [Amazon CloudWatch metrics for Amazon RDS Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.Cloudwatch.html) in the _Amazon Aurora User Guide_

## How Database Insights integrates with Performance Insights

Performance Insights is a database performance monitoring service.

Database Insights builds upon and extends the capabilities of Performance Insights. Database Insights adds monitoring, analysis, and optimization features.

To enable the Advanced mode of Database Insights, you must enable Performance Insights.

Database Insights imports Performance Insights counter metrics into CloudWatch automatically. The Advanced mode of Database Insights automatically retains 15 months of all metrics collected by Database Insights,
including Performance Insights metrics and CloudWatch metrics. This automatically happens for you when you enable Advanced mode in an instance, with no further configuration needed.
For information about Performance Insights counter metrics, see [Performance Insights counter metrics](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights_Counters.html)
in the _Amazon Aurora User Guide_.

## Pricing

For information about pricing, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Example telemetry event

Get started
