# Supported DB engines for DB instance classes

The following are DB engine–specific considerations for DB instance classes:

**Db2**

DB instance class support varies according to the version and edition of Db2. For
instance class support by version and edition, see [Amazon RDS for Db2 instance classes](db2-concepts-general-instanceclasses.md).

**Microsoft SQL Server**

DB instance class support varies according to the version and edition of SQL
Server. For instance class support by version and edition, see [DB instance class support for Microsoft SQL Server](sqlserver-concepts-general-instanceclasses.md).

**Oracle**

DB instance class support varies according to the Oracle Database version and
edition. RDS for Oracle supports additional memory-optimized instance classes. These
classes have names of the form
db.r5. `instance_size`.tpc `threads_per_core`.mem `ratio`.
For the vCPU count and memory allocation for each optimized class, see [Supported RDS for Oracle DB instance classes](oracle-concepts-instanceclasses.md#Oracle.Concepts.InstanceClasses.Supported).

**RDS Custom**

For information about the DB instance classes supported in RDS Custom, see [DB instance class support for RDS Custom for Oracle](custom-oracle-feature-support.md#custom-reqs-limits.instances) and [DB instance class support for RDS Custom for SQL Server](custom-reqs-limits-instancesms.md).

In the following table, you can find details about supported Amazon RDS DB instance classes for each
Amazon RDS DB engine. The cell for each engine contains one of the following values:

Yes

The instance class is supported for all versions of the DB engine.

No

The instance class isn't supported for the DB engine.

`specific-versions`

The instance class is supported only for the specified database versions of
the DB engine.

Amazon RDS periodically deprecates major and minor DB engine versions. Not all AWS Regions might
have support for earlier engine versions. For information about current supported versions,
see topics for the individual DB engines: [Db2\
versions](db2-concepts-versionmgmt-supported.md), [MariaDB\
versions](mariadb-concepts-versionmgmt.md#MariaDB.Concepts.VersionMgmt.Supported), [Microsoft SQL\
Server versions](sqlserver-concepts-general-versionsupport.md), [MySQL\
versions](mysql-concepts-versionmgmt.md), [Oracle versions](oracle-concepts-database-versions.md),
and [PostgreSQL\
versions](postgresql-concepts-general-dbversions.md).

###### Topics

- [Supported DB engines for general-purpose instance classes](#gen-purpose-inst-classes)

- [Supported DB engines for memory-optimized instance classes](#mem-opt-inst-classes)

- [Supported DB engines for compute-optimized instance classes](#compute-opt-inst-classes)

- [Supported DB engines for burstable-performance instance classes](#burstable-inst-classes)

- [Supported DB engines for Optimized Reads instance classes](#read-opt-inst-classes)

## Supported DB engines for general-purpose instance classes

The following tables show the supported databases and database versions for the
general-purpose instance classes.

**db.m8g – general-purpose instance classes powered by AWS**
**Graviton4 processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m8g.48xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.24xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.16xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.12xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.8xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.4xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.2xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m8g.largeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higher

**db.m7i – general-purpose instance classes powered by 4th**
**generation Intel Xeon Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m7i.48xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, Enterprise Edition onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.24xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, Enterprise Edition onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.16xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, Enterprise Edition onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.12xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, Enterprise Edition onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.8xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, Enterprise Edition onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.4xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.2xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.largeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.m7i.metal-48xlNoNoNoNoBYOL only, Enterprise Edition onlyNodb.m7i.metal-24xlNoNoNoNoBYOL only, Enterprise Edition onlyNo

**db.m7g – general-purpose instance classes powered by AWS**
**Graviton3 processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m7g.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.m7g.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.m7g.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.m7g.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.m7g.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.m7g.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.m7g.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versions

**db.m6g – general-purpose instance classes powered by AWS**
**Graviton2 processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m6g.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.m6g.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.m6g.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.m6g.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.m6g.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.m6g.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.m6g.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versions

**db.m6gd – general-purpose instance classes powered by AWS**
**Graviton2 processors and SSD storage**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m6gd.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, and 14 versions; 13.7 and higher 13 versions; and
13.4db.m6gd.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, and 14 versions; 13.7 and higher 13 versions; and
13.4db.m6gd.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, and 14 versions; 13.7 and higher 13 versions; and
13.4db.m6gd.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, and 14 versions; 13.7 and higher 13 versions; and
13.4db.m6gd.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, and 14 versions; 13.7 and higher 13 versions; and
13.4db.m6gd.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, and 14 versions; 13.7 and higher 13 versions; and
13.4db.m6gd.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, and 14 versions; 13.7 and higher 13 versions; and
13.4

**db.m6id – general-purpose instance classes powered by 3rd**
**generation Intel Xeon Scalable processors and SSD storage**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m6id.32xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.24xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.16xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.12xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.8xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.4xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.2xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.largeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6id.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.m6idn – general-purpose instance classes with 3rd**
**Generation Intel Xeon Scalable processors, SSD storage, and network**
**optimization**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m6idn.32xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.24xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.16xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.12xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.8xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.4xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.2xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.m6idn.largeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

**db.m6in – general-purpose instance classes powered by 3rd**
**generation Intel Xeon Scalable processors and network optimization**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m6in.32xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.24xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.16xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.12xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.8xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.4xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.2xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.largeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.m6in.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.m6i – general-purpose instance classes powered by 3rd**
**generation Intel Xeon Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m6i.32xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.24xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.16xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.12xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.8xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.4xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.2xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.largeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Oracle Database 19c

All available versions

db.m6i.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.m5d – general-purpose instance classes powered by Intel**
**Xeon Platinum processors and SSD storage**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m5d.24xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.m5d.16xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.m5d.12xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.m5d.8xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.m5d.4xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.m5d.2xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.m5d.xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.m5d.largeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4

**db.m5 – general-purpose instance classes 2.5 GHz Intel Xeon**
**Platinum processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m5.24xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.m5.16xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.m5.12xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.m5.8xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.m5.4xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.m5.2xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.m5.xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.m5.largeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

**db.m4 – general-purpose instance classes with Intel Xeon**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m4.16xlargeNoDeprecated

Deprecated

Deprecated

Deprecated

Deprecateddb.m4.10xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.m4.4xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.m4.2xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.m4.xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.m4.largeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecated

**db.m3 – general-purpose instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m3.2xlargeNoNo

Deprecated

Yes

Deprecated

Deprecateddb.m3.xlargeNoNo

Deprecated

Yes

Deprecated

Deprecateddb.m3.largeNoNo

Deprecated

Yes

Deprecated

Deprecateddb.m3.mediumNoNo

Deprecated

Yes

Deprecated

Deprecated

## Supported DB engines for memory-optimized instance classes

The following tables show the supported databases and database versions for the
memory-optimized instance classes.

**db.z1d – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.z1d.12xlargeNoNoYesNo

Yes

Nodb.z1d.6xlargeNoNoYesNo

Yes

Nodb.z1d.3xlargeNoNoYesNo

Yes

Nodb.z1d.2xlargeNoNoYesNo

Yes

Nodb.z1d.xlargeNoNoYesNo

Yes

Nodb.z1d.largeNoNoYesNo

Yes

No

**db.x2g – memory-optimized instance classes powered by AWS Graviton2 processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.x2g.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.x2g.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.x2g.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.x2g.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.x2g.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.x2g.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.x2g.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versions

**db.x2idn – memory-optimized instance classes powered by 3rd generation Intel Xeon Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.x2idn.32xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0Enterprise Edition onlyPostgreSQL 15 versions, 14.6, and 13.9db.x2idn.24xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0Enterprise Edition onlyPostgreSQL 15 versions, 14.6, and 13.9db.x2idn.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0Enterprise Edition onlyPostgreSQL 15 versions, 14.6, and 13.9db.x2idn.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.x2iedn – memory-optimized instance classes with local NVMe-based SSDs, powered by**
**3rd generation Intel Xeon Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.x2iedn.32xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4Enterprise and Standard Editions only, SQL Server 2014 12.00 and
higherMySQL 8.4 and 8.0Enterprise Edition onlyAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.x2iedn.24xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4Enterprise and Standard Editions only, SQL Server 2014 12.00 and
higherMySQL 8.4 and 8.0Enterprise Edition onlyAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.x2iedn.16xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4Enterprise and Standard Editions only, SQL Server 2014 12.00 and
higherMySQL 8.4 and 8.0Enterprise Edition onlyAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.x2iedn.8xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4Enterprise and Standard Editions only, SQL Server 2014 12.00 and
higherMySQL 8.4 and 8.0Enterprise Edition onlyAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.x2iedn.4xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4Enterprise and Standard Editions only, SQL Server 2014 12.00 and
higherMySQL 8.4 and 8.0Enterprise Edition and Standard Edition 2 (SE2)All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.x2iedn.2xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4Enterprise and Standard Editions only, SQL Server 2014 12.00 and
higherMySQL 8.4 and 8.0Enterprise Edition and Standard Edition 2 (SE2)All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.x2iedn.xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4Enterprise and Standard Editions only, SQL Server 2014 12.00 and
higherMySQL 8.4 and 8.0Enterprise Edition and Standard Edition 2 (SE2)All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.x2iedn.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.x2iezn – memory-optimized instance classes powered by 2nd generation Intel Xeon**
**Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.x2iezn.metalNoNoNoNoBYOL only, Enterprise Edition onlyNodb.x2iezn.8xlargeNoNoNoNoEnterprise Edition onlyNodb.x2iezn.6xlargeNoNoNoNoEnterprise Edition onlyNodb.x2iezn.4xlargeNoNoNoNoEnterprise Edition and Standard Edition 2 (SE2)Nodb.x2iezn.2xlargeNoNoNoNoEnterprise Edition and Standard Edition 2 (SE2)No

**db.x1e – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.x1e.32xlargeNoNoDeprecatedNoDeprecatedNodb.x1e.16xlargeNoNoDeprecatedNoDeprecatedNodb.x1e.8xlargeNoNoDeprecatedNoDeprecatedNodb.x1e.4xlargeNoNoDeprecatedNoDeprecatedNodb.x1e.2xlargeNoNoDeprecatedNoDeprecatedNodb.x1e.xlargeNoNoDeprecatedNoDeprecatedNo

**db.x1 – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.x1.32xlargeNoNoDeprecatedNoDeprecatedNodb.x1.16xlargeNoNoDeprecatedNoDeprecatedNo

**db.r8g – memory-optimized instance classes powered by AWS Graviton4**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r8g.48xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.24xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.16xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.12xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.8xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.4xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.2xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.xlargeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r8g.largeNoMariaDB 11.8.3 and higher, 11.4.3 and higher, 10.11.7 and higher, 10.6.13 and higher, 10.5.20 and higher, and 10.4.29 and higherNoMySQL 8.0.32 and higherNoPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higher

**db.r7i – memory-optimized instance classes preconfigured for high memory,**
**storage, and I/O**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r7i.8xlarge.tpc2.mem3xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.8xlarge.tpc2.mem2xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.6xlarge.tpc2.mem4xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.6xlarge.tpc2.mem2xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.4xlarge.tpc2.mem4xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.4xlarge.tpc2.mem3xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.4xlarge.tpc2.mem2xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.3xlarge.tpc2.mem4xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.2xlarge.tpc2.mem8xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.2xlarge.tpc2.mem4xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.xlarge.tpc2.mem4xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2Nodb.r7i.xlarge.tpc2.mem2xNoNoNoNoBYOL only, Enterprise Edition and Standard Edition 2No

**db.r7i – memory-optimized instance classes powered by 4th**
**generation Intel Xeon Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r7i.48xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.24xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.16xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.12xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.8xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL onlyPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.4xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.2xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.xlargeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.largeDb2 11.5MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.0.32 and higherBYOL only, all editionsPostgreSQL version 17.1 and higher, 16.1 and higher, 15.4 and higher, 14.9
and higher, and 13.11 and higherdb.r7i.metal-48xlNoNoNoNo

BYOL only, Enterprise Edition only

Nodb.r7i.metal-24xlNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.r7g – memory-optimized instance classes powered by AWS Graviton3**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r7g.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.r7g.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.r7g.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.r7g.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.r7g.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.r7g.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versionsdb.r7g.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.4 and higher 13 versions

**db.r6g – memory-optimized instance classes powered by AWS Graviton2**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6g.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r6g.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r6g.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r6g.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r6g.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r6g.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r6g.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versions

**db.r6gd – memory-optimized instance classes powered by AWS Graviton2**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6gd.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4

**db.r6id – memory-optimized instance classes powered by 3rd generation Intel Xeon**
**Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6id.32xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.24xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.16xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.12xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.8xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.4xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.2xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.largeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.r6idn – memory-optimized instance classes powered by 3rd generation**
**Intel Xeon Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6idn.32xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6idn.24xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6idn.16xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6idn.12xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6idn.8xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6idn.4xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6idn.2xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6idn.xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0No

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

**db.r6in – memory-optimized instance classes powered by 3rd generation Intel Xeon**
**Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6in.32xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.24xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.16xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.12xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.8xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.4xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.2xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.xlargeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.largeYes

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.3 and higher 14 versions, 13.7 and higher 13 versions,
12.11 and higher 12 versions, and 11.16 and higher 11 versions

db.r6in.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.r6i – memory-optimized instance classes preconfigured for high memory,**
**storage, and I/O**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6i.8xlarge.tpc2.mem4xNoNoNoNo

Enterprise Edition only

Nodb.r6i.8xlarge.tpc2.mem3xNoNoNoNo

Enterprise Edition only

Nodb.r6i.6xlarge.tpc2.mem4xNoNoNoNo

Enterprise Edition only

Nodb.r6i.4xlarge.tpc2.mem4xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.4xlarge.tpc2.mem3xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.4xlarge.tpc2.mem2xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.2xlarge.tpc2.mem8xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.2xlarge.tpc2.mem4xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.2xlarge.tpc1.mem2xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.xlarge.tpc2.mem4xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.xlarge.tpc2.mem2xNoNoNoNo

EE and SE2 BYOL

Nodb.r6i.large.tpc1.mem2xNoNoNoNo

EE and SE2 BYOL

No

**db.r6i – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6i.32xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.24xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.16xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.12xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.8xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.4xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.2xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.xlargeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.largeYesMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0

Yes

All PostgreSQL 17, 16, 15, and 14 versions, 13.4 and higher 13 versions, 12.8 and higher 12 versions, 11.13 and higher 11 versions, and 10.21 and higher 10 versionsdb.r6i.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

**db.r5d – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r5d.24xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.r5d.16xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.r5d.12xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.r5d.8xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.r5d.4xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.r5d.2xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.r5d.xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4db.r5d.largeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and 13.4

**db.r5b – memory-optimized instance classes preconfigured for high memory,**
**storage, and I/O**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r5b.8xlarge.tpc2.mem3xNoNoNoNoYesNodb.r5b.6xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5b.4xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5b.4xlarge.tpc2.mem3xNoNoNoNoYesNodb.r5b.4xlarge.tpc2.mem2xNoNoNoNoYesNodb.r5b.2xlarge.tpc2.mem8xNoNoNoNoYesNodb.r5b.2xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5b.2xlarge.tpc1.mem2xNoNoNoNoYesNodb.r5b.xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5b.xlarge.tpc2.mem2xNoNoNoNoYesNodb.r5b.large.tpc1.mem2xNoNoNoNoYesNo

**db.r5b – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r5b.24xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r5b.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r5b.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r5b.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0>YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r5b.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r5b.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r5b.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.r5b.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4YesMySQL 8.4 and 8.0YesAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versions

**db.r5 – memory-optimized instance classes preconfigured for high memory,**
**storage, and I/O**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r5.12xlarge.tpc2.mem2xNoNoNoNoYesNodb.r5.8xlarge.tpc2.mem3xNoNoNoNoYesNodb.r5.6xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5.4xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5.4xlarge.tpc2.mem3xNoNoNoNoYesNodb.r5.4xlarge.tpc2.mem2x NoNoNoNoYesNodb.r5.2xlarge.tpc2.mem8xNoNoNoNoYesNodb.r5.2xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5.2xlarge.tpc1.mem2xNoNoNoNoYesNodb.r5.xlarge.tpc2.mem4xNoNoNoNoYesNodb.r5.xlarge.tpc2.mem2xNoNoNoNoYesNodb.r5.large.tpc1.mem2xNoNoNoNoYesNo

**db.r5 – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r5.24xlargeNoYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.r5.16xlargeNoYesYesYesYes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.r5.12xlargeNoYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.r5.8xlargeNoYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.r5.4xlargeNoYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.r5.2xlargeNoYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.r5.xlargeNoYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

db.r5.largeNoYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, and 11 versions; 10.17 and higher 10 versions; and 9.6.22 and higher 9 versions

**db.r4 – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r4.16xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.r4.8xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.r4.4xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.r4.2xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.r4.xlargeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecateddb.r4.largeNoDeprecated

Deprecated

DeprecatedDeprecatedDeprecated

**db.r3 – memory-optimized instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r3.8xlarge\*\*NoAll MariaDB 10.6, 10.5, 10.4, and 10.3 versions

Deprecated

Yes

Deprecated

Deprecateddb.r3.4xlargeNoAll MariaDB 10.6, 10.5, 10.4, and 10.3 versions

Deprecated

Yes

Deprecated

Deprecateddb.r3.2xlargeNoAll MariaDB 10.6, 10.5, 10.4, and 10.3 versions

Deprecated

Yes

Deprecated

Deprecateddb.r3.xlargeNoAll MariaDB 10.6, 10.5, 10.4, and 10.3 versions

Deprecated

Yes

Deprecated

Deprecateddb.r3.largeNoAll MariaDB 10.6, 10.5, 10.4, and 10.3 versions

Deprecated

Yes

Deprecated

Deprecated

## Supported DB engines for compute-optimized instance classes

The following tables show the supported databases and database versions for the
compute-optimized instance classes.

**db.c6gd – compute-optimized instance classes (for Multi-AZ DB cluster deployments**
**only)**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.c6gd.16xlargeNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versionsdb.c6gd.12xlargeNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versionsdb.c6gd.8xlargeNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versionsdb.c6gd.4xlargeNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versionsdb.c6gd.2xlargeNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versionsdb.c6gd.xlargeNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versionsdb.c6gd.largeNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versionsdb.c6gd.mediumNoNoNoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, and 15 versions; 14.5 and higher 14 versions; 13.4 and
13.7 and higher 13 versions

## Supported DB engines for burstable-performance instance classes

The following tables show the supported databases and database versions for the
burstable-performance instance classes.

**db.t4g – burstable-performance instance classes powered by AWS Graviton2**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.t4g.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.t4g.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.t4g.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.t4g.mediumNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.t4g.smallNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versionsdb.t4g.microNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16, 15, 14, and 13 versions; and 12.7 and higher 12 versions

**db.t3 – burstable-performance instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.t3.2xlargeYesYesYesYesYesAll PostgreSQL 17, 16, 15, 14, 13, 12, 11, and 10 versions; and 9.6.22 and higher 9 versionsdb.t3.xlargeYesYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, 11, and 10 versions; and 9.6.22 and higher 9 versionsdb.t3.largeYesYesYesYesYesAll PostgreSQL 17, 16, 15, 14, 13, 12, 11, and 10 versions; and 9.6.22 and higher 9 versionsdb.t3.mediumYesYesYesYes

Yes

All PostgreSQL 17, 16, 15, 14, 13, 12, 11, and 10 versions; and 9.6.22 and higher 9 versionsdb.t3.smallYesYesYesYesYesAll PostgreSQL 17, 16, 15, 14, 13, 12, 11, and 10 versions; and 9.6.22 and higher 9 versionsdb.t3.microNoYesYesYesOnly on Oracle Database 12c Release 1 (12.1.0.2), which is deprecatedAll PostgreSQL 17, 16, 15, 14, 13, 12, 11, and 10 versions; and 9.6.22 and higher 9 versions

**db.t2 – burstable-performance instance classes**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.t2.2xlargeNoDeprecatedNoDeprecated

Deprecated

Deprecateddb.t2.xlargeNoDeprecatedNoDeprecated

Deprecated

Deprecateddb.t2.largeNoDeprecated

Deprecated

Deprecated

Deprecated

Deprecateddb.t2.mediumNoDeprecated

Deprecated

Deprecated

Deprecated

Deprecateddb.t2.smallNoDeprecated

Deprecated

Deprecated

Deprecated

Deprecateddb.t2.microNoDeprecated

Deprecated

Deprecated

Deprecated

Deprecated

## Supported DB engines for Optimized Reads instance classes

The following tables show the supported databases and database versions for the Optimized
Reads instance classes.

**db.m8gd – memory-optimized instance classes that support Optimized Reads and are powered by AWS Graviton4**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.m8gd.48xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNoMySQL 8.4 and 8.0No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.24xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.16xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.12xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.8xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.4xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.2xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.m8gd.largeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

**db.r8gd – memory-optimized instance classes that support Optimized Reads and are powered by AWS Graviton4**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r8gd.48xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNoMySQL 8.4 and 8.0No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.24xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.16xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.12xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.8xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.4xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.2xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.xlargeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

db.r8gd.largeNoMariaDB 10.5.27 and higher, 10.6.19 and higher, 10.11.9 and higher, 11.4.3
and higher, 11.8.3 and higherNo

MySQL 8.0.40 and higher, 8.4.3 and higher

No

PostgreSQL 13.18 and higher, 14.17 and higher, 15.10 and higher, 16.6 and
higher, 17.2 and higher, 18.1 and higher

**db.r6gd – memory-optimized instance classes that support Optimized Reads and are powered by AWS Graviton2**
**processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6gd.16xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.12xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.8xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.4xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.2xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.xlargeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4db.r6gd.largeNoMariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4NoMySQL 8.4 and 8.0NoAll PostgreSQL 17, 16 and 15 versions, 14.5 and higher 14 versions, 13.7 and higher 13 versions, and
13.4

**db.r6id – memory-optimized instance classes that support Optimized Reads and are powered by 3rd generation Intel Xeon**
**Scalable processors**

Instance classDb2MariaDBMicrosoft SQL ServerMySQLOraclePostgreSQLdb.r6id.32xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.24xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.16xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.12xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.8xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0EE and BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.4xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.2xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.xlargeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.largeNo

MariaDB 11.8, 11.4, 10.11, 10.6, 10.5, and 10.4

NoMySQL 8.4 and 8.0BYOL only

All PostgreSQL 17, 16, and 15 versions, 14.5 and higher 14 versions, and 13.7 and higher 13 versions

db.r6id.metalNoNoNoNo

BYOL only, Enterprise Edition only

No

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DB instance class types

Determining DB instance
class support in AWS Regions

All content copied from https://docs.aws.amazon.com/.
