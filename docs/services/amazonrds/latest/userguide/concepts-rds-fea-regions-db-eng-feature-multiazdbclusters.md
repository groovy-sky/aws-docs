# Supported Regions and DB engines for Multi-AZ DB clusters in Amazon RDS

A Multi-AZ DB cluster deployment in Amazon RDS provides a high availability deployment mode of Amazon RDS with two
readable standby DB instances. A Multi-AZ DB cluster has a writer DB instance and two reader DB instances in three separate
Availability Zones in the same Region. Multi-AZ DB clusters provide high availability,
increased capacity for read workloads, and lower write latency when compared to Multi-AZ DB instance
deployments. For more information, see [Multi-AZ DB cluster deployments for Amazon RDS](multi-az-db-clusters-concepts.md).

Multi-AZ DB clusters aren't available with the following engines:

- RDS for Db2

- RDS for MariaDB

- RDS for Oracle

- RDS for SQL Server

###### Topics

- [Multi-AZ DB clusters with RDS for MySQL](#Concepts.RDS_Fea_Regions_DB-eng.Feature.MultiAZDBClusters.my)

- [Multi-AZ DB clusters with RDS for PostgreSQL](#Concepts.RDS_Fea_Regions_DB-eng.Feature.MultiAZDBClusters.pg)

## Multi-AZ DB clusters with RDS for MySQL

The following Regions and engine versions are available for Multi-AZ DB clusters with
RDS for MySQL.

RegionRDS for MySQL 8.4RDS for MySQL 8.0US East (N. Virginia)

All available versions

All available versions

US East (Ohio)

All available versions

All available versions

US West (N. California)

Not available

Not available

US West (Oregon)

All available versions

All available versions

Africa (Cape Town)

All available versions

All available versions

Asia Pacific (Hong Kong)

All available versions

All available versions

Asia Pacific (Hyderabad)

All available versions

All available versions

Asia Pacific (Jakarta)

All available versions

All available versions

Asia Pacific (Malaysia)

All available versions

All available versions

Asia Pacific (Melbourne)

All available versions

All available versions

Asia Pacific (Mumbai)

All available versions

All available versions

Asia Pacific (New Zealand)Not availableNot availableAsia Pacific (Osaka)

All available versions

All available versions

Asia Pacific (Seoul)

All available versions

All available versions

Asia Pacific (Singapore)

All available versions

All available versions

Asia Pacific (Sydney)

All available versions

All available versions

Asia Pacific (Taipei)Not availableNot availableAsia Pacific (Thailand)

Not available

Not available

Asia Pacific (Tokyo)

All available versions

All available versions

Canada (Central)

All available versions

All available versions

Canada (Central)

All available versions

All available versions

Canada West (Calgary)

All available versions

All available versions

China (Beijing)

All available versions

All available versions

China (Ningxia)

All available versions

All available versions

Europe (Frankfurt)

All available versions

All available versions

Europe (Ireland)

All available versions

All available versions

Europe (London)

All available versions

All available versions

Europe (Milan)

All available versions

All available versions

Europe (Paris)

All available versions

All available versions

Europe (Spain)

All available versions

All available versions

Europe (Stockholm)

All available versions

All available versions

Europe (Zurich)

All available versions

All available versions

Israel (Tel Aviv)

All available versions

All available versions

Mexico (Central)

All available versions

All available versions

Middle East (Bahrain)

All available versions

All available versions

Middle East (UAE)

All available versions

All available versions

South America (São Paulo)

All available versions

All available versions

AWS GovCloud (US-East)

Not available

Not available

AWS GovCloud (US-West)

Not available

Not available

You can list the available versions in a Region for a given DB instance class
using the AWS CLI. Change the DB instance class to show the available engine versions for it.

For Linux, macOS, or Unix:

```nohighlight

aws rds describe-orderable-db-instance-options \
--engine mysql \
--db-instance-class db.r5d.large \
--query '*[]|[?SupportsClusters == `true`].[EngineVersion]'  \
--output text
```

For Windows:

```nohighlight

aws rds describe-orderable-db-instance-options ^
--engine mysql ^
--db-instance-class db.r5d.large ^
--query "*[]|[?SupportsClusters == `true`].[EngineVersion]"  ^
--output text
```

## Multi-AZ DB clusters with RDS for PostgreSQL

The following Regions and engine versions are available for Multi-AZ DB clusters with
RDS for PostgreSQL.

RegionRDS for PostgreSQL 18RDS for PostgreSQL 17RDS for PostgreSQL 16RDS for PostgreSQL 15RDS for PostgreSQL 14RDS for PostgreSQL 13US East (N. Virginia)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherUS East (Ohio)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherUS West (N. California)Not availableNot availableNot availableNot availableNot availableNot availableUS West (Oregon)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAfrica (Cape Town)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Hong Kong)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Hyderabad)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Jakarta)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Malaysia)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Melbourne)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Mumbai)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (New Zealand)Not availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Osaka)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Seoul)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Singapore)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Sydney)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAsia Pacific (Taipei)Not availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Thailand)Not availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Tokyo)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherCanada (Central)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherCanada West (Calgary)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherChina (Beijing)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherChina (Ningxia)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (Frankfurt)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (Ireland)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (London)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (Milan)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (Paris)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (Spain)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (Stockholm)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherEurope (Zurich)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherIsrael (Tel Aviv)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherMexico (Central)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherMiddle East (Bahrain)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherMiddle East (UAE)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherSouth America (São Paulo)All PostgreSQL 18 versionsAll PostgreSQL 17 versionsAll PostgreSQL 16 versionsAll PostgreSQL 15 versionsVersion 14.5 and higherVersion 13.4 and version 13.7 and higherAWS GovCloud (US-East)Not availableNot availableNot availableNot availableNot availableNot availableAWS GovCloud (US-West)Not availableNot availableNot availableNot availableNot availableNot available

You can list the available versions in a Region for a given DB instance class
using the AWS CLI. Change the DB instance class to show the available engine versions for it.

For Linux, macOS, or Unix:

```nohighlight

aws rds describe-orderable-db-instance-options \
--engine postgres \
--db-instance-class db.r5d.large \
--query '*[]|[?SupportsClusters == `true`].[EngineVersion]'  \
--output text
```

For Windows:

```nohighlight

aws rds describe-orderable-db-instance-options ^
--engine postgres ^
--db-instance-class db.r5d.large ^
--query "*[]|[?SupportsClusters == `true`].[EngineVersion]"  ^
--output text
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Kerberos authentication

Performance Insights
