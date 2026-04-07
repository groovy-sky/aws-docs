# Release calendars for Aurora PostgreSQL

The release calendars on this page can help you plan your major and minor version upgrades. For more information on Amazon Aurora upgrades, versioning,
and lifecycle, see [Amazon Aurora versions](../aurorauserguide/aurora-versionpolicy.md).

###### Topics

- [Release calendar for Aurora PostgreSQL major versions](#aurorapostgresql.major.versions.supported)

- [Release calendar for Aurora PostgreSQL minor versions](#aurorapostgresql.minor.versions.supported)

## Release calendar for Aurora PostgreSQL major versions

Aurora PostgreSQL major versions are available under standard support at least until
community end of life for the corresponding community version. You can continue running
a major version past its Aurora end of standard support date for a fee. For more
information, see [Using Amazon RDS Extended Support](../aurorauserguide/extended-support.md) and [Aurora pricing](https://aws.amazon.com/rds/aurora).

You can use the following dates to plan your testing and upgrade cycles.

###### Note

Dates with only a month and a year are approximate and are updated with an exact date when it's known.

You can also view information about support dates for major engine
versions by using the AWS CLI or the RDS API. For more information,
see [Viewing support dates for engine versions in Amazon RDS Extended Support](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/extended-support-viewing-support-dates.html).

PostgreSQL major versionCommunity release dateCommunity end of life dateAurora major versionAurora PostgreSQL LTS versionAurora release dateAurora end of standard support dateStart of RDS Extended Support year 1 pricingStart of RDS Extended Support year 3 pricingEnd of RDS Extended Support dateMinor versions eligible for RDS Extended SupportPostgreSQL 1118 October 2018November 2023Aurora PostgreSQL 3. Applies to PostgreSQL 11.12 and older versions only. For version 11.13 and higher versions, the Aurora
version is the same as the `major`. `minor` version of the
PostgreSQL community version, with a third digit in the `patch` location.Aurora PostgreSQL 11.926 November 201929 February 20241 April 20241 April 202631 March 2027Aurora PostgreSQL 11.9 and 11.21PostgreSQL 1214 November 2019November 2024Aurora PostgreSQL 4. Applies to PostgreSQL 12.7 and older versions only. For version 12.8 and higher versions, the Aurora
version is the same as the `major`. `minor` version of the
PostgreSQL community version, with a third digit in the `patch` location.Aurora PostgreSQL 12.923 December 202028 February 20251 March 20251 March 202729 February 2028Aurora PostgreSQL 12.9 and 12.22PostgreSQL 1324 September 2020November 2025Aurora PostgreSQL 13. For version 13.3 and higher versions, the Aurora version is the same
as the `major`. `minor` version of the
PostgreSQL community version, with a third digit in the
`patch` location when patches to Aurora are released.Aurora PostgreSQL 13.926 August 202128 February 20261 March 20261 March 202828 February 2029Aurora PostgreSQL 13.9 and 13.23PostgreSQL 1430 September 2021November 2026Aurora PostgreSQL 14.3 and higher. The Aurora version is the same as the
`major`. `minor` version of the
PostgreSQL community version, with a third digit in the
`patch` location when patches to Aurora are released.Aurora PostgreSQL 14.624 February 202228 February 20271 March 20271 March 202929 February 2030To be determinedPostgreSQL 1510 November 2022November 2027Aurora PostgreSQL 15.2 and higher. The Aurora version is the same as the
`major`. `minor` version of the
PostgreSQL community version, with a third digit in the
`patch` location when patches to Aurora are released.Aurora PostgreSQL 15.108 February 202329 February 20281 March 20281 March 203028 February 2031To be determinedPostgreSQL 1614 September 20239 November 2028Aurora PostgreSQL 16.1 and higher. The Aurora version is the same as the
`major`. `minor` version of the
PostgreSQL community version, with a third digit in the
`patch` location when patches to Aurora are released.16.831 January 202428 February 20291 March 20291 March 203128 February 2032To be determinedPostgreSQL 1720 February 2025November 2029Aurora PostgreSQL 17.4 and higher. The Aurora version is the same as the
`major`. `minor` version of the
PostgreSQL community version, with a third digit in the
`patch` location when patches to Aurora are released.17.71 May 202528 February 20301 March 20301 March 203228 February 2033To be determined

###### Note

RDS Extended Support charges only apply after a major version reaches end of standard support. RDS Extended Support for PostgreSQL 11 starts on March 1, 2024,
but will not be charged until April 1, 2024. Between March 1 and March 31, all PostgreSQL version 11 DB instances and clusters on RDS are covered under RDS Extended Support.

## Release calendar for Aurora PostgreSQL minor versions

Aurora currently supports the following minor versions of PostgreSQL.

In general, Aurora minor versions are released quarterly. The release schedule might vary to pick up additional features or fixes.

Amazon RDS Extended Support charges apply only to certain minor versions after a major version is eligible for Extended Support. For more information, see [Release calendar for Aurora PostgreSQL major versions](#aurorapostgresql.major.versions.supported).

###### Note

Dates with only a month and a year are approximate, and will be updated with an exact date when it’s known.

PostgreSQL minor engine versionCommunity release dateAurora release dateAurora end of standard support date**17**

17.7 (LTS)

13 November 2025

18 December 2025

28 February 2030

17.6

14 August 2025

25 November 2025

June 2027

17.5

8 May 2025

30 June 2025

December 2026

17.4

20 February 2025

May 1 2025

November 2026

**16**

16.11

13 November 2025

18 December 2025

July 2027

16.10

14 August 2025

25 November 2025

June 2027

16.9

8 May 2025

30 June 2025

December 2026

16.8 (LTS)

20 February 2025

April 8 2025

28 February 2029

16.6

21 November 2024

13 December 2024

May 2026

16.4

08 August 2024

30 September 2024

May 2026

**15**

15.15

13 November 2025

18 December 2025

July 2027

15.14

14 August 2025

25 November 2025

June 2027

15.13

8 May 2025

30 June 2025

December 2026

15.12

20 February 2025

April 8 2025

November 2026

15.10 (LTS)

21 November 2024

13 December 2024

29 February 2028

15.8

08 August 2024

30 September 2024

May 2026

**14**

14.20

13 November 2025

18 December 2025

28 February 2027

14.19

14 August 2025

25 November 2025

28 February 2027

14.18

8 May 2025

30 June 2025

December 2026

14.17

20 February 2025

April 8 2025

November 2026

14.15

21 November 2024

13 December 2024

May 2026

14.13

08 August 2024

30 September 2024

May 2026

14.6 (LTS)

10 November 2022

23 January 2023

28 February 2027

**13**

13.23\*

13 November 2025

18 December 2025

28 February 2026

13.9 (LTS)\*

10 November 2022

23 January 2023

28 February 2026

**12**

12.22\*

21 November 2024

13 December 2024

28 February 2025

12.9\* (LTS)

11 November 2021

25 February 2022

28 February 2025

**11**

11.21\*

10 August 2023

7 September 2023

29 February 2024

11.9\* (LTS)

13 August 2020

11 December 2020

29 February 2024

\\* Amazon RDS Extended Support eligible minor engine version. For more information, see [Using Amazon RDS Extended Support](../aurorauserguide/extended-support.md).

LTS - Aurora PostgreSQL long-term support (LTS) releases. For more information see [Aurora PostgreSQL long-term support (LTS) releases](../aurorauserguide/aurorapostgresql-updates-lts.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Aurora PostgreSQL release notes

Aurora PostgreSQL updates
