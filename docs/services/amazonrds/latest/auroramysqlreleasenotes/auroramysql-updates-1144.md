# Aurora MySQL database engine updates: 2018-03-13 (version 1.14.4) (Deprecated)

**Version:** 1.14.4

Aurora MySQL 1.14.4 is generally available. You can create new DB clusters in Aurora 1.14.4,
using the AWS CLI or the Amazon RDS API and specifying the engine version. You have the
option, but are not required, to upgrade existing 1.14.x DB clusters to Aurora 1.14.4.

With version 1.14.4 of Aurora, we are using a cluster-patching model where all nodes in an Aurora
DB cluster are patched at the same time. We support zero-downtime patching, which works on a best-effort
basis to preserve client connections through the patching process. For more information, see
[Maintaining an Amazon Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

If you have any questions or concerns, AWS Support is available on the
community forums and through [AWS Support](https://aws.amazon.com/support). For more
information, see [Maintaining an Amazon Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

## Zero-downtime patching

The zero-downtime patching (ZDP) feature attempts, on a best-effort basis, to preserve client connections through an engine
patch. For more information about ZDP, see
[Using zero-downtime patching](../aurorauserguide/auroramysql-updates-patching.md#AuroraMySQL.Updates.ZDP) in the _Amazon Aurora User Guide_.

## New features

- Aurora MySQL now supports db.r4 instance classes.

## Improvements

- Fixed an issue where `LOST_EVENTS` were generated when writing large binlog events.

## Integration of MySQL bug fixes

- Ignorable events don't work and are not tested (Bug #74683)

- NEW->OLD ASSERT FAILURE 'GTID\_MODE > 0' (Bug #20436436)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2017-10-24 (version 1.15) (Deprecated)

Aurora MySQL updates: 2017-09-22 (version 1.14.1) (Deprecated)

All content copied from https://docs.aws.amazon.com/.
