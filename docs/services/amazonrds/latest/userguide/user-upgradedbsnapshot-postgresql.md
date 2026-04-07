# Upgrading a PostgreSQL DB snapshot engine version

With Amazon RDS, you can create a storage volume DB snapshot of your PostgreSQL DB instance.
When you create a DB snapshot, the snapshot is based on the engine version used by your
Amazon RDS instance. You can
upgrade the engine version for your DB snapshots.

After restoring a DB snapshot upgraded to a new engine version,
make sure to test that the upgrade was successful.
For more information about a major version upgrade, see
[Upgrades of the RDS for PostgreSQL DB engine](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.html).
To learn how to restore a DB snapshot, see [Restoring to a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RestoreFromSnapshot.html).

You can upgrade manual DB snapshots that are either encrypted or not encrypted.

To view the available engine versions for your RDS for PostgreSQL DB snapshot, use the following AWS CLI example.

```nohighlight

aws rds describe-db-engine-versions --engine postgres  --engine-version example-engine-version --query "DBEngineVersions[*].ValidUpgradeTarget[*].{EngineVersion:EngineVersion}" --output text --include-all
```

For more information about available engine versions for RDS for PostgreSQL DB snapshots, see [Choosing a major version for an RDS for PostgreSQL upgrade](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.MajorVersion.html).

###### Note

You can't upgrade automated DB snapshots that are created during the automated backup process.

###### To upgrade a DB snapshot

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Snapshots**.

3. Choose the snapshot that you want to upgrade.

4. For **Actions**, choose **Upgrade**
**snapshot**. The **Upgrade snapshot** page
    appears.

5. Choose the **New engine version** to upgrade to.

6. Choose **Save changes** to upgrade the snapshot.

During the upgrade process, all snapshot actions are disabled for this DB
    snapshot. Also, the DB snapshot status changes from
    **available** to **upgrading**, and
    then changes to **active** upon completion. If the DB
    snapshot can't be upgraded because of snapshot corruption issues, the status
    changes to **unavailable**. You can't recover the snapshot
    from this state.

###### Note

If the DB snapshot upgrade fails, the snapshot is rolled back to the
original state with the original version.

To upgrade a DB snapshot to a new database engine version, use the AWS CLI [modify-db-snapshot](https://docs.aws.amazon.com/cli/latest/reference/rds/modify-db-snapshot.html) command.

###### Parameters

- `--db-snapshot-identifier` – The identifier of the DB
snapshot to upgrade. The identifier must be a unique Amazon Resource Name (ARN).
For more information, see [Amazon Resource Names (ARNs) in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html).

- `--engine-version` – The engine version to upgrade the DB
snapshot to.

###### Example

For Linux, macOS, or Unix:

```nohighlight

aws rds modify-db-snapshot \
    --db-snapshot-identifier my_db_snapshot \
    --engine-version new_version
```

For Windows:

```nohighlight

aws rds modify-db-snapshot ^
    --db-snapshot-identifier my_db_snapshot ^
    --engine-version new_version
```

To upgrade a DB snapshot to a new database engine version, call the Amazon RDS API [ModifyDBSnapshot](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_ModifyDBSnapshot.html) operation.

- `DBSnapshotIdentifier` – The identifier of the DB snapshot
to upgrade. The identifier must be a unique Amazon Resource Name (ARN). For more
information, see [Amazon Resource Names (ARNs) in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html).

- `EngineVersion` – The engine version to upgrade the DB
snapshot to.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Monitoring with events

RDS for PostgreSQL read
replicas
