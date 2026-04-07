# Backing up and restoring an Amazon RDS Custom for SQL Server DB instance

Like Amazon RDS, RDS Custom creates and saves automated backups of your RDS Custom for SQL Server DB instance when backup retention is enabled.
You can also back up your DB instance manually. The automated backups are comprised of snapshot backups and transaction log backups.
Snapshot backups are taken for the entire storage volume of DB instance during your specified backup window.
Transaction log backups are taken for the PITR-eligible databases on a regular interval period.
RDS Custom saves the automated backups of your DB instance according to your specified backup retention period.
You can use automated backups to recover your DB instance to a point in time within the backup retention period.

You can also take snapshot backups manually. You can create a new DB instance from these snapshot backups at any time.
For more information about manually creating a DB snapshot, see [Creating an RDS Custom for SQL Server snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-backup-sqlserver.creating.html).

Although snapshot backups serve operationally as full backups, you are billed only for incremental storage use.
The first snapshot of an RDS Custom DB instance contains the data for the full DB instance. Subsequent snapshots of the same
database are incremental, which means that only the data that has changed after your most recent snapshot is saved.

###### Topics

- [Creating an RDS Custom for SQL Server snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-backup-sqlserver.creating.html)

- [Restoring from an RDS Custom for SQL Server DB snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-backup-sqlserver.restoring.html)

- [Restoring an RDS Custom for SQL Server instance to a point in time](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-backup.pitr-sqs.html)

- [Deleting an RDS Custom for SQL Server snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-backup-sqlserver.deleting.html)

- [Deleting RDS Custom for SQL Server automated backups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-backup-sqlserver.deleting-backups.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Failover process

Creating an RDS Custom for SQL Server snapshot
