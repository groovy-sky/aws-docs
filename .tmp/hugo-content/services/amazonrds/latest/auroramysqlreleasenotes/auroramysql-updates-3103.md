# Aurora MySQL database engine updates 2026-01-02 (version 3.10.3, compatible with MySQL 8.0.42)

**Version:** 3.10.3

Aurora MySQL 3.10.3 is generally available. Aurora MySQL 3.10 versions are compatible with MySQL 8.0.42. For more information on the community changes that have occurred, see [MySQL 8.0 Release Notes](https://dev.mysql.com/doc/relnotes/mysql/8.0/en).

For details of the new features in Aurora MySQL version 3, see [Aurora MySQL version 3 compatible with MySQL 8.0](../aurorauserguide/auroramysql-mysql80.md). For differences between Aurora MySQL version 3 and Aurora MySQL version 2, see [Comparison of Aurora MySQL version 2 and Aurora MySQL version 3](../aurorauserguide/auroramysql-compare-v2-v3.md). For a comparison of Aurora MySQL version 3 and MySQL 8.0 Community Edition, see [Comparison of Aurora MySQL version 3 and MySQL 8.0 Community Edition](../aurorauserguide/auroramysql-compare-80-v3.md) in the _Amazon Aurora User Guide_.

You can perform an in-place upgrade leveraging [Zero Downtime Patching (ZDP)](../aurorauserguide/auroramysql-updates-zdp.md), restore a snapshot, or initiate a managed blue/green upgrade using [Amazon RDS Blue/Green Deployments](../aurorauserguide/blue-green-deployments-overview.md) from any currently supported Aurora MySQL version 2 cluster into an Aurora MySQL version 3.10.3 cluster.

For information on planning an upgrade to Aurora MySQL version 3, see [Planning a major version upgrade for an Aurora MySQL cluster](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Planning). For general information about Aurora MySQL upgrades, see [Upgrading Aurora MySQL DB clusters](../aurorauserguide/auroramysql-updates-upgrading.md) in the _Amazon Aurora User Guide_.

For troubleshooting information, see [Troubleshooting for Aurora MySQL in-place upgrade](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Troubleshooting) in the _Amazon Aurora User Guide_.

If you have any questions or concerns, Support is available on the community forums and through [Support](https://aws.amazon.com/support). For more information, see [Maintaining an Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

## Improvements

**General improvements:**

- Reduced CPU usage overhead while establishing Encryption in Transit between the database instance and the storage layer.

- Fixed an issue which can cause some SQL statements to not get logged in the audit log.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2026-03-30 (version 3.10.4, compatible with MySQL 8.0.42)

Aurora MySQL updates: 2025-11-20 (version 3.10.2, compatible with MySQL 8.0.42)

All content copied from https://docs.aws.amazon.com/.
