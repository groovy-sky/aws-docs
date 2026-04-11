# Amazon RDS DB engine, Region, and instance class support for Performance Insights

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

The following table provides Amazon RDS DB engines that support Performance Insights.

###### Note

For Amazon Aurora, see [Amazon Aurora DB engine support for Performance Insights](../aurorauserguide/user-perfinsights-overview-engines.md) in _Amazon Aurora User Guide_.

Amazon RDS DB engine

Supported engine versions and RegionsInstance class restrictions

Amazon RDS for MariaDB

For more information on version and Region availability of Performance Insights with RDS for MariaDB, see
[Supported Regions and DB engines for Performance Insights in Amazon RDS](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md).

Performance Insights isn't supported for the following instance classes:

- db.t2.micro

- db.t2.small

- db.t3.micro

- db.t3.small

- db.t4g.micro

- db.t4g.small

RDS for MySQL

For more information on version and Region availability of Performance Insights with RDS for MySQL, see
[Supported Regions and DB engines for Performance Insights in Amazon RDS](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md).

Performance Insights isn't supported for the following instance classes:

- db.t2.micro

- db.t2.small

- db.t3.micro

- db.t3.small

- db.t4g.micro

- db.t4g.small

Amazon RDS for Microsoft SQL Server

For more information on version and Region availability of Performance Insights with RDS for SQL Server, see
[Supported Regions and DB engines for Performance Insights in Amazon RDS](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md).

N/A

Amazon RDS for PostgreSQL

For more information on version and Region availability of Performance Insights with RDS for PostgreSQL, see
[Supported Regions and DB engines for Performance Insights in Amazon RDS](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md).

N/A

Amazon RDS for Oracle

For more information on version and Region availability of Performance Insights with RDS for Oracle, see
[Supported Regions and DB engines for Performance Insights in Amazon RDS](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md).

N/A

## Amazon RDS DB engine, Region, and instance class support for Performance Insights features

The following table provides Amazon RDS DB engines that support Performance Insights features.

Feature

[Pricing tier](https://aws.amazon.com/rds/performance-insights/pricing)

[Supported regions](concepts-regionsandavailabilityzones.md#Concepts.RegionsAndAvailabilityZones.Regions)

[Supported DB engines](welcome.md#Welcome.Concepts.DBInstance)

[Supported instance classes](concepts-dbinstanceclass.md#Concepts.DBInstanceClass.Types)

[SQL statistics for Performance Insights](sql-statistics.md)

All

All

All

All

[Analyzing Oracle execution plans using the Performance Insights dashboard for Amazon RDS](user-perfinsights-usingdashboard-accessplans.md)

All

All

RDS for Oracle

All

[Analyzing database performance for a period of time](user-perfinsights-usingdashboard-analyzeperformancetimeperiod.md)

Paid tier only

All

RDS for PostgreSQL

All

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

All

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Maximum CPU

Pricing and data retention for Performance Insights

All content copied from https://docs.aws.amazon.com/.
