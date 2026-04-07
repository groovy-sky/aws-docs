# Amazon Aurora DB engine, Region, and instance class support for Performance Insights

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
For more information about Database Insights, see [Monitoring Amazon Aurora databases with CloudWatch Database Insights](user-databaseinsights.md).
For pricing information, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

The following table provides Amazon Aurora DB engines that support Performance Insights.

Amazon Aurora DB engine

Supported engine versions and RegionsInstance class restrictions

Amazon Aurora MySQL-Compatible Edition

For more information on version and Region availability of Performance Insights with Aurora MySQL, see
[Performance Insights with Aurora MySQL](concepts-aurora-fea-regions-db-eng-feature-perfinsights.md#Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.amy).

Performance Insights has the following engine class restrictions:

- db.t2 – Not supported

- db.t3 – Not supported

- db.t4g.micro and db.t4g.small – Not supported

Amazon Aurora PostgreSQL-Compatible Edition

For more information on version and Region availability of Performance Insights with Aurora PostgreSQL, see
[Performance Insights with Aurora PostgreSQL](concepts-aurora-fea-regions-db-eng-feature-perfinsights.md#Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.apg).

N/A

## Amazon Aurora DB engine, Region, and instance class support for Performance Insights features

The following table provides Amazon Aurora DB engines that support Performance Insights features.

Feature

[Pricing tier](https://aws.amazon.com/rds/performance-insights/pricing)

[Supported regions](concepts-regionsandavailabilityzones.md#Concepts.RegionsAndAvailabilityZones.Regions)

Supported DB engines

[Supported instance classes](concepts-dbinstanceclass.md#Concepts.DBInstanceClass.Types)

[SQL statistics for Performance Insights](sql-statistics.md)

All

All

All

All

[Analyzing database performance for a period of time](user-perfinsights-usingdashboard-analyzeperformancetimeperiod.md)

Paid tier only

All

All

All except db.serverless (Aurora Serverless v2)

[Viewing Performance Insights proactive recommendations](user-perfinsights-insightsrecommendationviewdetails.md)

Paid tier only

- US East (Ohio)

- US East (N. Virginia)

- US West (N. California)

- US West (Oregon)

- Asia Pacific (Mumbai)

- Asia Pacific (Seoul)

- Asia Pacific (Singapore)

- Asia Pacific (Sydney)

- Asia Pacific (Tokyo)

- Canada (Central)

- Europe (Frankfurt)

- Europe (Ireland)

- Europe (London)

- Europe (Paris)

- Europe (Stockholm)

- South America (São Paulo)

All

All except db.serverless (Aurora Serverless v2)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Maximum CPU

Pricing and data retention for Performance Insights
