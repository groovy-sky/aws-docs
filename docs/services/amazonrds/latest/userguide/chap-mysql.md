# Amazon RDS for MySQL

Amazon RDS supports several versions of MySQL for DB instances. For complete information about the supported versions, see [MySQL on Amazon RDS versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Concepts.VersionMgmt.html).

To create an Amazon RDS for MySQL DB instance, use the Amazon RDS management tools or interfaces. You
can then do the following:

- Resize your DB instance

- Authorize connections to your DB instance

- Create and restore from backups or snapshots

- Create Multi-AZ secondaries

- Create read replicas

- Monitor the performance of your DB instance

To store and access the data in your DB instance, you use standard MySQL utilities and applications.

Amazon RDS for MySQL is compliant with many industry standards. For example, you can use
RDS for MySQL databases to build HIPAA-compliant applications. You can use RDS for MySQL databases
to store healthcare related information, including protected health information (PHI) under
a Business Associate Agreement (BAA) with AWS. Amazon RDS for MySQL also meets Federal Risk and
Authorization Management Program (FedRAMP) security requirements. In addition,
Amazon RDS for MySQL has received a FedRAMP Joint Authorization Board (JAB) Provisional Authority
to Operate (P-ATO) at the FedRAMP HIGH Baseline within the AWS GovCloud (US) Regions. For more
information on supported compliance standards, see [AWS cloud compliance](https://aws.amazon.com/compliance).

For information about the features in each version of MySQL, see [The main features of\
MySQL](https://dev.mysql.com/doc/refman/8.0/en/features.html) in the MySQL documentation.

Before creating a DB instance, complete the steps in [Setting up your Amazon RDS environment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html). When you create a DB instance, the RDS master user
gets DBA privileges, with some limitations. Use this account for
administrative tasks such as creating additional database accounts.

You can create the following:

- DB instances

- DB snapshots

- Point-in-time restores

- Automated backups

- Manual backups

You can use DB instances running MySQL inside a virtual private cloud (VPC) based on
Amazon VPC. You can also add features to your MySQL DB instance by turning on various options.
Amazon RDS supports Multi-AZ deployments for MySQL as a high-availability, failover
solution.

###### Important

To deliver a managed service experience, Amazon RDS doesn't provide shell access to DB instances. It also
restricts access to certain system procedures and tables that need advanced privileges. You can access your
database using standard SQL clients such as the mysql client. However, you can't access the host directly by using
Telnet or Secure Shell (SSH).

###### Topics

- [MySQL feature support on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Concepts.FeatureSupport.html)

- [MySQL on Amazon RDS versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Concepts.VersionMgmt.html)

- [Connecting to your MySQL DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToInstance.html)

- [Securing MySQL DB instance connections](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/securing-mysql-connections.html)

- [Improving query performance for RDS for MySQL with Amazon RDS Optimized Reads](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-optimized-reads.html)

- [Improving write performance with RDS Optimized Writes for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-optimized-writes.html)

- [Upgrades of the RDS for MySQL DB engine](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.MySQL.html)

- [Upgrading a MySQL DB snapshot engine version](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-upgrade-snapshot.html)

- [Importing data into an Amazon RDS for MySQL DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Importing.Other.html)

- [Working with MySQL replication in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MySQL.Replication.html)

- [Configuring active-active clusters for RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-active-active-clusters.html)

- [Exporting data from a MySQL DB instance by using replication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Procedural.Exporting.NonRDSRepl.html)

- [Options for MySQL DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Options.html)

- [Parameters for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Parameters.html)

- [Common DBA tasks for MySQL DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.CommonDBATasks.html)

- [Local time zone for MySQL DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Concepts.LocalTimeZone.html)

- [Known issues and limitations for Amazon RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.KnownIssuesAndLimitations.html)

- [RDS for MySQL stored procedure reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.SQLRef.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Working with
trace and dump files

MySQL feature support
