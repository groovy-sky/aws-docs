# Common management tasks for Amazon RDS for PostgreSQL

The following are the common management tasks you perform with an Amazon RDS for PostgreSQL
DB instance, with links to relevant documentation for each task.

Task areaRelevant documentation

**Setting up Amazon RDS for first-time**
**use**

Before you can create your DB instance, make sure to complete a
few prerequisites. For example, DB instances are created by default
with a firewall that prevents access to it. So you need to create a
security group with the correct IP addresses and network
configuration to access the DB instance.

[Setting up your Amazon RDS environment](chap-settingup.md)

**Understanding Amazon RDS DB**
**instances**

If you are creating a DB instance for production purposes, you
should understand how instance classes, storage types, and
Provisioned IOPS work in Amazon RDS.

[DB instance classes](concepts-dbinstanceclass.md)

[Amazon RDS storage types](chap-storage.md#Concepts.Storage)

[Provisioned IOPS SSD storage](chap-storage.md#USER_PIOPS)

**Finding available PostgreSQL**
**versions**

Amazon RDS supports several versions of PostgreSQL.

[Available PostgreSQL database versions](postgresql-concepts-general-dbversions.md)

**Setting up high availability and failover**
**support**

A production DB instance should use Multi-AZ deployments. Multi-AZ
deployments provide increased availability, data durability, and
fault tolerance for DB instances.

[Configuring and managing a Multi-AZ deployment for Amazon RDS](concepts-multiaz.md)

**Understanding the Amazon Virtual Private Cloud (VPC)**
**network**

If your AWS account has a default VPC, then your DB instance is
automatically created inside the default VPC. In some cases, your
account might not have a default VPC, and you might want the DB
instance in a VPC. In these cases, create the VPC and subnet groups
before you create the DB instance.

[Working with a DB instance in a VPC](user-vpc-workingwithrdsinstanceinavpc.md)

**Importing data into Amazon RDS**
**PostgreSQL**

You can use several different tools to import data into your
PostgreSQL DB instance on Amazon RDS.

[Importing data into PostgreSQL on Amazon RDS](postgresql-procedural-importing.md)

**Setting up read-only read replicas (primary**
**and standbys)**

RDS for PostgreSQL supports read replicas in both the same AWS Region
and in a different AWS Region from the primary instance.

[Working with DB instance read replicas](user-readrepl.md)

[Working with read replicas for Amazon RDS for PostgreSQL](user-postgresql-replication-readreplicas.md)

[Creating a read replica in a different AWS Region](user-readrepl-xrgn.md)

**Understanding security**
**groups**

By default, DB instances are created with a firewall that prevents
access to them. To provide access through that firewall, you edit
the inbound rules for the VPC security group associated with the VPC
hosting the DB instance.

[Controlling access with security groups](overview-rdssecuritygroups.md)

**Setting up parameter groups and**
**features**

To change the default parameters for your DB instance, create a
custom DB parameter group and change settings to that. If you do
this before creating your DB instance, you can choose your custom DB
parameter group when you create the instance.

[Parameter groups for Amazon RDS](user-workingwithparamgroups.md)

**Connecting to your PostgreSQL DB**
**instance**

After creating a security group and associating it to a DB
instance, you can connect to the DB instance using any standard SQL
client application such as `psql` or
`pgAdmin`.

[Connecting to a DB instance running the PostgreSQL database engine](user-connecttopostgresqlinstance.md)

[Using SSL with a PostgreSQL DB instance](postgresql-concepts-general-ssl.md)

**Backing up and restoring your DB**
**instance**

You can configure your DB instance to take automated backups, or
take manual snapshots, and then restore instances from the backups
or snapshots.

[Backing up, restoring, and exporting data](chap-commontasks-backuprestore.md)

**Monitoring the activity and performance of**
**your DB instance**

You can monitor a PostgreSQL DB instance by using CloudWatch Amazon RDS
metrics, events, and enhanced monitoring.

[Viewing metrics in the Amazon RDS console](user-monitoring.md)

[Viewing Amazon RDS events](user-listevents.md)

**Upgrading the PostgreSQL database**
**version**

You can do both major and minor version upgrades for your
PostgreSQL DB instance.

[Upgrades of the RDS for PostgreSQL DB engine](user-upgradedbinstance-postgresql.md)

[Choosing a major version for an RDS for PostgreSQL upgrade](user-upgradedbinstance-postgresql-majorversion.md)

**Working with log files**

You can access the log files for your PostgreSQL DB instance.

[RDS for PostgreSQL database log files](user-logaccess-concepts-postgresql.md)

**Understanding the best practices for**
**PostgreSQL DB instances**

Find some of the best practices for working with PostgreSQL on
Amazon RDS.

[Best practices for working with PostgreSQL](chap-bestpractices.md#CHAP_BestPractices.PostgreSQL)

Following is a list of other sections in this guide that can help you understand and
use important features of RDS for PostgreSQL:

- [Understanding PostgreSQL roles and permissions](appendix-postgresql-commondbatasks-roles.md)

- [Controlling user access to the PostgreSQL database](appendix-postgresql-commondbatasks-access.md)

- [Working with parameters on your RDS for PostgreSQL DB instance](appendix-postgresql-commondbatasks-parameters.md)

- [Understanding\
logging mechanisms supported by RDS for PostgreSQL](appendix-postgresql-commondbatasks.md#Appendix.PostgreSQL.CommonDBATasks.Auditing)

- [Working with PostgreSQL autovacuum on Amazon RDS for PostgreSQL](appendix-postgresql-commondbatasks-autovacuum.md)

- [Using a custom DNS server for outbound network access](appendix-postgresql-commondbatasks-customdns.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PostgreSQL on Amazon RDS

Working with
the Database Preview environment

All content copied from https://docs.aws.amazon.com/.
