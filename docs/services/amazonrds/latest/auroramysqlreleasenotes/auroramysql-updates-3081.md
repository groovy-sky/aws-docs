# Aurora MySQL database engine updates 2025-02-17 (version 3.08.1, compatible with MySQL 8.0.39)

**Version:** 3.08.1

Aurora MySQL 3.08.1 is generally available. Aurora MySQL 3.08 versions are compatible with MySQL 8.0.39. For more information on the
community changes that have occurred, see [MySQL 8.0 Release Notes](https://dev.mysql.com/doc/relnotes/mysql/8.0/en).

For details of the new features in Aurora MySQL version 3, see [Aurora MySQL version 3 compatible with MySQL 8.0](../aurorauserguide/auroramysql-mysql80.md). For
differences between Aurora MySQL version 3 and Aurora MySQL version 2, see [Comparison of Aurora MySQL version 2 and Aurora MySQL version\
3](../aurorauserguide/auroramysql-compare-v2-v3.md). For a comparison of Aurora MySQL version 3 and MySQL 8.0 Community Edition, see [Comparison of Aurora MySQL version 3 and MySQL 8.0 Community\
Edition](../aurorauserguide/auroramysql-compare-80-v3.md) in the _Amazon Aurora User Guide_.

Currently supported Aurora MySQL releases are 2.11.\*, 2.12.\*, 3.04.\*, 3.05.\*, 3.06.\*, 3.07.\*, and 3.08.\*.

You can perform an in-place upgrade, restore a snapshot, or initiate a managed blue/green upgrade using
[Amazon RDS Blue/Green Deployments](../aurorauserguide/blue-green-deployments-overview.md) from any currently
supported Aurora MySQL version 2 cluster into an Aurora MySQL version 3.08.1 cluster.

For information on planning an upgrade to Aurora MySQL version 3, see
[Planning \
a major version upgrade for an Aurora MySQL cluster](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Planning). For general information about Aurora MySQL upgrades, see
[Upgrading Amazon Aurora MySQL DB clusters](../aurorauserguide/auroramysql-updates-upgrading.md)
in the _Amazon Aurora User Guide_.

For troubleshooting information, see [Troubleshooting for Aurora MySQL \
in-place upgrade](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Troubleshooting) in the _Amazon Aurora User Guide_.

If you have any questions or concerns, AWS Support is available on the community forums and through
[AWS Support](https://aws.amazon.com/support). For more information, see
[Maintaining an Amazon Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in
the _Amazon Aurora User Guide_.

## Improvements

**Availability improvements:**

- Made improvements that reduce the DB instance restart time during database restarts, patch or minor version upgrades, and failovers.

**Upgrades and migrations**

- Fixed an issue that can cause the failure of Zero Downtime Patching (ZDP) while attempting to restore the roles and grants of dropped users. For more information about the `DROP USER` operation, see [DROP USER Statement](https://dev.mysql.com/doc/refman/8.0/en/drop-user.html) in the MySQL documentation.

## Integration of MySQL Community Edition bug fixes

This release includes all community bug fixes up to and including 8.0.39, in addition to the following. For more information, see
[MySQL\
bugs fixed by Aurora MySQL 3.x database engine updates](auroramysql-updates-mysqlbugs.md#AuroraMySQL.Updates.MySQLBugs.v3).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2025-04-07 (version 3.08.2, compatible with MySQL 8.0.39)

Aurora MySQL updates: 2024-11-18 (version 3.08.0, compatible with MySQL 8.0.39)

All content copied from https://docs.aws.amazon.com/.
