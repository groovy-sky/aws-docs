# Amazon RDS on AWS Outposts support for Amazon RDS features

The following table describes the Amazon RDS features supported by Amazon RDS on AWS Outposts.

FeatureSupportedNotesMore information

DB instance provisioning

Yes

You can only create DB instances for RDS for SQL Server, RDS for MySQL, RDS for PostgreSQL, or RDS for Oracle DB engines. The
following versions are supported:

- Microsoft SQL Server:

- 16.00.4085.2.v1 and higher 2022 versions

- 15.00.4043.16.v1 and higher 2019 versions

- 14.00.3294.2.v1 and higher 2017 versions

- 13.00.5820.21.v1 and higher 2016 versions

- MySQL 8.0 and 8.4 versions

- All PostgreSQL 16 & 15 & 14 & 13 versions, and PostgreSQL version 12.5 and higher PostgreSQL 12 versions

- All Oracle versions

[Creating DB instances for Amazon RDS on AWS Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.creating.html)

Connect to a Microsoft SQL Server DB instance with Microsoft SQL Server Management Studio

Yes

Some TLS versions and encryption ciphers might not be secure. To turn them off, follow the instructions in
[Configuring SQL Server security protocols and ciphers](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Ciphers.html).

[Connecting to your Microsoft SQL Server DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToMicrosoftSQLServerInstance.html)

Modifying the master user password

Yes

None

[Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md)

Renaming a DB instance

Yes

None

[Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md)

Rebooting a DB instance

Yes

None

[Rebooting a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RebootInstance.html)

Stopping a DB instance

Yes

None

[Stopping an Amazon RDS DB instance temporarily](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StopInstance.html)

Starting a DB instance

Yes

None

[Starting an Amazon RDS DB instance that was previously stopped](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StartInstance.html)

Multi-AZ deployments

Yes

Multi-AZ deployments are supported on MySQL, PostgreSQL, and Oracle DB instances.

Multi-AZ deployments do not support Direct VPC Routing (DVR).

[Creating DB instances for Amazon RDS on AWS Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.creating.html)

[Configuring and managing a Multi-AZ deployment for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html)

DB parameter groups

Yes

None

[Parameter groups for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html)

Read replicas

Yes

Read replicas are supported for MySQL, PostgreSQL, and Oracle DB instances.

Read replicas do not support Direct VPC Routing (DVR).

[Creating read replicas for Amazon RDS on AWS Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.rr.html)

Encryption at rest

Yes

RDS on Outposts doesn't support unencrypted DB instances.

[Encrypting Amazon RDS resources](overview-encryption.md)

AWS Identity and Access Management (IAM) database authentication

No

None

[IAM database authentication for MariaDB, MySQL, and PostgreSQL](usingwithrds-iamdbauth.md)

Associating an IAM role with a DB instance

No

None

[add-role-to-db-instance](https://docs.aws.amazon.com/cli/latest/reference/rds/add-role-to-db-instance.html) AWS CLI command

[AddRoleToDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_AddRoleToDBInstance.html) RDS API operation

Kerberos authentication

No

None

[Kerberos authentication](database-authentication.md#kerberos-authentication)

Tagging Amazon RDS resources

Yes

None

[Tagging Amazon RDS resources](user-tagging.md)

Option groups

Yes

The following RDS for Oracle options are not supported:

- [Amazon EFS integration](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-efs-integration.html)

- [Amazon S3 integration](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html)

[Working with option groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html)

Modifying the maintenance window

Yes

None

[Maintaining a DB instance](user-upgradedbinstance-maintenance.md)

Automatic minor version upgrade

Yes

None

[Automatically upgrading the minor engine version](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Upgrading.html#USER_UpgradeDBInstance.Upgrading.AutoMinorVersionUpgrades)

Modifying the backup window

Yes

None

[Introduction to backups](user-workingwithautomatedbackups.md)

[Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md)

Changing the DB instance class

Yes

None

[Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md)

Changing the allocated storage

Yes

None

[Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md)

Storage autoscaling

Yes

None

[Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.Autoscaling.html)

Manual and automatic DB instance snapshots

Yes

You can store automated backups and manual snapshots in your
AWS Region. Or you can store them locally on your Outpost.

Local backups are supported on MySQL, PostgreSQL, and Oracle DB instances.

To store backups on your Outpost, make sure that you have Amazon S3 on Outposts configured.

Local backups are not supported for Multi-AZ instance deployments.

[Creating DB instances for Amazon RDS on AWS Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.creating.html)

[Amazon S3 on Outposts](https://aws.amazon.com/s3/outposts)

[Creating a DB snapshot for a Single-AZ DB instance for Amazon RDS](user-createsnapshot.md)

Restoring from a DB snapshot

Yes

You can store automated backups and manual snapshots for the restored DB instance in the parent AWS Region
or locally on your Outpost.

[Considerations for restoring DB instances on Amazon RDS on AWS Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.restoring.html)

[Restoring to a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RestoreFromSnapshot.html)

Restoring a DB instance from Amazon S3

No

None

[Restoring a backup into an Amazon RDS for MySQL DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.html)

Exporting snapshot data to Amazon S3

No

None

[Exporting DB snapshot data to Amazon S3 for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ExportSnapshot.html)

Point-in-time recovery

Yes

You can store automated backups and manual snapshots for the restored DB instance in the parent AWS Region
or locally on your Outpost, with one exception.

[Considerations for restoring DB instances on Amazon RDS on AWS Outposts](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.restoring.html)

[Restoring a DB instance to a specified time for Amazon RDS](user-pit.md)

Enhanced monitoring

No

None

[Monitoring OS metrics with Enhanced Monitoring](user-monitoring-os.md)

Amazon CloudWatch monitoring

Yes

You can view the same set of metrics that are available for your databases in the AWS Region.

[Monitoring Amazon RDS metrics with Amazon CloudWatch](monitoring-cloudwatch.md)

Publishing database engine logs to CloudWatch Logs

Yes

None

[Publishing database logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.Procedural.UploadtoCloudWatch.html)

Event notification

Yes

None

[Working with Amazon RDS event notification](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html)

Amazon RDS Performance Insights

No

None

[Monitoring DB load with Performance Insights on Amazon RDS](user-perfinsights.md)

Viewing or downloading database logs

No

RDS on Outposts doesn't support viewing database logs using the console or describing database logs using
the AWS CLI or RDS API.

RDS on Outposts doesn't support downloading database logs using the console or downloading database logs
using the AWS CLI or RDS API.

[Monitoring Amazon RDS log files](user-logaccess.md)

Amazon RDS Proxy

No

None

[Amazon RDS Proxy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy.html)

Stored procedures for Amazon RDS for MySQL

Yes

None

[RDS for MySQL stored procedure reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.SQLRef.html)

Replication with external databases for RDS for MySQL

No

None

[Configuring binary log file position replication with an external source instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.External.Repl.html)

Native backup and restore for Amazon RDS for Microsoft SQL Server

Yes

None

[Importing and exporting SQL Server databases using native backup and restore](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Procedural.Importing.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon RDS on AWS Outposts

Supported DB instance classes
