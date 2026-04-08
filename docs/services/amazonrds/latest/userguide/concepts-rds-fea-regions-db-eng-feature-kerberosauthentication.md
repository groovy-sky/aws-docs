# Supported Regions and DB engines for Kerberos authentication in Amazon RDS

By using Kerberos authentication in Amazon RDS, you can support external authentication of
database users using Kerberos and Microsoft Active Directory. Using Kerberos and Active
Directory provides the benefits of single sign-on and centralized authentication of database
users.

Kerberos authentication isn't available with the following engines:

- RDS for MariaDB

Although most AWS Regions are active by default for your AWS account, certain Regions
are activated only when you manually select them. These Regions are referred to as
_opt-in Regions_. In contrast, Regions that are active by default, as soon
as your AWS account is created, are referred to as _commercial_
_Regions_, or simply, _Regions_. For opt-in Regions,
you must use a regionalized service principal of the form
`directoryservice.rds.region_name.amazonaws.com`. For
example, for Africa (Cape Town), you must add service principal
`directoryservice.rds.af-south-1.amazonaws.com` to your trust policy. For
more information, see [Kerberos authentication](database-authentication.md#kerberos-authentication).

###### Topics

- [Kerberos authentication with RDS for Db2](#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.db2)

- [Kerberos authentication with RDS for MySQL](#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.my)

- [Kerberos authentication with RDS for Oracle](#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.ora)

- [Kerberos authentication with RDS for PostgreSQL](#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.pg)

- [Kerberos authentication with RDS for SQL Server](#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.sq)

## Kerberos authentication with RDS for Db2

The following Regions and engine versions are available for Kerberos authentication with
RDS for Db2.

RegionRDS for Db2 11.5US East (N. Virginia)All versionsUS East (Ohio)All versionsUS West (N. California)All versionsUS West (Oregon)All versionsAfrica (Cape Town)All versionsAsia Pacific (Hong Kong)Not availableAsia Pacific (Hyderabad)All versionsAsia Pacific (Jakarta)All versionsAsia Pacific (Malaysia)Not availableAsia Pacific (Melbourne)All versionsAsia Pacific (Mumbai)All versionsAsia Pacific (New Zealand)Not availableAsia Pacific (Osaka)Not availableAsia Pacific (Seoul)All versionsAsia Pacific (Singapore)All versionsAsia Pacific (Sydney)All versionsAsia Pacific (Taipei)Not availableAsia Pacific (Thailand)Not availableAsia Pacific (Tokyo)All versionsCanada (Central)All versionsCanada West (Calgary)Not availableChina (Beijing)Not availableChina (Ningxia)Not availableEurope (Frankfurt)All versionsEurope (Ireland)All versionsEurope (London)All versionsEurope (Milan)All versionsEurope (Paris)Not availableEurope (Spain)All versionsEurope (Stockholm)All versionsEurope (Zurich)All versionsIsrael (Tel Aviv)All versionsMexico (Central)Not availableMiddle East (Bahrain)All versionsMiddle East (UAE)All versionsSouth America (São Paulo)All versionsAWS GovCloud (US-East)Not availableAWS GovCloud (US-West)Not available

## Kerberos authentication with RDS for MySQL

The following Regions and engine versions are available for Kerberos authentication with
RDS for MySQL.

RegionRDS for MySQL 8.4RDS for MySQL 8.0RDS for MySQL 5.7 (under RDS Extended Support)US East (N. Virginia)

All versions

All versions

All versions

US East (Ohio)

All versions

All versions

All versions

US West (N. California)

All versions

All versions

All versions

US West (Oregon)

All versions

All versions

All versions

Africa (Cape Town)

All versions

All versions

All versions

Asia Pacific (Hong Kong)

All versions

All versions

All versions

Asia Pacific (Hyderabad)

All versions

All versions

All versions

Asia Pacific (Jakarta)

All versions

All versions

All versions

Asia Pacific (Malaysia)

Not available

Not available

Not available

Asia Pacific (Melbourne)

All versions

All versions

All versions

Asia Pacific (Mumbai)

All versions

All versions

All versions

Asia Pacific (New Zealand)Not availableNot availableNot availableAsia Pacific (Osaka)

All versions

All versions

All versions

Asia Pacific (Seoul)

All versions

All versions

All versions

Asia Pacific (Singapore)

All versions

All versions

All versions

Asia Pacific (Sydney)

All versions

All versions

All versions

Asia Pacific (Taipei)Not availableNot availableNot availableAsia Pacific (Thailand)Not availableNot availableNot availableAsia Pacific (Tokyo)

All versions

All versions

All versions

Canada (Central)

All versions

All versions

All versions

Canada West (Calgary)

Not available

Not available

Not available

China (Beijing)

All versions

All versions

All versions

China (Ningxia)

All versions

All versions

All versions

Europe (Frankfurt)

All versions

All versions

All versions

Europe (Ireland)

All versions

All versions

All versions

Europe (London)

All versions

All versions

All versions

Europe (Milan)

All versions

All versions

All versions

Europe (Paris)

All versions

All versions

All versions

Europe (Spain)

All versions

All versions

All versions

Europe (Stockholm)

All versions

All versions

All versions

Europe (Zurich)

All versions

All versions

All versions

Israel (Tel Aviv)

All versions

All versions

All versions

Mexico (Central)Not availableNot availableNot availableMiddle East (Bahrain)

All versions

All versions

All versions

Middle East (UAE)

All versions

All versions

All versions

South America (São Paulo)

All versions

All versions

All versions

AWS GovCloud (US-East)

All versions

All versions

All versions

AWS GovCloud (US-West)

All versions

All versions

All versions

## Kerberos authentication with RDS for Oracle

The following Regions and engine versions are available for Kerberos authentication with
RDS for Oracle.

RegionRDS for Oracle 21cRDS for Oracle 19cUS East (N. Virginia)All versionsAll versionsUS East (Ohio)All versionsAll versionsUS West (N. California)All versionsAll versionsUS West (Oregon)All versionsAll versionsAfrica (Cape Town) (opt-in Region)All versionsAll versionsAsia Pacific (Hong Kong) (opt-in Region)All versionsAll versionsAsia Pacific (Hyderabad) (opt-in Region)All versionsAll versionsAsia Pacific (Jakarta) (opt-in Region)All versionsAll versionsAsia Pacific (Malaysia)Not availableNot availableAsia Pacific (Melbourne) (opt-in Region)All versionsAll versionsAsia Pacific (Mumbai)All versionsAll versionsAsia Pacific (New Zealand)Not availableNot availableAsia Pacific (Osaka)Not availableNot availableAsia Pacific (Seoul)All versionsAll versionsAsia Pacific (Singapore)All versionsAll versionsAsia Pacific (Sydney)All versionsAll versionsAsia Pacific (Taipei)Not availableNot availableAsia Pacific (Thailand)Not availableNot availableAsia Pacific (Tokyo)All versionsAll versionsCanada (Central)All versionsAll versionsCanada West (Calgary)Not availableNot availableChina (Beijing)Not availableNot availableChina (Ningxia)Not availableNot availableEurope (Frankfurt)All versionsAll versionsEurope (Ireland)All versionsAll versionsEurope (London)All versionsAll versionsEurope (Milan) (opt-in Region)All versionsAll versionsEurope (Paris)Not availableNot availableEurope (Spain) (opt-in Region)All versionsAll versionsEurope (Stockholm)All versionsAll versionsEurope (Zurich) (opt-in Region)All versionsAll versionsIsrael (Tel Aviv) (opt-in Region)All versionsAll versionsMexico (Central)Not availableNot availableMiddle East (Bahrain) (opt-in Region)All versionsAll versionsMiddle East (UAE) (opt-in Region)All versionsAll versionsSouth America (São Paulo)All versionsAll versionsAWS GovCloud (US-East)All versionsAll versionsAWS GovCloud (US-West)All versionsAll versions

## Kerberos authentication with RDS for PostgreSQL

The following Regions and engine versions are available for Kerberos authentication with
RDS for PostgreSQL.

RegionRDS for PostgreSQL 18RDS for PostgreSQL 17RDS for PostgreSQL 16RDS for PostgreSQL 15RDS for PostgreSQL 14RDS for PostgreSQL 13RDS for PostgreSQL 12RDS for PostgreSQL 11RDS for PostgreSQL 10US East (N. Virginia)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsUS East (Ohio)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsUS West (N. California)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsUS West (Oregon)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAfrica (Cape Town)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Hong Kong)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Hyderabad)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Jakarta)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Malaysia)Not availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Melbourne)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Mumbai)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (New Zealand)Not availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Osaka)Not availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Seoul)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Singapore)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Sydney)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAsia Pacific (Taipei)Not availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Thailand)Not availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Tokyo)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsCanada (Central)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsCanada West (Calgary)Not availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableChina (Beijing)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsChina (Ningxia)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Frankfurt)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Ireland)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (London)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Milan)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Paris)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Spain)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Stockholm)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsEurope (Zurich)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsIsrael (Tel Aviv)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsMexico (Central)Not availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableNot availableMiddle East (Bahrain)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsMiddle East (UAE)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsSouth America (São Paulo)All versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAll versionsAWS GovCloud (US-East)

