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

[Creating DB instances for Amazon RDS on AWS Outposts](rds-on-outposts-creating.md)

Connect to a Microsoft SQL Server DB instance with Microsoft SQL Server Management Studio

Yes

Some TLS versions and encryption ciphers might not be secure. To turn them off, follow the instructions in
[Configuring SQL Server security protocols and ciphers](sqlserver-ciphers.md).

[Connecting to your Microsoft SQL Server DB instance](user-connecttomicrosoftsqlserverinstance.md)

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

[Rebooting a DB instance](user-rebootinstance.md)

Stopping a DB instance

Yes

None

[Stopping an Amazon RDS DB instance temporarily](user-stopinstance.md)

Starting a DB instance

Yes

None

[Starting an Amazon RDS DB instance that was previously stopped](user-startinstance.md)

Multi-AZ deployments

Yes

Multi-AZ deployments are supported on MySQL, PostgreSQL, and Oracle DB instances.

Multi-AZ deployments do not support Direct VPC Routing (DVR).

[Creating DB instances for Amazon RDS on AWS Outposts](rds-on-outposts-creating.md)

[Configuring and managing a Multi-AZ deployment for Amazon RDS](concepts-multiaz.md)

DB parameter groups

Yes

None

[Parameter groups for Amazon RDS](user-workingwithparamgroups.md)

Read replicas

Yes

Read replicas are supported for MySQL, PostgreSQL, and Oracle DB instances.

Read replicas do not support Direct VPC Routing (DVR).

[Creating read replicas for Amazon RDS on AWS Outposts](rds-on-outposts-rr.md)

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

[add-role-to-db-instance](../../../cli/latest/reference/rds/add-role-to-db-instance.md) AWS CLI command

[AddRoleToDBInstance](../../../../reference/amazonrds/latest/apireference/api-addroletodbinstance.md) RDS API operation

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

- [Amazon EFS integration](oracle-efs-integration.md)

- [Amazon S3 integration](oracle-s3-integration.md)

[Working with option groups](user-workingwithoptiongroups.md)

Modifying the maintenance window

Yes

None

[Maintaining a DB instance](user-upgradedbinstance-maintenance.md)

Automatic minor version upgrade

Yes

None

[Automatically upgrading the minor engine version](user-upgradedbinstance-upgrading.md#USER_UpgradeDBInstance.Upgrading.AutoMinorVersionUpgrades)

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

[Managing capacity automatically with Amazon RDS storage autoscaling](user-piops-autoscaling.md)

Manual and automatic DB instance snapshots

Yes

You can store automated backups and manual snapshots in your
AWS Region. Or you can store them locally on your Outpost.

Local backups are supported on MySQL, PostgreSQL, and Oracle DB instances.

To store backups on your Outpost, make sure that you have Amazon S3 on Outposts configured.

Local backups are not supported for Multi-AZ instance deployments.

[Creating DB instances for Amazon RDS on AWS Outposts](rds-on-outposts-creating.md)

[Amazon S3 on Outposts](https://aws.amazon.com/s3/outposts)

[Creating a DB snapshot for a Single-AZ DB instance for Amazon RDS](user-createsnapshot.md)

Restoring from a DB snapshot

Yes

You can store automated backups and manual snapshots for the restored DB instance in the parent AWS Region
or locally on your Outpost.

[Considerations for restoring DB instances on Amazon RDS on AWS Outposts](rds-on-outposts-restoring.md)

[Restoring to a DB instance](user-restorefromsnapshot.md)

Restoring a DB instance from Amazon S3

No

None

[Restoring a backup into an Amazon RDS for MySQL DB instance](mysql-procedural-importing.md)

Exporting snapshot data to Amazon S3

No

None

[Exporting DB snapshot data to Amazon S3 for Amazon RDS](user-exportsnapshot.md)

Point-in-time recovery

Yes

You can store automated backups and manual snapshots for the restored DB instance in the parent AWS Region
or locally on your Outpost, with one exception.

[Considerations for restoring DB instances on Amazon RDS on AWS Outposts](rds-on-outposts-restoring.md)

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

[Publishing database logs to Amazon CloudWatch Logs](user-logaccess-procedural-uploadtocloudwatch.md)

Event notification

Yes

None

[Working with Amazon RDS event notification](user-events.md)

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

[Amazon RDS Proxy](rds-proxy.md)

Stored procedures for Amazon RDS for MySQL

Yes

None

[RDS for MySQL stored procedure reference](appendix-mysql-sqlref.md)

Replication with external databases for RDS for MySQL

No

None

[Configuring binary log file position replication with an external source instance](mysql-procedural-importing-external-repl.md)

Native backup and restore for Amazon RDS for Microsoft SQL Server

Yes

None

[Importing and exporting SQL Server databases using native backup and restore](sqlserver-procedural-importing.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon RDS on AWS Outposts

Supported DB instance classes

All content copied from https://docs.aws.amazon.com/.
