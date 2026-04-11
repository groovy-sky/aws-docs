# Modifying an Amazon Aurora DB cluster

You can change the settings of a DB cluster to accomplish tasks such as changing its backup retention period or its database
port. You can also modify DB instances in a DB cluster to accomplish tasks such as changing its DB instance class or enabling
Performance Insights for it. This topic guides you through modifying an Aurora DB cluster and its DB instances, and describes the
settings for each.

We recommend that you test any changes on a test DB cluster or DB instance before modifying a production DB cluster or DB
instance, so that you fully understand the impact of each change. This is especially important when upgrading database
versions.

###### Topics

- [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster)

- [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance)

- [Changing the password for the database master user](#Aurora.Modifying.Password)

- [Settings for Amazon Aurora](#Aurora.Modifying.Settings)

- [Settings that don't apply to Amazon Aurora DB clusters](#Aurora.Modifying.SettingsNotApplicableDBClusters)

- [Settings that don't apply to Amazon Aurora DB instances](#Aurora.Modifying.SettingsNotApplicable)

## Modifying the DB cluster by using the console, CLI, and API

You can modify a DB cluster using the AWS Management Console, the AWS CLI, or the RDS API.

###### Note

Most modifications can be applied immediately or during the next scheduled maintenance window. Some modifications, such as
turning on deletion protection, are applied immediately—regardless of when you choose to apply them.

Changing the master password in the AWS Management Console is always applied immediately.

If you're using SSL endpoints and change the DB cluster identifier, stop and restart the DB cluster to update the SSL
endpoints. For more information, see [Stopping and starting an Amazon Aurora DB cluster](aurora-cluster-stop-start.md).

###### To modify a DB cluster

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Databases**, and then select the DB cluster that you
    want to modify.

3. Choose **Modify**.
    The **Modify DB cluster** page appears.

4. Change any of the settings that you want.
    For information about each setting, see
    [Settings for Amazon Aurora](#Aurora.Modifying.Settings).

###### Note

In the AWS Management Console, some instance level changes only apply to the current DB instance, while others
apply to the entire DB cluster. For information about whether a setting applies to the DB instance or the DB
cluster, see the scope for the setting in [Settings for Amazon Aurora](#Aurora.Modifying.Settings). To change a setting that modifies the entire DB cluster at
the instance level in the AWS Management Console, follow the instructions in [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

5. When all the changes are as you want them, choose **Continue** and check the summary of
    modifications.

6. To apply the changes immediately, select **Apply immediately**.

7. On the confirmation page, review your changes.
    If they are correct, choose **Modify cluster**
    to save your changes.

Alternatively, choose **Back** to edit your changes,
    or choose **Cancel** to cancel your changes.

To modify a DB cluster using the AWS CLI, call the
[modify-db-cluster](../../../cli/latest/reference/rds/modify-db-cluster.md) command.
Specify the DB cluster identifier, and the values for the settings that you want to modify.
For information about each setting, see
[Settings for Amazon Aurora](#Aurora.Modifying.Settings).

###### Note

Some settings only apply to DB instances. To change those settings,
follow the instructions in [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

###### Example

The following command modifies `mydbcluster`
by setting the backup retention period to 1 week (7 days).

For Linux, macOS, or Unix:

```nohighlight

aws rds modify-db-cluster \
    --db-cluster-identifier mydbcluster \
    --backup-retention-period 7

```

For Windows:

```nohighlight

aws rds modify-db-cluster ^
    --db-cluster-identifier mydbcluster ^
    --backup-retention-period 7

```

To modify a DB cluster using the Amazon RDS API, call the
[ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) operation.
Specify the DB cluster identifier, and the values for the settings that you want to modify.
For information about each parameter, see
[Settings for Amazon Aurora](#Aurora.Modifying.Settings).

###### Note

Some settings only apply to DB instances. To change those settings,
follow the instructions in [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

## Modifying a DB instance in a DB cluster

You can modify a DB instance in a DB cluster using the AWS Management Console, the AWS CLI, or the RDS API.

When you modify a DB instance, you can apply the changes immediately.
To apply changes immediately, you select the **Apply Immediately** option in the AWS Management Console,
you use the `--apply-immediately` parameter when calling the AWS CLI,
or you set the `ApplyImmediately` parameter to `true` when using the Amazon RDS API.

If you don't choose to apply changes immediately, the changes are deferred until the next maintenance window. During
the next maintenance window, any of these deferred changes are applied. If you choose to apply changes immediately, your new
changes and any previously deferred changes are applied.

To see the modifications that are pending for the next maintenance window, use the [describe-db-clusters](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/describe-db-clusters.html) AWS CLI command and check the `PendingModifiedValues` field.

###### Important

If any of the deferred modifications require downtime, choosing **Apply immediately** can cause
unexpected downtime for the DB instance. There is no downtime for the other DB instances in the DB cluster.

Modifications that you defer aren't listed in the output of the `describe-pending-maintenance-actions`
CLI command. Maintenance actions only include system upgrades that you schedule for the next maintenance window.

###### To modify a DB instance in a DB cluster

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Databases**, and then select the DB instance that you want
    to modify.

3. For **Actions**, choose **Modify**. The **Modify DB**
**instance** page appears.

4. Change any of the settings that you want. For information about each setting, see [Settings for Amazon Aurora](#Aurora.Modifying.Settings).

###### Note

Some settings apply to the entire DB cluster and must be changed at the cluster level. To change those
settings, follow the instructions in [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

In the AWS Management Console, some instance level changes only apply to the current DB instance, while others apply
to the entire DB cluster. For information about whether a setting applies to the DB instance or the DB
cluster, see the scope for the setting in [Settings for Amazon Aurora](#Aurora.Modifying.Settings).

5. When all the changes are as you want them, choose **Continue** and check the summary of
    modifications.

6. To apply the changes immediately, select **Apply immediately**.

7. On the confirmation page, review your changes. If they are correct, choose **Modify DB**
**instance** to save your changes.

Alternatively, choose **Back** to edit your changes, or choose **Cancel** to
    cancel your changes.

To modify a DB instance in a DB cluster by using the AWS CLI, call the [modify-db-instance](../../../cli/latest/reference/rds/modify-db-instance.md) command. Specify the DB instance
identifier, and the values for the settings that you want to modify. For information about each parameter, see [Settings for Amazon Aurora](#Aurora.Modifying.Settings).

###### Note

Some settings apply to the entire DB cluster. To change those settings,
follow the instructions in [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

###### Example

The following code modifies `mydbinstance`
by setting the DB instance class to `db.r4.xlarge`.
The changes are applied during the next maintenance window
by using `--no-apply-immediately`.
Use `--apply-immediately` to apply the changes immediately.

For Linux, macOS, or Unix:

```nohighlight

aws rds modify-db-instance \
    --db-instance-identifier mydbinstance \
    --db-instance-class db.r4.xlarge \
    --no-apply-immediately
```

For Windows:

```nohighlight

aws rds modify-db-instance ^
    --db-instance-identifier mydbinstance ^
    --db-instance-class db.r4.xlarge ^
    --no-apply-immediately
```

To modify a DB instance by using the Amazon RDS API, call the [ModifyDBInstance](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) operation. Specify the DB instance identifier, and the values for the settings that you
want to modify. For information about each parameter, see [Settings for Amazon Aurora](#Aurora.Modifying.Settings).

###### Note

Some settings apply to the entire DB cluster. To change those settings,
follow the instructions in [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

## Changing the password for the database master user

You can use the AWS Management Console or the AWS CLI to change the master user password.

You modify the writer DB instance to change the master user password using the AWS Management Console.

###### To change the master user password

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Databases**, and then select the DB instance that you want
    to modify.

3. For **Actions**, choose **Modify**.

The **Modify DB instance** page appears.

4. Enter a **New master password**.

5. For **Confirm master password**, enter the same new password.

![Enter a new master user password and confirm it.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/aur_new_master_password.png)

6. Choose **Continue** and check the summary of modifications.

###### Note

Password changes are always applied immediately.

7. On the confirmation page, choose **Modify DB instance**.

You call the [modify-db-cluster](../../../cli/latest/reference/rds/modify-db-cluster.md) command to change the
master user password using the AWS CLI. Specify the DB cluster identifier and the new password, as shown in the following
examples.

You don't need to specify `--apply-immediately|--no-apply-immediately`, because password changes are always
applied immediately.

For Linux, macOS, or Unix:

```nohighlight

aws rds modify-db-cluster \
    --db-cluster-identifier mydbcluster \
    --master-user-password mynewpassword

```

For Windows:

```nohighlight

aws rds modify-db-cluster ^
    --db-cluster-identifier mydbcluster ^
    --master-user-password mynewpassword

```

## Settings for Amazon Aurora

The following table contains details about which settings you can modify, the methods for modifying the setting,
and the scope of the setting. The scope determines whether the setting applies to the entire DB cluster or if it
can be set only for specific DB instances.

###### Note

Additional settings are available if you are modifying an Aurora Serverless v1 or Aurora Serverless v2 DB cluster. For
information about these settings, see [Modifying an Aurora Serverless v1 DB cluster](aurora-serverless-modifying.md) and [Managing Aurora Serverless v2 DB clusters](aurora-serverless-v2-administration.md).

Some settings aren't available for Aurora Serverless v1 and Aurora Serverless v2 because of their limitations. For more
information, see [Limitations of Aurora Serverless v1](aurora-serverless.md#aurora-serverless.limitations) and [Requirements and limitations for Aurora Serverless v2](aurora-serverless-v2-requirements.md).

Setting and descriptionMethodScopeDowntime notes

**Auto minor version upgrade**

Whether you want the DB instance to receive preferred minor engine version upgrades automatically
when they become available. Upgrades are installed only during your scheduled maintenance window.

For more information about engine updates, see
[Database engine updates for Amazon Aurora PostgreSQL](aurorapostgresql-updates.md)
and [Database engine updates for Amazon Aurora MySQL](auroramysql-updates.md).
For more information about the **Auto minor version upgrade** setting
for Aurora MySQL, see [Enabling automatic upgrades between minor Aurora MySQL versions](auroramysql-updates-amvu.md).

###### Note

This setting is enabled by default. For each new cluster, choose the appropriate value for this setting
based on its importance, expected lifetime, and the amount of verification testing that you do after
each upgrade.

When you change this setting, perform this modification for every DB instance in your Aurora cluster. If any DB
instance in your cluster has this setting turned off, the cluster isn't automatically upgraded.

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [modify-db-instance](../../../cli/latest/reference/rds/modify-db-instance.md) and set the
`--auto-minor-version-upgrade|--no-auto-minor-version-upgrade` option.

Using the RDS API, call [ModifyDBInstance](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `AutoMinorVersionUpgrade` parameter.

The entire DB cluster

An outage doesn't occur during this change. Outages do occur during future maintenance windows when
Aurora applies automatic upgrades.

**Backup retention period**

The number of days that
automatic backups are retained.
The minimum value is `1`.

For more information, see
[Backups](aurora-managing-backups.md#Aurora.Managing.Backups.Backup).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [modify-db-cluster](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--backup-retention-period` option.

Using the RDS API, call [ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `BackupRetentionPeriod` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**Backup window (Start time)**

The time range during which automated backups of your database
occurs. The backup window is a start time in Universal Coordinated
Time (UTC), and a duration in hours.

Aurora backups are continuous and incremental, but the backup window is used to create
a daily system backup that is preserved within the backup retention period. You can
copy it to preserve it outside of the retention period.

The maintenance window and the backup window for the DB cluster can't overlap.

For more information, see [Backup window](aurora-managing-backups.md#Aurora.Managing.Backups.BackupWindow).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [modify-db-cluster](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--preferred-backup-window` option.

Using the RDS API, call [ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `PreferredBackupWindow` parameter.

The entire DB cluster.

An outage doesn't occur during this change.

**Capacity settings**

The scaling properties of an Aurora Serverless v1 DB cluster. You can only modify scaling properties for DB
clusters in `serverless` DB engine mode.

For information about Aurora Serverless v1, see [Using Amazon Aurora Serverless v1](aurora-serverless.md).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [modify-db-cluster](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--scaling-configuration` option.

Using the RDS API, call [ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `ScalingConfiguration` parameter.

The entire DB cluster

An outage doesn't occur during this change.

The change occurs immediately. This setting ignores the apply immediately setting.

**Certificate authority**

The certificate authority (CA) for the server certificate used by the DB instance.

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [modify-db-instance](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--ca-certificate-identifier`
option.

Using the RDS API, call [ModifyDBInstance](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `CACertificateIdentifier` parameter.

Only the specified DB instance

An outage only occurs if the DB engine doesn't support rotation without restart.
You can use the [describe-db-engine-versions](../../../cli/latest/reference/rds/describe-db-engine-versions.md) AWS CLI command to determine whether
the DB engine supports rotation without restart.

**Cluster storage configuration**

The storage type for the DB cluster: **Aurora I/O-Optimized** or **Aurora Standard**.

For more information, see [Storage configurations for Amazon Aurora DB clusters](aurora-overview-storagereliability.md#aurora-storage-type).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [modify-db-cluster](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--storage-type` option.

Using the RDS API, call [ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `StorageType` parameter.

The entire DB cluster

Changing the storage type of an Aurora PostgreSQL DB cluster with Optimized Reads instance classes causes an outage.
This does not occur when changing storage types for clusters with other instance class types.
For more information on the DB instance class types, see [DB instance class types](concepts-dbinstanceclass-types.md).

**Copy tags to snapshots**

Select to specify that tags defined for this DB cluster are copied
to DB snapshots created from this DB cluster. For more information,
see [Tagging Amazon Aurora andAmazon RDS resources](user-tagging.md).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [https://docs.aws.amazon.com/cli/latest/reference/rds/modify-db-cluster.html](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--copy-tags-to-snapshot` or
`--no-copy-tags-to-snapshot` option.

Using the RDS API, call [https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API\_ModifyDBCluster.html](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `CopyTagsToSnapshot` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**Data API**

You can access Aurora Serverless v1 with web services–based applications, including AWS Lambda and
AWS AppSync.

This setting only applies to an Aurora Serverless v1 DB cluster.

For more information, see
[Using the Amazon RDS Data API](data-api.md).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--enable-http-endpoint`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `EnableHttpEndpoint` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**Database authentication**

The database authentication you want to use.

For
MySQL:

- Choose **Password authentication** to
authenticate database users with database passwords
only.

- Choose **Password and IAM database**
**authentication** to authenticate database users
with database passwords and user credentials through IAM
users and roles. For more information, see [IAM database authentication](usingwithrds-iamdbauth.md).

For PostgreSQL:

- Choose **IAM database authentication** to
authenticate database users with database passwords and user
credentials through users and roles. For more
information, see [IAM database authentication](usingwithrds-iamdbauth.md).

- Choose **Kerberos authentication** to
authenticate database passwords and user credentials using
Kerberos authentication. For more information, see [Using Kerberos authentication with Aurora PostgreSQL](postgresql-kerberos.md).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [modify-db-cluster](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the following options:

- For IAM authentication, set the
`--enable-iam-database-authentication|--no-enable-iam-database-authentication`
option.

- For Kerberos authentication, set the `--domain`
and `--domain-iam-role-name` options.

Using the RDS API, call [ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the following parameters:

- For IAM authentication, set the
`EnableIAMDatabaseAuthentication`
parameter.

- For Kerberos authentication, set the `Domain`
and `DomainIAMRoleName` parameters.

The entire DB cluster

An outage doesn't occur during this change.

**Database port**

The port that you want
to use to access the DB cluster.

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--port`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `Port` parameter.

The entire DB cluster

An outage occurs during this change. All of the DB instances in the DB cluster are rebooted immediately.

**DB cluster identifier**

The DB cluster identifier. This value is stored as a lowercase string.

When you change the DB cluster identifier, the DB cluster endpoints change. The endpoints of the DB
instances in the DB cluster don't change.

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--new-db-cluster-identifier`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `NewDBClusterIdentifier` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**DB cluster parameter group**

The DB cluster parameter group that you want
associated with the DB cluster.

For more information, see
[Parameter groups for Amazon Aurora](user-workingwithparamgroups.md).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--db-cluster-parameter-group-name`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `DBClusterParameterGroupName` parameter.

The entire DB cluster

An outage doesn't occur during this change. When you change the parameter group, changes to some parameters
are applied to the DB instances in the DB cluster immediately without a reboot. Changes to other parameters are
applied only after the DB instances in the DB cluster are rebooted.

**DB instance class**

The DB instance class that you want to use.

For more information,
see [Amazon AuroraDB instance classes](concepts-dbinstanceclass.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--db-instance-class`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `DBInstanceClass` parameter.

Only the specified DB instance

An outage occurs during this change.

**DB instance identifier**

The DB instance identifier.
This value is stored as a lowercase string.

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--new-db-instance-identifier`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `NewDBInstanceIdentifier` parameter.

Only the specified DB instance

Downtime occurs during this change.

RDS restarts the DB instance to update the following:

- Aurora MySQL – `SERVER_ID` column in the
`information_schema.replica_host_status`
table

- Aurora PostgreSQL – `server_id` column in
the `aurora_replica_status()` function

**DB parameter group**

The DB parameter group that you want
associated with the DB instance.

For more information, see
[Parameter groups for Amazon Aurora](user-workingwithparamgroups.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--db-parameter-group-name`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `DBParameterGroupName` parameter.

Only the specified DB instance

An outage doesn't occur during this change.

When you associate a new DB parameter group with a DB instance, the modified static and dynamic parameters are applied
only after the DB instance is rebooted. However, if you modify dynamic parameters in the DB parameter group
after you associate it with the DB instance, these changes are applied immediately without a reboot.

For more information, see [Parameter groups for Amazon Aurora](user-workingwithparamgroups.md) and [Rebooting an Amazon Aurora DB cluster or Amazon Aurora DB instance](user-rebootcluster.md).

**Deletion protection**

**Enable deletion protection** to prevent your DB
cluster from being deleted. For more information, see [Deletion protection for Aurora clusters](user-deletecluster.md#USER_DeletionProtection).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the
`--deletion-protection|--no-deletion-protection` option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the
`DeletionProtection` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**Engine version**

The version of the DB engine that you want to use. Before you upgrade your production DB cluster,
we recommend that you test the upgrade process on a test DB cluster to verify its duration and to
validate your applications.

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--engine-version`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `EngineVersion` parameter.

The entire DB cluster

An outage occurs during this change.

**Enhanced monitoring**

**Enable enhanced monitoring** to enable gathering metrics
in real time for the operating system that your DB instance runs on.

For more information, see
[Monitoring OS metrics with Enhanced Monitoring](user-monitoring-os.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--monitoring-role-arn`
and `--monitoring-interval` options.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `MonitoringRoleArn` and
`MonitoringInterval` parameters.

Only the specified DB instance

An outage doesn't occur during this change.

**Log exports**

Select the log types to publish to Amazon CloudWatch Logs.

For more information, see
[AuroraMySQL database log files](user-logaccess-concepts-mysql.md).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--cloudwatch-logs-export-configuration`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `CloudwatchLogsExportConfiguration` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**Maintenance window**

The time range during which system maintenance occurs.
System maintenance includes upgrades, if applicable.
The maintenance window is a start time
in Universal Coordinated Time (UTC),
and a duration in hours.

If you set the window to the current time,
there must be at least 30 minutes
between the current time and end of the window
to ensure any pending changes are applied.

You can set the maintenance window independently for the DB cluster and
for each DB instance in the DB cluster. When the scope of a modification is
the entire DB cluster, the modification is performed during the DB cluster
maintenance window. When the scope of a modification is
the a DB instance, the modification is performed during
maintenance window of that DB instance.

The maintenance window and the backup window for the DB cluster can't overlap.

For more information, see
[Amazon RDS maintenance window](user-upgradedbinstance-maintenance.md#Concepts.DBMaintenance).

To change the maintenance window for the DB cluster using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

To change the maintenance window for a DB instance using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

To change the maintenance window for the DB cluster using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--preferred-maintenance-window`
option.

To change the maintenance window for a DB instance using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--preferred-maintenance-window`
option.

To change the maintenance window for the DB cluster using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `PreferredMaintenanceWindow` parameter.

To change the maintenance window for a DB instance using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `PreferredMaintenanceWindow` parameter.

The entire DB cluster or a single DB instance

If there are one or more pending actions that cause an outage, and the maintenance window is changed to include the current time,
then those pending actions are applied immediately, and an outage occurs.

**Manage master credentials in AWS Secrets Manager**

Select **Manage master credentials in AWS Secrets Manager** to manage the master user
password in a secret in Secrets Manager.

Optionally, choose a KMS key to use to protect the secret.
Choose from the KMS keys in your account, or enter the key from a
different account.

For more information, see [Password management with Amazon Aurora and AWS Secrets Manager](rds-secrets-manager.md).

If Aurora is already managing the master user password for the DB cluster, you can rotate
the master user password by choosing **Rotate secret immediately**.

For more information, see [Password management with Amazon Aurora and AWS Secrets Manager](rds-secrets-manager.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the
`--manage-master-user-password | --no-manage-master-user-password`
and `--master-user-secret-kms-key-id` options. To rotate the master user
password immediately, set the `--rotate-master-user-password` option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `ManageMasterUserPassword`
and `MasterUserSecretKmsKeyId` parameters. To rotate the master user
password immediately, set the `RotateMasterUserPassword` parameter to
`true`.

The entire DB cluster

An outage doesn't occur during this change.

**Network type**

The IP addressing protocols supported by the DB cluster.

**IPv4** to specify that resources can communicate with the DB cluster only
over the IPv4 addressing protocol.

**Dual-stack mode** to specify that resources can communicate with the
DB cluster over IPv4, IPv6, or both. Use dual-stack mode if you have any resources that
must communicate with your DB cluster over the IPv6 addressing protocol. To use dual-stack
mode, make sure at least two subnets spanning two Availability Zones that support both the IPv4
and IPv6 network protocol. Also, make sure you associate an IPv6 CIDR block with subnets in the
DB subnet group you specify.

For more information, see [Amazon Aurora IP addressing](user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.IP_addressing).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--network-type`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `NetworkType` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**New master password**

The password for your master user.

- For Aurora MySQL, the password must contain 8–41 printable ASCII characters.

- For Aurora PostgreSQL, it must contain 8–99 printable ASCII characters.

- It can't contain `/`, `"`, `@`, or a space.

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--master-user-password`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `MasterUserPassword` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**Performance Insights**

Whether to enable Performance Insights, a tool that monitors your DB instance
load so that you can analyze and troubleshoot your database performance.

For more information, see
[Monitoring DB load with Performance Insights on Amazon Aurora](user-perfinsights.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--enable-performance-insights|--no-enable-performance-insights`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `EnablePerformanceInsights` parameter.

Only the specified DB instance

An outage doesn't occur during this change.

**Performance Insights AWS KMS key**

The AWS KMS key identifier for encryption of Performance Insights data. The KMS key identifier is the
Amazon Resource Name (ARN), key identifier, or key alias for the KMS key.

For more information, see
[Turning Performance Insights on and off for Aurora](user-perfinsights-enabling.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--performance-insights-kms-key-id`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `PerformanceInsightsKMSKeyId` parameter.

Only the specified DB instance

An outage doesn't occur during this change.

**Performance Insights retention period**

The amount of time, in days, to retain Performance Insights data.
The retention setting is **Default (7 days)**. To retain your performance
data for longer, specify 1–24 months. For more information about retention periods, see
[Pricing and data retention for Performance Insights](user-perfinsights-overview-cost.md).

For more information, see
[Turning Performance Insights on and off for Aurora](user-perfinsights-enabling.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--performance-insights-retention-period`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `PerformanceInsightsRetentionPeriod` parameter.

Only the specified DB instance

An outage doesn't occur during this change.

**Promotion tier**

A value that specifies the order in which an Aurora Replica is promoted to the primary instance in a DB
cluster, after a failure of the existing primary instance.

For more information, see
[Fault tolerance for an Aurora DB cluster](concepts-aurorahighavailability.md#Aurora.Managing.FaultTolerance).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--promotion-tier`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `PromotionTier` parameter.

Only the specified DB instance

An outage doesn't occur during this change.

**Public access**

**Publicly accessible** to give the DB instance a public IP
address, meaning that it's accessible outside the VPC. To be
publicly accessible, the DB instance also has to be in a public
subnet in the VPC.

**Not publicly accessible** to make the DB instance accessible only
from inside the VPC.

For more information, see [Hiding a DB cluster in a VPC from the internet](user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.Hiding).

To connect to a DB instance from outside of its Amazon VPC, the DB instance must be publicly accessible, access must be granted
using the inbound rules of the DB instance's security group, and other requirements must be met. For more information,
see [Can't connect to Amazon RDS DB instance](chap-troubleshooting.md#CHAP_Troubleshooting.Connecting).

If your DB instance is isn't publicly accessible, you can also use an AWS Site-to-Site VPN connection or
an Direct Connect connection to access it from a private network. For more information, see
[Internetwork traffic privacy](inter-network-traffic-privacy.md).

Using the AWS Management Console, [Modifying a DB instance in a DB cluster](#Aurora.Modifying.Instance).

Using the AWS CLI, run [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and set the `--publicly-accessible|--no-publicly-accessible`
option.

Using the RDS API, call [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) and set the `PubliclyAccessible` parameter.

Only the specified DB instance

An outage doesn't occur during this change.

**Serverless v2 capacity settings**

The database capacity of an Aurora Serverless v2 DB cluster, measured in Aurora Capacity Units
(ACUs).

For more information, see [Setting the Aurora Serverless v2 capacity range for a cluster](aurora-serverless-v2-administration.md#aurora-serverless-v2-setting-acus).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--serverless-v2-scaling-configuration`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `ServerlessV2ScalingConfiguration`
parameter.

The entire DB cluster

An outage doesn't occur during this change.

The change occurs immediately. This setting ignores the apply immediately setting.

**Security group**

The security group
you want associated with the DB cluster.

For more information, see
[Controlling access with security groups](overview-rdssecuritygroups.md).

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--vpc-security-group-ids`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `VpcSecurityGroupIds` parameter.

The entire DB cluster

An outage doesn't occur during this change.

**Target Backtrack window**

The amount of time you want to be able to backtrack your DB cluster, in seconds. This
setting is available only for Aurora MySQL and only if the DB cluster was created with Backtrack enabled.

Using the AWS Management Console, [Modifying the DB cluster by using the console, CLI, and API](#Aurora.Modifying.Cluster).

Using the AWS CLI, run [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and set the `--backtrack-window`
option.

Using the RDS API, call [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) and set the `BacktrackWindow` parameter.

The entire DB cluster

An outage doesn't occur during this change.

## Settings that don't apply to Amazon Aurora DB clusters

The following settings in the AWS CLI command [`modify-db-cluster`](../../../cli/latest/reference/rds/modify-db-cluster.md) and the RDS API operation [`ModifyDBCluster`](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md) don't apply to Amazon Aurora DB clusters.

###### Note

You can't use the AWS Management Console to modify these settings for Aurora DB
clusters.

AWS CLI settingRDS API setting

`--allocated-storage`

`AllocatedStorage`

`--auto-minor-version-upgrade | --no-auto-minor-version-upgrade`

`AutoMinorVersionUpgrade`

`--db-cluster-instance-class`

`DBClusterInstanceClass`

`--enable-performance-insights | --no-enable-performance-insights`

`EnablePerformanceInsights`

`--iops`

`Iops`

`--monitoring-interval`

`MonitoringInterval`

`--monitoring-role-arn`

`MonitoringRoleArn`

`--option-group-name`

`OptionGroupName`

`--performance-insights-kms-key-id`

`PerformanceInsightsKMSKeyId`

`--performance-insights-retention-period`

`PerformanceInsightsRetentionPeriod`

## Settings that don't apply to Amazon Aurora DB instances

The following settings in the AWS CLI command [`modify-db-instance`](../../../cli/latest/reference/rds/modify-db-instance.md) and the RDS API operation [`ModifyDBInstance`](../../../../reference/amazonrds/latest/apireference/api-modifydbinstance.md) don't apply to Amazon Aurora DB instances.

###### Note

You can't use the AWS Management Console to modify these settings for Aurora DB
instances.

AWS CLI settingRDS API setting

`--allocated-storage`

`AllocatedStorage`

`--allow-major-version-upgrade|--no-allow-major-version-upgrade`

`AllowMajorVersionUpgrade`

`--copy-tags-to-snapshot|--no-copy-tags-to-snapshot`

`CopyTagsToSnapshot`

`--domain`

`Domain`

`--db-security-groups`

`DBSecurityGroups`

`--db-subnet-group-name`

`DBSubnetGroupName`

`--domain-iam-role-name`

`DomainIAMRoleName`

`--multi-az|--no-multi-az`

`MultiAZ`

`--iops`

`Iops`

`--license-model`

`LicenseModel`

`--network-type`

`NetworkType`

`--option-group-name`

`OptionGroupName`

`--processor-features`

`ProcessorFeatures`

`--storage-type`

`StorageType`

`--tde-credential-arn`

`TdeCredentialArn`

`--tde-credential-password`

`TdeCredentialPassword`

`--use-default-processor-features|--no-use-default-processor-features`

`UseDefaultProcessorFeatures`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Connecting a Lambda function

Adding Aurora Replicas

All content copied from https://docs.aws.amazon.com/.
