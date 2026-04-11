# Deleting a DB snapshot for Amazon RDS

You can delete DB snapshots managed by Amazon RDS when you no longer need them.

###### Note

To delete backups managed by AWS Backup, use the AWS Backup console. For information about
AWS Backup, see the [_AWS Backup Developer Guide_](../../../aws-backup/latest/devguide.md).

## Deleting a DB snapshot

You can delete a manual, shared, or public DB snapshot using the AWS Management Console, the AWS CLI, or the RDS API.

To delete a shared or public snapshot, you must sign in to the AWS account that owns the snapshot.

If you have automated DB snapshots that you want to delete without deleting the DB instance,
change the backup retention period for the DB instance to 0. The automated snapshots are deleted when
the change is applied. You can apply the change immediately if you don't want to wait until the
next maintenance period. After the change is complete, you can then re-enable automatic backups
by setting the backup retention period to a number greater than 0. For information about modifying a
DB instance, see [Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md).

Retained automated backups and manual snapshots incur billing charges until they're deleted. For more information, see
[Retention costs](user-workingwithautomatedbackups-retaining.md#USER_WorkingWithAutomatedBackups.RetentionCosts).

If you deleted a DB instance, you can delete its automated DB snapshots by removing the automated
backups for the DB instance. For information about automated backups, see
[Introduction to backups](user-workingwithautomatedbackups.md).

###### To delete a DB snapshot

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Snapshots**.

The **Manual snapshots** list appears.

3. Choose the DB snapshot that you want to delete.

4. For **Actions**, choose **Delete snapshot**.

![Delete snapshot](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/delete-snapshot.png)

5. Choose **Delete** on the confirmation page.

![Confirm snapshot deletion](https://docs.aws.amazon.com/images/AmazonRDS/latest/UserGuide/images/delete-snapshot-confirm.png)

You can delete a DB snapshot by using the AWS CLI command
[delete-db-snapshot](../../../cli/latest/reference/rds/delete-db-snapshot.md).

The following options are used to delete a DB snapshot.

- `--db-snapshot-identifier` –
The identifier for the DB snapshot.

###### Example

The following code deletes the `mydbsnapshot`
DB snapshot.

For Linux, macOS, or Unix:

```nohighlight

aws rds delete-db-snapshot \
    --db-snapshot-identifier mydbsnapshot

```

For Windows:

```nohighlight

aws rds delete-db-snapshot ^
    --db-snapshot-identifier mydbsnapshot

```

You can delete a DB snapshot by using the Amazon RDS API operation [DeleteDBSnapshot](../../../../reference/amazonrds/latest/apireference/api-deletedbsnapshot.md).

The following parameters are used to delete a DB snapshot.

- `DBSnapshotIdentifier` –
The identifier for the DB snapshot.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Creating a Multi-AZ DB cluster snapshot

Restoring to a DB instance

All content copied from https://docs.aws.amazon.com/.
