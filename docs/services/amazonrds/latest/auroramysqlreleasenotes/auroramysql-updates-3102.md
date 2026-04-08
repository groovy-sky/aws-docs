# Aurora MySQL database engine updates 2025-11-20 (version 3.10.2, compatible with MySQL 8.0.42)

**Version:** 3.10.2

Aurora MySQL 3.10.2 is generally available. Aurora MySQL 3.10 versions are compatible with MySQL 8.0.42. For more information on the community changes that have occurred, see [MySQL 8.0 Release Notes](https://dev.mysql.com/doc/relnotes/mysql/8.0/en).

For details of the new features in Aurora MySQL version 3, see [Aurora MySQL version 3 compatible with MySQL 8.0](../aurorauserguide/auroramysql-mysql80.md). For differences between Aurora MySQL version 3 and Aurora MySQL version 2, see [Comparison of Aurora MySQL version 2 and Aurora MySQL version 3](../aurorauserguide/auroramysql-compare-v2-v3.md). For a comparison of Aurora MySQL version 3 and MySQL 8.0 Community Edition, see [Comparison of Aurora MySQL version 3 and MySQL 8.0 Community Edition](../aurorauserguide/auroramysql-compare-80-v3.md) in the _Amazon Aurora User Guide_.

You can perform an in-place upgrade leveraging [Zero Downtime Patching (ZDP)](../aurorauserguide/auroramysql-updates-zdp.md), restore a snapshot, or initiate a managed blue/green upgrade using [Amazon RDS Blue/Green Deployments](../aurorauserguide/blue-green-deployments-overview.md) from any currently supported Aurora MySQL version 2 cluster into an Aurora MySQL version 3.10.2 cluster.

For information on planning an upgrade to Aurora MySQL version 3, see [Planning a major version upgrade for an Aurora MySQL cluster](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Planning). For general information about Aurora MySQL upgrades, see [Upgrading Aurora MySQL DB clusters](../aurorauserguide/auroramysql-updates-upgrading.md) in the _Amazon Aurora User Guide_.

For troubleshooting information, see [Troubleshooting for Aurora MySQL in-place upgrade](../aurorauserguide/auroramysql-updates-majorversionupgrade.md#AuroraMySQL.Upgrading.Troubleshooting) in the _Amazon Aurora User Guide_.

If you have any questions or concerns, Support is available on the community forums and through [Support](https://aws.amazon.com/support). For more information, see [Maintaining an Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

## Improvements

**Availability improvements:**

- Fixed an issue which can cause engine restart during Zero-downtime patching (ZDP)/Zero-downtime restart (ZDR), when preserving SSL/TLS connections with active transactions.

- Fixed an issue which, can cause an engine restart when running `KILL <query-id>` after running `EXPLAIN FOR CONNECTION <query-id>` on a running Parallel Query.

- Fixed an issue which could cause writer to become unresponsive when write forwarding is being disabled while two or more replicas are forwarding requests.

- Fixed an issue which could cause the writer instance to restart when global write forwarding or local write forwarding is being disabled.

**General improvements:**

- Fixed an issue causing inaccurate tracking of Parallel Query requests while running `EXPLAIN ANALYZE` statements where `Aurora_pq_request_in_progress` counter was not updated accurately.

- Automatically disable `aurora_oom_response` actions (except print, if configured) when `aurora_oom_response` fails to resolve memory pressure after a threshold time (in the order of few mins).

- Improved write IOPS performance when system variable `innodb_flush_log_at_trx_commit` is set to `0`.

- Fixed an issue with replicas incorrectly being restarted when joining the writer.

- Fixed an issue where writes to the database may stall while executing a long running transaction leading to a database restart or a major version upgrade to fail.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2026-01-02 (version 3.10.3, compatible with MySQL 8.0.42)

Aurora MySQL updates: 2025-09-30 (version 3.10.1, compatible with MySQL 8.0.42)

All content copied from https://docs.aws.amazon.com/.