All versions

All versions

All versions

All versions

All versions

All versions

All versions

All versions

All versions

AWS GovCloud (US-West)

All versions

All versions

All versions

All versions

All versions

All versions

All versions

All versions

All versions

## Kerberos authentication with RDS for SQL Server

The following Regions and engine versions are available for Kerberos authentication with
RDS for SQL Server.

RegionRDS for SQL Server 2022RDS for SQL Server 2019RDS for SQL Server 2017RDS for SQL Server 2016US East (N. Virginia)All versionsAll versionsAll versionsAll versionsUS East (Ohio)All versionsAll versionsAll versionsAll versionsUS West (N. California)All versionsAll versionsAll versionsAll versionsUS West (Oregon)All versionsAll versionsAll versionsAll versionsAfrica (Cape Town)All versionsAll versionsAll versionsAll versionsAsia Pacific (Hong Kong)All versionsAll versionsAll versionsAll versionsAsia Pacific (Hyderabad)All versionsAll versionsAll versionsAll versionsAsia Pacific (Malaysia)Not availableNot availableNot availableNot availableAsia Pacific (Melbourne)All versionsAll versionsAll versionsAll versionsAsia Pacific (Mumbai)All versionsAll versionsAll versionsAll versionsAsia Pacific (New Zealand)Not availableNot availableNot availableNot availableAsia Pacific (Osaka)All versionsAll versionsAll versionsAll versionsAsia Pacific (Seoul)All versionsAll versionsAll versionsAll versionsAsia Pacific (Singapore)All versionsAll versionsAll versionsAll versionsAsia Pacific (Sydney)All versionsAll versionsAll versionsAll versionsAsia Pacific (Taipei)Not availableNot availableNot availableNot availableAsia Pacific (Thailand)Not availableNot availableNot availableNot availableAsia Pacific (Tokyo)All versionsAll versionsAll versionsAll versionsCanada (Central)All versionsAll versionsAll versionsAll versionsCanada West (Calgary)Not availableNot availableNot availableNot availableChina (Beijing)All versionsAll versionsAll versionsAll versionsChina (Ningxia)All versionsAll versionsAll versionsAll versionsEurope (Frankfurt)All versionsAll versionsAll versionsAll versionsEurope (Ireland)All versionsAll versionsAll versionsAll versionsEurope (London)All versionsAll versionsAll versionsAll versionsEurope (Milan)All versionsAll versionsAll versionsAll versionsEurope (Paris)All versionsAll versionsAll versionsAll versionsEurope (Spain)All versionsAll versionsAll versionsAll versionsEurope (Stockholm)All versionsAll versionsAll versionsAll versionsEurope (Zurich)All versionsAll versionsAll versionsAll versionsIsrael (Tel Aviv)Not availableNot availableNot availableNot availableMexico (Central)Not availableNot availableNot availableNot availableMiddle East (Bahrain)All versionsAll versionsAll versionsAll versionsMiddle East (UAE)All versionsAll versionsAll versionsAll versionsSouth America (São Paulo)All versionsAll versionsAll versionsAll versionsAWS GovCloud (US-East)All versionsAll versionsAll versionsAll versionsAWS GovCloud (US-West)All versionsAll versionsAll versionsAll versions

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IAM database authentication

Multi-AZ DB clusters

All content copied from https://docs.aws.amazon.com/.
