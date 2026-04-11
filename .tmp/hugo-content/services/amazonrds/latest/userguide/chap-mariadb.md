# Amazon RDS for MariaDB

Amazon RDS supports several versions of MariaDB for DB instances. For complete information about the supported versions, see [MariaDB on Amazon RDS versions](mariadb-concepts-versionmgmt.md).

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

Before creating a DB instance, complete the steps in [Setting up your Amazon RDS environment](chap-settingup.md). When you create a DB instance, the RDS master user
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

- [MariaDB feature support on Amazon RDS](mariadb-concepts-featuresupport.md)

- [MariaDB on Amazon RDS versions](mariadb-concepts-versionmgmt.md)

- [Connecting to your MariaDB DB instance](user-connecttomariadbinstance.md)

- [Securing MariaDB DB instance connections](securing-mariadb-connections.md)

- [Improving query performance for RDS for MariaDB with Amazon RDS Optimized Reads](rds-optimized-reads-mariadb.md)

- [Improving write performance with Amazon RDS Optimized Writes for MariaDB](rds-optimized-writes-mariadb.md)

- [Upgrades of the MariaDB DB engine](user-upgradedbinstance-mariadb.md)

- [Upgrading a MariaDB DB snapshot engine version](mariadb-upgrade-snapshot.md)

- [Importing data into an Amazon RDS for MariaDB DB instance](mariadb-procedural-importing.md)

- [Working with MariaDB replication in Amazon RDS](user-mariadb-replication.md)

- [Options for MariaDB database engine](appendix-mariadb-options.md)

- [Parameters for MariaDB](appendix-mariadb-parameters.md)

- [Migrating data from a MySQL DB snapshot to a MariaDB DB instance](user-migrate-mariadb.md)

- [MariaDB on Amazon RDS SQL reference](appendix-mariadb-sqlref.md)

- [Local time zone for MariaDB DB instances](mariadb-concepts-localtimezone.md)

- [Known issues and limitations for RDS for MariaDB](chap-mariadb-limitations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting

MariaDB feature support

All content copied from https://docs.aws.amazon.com/.
