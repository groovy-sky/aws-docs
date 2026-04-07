# Supported Regions and Aurora DB engines for Performance Insights

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

Performance Insights expands on existing Amazon RDS monitoring features to illustrate
and help you analyze your database performance. With the Performance Insights dashboard,
you can visualize the database load on your Amazon RDS DB instance load and filter the
load by waits, SQL statements, hosts, or users. For more information, see [Overview of Performance Insights on Amazon Aurora](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.Overview.html).

For the region, DB engine, and instance class support information for Performance Insights features,
see [Amazon Aurora DB engine, Region, and instance class support for Performance Insights features](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.Overview.Engines.html#USER_PerfInsights.Overview.PIfeatureEngnRegSupport).

###### Topics

- [Performance Insights with Aurora MySQL](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.amy)

- [Performance Insights with Aurora PostgreSQL](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.apg)

- [Performance Insights with Aurora Serverless](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.serverless)

## Performance Insights with Aurora MySQL

###### Note

Engine version support is different for Performance Insights with Aurora MySQL
if you have parallel query turned on. For more information on parallel query,
see [Parallel query for Amazon Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html).

###### Topics

- [Performance Insights with Aurora MySQL and parallel query turned off](#Feature.PerfInsights.regions.amy.pq)

- [Performance Insights with Aurora MySQL and parallel query turned on](#Feature.PerfInsights.regions.amy.pqoff)

### Performance Insights with Aurora MySQL and parallel query turned off

The following Regions and engine versions are available for Performance
Insights with Aurora MySQL and parallel query turned off.

RegionAurora MySQL version 3Aurora MySQL version 2US East (N. Virginia)All versionsAll versionsUS East (Ohio)All versionsAll versionsUS West (N. California)All versionsAll versionsUS West (Oregon)All versionsAll versionsAfrica (Cape Town)All versionsAll versionsAsia Pacific (Hong Kong)All versionsAll versionsAsia Pacific (Hyderabad)All versionsAll versionsAsia Pacific (Jakarta)All versionsAll versionsAsia Pacific (Malaysia)All versionsAll versionsAsia Pacific (Melbourne)All versionsAll versionsAsia Pacific (Mumbai)All versionsAll versionsAsia Pacific (New Zealand)All versionsAll versionsAsia Pacific (Osaka)All versionsAll versionsAsia Pacific (Seoul)All versionsAll versionsAsia Pacific (Singapore)All versionsAll versionsAsia Pacific (Sydney)All versionsAll versionsAsia Pacific (Taipei)All versionsAll versionsAsia Pacific (Thailand)All versionsAll versionsAsia Pacific (Tokyo)All versionsAll versionsCanada (Central)All versionsAll versionsCanada West (Calgary)All versionsAll versionsChina (Beijing)All versionsAll versionsChina (Ningxia)All versionsAll versionsEurope (Frankfurt)All versionsAll versionsEurope (Ireland)All versionsAll versionsEurope (London)All versionsAll versionsEurope (Milan)All versionsAll versionsEurope (Paris)All versionsAll versionsEurope (Spain)All versionsAll versionsEurope (Stockholm)All versionsAll versionsEurope (Zurich)All versionsAll versionsIsrael (Tel Aviv)All versionsAll versionsMexico (Central)All versionsAll versionsMiddle East (Bahrain)All versionsAll versionsMiddle East (UAE)All versionsAll versionsSouth America (São Paulo)All versionsAll versionsAWS GovCloud (US-East)All versionsAll versionsAWS GovCloud (US-West)All versionsAll versions

### Performance Insights with Aurora MySQL and parallel query turned on

The following Regions and engine versions are available for Performance
Insights with Aurora MySQL and parallel query turned on.

RegionAurora MySQL version 3Aurora MySQL version 2US East (N. Virginia)Not availableVersion 2.09.0 and higherUS East (Ohio)Not availableVersion 2.09.0 and higherUS West (N. California)Not availableVersion 2.09.0 and higherUS West (Oregon)Not availableVersion 2.09.0 and higherAfrica (Cape Town)Not availableVersion 2.09.0 and higherAsia Pacific (Hong Kong)Not availableVersion 2.09.0 and higherAsia Pacific (Hyderabad)Not availableAll versionsAsia Pacific (Jakarta)Not availableVersion 2.09.0 and higherAsia Pacific (Malaysia)Not availableVersion 2.09.0 and higherAsia Pacific (Melbourne)Not availableVersion 2.09.0 and higherAsia Pacific (Mumbai)Not availableVersion 2.09.0 and higherAsia Pacific (New Zealand)Not availableVersion 2.09.0 and higherAsia Pacific (Osaka)Not availableVersion 2.09.0 and higherAsia Pacific (Seoul)Not availableVersion 2.09.0 and higherAsia Pacific (Singapore)Not availableVersion 2.09.0 and higherAsia Pacific (Sydney)Not availableVersion 2.09.0 and higherAsia Pacific (Taipei)Not availableVersion 2.09.0 and higherAsia Pacific (Thailand)Not availableVersion 2.09.0 and higherAsia Pacific (Tokyo)Not availableVersion 2.09.0 and higherCanada (Central)Not availableVersion 2.09.0 and higherCanada West (Calgary)Not availableVersion 2.09.0 and higherChina (Beijing)Not availableVersion 2.09.0 and higherChina (Ningxia)Not availableVersion 2.09.0 and higherEurope (Frankfurt)Not availableVersion 2.09.0 and higherEurope (Ireland)Not availableVersion 2.09.0 and higherEurope (London)Not availableVersion 2.09.0 and higherEurope (Milan)Not availableVersion 2.09.0 and higherEurope (Paris)Not availableVersion 2.09.0 and higherEurope (Spain)Not availableVersion 2.09.0 and higherEurope (Stockholm)Not availableVersion 2.09.0 and higherEurope (Zurich)Not availableVersion 2.09.0 and higherIsrael (Tel Aviv)Not availableVersion 2.09.0 and higherMexico (Central)Not availableVersion 2.09.0 and higherMiddle East (Bahrain)Not availableVersion 2.09.0 and higherMiddle East (UAE)Not availableVersion 2.09.0 and higherSouth America (São Paulo)Not availableVersion 2.09.0 and higherAWS GovCloud (US-East)Not availableVersion 2.09.0 and higherAWS GovCloud (US-West)Not availableVersion 2.09.0 and higher

## Performance Insights with Aurora PostgreSQL

The following Regions and engine versions are available for Performance Insights
with Aurora PostgreSQL.

RegionAurora PostgreSQL 17Aurora PostgreSQL 16Aurora PostgreSQL 15Aurora PostgreSQL 14Aurora PostgreSQL 13Aurora PostgreSQL 12Aurora PostgreSQL 11Aurora PostgreSQL 10US East (N. Virginia)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsUS East (Ohio)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsUS West (N. California)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsUS West (Oregon)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAfrica (Cape Town)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Hong Kong)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Hyderabad)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Jakarta)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Malaysia)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Melbourne)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Mumbai)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (New Zealand)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Osaka)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Seoul)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Singapore)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Sydney)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Taipei)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Thailand)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Tokyo)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsCanada (Central)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsCanada West (Calgary)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsChina (Beijing)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsChina (Ningxia)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Frankfurt)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Ireland)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (London)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Milan)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Paris)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Spain)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Stockholm)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Zurich)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsIsrael (Tel Aviv)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsMexico (Central)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsMiddle East (Bahrain)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsMiddle East (UAE)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsSouth America (São Paulo)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAWS GovCloud (US-East)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAWS GovCloud (US-West)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versions

## Performance Insights with Aurora Serverless

Aurora Serverless v2 supports Performance Insights for all MySQL-compatible and
PostgreSQL-compatible versions. We recommend that you set the minimum capacity to at
least 2 Aurora capacity units (ACUs).

Aurora Serverless v1 doesn't support Performance Insights.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Aurora machine learning

Zero-ETL integrations
