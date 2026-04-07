# Aurora Serverless v1

###### Important

AWS has announced the end-of-life date for Aurora Serverless v1: March 31st,
2025\. We strongly recommend upgrading any Aurora Serverless v1 DB clusters to
Aurora Serverless v2 before that date. The upgrade can involve a change in the major
version number of the database engine. Thus, it's important to plan, test, and
implement this switchover before the end-of-life date. Starting January 8th, 2025,
customers will no longer be able to create new Aurora Serverless v1 clusters or
instances with either the AWS Management Console or the CLI. For information about the migration
process, see [Upgrading from an Aurora Serverless v1 cluster to Aurora Serverless v2](aurora-serverless-v2-upgrade.md#aurora-serverless-v2.upgrade-from-serverless-v1-procedure).

Aurora Serverless v2 scales more quickly and in a more granular way.
Aurora Serverless v2 also has more compatibility with other Aurora features such as
reader DB instances. You can learn about Aurora Serverless v2 in [Using Aurora Serverless v2](aurora-serverless-v2.md).

Aurora Serverless v1 is an on-demand, auto-scaling feature designed to be a
cost-effective approach to running intermittent or unpredictable workloads on Amazon Aurora.
It automatically starts up, shuts down, and scales capacity up or down, as needed by
your applications, using a single DB instance in each cluster. For more information, see
[Using Amazon Aurora Serverless v1](aurora-serverless.md).

###### Topics

- [Aurora Serverless v1 with Aurora MySQL](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.ServerlessV1.amy)

- [Aurora Serverless v1 with Aurora PostgreSQL](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.ServerlessV1.apg)

## Aurora Serverless v1 with Aurora MySQL

The following Regions and engine versions are available for Aurora Serverless v1
with Aurora MySQL.

RegionAurora MySQL version 3Aurora MySQL version 2US East (N. Virginia)Not availableVersion 2.11.4US East (Ohio)Not availableVersion 2.11.4US West (N. California)Not availableVersion 2.11.4US West (Oregon)Not availableVersion 2.11.4Africa (Cape Town)Not availableNot availableAsia Pacific (Hong Kong)Not availableNot availableAsia Pacific (Hyderabad)Not availableNot availableAsia Pacific (Jakarta)Not availableNot availableAsia Pacific (Malaysia)Not availableNot availableAsia Pacific (Melbourne)Not availableNot availableAsia Pacific (Mumbai)Not availableVersion 2.11.4Asia Pacific (Osaka)Not availableNot availableAsia Pacific (Seoul)Not availableVersion 2.11.4Asia Pacific (Singapore)Not availableVersion 2.11.4Asia Pacific (Sydney)Not availableVersion 2.11.4Asia Pacific (Thailand)Not availableNot availableAsia Pacific (Tokyo)Not availableVersion 2.11.4Canada (Central)Not availableVersion 2.11.4Canada West (Calgary)Not availableNot availableChina (Beijing)Not availableNot availableChina (Ningxia)Not availableVersion 2.11.4Europe (Frankfurt)Not availableVersion 2.11.4Europe (Ireland)Not availableVersion 2.11.4Europe (London)Not availableVersion 2.11.4Europe (Milan)Not availableNot availableEurope (Paris)Not availableVersion 2.11.4Europe (Spain)Not availableNot availableEurope (Stockholm)Not availableNot availableEurope (Zurich)Not availableNot availableIsrael (Tel Aviv)Not availableNot availableMiddle East (Bahrain)Not availableNot availableMiddle East (UAE)Not availableNot availableSouth America (São Paulo)Not availableNot availableAWS GovCloud (US-East)Not availableNot availableAWS GovCloud (US-West)Not availableNot available

## Aurora Serverless v1 with Aurora PostgreSQL

The following Regions and engine versions are available for Aurora Serverless v1
with Aurora PostgreSQL.

RegionAurora PostgreSQL 13US East (N. Virginia)Version 13.12US East (Ohio)Version 13.12US West (N. California)Version 13.12US West (Oregon)Version 13.12Africa (Cape Town)Not availableAsia Pacific (Hong Kong)Not availableAsia Pacific (Hyderabad)Not availableAsia Pacific (Jakarta)Not availableAsia Pacific (Malaysia)Not availableAsia Pacific (Melbourne)Not availableAsia Pacific (Mumbai)Version 13.12Asia Pacific (Osaka)Not availableAsia Pacific (Seoul)Version 13.12Asia Pacific (Singapore)Version 13.12Asia Pacific (Sydney)Version 13.12Asia Pacific (Thailand)Not availableAsia Pacific (Tokyo)Version 13.12Canada (Central)Version 13.12Canada West (Calgary)Not availableChina (Beijing)Not availableChina (Ningxia)Not availableEurope (Frankfurt)Version 13.12Europe (Ireland)Version 13.12Europe (London)Version 13.12Europe (Milan)Not availableEurope (Paris)Version 13.12Europe (Spain)Not availableEurope (Stockholm)Not availableEurope (Zurich)Not availableIsrael (Tel Aviv)Not availableMiddle East (Bahrain)Not availableMiddle East (UAE)Not availableSouth America (São Paulo)Not availableAWS GovCloud (US-East)Not availableAWS GovCloud (US-West)Not available

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Aurora Serverless v2

RDS Data API
