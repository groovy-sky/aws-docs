# Managing automated backups

This section shows how to manage automated backups for DB instances and Multi-AZ DB clusters.

###### Topics

- [Backup window](#USER_WorkingWithAutomatedBackups.BackupWindow)

- [Backup retention period](user-workingwithautomatedbackups-backupretention.md)

- [Enabling automated backups](user-workingwithautomatedbackups-enabling.md)

- [Retaining automated backups](user-workingwithautomatedbackups-retaining.md)

- [Deleting retained automated backups](user-workingwithautomatedbackups-deleting.md)

- [Automated backups with unsupported MySQL storage engines](overview-backupdevicerestrictions.md)

- [Automated backups with unsupported MariaDB storage engines](overview-backupdevicerestrictionsmariadb.md)

- [Replicating automated backups to another AWS Region](user-replicatebackups.md)

## Backup window

Automated backups occur daily during the preferred backup window. If the backup
requires more time than allotted to the backup window, the backup continues after the
window ends until it finishes. The backup window can't overlap with the weekly
maintenance window for the DB instance or Multi-AZ DB cluster.

During the automatic backup window, storage I/O might be suspended briefly while the
backup process initializes (typically under a few seconds). You might experience
elevated latencies for a few minutes during backups for Multi-AZ deployments. For
MariaDB, MySQL, Oracle, and PostgreSQL, I/O activity isn't suspended on your primary
during backup for Multi-AZ deployments because the backup is taken from the standby. For
SQL Server, I/O activity is suspended briefly during backup for both Single-AZ and
Multi-AZ deployments because the backup is taken from the primary.

For Db2, I/O activity is also suspended briefly during backup even though the backup is
taken from the standby.

Automated backups might occasionally be skipped if the DB instance or cluster has a
heavy workload at the time a backup is supposed to start. If a backup is skipped, you
can still do a point-in-time-recovery (PITR), and a backup is still attempted during the
next backup window. For more information on PITR, see [Restoring a DB instance to a specified time for Amazon RDS](user-pit.md).

If you don't specify a preferred backup window when you create the DB instance or
Multi-AZ DB cluster, Amazon RDS assigns a default 30-minute backup window. This window is
selected at random from an 8-hour block of time for each AWS Region. The following
table lists the time blocks for each AWS Region from which the default backup windows
are assigned.

Region NameRegionTime BlockUS East (N. Virginia)us-east-103:00–11:00 UTCUS East (Ohio)us-east-203:00–11:00 UTCUS West (N. California)us-west-106:00–14:00 UTCUS West (Oregon)us-west-206:00–14:00 UTCAfrica (Cape Town)af-south-103:00–11:00 UTCAsia Pacific (Hong Kong)ap-east-106:00–14:00 UTCAsia Pacific (Hyderabad)ap-south-206:30–14:30 UTCAsia Pacific (Jakarta)ap-southeast-308:00–16:00 UTCAsia Pacific (Malaysia)ap-southeast-509:00–17:00 UTCAsia Pacific (Melbourne)ap-southeast-411:00–19:00 UTCAsia Pacific (Mumbai)ap-south-116:30–00:30 UTCAsia Pacific (New Zealand)ap-southeast-613:00–21:00 UTCAsia Pacific (Osaka)ap-northeast-300:00–08:00 UTCAsia Pacific (Seoul)ap-northeast-213:00–21:00 UTCAsia Pacific (Singapore)ap-southeast-114:00–22:00 UTCAsia Pacific (Sydney)ap-southeast-212:00–20:00 UTCAsia Pacific (Taipei)ap-east-29:00–17:00 UTCAsia Pacific (Thailand)ap-southeast-78:00–16:00 UTCAsia Pacific (Tokyo)ap-northeast-113:00–21:00 UTCCanada (Central)ca-central-103:00–11:00 UTCCanada West (Calgary)ca-west-118:00–02:00 UTCChina (Beijing)cn-north-106:00–14:00 UTCChina (Ningxia)cn-northwest-106:00–14:00 UTCEurope (Frankfurt)eu-central-120:00–04:00 UTCEurope (Ireland)eu-west-122:00–06:00 UTCEurope (London)eu-west-222:00–06:00 UTCEurope (Milan)eu-south-102:00–10:00 UTCEurope (Paris)eu-west-307:29–14:29 UTCEurope (Spain)eu-south-202:00–10:00 UTCEurope (Stockholm)eu-north-123:00–07:00 UTCEurope (Zurich)eu-central-202:00–10:00 UTCIsrael (Tel Aviv)il-central-103:00–11:00 UTCMexico (Central)mx-central-119:00–03:00 UTCMiddle East (Bahrain)me-south-106:00–14:00 UTCMiddle East (UAE)me-central-105:00–13:00 UTCSouth America (São Paulo)sa-east-123:00–07:00 UTCAWS GovCloud (US-East)us-gov-east-117:00–01:00 UTCAWS GovCloud (US-West)us-gov-west-106:00–14:00 UTC

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Introduction to backups

Backup retention period

All content copied from https://docs.aws.amazon.com/.
