# Upgrading the minor version or patch level of an Aurora MySQL DB cluster

You can use the following methods to upgrade the minor version of a DB cluster or to patch a DB cluster:

- [Upgrading Aurora MySQL by modifying the engine version](auroramysql-updates-patching-modifyengineversion.md)
(for Aurora MySQL version 2 and 3)

- [Enabling automatic upgrades between minor Aurora MySQL versions](auroramysql-updates-amvu.md)

For information about how zero-downtime patching can reduce interruptions during the upgrade process, see
[Using zero-downtime patching](auroramysql-updates-zdp.md).

For information about performing a minor version upgrade for your Aurora MySQL DB cluster, see the following topics.

###### Topics

- [Before performing a minor version upgrade](#USER_UpgradeDBInstance.PostgreSQL.BeforeMinor)

- [Minor version upgrade prechecks for Aurora MySQL](#AuroraMySQL.minor-upgrade-prechecks)

- [Upgrading Aurora MySQL by modifying the engine version](auroramysql-updates-patching-modifyengineversion.md)

- [Enabling automatic upgrades between minor Aurora MySQL versions](auroramysql-updates-amvu.md)

- [Using zero-downtime patching](auroramysql-updates-zdp.md)

- [Alternative blue/green upgrade technique](#AuroraMySQL.UpgradingMinor.BlueGreen)

## Before performing a minor version upgrade

We recommend that you perform the following actions to reduce the downtime during a minor version upgrade:

- The Aurora DB cluster maintenance should be performed during a period of low traffic. Use Performance Insights to identify these time periods in order to configure the maintenance windows correctly.
For more information on Performance Insights, see [Monitoring DB load with Performance Insights on Amazon RDS](../userguide/user-perfinsights.md).
For more information on DB cluster maintenance window, [Adjusting the preferred DB cluster maintenance window](user-upgradedbinstance-maintenance.md#AdjustingTheMaintenanceWindow.Aurora).

- Use AWS SDKs that support exponential backoff and jitter as a best practice.
For more information, see [Exponential Backoff And Jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter).

## Minor version upgrade prechecks for Aurora MySQL

When you start a minor version upgrade, Amazon Aurora runs prechecks automatically.

These prechecks are mandatory. You can't choose to skip them. The prechecks provide the following benefits:

- They enable you to avoid unplanned downtime during the upgrade.

- If there are incompatibilities, Amazon Aurora prevents the upgrade and provides a log for you to learn about them. You
can then use the log to prepare your database for the upgrade by reducing the incompatibilities. For detailed
information about removing incompatibilities, see [Preparing \
your installation for upgrade](https://dev.mysql.com/doc/refman/8.0/en/upgrade-prerequisites.html) in the MySQL documentation.

The prechecks run before the DB instance is stopped for the upgrade, meaning that they don't cause any downtime when
they run. If the prechecks find an incompatibility, Aurora automatically cancels the upgrade before the DB instance is
stopped. Aurora also generates an event for the incompatibility. For more information about Amazon Aurora events, see
[Working with Amazon RDS event notification](user-events.md).

Aurora records detailed information about each incompatibility in the log file `PrePatchCompatibility.log`. In
most cases, the log entry includes a link to the MySQL documentation for correcting the incompatibility. For more
information about viewing log files, see [Viewing and listing database log files](user-logaccess-procedural-viewing.md).

Due to the nature of the prechecks, they analyze the objects in your database. This analysis results in resource
consumption and increases the time for the upgrade to complete.

## Alternative blue/green upgrade technique

In some situations, your top priority is to perform an immediate switchover from the old
cluster to an upgraded one. In such situations, you can use a multistep process that runs the
old and new clusters side-by-side. Here, you replicate data from the old cluster to the new
one until you are ready for the new cluster to take over. For details, see
[Using Amazon Aurora Blue/Green Deployments for database updates](blue-green-deployments.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Upgrading Amazon Aurora MySQL DB clusters

Modifying the engine version

All content copied from https://docs.aws.amazon.com/.
