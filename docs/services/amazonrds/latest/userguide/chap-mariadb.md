# Amazon RDS for MariaDB

Amazon RDS supports several versions of MariaDB for DB instances. For complete information about the supported versions, see [MariaDB on Amazon RDS versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MariaDB.Concepts.VersionMgmt.html).

To create a MariaDB DB instance, use the Amazon RDS management tools or interfaces. You can
then use the Amazon RDS tools to perform management actions for the DB instance. These include
actions such as the following:

- Reconfiguring or resizing the DB instance

- Authorizing connections to the DB instance

- Creating and restoring from backups or snapshots

- Creating Multi-AZ secondaries

- Creating read replicas

- Monitoring the performance of your DB instance

To store and access the data in your DB instance, use standard MariaDB utilities and
applications.

MariaDB is available in all of the AWS Regions.
For more information about AWS Regions, see
[Regions, Availability Zones, and Local Zones](concepts-regionsandavailabilityzones.md).

You can use Amazon RDS for MariaDB databases to build HIPAA-compliant applications.
You can store healthcare-related information,
including protected health information (PHI),
under a Business Associate Agreement (BAA) with AWS.
For more information, see
[HIPAA compliance](https://aws.amazon.com/compliance/hipaa-compliance).
AWS Services in Scope have been fully assessed by a third-party auditor
and result in a certification, attestation of compliance, or Authority to Operate (ATO).
For more information, see
[AWS services in scope by compliance program](https://aws.amazon.com/compliance/services-in-scope).

Before creating a DB instance, complete the steps in [Setting up your Amazon RDS environment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html). When you create a DB instance, the RDS master user
gets DBA privileges, with some limitations. Use this account for
administrative tasks such as creating additional database accounts.

You can create the following:

- DB instances

- DB snapshots

- Point-in-time restores

- Automated backups

- Manual backups

You can use DB instances running MariaDB inside a virtual private cloud (VPC) based on Amazon VPC. You can
also add features to your MariaDB DB instance by enabling various options. Amazon RDS supports
Multi-AZ deployments for MariaDB as a high-availability, failover solution.

###### Important

To deliver a managed service experience, Amazon RDS doesn't provide shell access to DB instances. It also
restricts access to certain system procedures and tables that need advanced privileges. You can access your
database using standard SQL clients such as the mysql client. However, you can't access the host directly by using
Telnet or Secure Shell (SSH).

###### Topics

- [MariaDB feature support on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MariaDB.Concepts.FeatureSupport.html)

- [MariaDB on Amazon RDS versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MariaDB.Concepts.VersionMgmt.html)

- [Connecting to your MariaDB DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToMariaDBInstance.html)

- [Securing MariaDB DB instance connections](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/securing-mariadb-connections.html)

- [Improving query performance for RDS for MariaDB with Amazon RDS Optimized Reads](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-optimized-reads-mariadb.html)

- [Improving write performance with Amazon RDS Optimized Writes for MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-optimized-writes-mariadb.html)

- [Upgrades of the MariaDB DB engine](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.MariaDB.html)

- [Upgrading a MariaDB DB snapshot engine version](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mariadb-upgrade-snapshot.html)

- [Importing data into an Amazon RDS for MariaDB DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MariaDB.Procedural.Importing.html)

- [Working with MariaDB replication in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_MariaDB.Replication.html)

- [Options for MariaDB database engine](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Options.html)

- [Parameters for MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Parameters.html)

- [Migrating data from a MySQL DB snapshot to a MariaDB DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Migrate_MariaDB.html)

- [MariaDB on Amazon RDS SQL reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.SQLRef.html)

- [Local time zone for MariaDB DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MariaDB.Concepts.LocalTimeZone.html)

- [Known issues and limitations for RDS for MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.Limitations.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Troubleshooting

MariaDB feature support
