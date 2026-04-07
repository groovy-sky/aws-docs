# Deleting a DB cluster snapshot

You can delete DB cluster snapshots managed by Amazon RDS when you no longer need them.

###### Note

To delete backups managed by AWS Backup, use the AWS Backup console. For information about
AWS Backup, see the [_AWS Backup Developer Guide_](../../../aws-backup/latest/devguide.md).

## Deleting a DB cluster snapshot

You can delete a DB cluster snapshot using the console, the AWS CLI, or the RDS API.

To delete a shared or public snapshot, you must sign in to the AWS account that owns the
snapshot.

###### To delete a DB cluster snapshot

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Snapshots**.

3. Choose the DB cluster snapshot that you want to delete.

4. For **Actions**, choose **Delete snapshot**.

5. Choose **Delete** on the confirmation page.

You can delete a DB cluster snapshot by using the AWS CLI command
[delete-db-cluster-snapshot](https://docs.aws.amazon.com/cli/latest/reference/rds/delete-db-cluster-snapshot.html).

The following options are used to delete a DB cluster snapshot.

- `--db-cluster-snapshot-identifier` –
The identifier for the DB cluster snapshot.

###### Example

The following code deletes the `mydbclustersnapshot`
DB cluster snapshot.

For Linux, macOS, or Unix:

```nohighlight

aws rds delete-db-cluster-snapshot \
    --db-cluster-snapshot-identifier mydbclustersnapshot

```

For Windows:

```nohighlight

aws rds delete-db-cluster-snapshot ^
    --db-cluster-snapshot-identifier mydbclustersnapshot

```

You can delete a DB cluster snapshot by using the Amazon RDS API operation [DeleteDBClusterSnapshot](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DeleteDBClusterSnapshot.html).

The following parameters are used to delete a DB cluster snapshot.

- `DBClusterSnapshotIdentifier` –
The identifier for the DB cluster snapshot.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Point-in-time recovery using AWS Backup

Tutorial: Restore a DB cluster from a snapshot
