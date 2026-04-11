# Release calendars for Amazon Aurora MySQL

The release calendars on this page can help you plan your major and minor version upgrades. For more information on Amazon Aurora
upgrades, versioning, and lifecycle, see [Amazon Aurora versions](../aurorauserguide/aurora-versionpolicy.md).

###### Topics

- [Release calendar for Aurora MySQL major versions](#AuroraMySQL.release-calendars.major)

- [Release calendar for Aurora MySQL minor versions](#AuroraMySQL.release-calendars.minor)

## Release calendar for Aurora MySQL major versions

Aurora MySQL major versions are available under standard support at least until community end of life for the corresponding community version.
You can continue running a major version past its Aurora end of standard support date for a fee. For more information, see
[Using Amazon RDS Extended Support](../aurorauserguide/extended-support.md) and [Amazon Aurora pricing](https://aws.amazon.com/rds/aurora/pricing).

Aurora MySQL currently supports the following major versions.

###### Note

You can also view information about support dates for major engine versions by
using
the AWS CLI or RDS API. For more information, see [Viewing support dates for engine versions in Amazon RDS Extended Support](../aurorauserguide/extended-support-viewing-support-dates.md).

Community major versionAurora major versionCommunity end of life dateAurora end of standard support dateRDS start of Extended Support year 1 pricing dateRDS start of Extended Support year 3 pricing dateRDS end of Extended Support dateMinor versions eligible for RDS Extended SupportMySQL 8.0Aurora MySQL version 3April 202630 April 20281 May 2028Not applicable31 July 2029To be determinedMySQL 5.7Aurora MySQL version 2October 202331 October 20241 December 2024Not applicable28 February 2027Aurora MySQL 2.11 and 2.12MySQL 5.6 (deprecated)Aurora MySQL version 1 (deprecated)5 February 202128 February 2023Not applicableNot applicableNot applicableNot applicable

###### Note

Amazon RDS Extended Support for Aurora MySQL version 2 starts on November 1, 2024, but you won't be charged until December 1, 2024.
Between November 1 and November 30, 2024, all Aurora MySQL version 2 DB clusters are covered under Amazon RDS Extended Support.

## Release calendar for Aurora MySQL minor versions

Aurora MySQL currently supports the following minor versions.

In general, Aurora minor versions are released quarterly. The release schedule might vary to pick up additional features or fixes.

Minor versions can reach end of standard support before corresponding major versions do. For example, version 3.07 will reach its
end of standard support date in August 2025, while major version 3 will reach its end of standard support on 30 April 2028.
Amazon RDS will support additional 3.\* minor versions released between these dates.

Amazon RDS Extended Support charges apply only to certain minor versions after a major version is eligible for Extended Support. For more information about major versions eligible for Extended Support, see [Release calendar for Aurora MySQL major versions](#AuroraMySQL.release-calendars.major)
.

Aurora MySQL versionAurora MySQL release dateAurora MySQL end of standard support date

**3.12** (Compatible with Community MySQL 8.0.44)

February 17, 2026

February 17, 2027

**3.11** (Compatible with Community MySQL 8.0.43)

November 13, 2025

November 13, 2026

**3.10** (Compatible with Community MySQL 8.0.42) (LTS)

July 31, 2025

April 30, 2028

**3.09** (Compatible with Community MySQL 8.0.40)

May 14, 2025

May 31, 2026

**3.08** (Compatible with Community MySQL 8.0.39)

November 18, 2024

May 31, 2026

**3.04** (Compatible with Community MySQL 8.0.28) (LTS)

July 31, 2023

October 31, 2026

**2.121** (Compatible with Community MySQL 5.7.40 or
5.7.442)

July 25, 2023

October 31, 2024

**2.111** (Compatible with Community MySQL 5.7.12)

October 25, 2022

October 31, 2024

LTS – Aurora MySQL long-term support (LTS) versions. For more information, see
[Long-term support (LTS) and beta releases for Amazon Aurora MySQL](../aurorauserguide/auroramysql-update-specialversions.md).

1 This minor version will continue to be available when the major version is in Amazon RDS Extended Support.

2 Aurora MySQL 2.12 versions through 2.12.1 are compatible with MySQL version 5.7.40, and versions 2.12.2 and higher
are compatible with MySQL version 5.7.44.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL release notes

Aurora MySQL version 3

All content copied from https://docs.aws.amazon.com/.
