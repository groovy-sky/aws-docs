# Aurora MySQL database engine updates 2026-03-30 (version 3.10.4, compatible with MySQL 8.0.42)

**Version:** 3.10.4

Aurora MySQL 3.10.4 is generally available. Aurora MySQL 3.10 versions are compatible with MySQL 8.0.42. For more information on the community changes that have occurred, see [MySQL 8.0 Release Notes](https://dev.mysql.com/doc/relnotes/mysql/8.0/en).

For details of the new features in Aurora MySQL version 3, see [Aurora MySQL version 3 compatible with MySQL 8.0](../aurorauserguide/auroramysql-mysql80.md). For differences between Aurora MySQL version 3 and Aurora MySQL version 2, see [Comparison of Aurora MySQL version 2 and Aurora MySQL version 3](../aurorauserguide/auroramysql-compare-v2-v3.md). For a comparison of Aurora MySQL version 3 and MySQL 8.0 Community Edition, see [Comparison of Aurora MySQL version 3 and MySQL 8.0 Community Edition](../aurorauserguide/auroramysql-compare-80-v3.md) in the _Amazon Aurora User Guide_.

You can perform an in-place upgrade leveraging [Zero Downtime Patching (ZDP)](../aurorauserguide/auroramysql-updates-zdp.md), restore a snapshot, or initiate a managed blue/green upgrade using [Amazon RDS Blue/Green Deployments](../aurorauserguide/blue-green-deployments-overview.md) from any currently supported Aurora MySQL version 2 cluster into an Aurora MySQL version 3.10.4 cluster.

For information on planning an upgrade to Aurora MySQL version 3, see [Planning a major version upgrade for an Aurora MySQL cluster](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Planning). For general information about Aurora MySQL upgrades, see [Upgrading Aurora MySQL DB clusters](../aurorauserguide/auroramysql-updates-upgrading.md) in the _Amazon Aurora User Guide_.

For troubleshooting information, see [Troubleshooting for Aurora MySQL in-place upgrade](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Troubleshooting) in the _Amazon Aurora User Guide_.

If you have any questions or concerns, Support is available on the community forums and through [Support](https://aws.amazon.com/support). For more information, see [Maintaining an Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

## Improvements

**General improvements:**

- Fixed an issue where the first Zero-downtime restart (ZDR) after enabling Global Write Forwarding could fail to preserve existing connections on the writer.

**Availability improvements:**

- Fixed an issue which can cause a database instance restart when the lock table is full during semi-consistent read.

- Fixed an issue in binlog recovery that could cause binlog replica instance to restart when using \`aurora\_in\_memory\_relaylog\` with multi-threaded replication.

- Fixed an issue which could cause the writer to restart when a reader restarts while local or global write forwarding is enabled.

- Fixed an issue with improper handling of empty responses during connection loss with storage nodes.

## Integration of MySQL Community Edition bug fixes

This release includes all community bug fixes up to and including 8.0.42. For more information, see [MySQL bugs fixed by Aurora MySQL 3.x database engine updates](auroramysql-updates-mysqlbugs.md#AuroraMySQL.Updates.MySQLBugs.v3).

- InnoDB: Fixed an issue relating to DELETE operations. (Bug #37478594)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2025-11-13 (version 3.11.0, compatible with MySQL 8.0.43)

Aurora MySQL updates: 2026-01-02 (version 3.10.3, compatible with MySQL 8.0.42)

All content copied from https://docs.aws.amazon.com/.
