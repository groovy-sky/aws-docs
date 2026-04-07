# Amazon Aurora DB engine, Region, and instance class support for Database Insights

The following table provides Amazon Aurora DB engines that support Database Insights.

Amazon Aurora DB engine

Supported engine versions and RegionsInstance class restrictions

Amazon Aurora MySQL-Compatible Edition

For more information on version and Region availability of Database Insights with Aurora MySQL, see
[Performance Insights with Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.html#Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.amy).

Database Insights has the following engine class restrictions:

- db.t2 – Not supported

- db.t3 – Not supported

- db.t4g.micro and db.t4g.small – Not supported

Amazon Aurora PostgreSQL-Compatible Edition

For more information on version and Region availability of Database Insights with Aurora PostgreSQL, see
[Performance Insights with Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.html#Concepts.Aurora_Fea_Regions_DB-eng.Feature.PerfInsights.apg).

Not applicable

Aurora PostgreSQL Limitless Database

For more information about using Database Insights with Aurora PostgreSQL Limitless Database, see
[Monitoring Aurora PostgreSQL Limitless Database with CloudWatch Database Insights](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/limitless-monitoring.cwdbi.html).

Not applicable

Database Insights supports Amazon Aurora Serverless v2.

## Amazon Aurora DB engine, Region, and instance class support for Database Insights features

The following table provides Amazon Aurora DB engines that support Database Insights features.

Feature

[Pricing tier](https://aws.amazon.com/rds/performance-insights/pricing)

[Supported regions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html#Concepts.RegionsAndAvailabilityZones.Regions)

Supported DB engines

[Supported instance classes](concepts-dbinstanceclass.md#Concepts.DBInstanceClass.Types)

[SQL statistics for Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/sql-statistics.html)

All

All

All

All

[Analyzing database performance for a period of time](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.UsingDashboard.AnalyzePerformanceTimePeriod.html)

Paid tier only

All

All

All except db.serverless (Aurora Serverless v2)

[Viewing Performance Insights proactive recommendations](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.InsightsRecommendationViewDetails.html)

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

## Amazon Aurora Region support for Database Insights

Aurora supports Database Insights in the following AWS Regions.

- US East (N. Virginia)

- US East (Ohio)

- US West (N. California)

- US West (Oregon)

- Africa (Cape Town)

- Asia Pacific (Hong Kong)

- Asia Pacific (Hyderabad)

- Asia Pacific (Jakarta)

- Asia Pacific (Malaysia)

- Asia Pacific (Melbourne)

- Asia Pacific (Mumbai)

- Asia Pacific (Osaka)

- Asia Pacific (Seoul)

- Asia Pacific (Singapore)

- Asia Pacific (Sydney)

- Asia Pacific (Tokyo)

- Canada (Central)

- Canada West (Calgary)

- Europe (Frankfurt)

- Europe (Ireland)

- Europe (London)

- Europe (Milan)

- Europe (Paris)

- Europe (Spain)

- Europe (Stockholm)

- Europe (Zurich)

- Israel (Tel Aviv)

- Middle East (Bahrain)

- Middle East (UAE)

- South America (São Paulo)

- AWS GovCloud (US-East)

- AWS GovCloud (US-West)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Monitoring with Database Insights

Turning on the Advanced mode
