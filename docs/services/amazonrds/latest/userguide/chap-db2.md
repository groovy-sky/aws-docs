# Amazon RDS for Db2

Amazon RDS supports DB instances that run the following editions of IBM Db2:

- Db2 Standard Edition

- Db2 Advanced Edition

Amazon RDS supports DB instances that run the following versions of Db2:

- Db2 11.5

For more information about minor version support, see [Db2 on Amazon RDS versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Db2.Concepts.VersionMgmt.html).

Before creating a DB instance, complete the steps in the [Setting up your Amazon RDS environment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html) section of this user guide. When you create a DB
instance using your master user, the user gets `DBADM` authority, with some
limitations. Use this user for administrative tasks such as creating additional database
accounts. You can't use `SYSADM`, `SYSCTRL`, `SYSMAINT`
instance-level authority, or `SECADM` database-level authority.

You can create the following:

- DB instances

- DB snapshots

- Point-in-time restores

- Automated storage backups

- Manual storage backups

You can use DB instances running Db2 inside a virtual private cloud (VPC). You can also
add features to your Amazon RDS for Db2 DB instance by enabling various options. Amazon RDS supports
Multi-AZ deployments for RDS for Db2 as a high availability, failover solution.

###### Important

To deliver a managed service experience, Amazon RDS doesn't provide shell access to DB
instances. It also restricts access to certain system procedures and tables that need
elevated privileges. You can access your database using standard SQL clients such as IBM Db2 CLP. However, you
can't access the host directly by using Telnet or Secure Shell (SSH).

###### Topics

- [Overview of Db2 on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-overview.html)

- [Prerequisites for creating an Amazon RDS for Db2 DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-db-instance-prereqs.html)

- [Multiple databases on an Amazon RDS for Db2 DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-multiple-databases.html)

- [Connecting to your Db2 DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToDb2DBInstance.html)

- [Securing Amazon RDS for Db2 DB instance connections](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Db2.Concepts.RestrictedDBAPrivileges.html)

- [Administering your Amazon RDS for Db2 DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-administering-db-instance.html)

- [Integrating an Amazon RDS for Db2 DB instance with Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-s3-integration.html)

- [Migrating data to Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-migrating-data-to-rds.html)

- [Amazon RDS for Db2 federation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-federation.html)

- [Working with replicas for Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-replication.html)

- [Options for Amazon RDS for Db2 DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Db2.Options.html)

- [External stored procedures for Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-external-stored-procedures.html)

- [Known issues and limitations for Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-known-issues-limitations.html)

- [Amazon RDS for Db2 stored procedure reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-stored-procedures.html)

- [Amazon RDS for Db2 user-defined function reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-user-defined-functions.html)

- [Troubleshooting for Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-troubleshooting.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Troubleshooting zero-ETL integrations

Db2 overview
