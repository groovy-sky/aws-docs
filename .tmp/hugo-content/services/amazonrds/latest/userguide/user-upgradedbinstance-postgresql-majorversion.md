# Choosing a major version for an RDS for PostgreSQL upgrade

Major version upgrades can contain changes that are not backward-compatible
with previous versions of the database. New functionality can cause your
existing applications to stop working correctly. For this reason, Amazon RDS
doesn't apply major version upgrades automatically. To perform a major
version upgrade, you modify your database manually. Make sure that you
thoroughly test any upgrade to verify that your applications work correctly
before applying the upgrade to your production databases. When you do a
PostgreSQL major version upgrade, we recommend that you follow the steps
described in [How to perform a major version upgrade for RDS for PostgreSQL](user-upgradedbinstance-postgresql-majorversion-process.md).

When you upgrade a PostgreSQL Single-AZ DB instance or Multi-AZ DB instance deployment to
its next major version, any read replicas associated with the database are
also upgraded to that next major version. In some cases, you can skip to a
higher major version when upgrading. If your upgrade skips a major version,
the read replicas are also upgraded to that target major version. Upgrades
to version 11 that skip other major versions have certain limitations. You
can find the details in the steps described in [How to perform a major version upgrade for RDS for PostgreSQL](user-upgradedbinstance-postgresql-majorversion-process.md).

Most PostgreSQL extensions aren't upgraded during a PostgreSQL engine
upgrade. These must be upgraded separately. For more information, see [Upgrading PostgreSQL extensions in RDS for PostgreSQL databases](user-upgradedbinstance-postgresql-extensionupgrades.md).

You can find out which major versions are available for your RDS for PostgreSQL
database by running the following AWS CLI query:

```nohighlight

aws rds describe-db-engine-versions --engine postgres  --engine-version your-version --query "DBEngineVersions[*].ValidUpgradeTarget[*].{EngineVersion:EngineVersion}" --output text
```

The following table summarizes the results of this query for all available
versions. An asterisk (\*) on the version number means that version is no
longer supported. If your current version is unsupported, we recommend that
you upgrade to the newest minor version upgrade target or to one of the
other available upgrade targets for that version.

Current source versionUpgrade targets17.6

None

17.5

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176)

17.4

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175)

17.3\*, 17.2

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

17.1\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

16.10

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176)

16.9

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610)

16.8

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169)

16.7\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

16.7

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

16.6

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

16.5\*, 16.4

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166)

16.3

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164)

16.2\*, 16.1\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

15.14

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610)

15.13

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514)

15.12, 15.11\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513)

15.10

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512)

15.9\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510)

15.8

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510)

15.7

[16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version167), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version165), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.11](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1511), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version159), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158)

15.6\*, 15.5\*, 15.4\*, 15.3\*, 15.2\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

14.19

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514)

14.18

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419)

14.17, 14.16\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418)

14.15

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417)

14.14\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415)

14.13

[16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164)

[15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.11](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1511), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version159), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158)

[14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1416), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1414)

14.12

[16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.11](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1511), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version159), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

[14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1416), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1414), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413)

14.11\*, 14.10\*, 14.9\*, 14.8\*, 14.7\*, 14.6\*,
14.5\*, 14.4\*, 14.3\*, 14.2\*, 14.1\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413), [14.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1412)

13.22

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419)

13.21

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322)

13.20, 13.19\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321)

13.18, 13.17\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320)

13.16

[16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164)

[15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158)

[14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1414), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413)

[13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1319), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318), [13.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1317)

13.15

[16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

[14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1414), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413), [14.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1412)

[13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1319), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318), [13.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1317), [13.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1316)

13.14\*, 13.13\*, 13.12\*, 13.11\*, 13.10\*,
13.9\*, 13.8\*, 13.7\*, 13.6\*, 13.5\*, 13.4\*, 13.3\*,
13.2\*, 13.1\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413), [14.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1412)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318), [13.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1316), [13.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1315)

12.22-rds.20250508

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321)

12.22-rds.20250220

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320)

[12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

12.22, 12.21\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318)

[12.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250220), [12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

12.20\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318), [13.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1316)

[12.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222), [12.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250220), [12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

12.19\*, 12.18\*, 12.17\*, 12.16\*, 12.15\*,
12.14\*, 12.13\*, 12.12\*, 12.11\*, 12.10\*, 12.9\*,
12.8\*, 12.7\*, 12.6\*, 12.5\*, 12.4\*, 12.3\*,
12.2\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413), [14.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1412)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318), [13.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1316), [13.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1315)

[12.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222), [12.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250220), [12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

11.22-rds.20250508

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321)

[12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

11.22-rds.20250220

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320)

[12.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250220), [12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

[11.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20250508)

11.22-rds.20240509

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413), [14.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1412)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318), [13.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1316), [13.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1315)

[12.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222), [12.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250220), [12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

[11.22-rds.20240808](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20240808), [11.22-rds.20241121](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20241121), [11.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20250220), [11.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20250508)

11.22, 11.21\*, 11.20\*, 11.19\*, 11.18\*,
11.17\*, 11.16\*, 11.15\*, 11.14\*, 11.13\*, 11.12\*,
11.11\*, 11.10\*, 11.9\*, 11.8\*, 11.7\*, 11.6\*, 11.5\*,
11.4\*, 11.2\*, 11.1\*

[17.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version176), [17.5](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version175), [17.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version174), [17.2](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version172)

[16.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1610), [16.9](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version169), [16.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version168), [16.6](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version166), [16.4](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version164), [16.3](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version163)

[15.14](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1514), [15.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1513), [15.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1512), [15.10](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1510), [15.8](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version158), [15.7](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version157)

[14.19](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1419), [14.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1418), [14.17](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1417), [14.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1415), [14.13](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1413), [14.12](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1412)

[13.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1322), [13.21](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1321), [13.20](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1320), [13.18](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1318), [13.16](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1316), [13.15](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1315)

[12.22](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222), [12.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250220), [12.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1222rds20250508)

[11.22-rds.20240418](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20240418), [11.22-rds.20240509](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20240509), [11.22-rds.20240808](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20240808), [11.22-rds.20241121](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20241121), [11.22-rds.20250220](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20250220), [11.22-rds.20250508](../postgresqlreleasenotes/postgresql-versions.md#postgresql-versions-version1122rds20250508)

\\* This version is no longer supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RDS version numbers

How to perform a major version upgrade

All content copied from https://docs.aws.amazon.com/.
