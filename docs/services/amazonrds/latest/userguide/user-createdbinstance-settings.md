# Settings for DB instances

In the following table, you can find details about settings that you choose when you
create a DB instance. The table also shows the DB engines for which each setting is
supported.

You can create a DB instance using the console, the [create-db-instance](https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-instance.html) CLI
command, or the [CreateDBInstance](../../../../reference/amazonrds/latest/apireference/api-createdbinstance.md) RDS API operation.

Console settingSetting descriptionCLI option and RDS API parameterSupported DB engines

**Allocated storage**

The amount of storage to allocate for your DB instance (in gibibytes).
In some cases, allocating a higher amount of storage for your DB instance
than the size of your database can improve I/O performance.

For more information, see [Amazon RDS DB instance storage](chap-storage.md).

**CLI option:**

`--allocated-storage`

**API parameter:**

`AllocatedStorage`

All

**Architecture settings**

If you choose **Oracle multitenant**
**architecture**, RDS for Oracle creates a container database
(CDB). If you don't choose this option, RDS for Oracle creates a non-CDB.
A non-CDB uses the traditional Oracle database architecture. A CDB
can contain pluggable databases (PDBs) whereas a non-CDB
cannot.

Oracle Database 21c uses the CDB architecture only. Oracle
Database 19c can use either the CDB or non-CDB architecture.
Releases lower than Oracle Database 19c use the non-CDB architecture
only.

For more information, see [Overview of RDS for Oracle CDBs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Oracle.Concepts.CDBs.html).

**CLI option:**

`--engine oracle-ee-cdb` (Oracle multitenant)

`--engine oracle-se2-cdb` (Oracle multitenant)

`--engine oracle-ee` (traditional)

`--engine oracle-se2` (traditional)

**API parameter:**

`Engine`

Oracle

**Architecture configuration**

These settings are only valid when you choose **Oracle**
**multitenant architecture** for **Architecture**
**settings**. Choose either of the following additional
settings:

- With the **Multi-tenant configuration**,
your RDS for Oracle CDB instance can contain 1–30 tenant databases, depending on the database edition and any required option licenses. In the context of
an Oracle database, a tenant database is a PDB. Application
PDBs and proxy PDBs aren't supported.

Your DB instance is created with 1 initial tenant database.
Choose values for **Tenant database name**,
**Tenant database master username**,
**Tenant database master password**,
and **Tenant database character**
**set**.

The multi-tenant configuration is permanent. Thus, you
can't convert the multi-tenant configuration back to the
single-tenant configuration. The minimum supported release
update (RU) for the multi-tenant configuration is
19.0.0.0.ru-2022-01.rur-2022.r1.

###### Note

The Amazon RDS configuration is called "multi-tenant" rather than "multitenant" because it is a
capability of Amazon RDS, not just the Oracle DB engine. Similarly, the RDS term "tenant"
refers to any tenant in an RDS configuration, not just Oracle PDBs. In the RDS documentation,
the unhyphenated term "Oracle multitenant" refers exclusively to the Oracle database CDB
architecture, which is compatible with both on-premises and RDS deployments.

- With the **Single-tenant configuration**,
your RDS for Oracle CDB contains 1 PDB. This is the default
configuration when you create a CDB. You can't delete the
initial PDB or add more PDBs. You can later convert the
single-tenant configuration of your CDB to the multi-tenant
configuration, but you can't then convert back to the
single-tenant configuration.

Regardless of which configuration you choose, your CDB contains a
single initial PDB. In the multi-tenant configuration, you can
create more PDBs later using RDS APIs.

For more information, see [Overview of RDS for Oracle CDBs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Oracle.Concepts.CDBs.html).

**CLI option:**

`--multi-tenant` (multi-tenant configuration)

`--no-multi-tenant` (single-tenant
configuration)

**API parameter:**

`MultiTenant`

Oracle

**Auto minor version upgrade**

Choose **Enable auto minor version upgrade** to enable your DB instance to
receive preferred minor DB engine version upgrades automatically when they become available. This
is the default behavior. Amazon RDS performs automatic minor version upgrades in the maintenance window.
If you don't choose **Enable auto minor version upgrade**,
your DB instance isn't upgraded automatically when new minor versions become available.

For more information, see [Automatically upgrading the minor engine version](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Upgrading.html#USER_UpgradeDBInstance.Upgrading.AutoMinorVersionUpgrades).

**CLI option:**

`--auto-minor-version-upgrade`

`--no-auto-minor-version-upgrade`

**API parameter:**

`AutoMinorVersionUpgrade`

All**Availability zone**

The Availability Zone for your DB instance. Use the default value of
**No Preference** unless you want to specify an
Availability Zone.

For more information, see [Regions, Availability Zones, and Local Zones](concepts-regionsandavailabilityzones.md).

**CLI option:**

`--availability-zone`

**API parameter:**

`AvailabilityZone`

All

**AWS KMS key**

Only available if **Encryption** is set to
**Enable encryption**. Choose the
AWS KMS key to use for encrypting this DB instance. For more
information, see [Encrypting Amazon RDS resources](overview-encryption.md).

**CLI option:**

`--kms-key-id`

**API parameter:**

`KmsKeyId`

All**AWS License Manager configuration**

Enter a name for an AWS License Manager license configuration. The name must be 100 characters or
less, and only include a-z, A-Z, and 0-9.

For more information, see [Integrating with AWS License Manager](db2-licensing.md#db2-lms-integration).

**CLI option:**

For more information, see [AWS License Manager CLI](db2-licensing.md#db2-lms-integration.cli).

**API parameter:**

For more information, see [AWS License Manager API](db2-licensing.md#db2-lms-integration.api).

Db2**Backup replication**

Choose **Enable replication in another AWS**
**Region** to create backups in an additional Region for
disaster recovery.

Then choose the **Destination Region** for the
additional backups.

Not available when creating a DB instance. For information on enabling
cross-Region backups using the AWS CLI or RDS API, see [Enabling cross-Region automated backups for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AutomatedBackups.Replicating.Enable.html).

Db2

Oracle

PostgreSQL

SQL Server

**Backup retention period**

The number of days that you want automatic backups of your DB
instance to be retained. For any nontrivial DB instance, set this value to
`1` or greater.

For more information, see [Introduction to backups](user-workingwithautomatedbackups.md).

**CLI option:**

`--backup-retention-period`

**API parameter:**

`BackupRetentionPeriod`

All**Backup target**

Choose **AWS Cloud** to store automated backups
and manual snapshots in the parent AWS Region. Choose
**Outposts (on-premises)** to store them
locally on your Outpost.

This option setting applies only to RDS on Outposts. For more
information, see [Creating DB instances for Amazon RDS on AWS Outposts](rds-on-outposts-creating.md).

**CLI option:**

`--backup-target`

**API parameter:**

`BackupTarget`

MySQL, PostgreSQL, SQL Server**Backup window**

The time period during which Amazon RDS automatically takes a backup of
your DB instance. Unless you have a specific time that you want to have
your database backed up, use the default of **No**
**Preference**.

For more information, see [Introduction to backups](user-workingwithautomatedbackups.md).

**CLI option:**

`--preferred-backup-window`

**API parameter:**

`PreferredBackupWindow`

All

**Certificate authority**

The certificate authority (CA) for the server certificate used by
the DB instance.

For more information, see [Using SSL/TLS to encrypt a connection to a DB instance or cluster](usingwithrds-ssl.md).

**CLI option:**

`--ca-certificate-identifier`

**RDS API parameter:**

`CACertificateIdentifier`

All

**Character set**

The character set for your DB instance. The default value of
**AL32UTF8** for the DB character set is for
the Unicode 5.0 UTF-8 Universal character set. You can't change
the DB character set after you create the DB instance.

In a single-tenant configuration, a non-default DB character set
affects only the PDB, not the CDB. For more information, see [Single-tenant configuration of the CDB architecture](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Oracle.Concepts.CDBs.html#Oracle.Concepts.single-tenant).

The DB character set is different from the national character set,
which is called the NCHAR character set. Unlike the DB character
set, the NCHAR character set specifies the encoding for NCHAR data
types (NCHAR, NVARCHAR2, and NCLOB) columns without affecting
database metadata.

For more information, see [RDS for Oracle character sets](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html).

**CLI option:**

`--character-set-name`

**API parameter:**

`CharacterSetName`

Oracle**Collation**

A server-level collation for your DB instance.

For more information, see [Server-level collation for Microsoft SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Collation.html#Appendix.SQLServer.CommonDBATasks.Collation.Server).

**CLI option:**

`--character-set-name`

**API parameter:**

`CharacterSetName`

SQL Server**Copy tags to snapshots**

This option copies any DB instance tags to a DB snapshot when you create
a snapshot.

For more information, see [Tagging Amazon RDS resources](user-tagging.md).

**CLI option:**

`--copy-tags-to-snapshot`

`--no-copy-tags-to-snapshot`

**RDS API parameter:**

`CopyTagsToSnapshot`

All**Database authentication**

The database authentication option that you want to use.

Choose **Password authentication** to
authenticate database users with database passwords only.

Choose **Password and IAM DB authentication** to
authenticate database users with database passwords and user
credentials through users and roles. For more information, see [IAM database authentication for MariaDB, MySQL, and PostgreSQL](usingwithrds-iamdbauth.md). This option is only
supported for MySQL and PostgreSQL.

Choose **Password and Kerberos authentication**
to authenticate database users with database passwords and Kerberos
authentication through an AWS Managed Microsoft AD created with Directory Service. Next,
choose the directory or choose **Create a new**
**Directory**.

For more information, see one of the following:

- [Using Kerberos authentication for Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-kerberos.html)

- [Using Kerberos authentication for Amazon RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-kerberos.html)

- [Configuring Kerberos authentication for Amazon RDS for Oracle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos.html)

- [Using Kerberos authentication with Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/postgresql-kerberos.html)

**_IAM:_**

**CLI option:**

`--enable-iam-database-authentication`

`--no-enable-iam-database-authentication`

**RDS API parameter:**

`EnableIAMDatabaseAuthentication`

**_Kerberos:_**

**CLI option:**

`--domain`

`--domain-iam-role-name`

**RDS API parameter:**

`Domain`

`DomainIAMRoleName`

Varies by authentication type

**Database management type**

Choose **Amazon RDS** if you don't need to customize
your environment.

Choose **Amazon RDS Custom** if you want to customize
the database, OS, and infrastructure. For more information, see
[Amazon RDS Custom](rds-custom.md).

For the CLI and API, you specify the database engine type.

Oracle

SQL Server

**Database port**

The port that you want to access the DB instance through. The default
port is shown.

###### Note

The firewalls at some companies block connections to the
default MariaDB, MySQL, and PostgreSQL ports. If your company
firewall blocks the default port, enter another port for your DB
instance.

**CLI option:**

`--port`

**RDS API parameter:**

`Port`

All**DB engine version**

The version of database engine that you want to use.

**CLI option:**

`--engine-version`

**RDS API parameter:**

`EngineVersion`

All**DB instance class**

The configuration for your DB instance. For example, a
**db.t3.small** DB instance class has 2 GiB memory, 2
vCPUs, 1 virtual core, a variable ECU, and a moderate I/O
capacity.

If possible, choose a DB instance class large enough that a typical
query working set can be held in memory. When working sets are held
in memory, the system can avoid writing to disk, which improves
performance. For more information, see [DB instance classes](concepts-dbinstanceclass.md).

In RDS for Oracle, you can select **Include additional**
**memory configurations**. These configurations are
optimized for a high ratio of memory to vCPU. For example,
**db.r5.6xlarge.tpc2.mem4x** is a db.r5.8x DB
instance that has 2 threads per core (tpc2) and 4x the memory of a
standard db.r5.6xlarge DB instance. For more information, see [RDS for Oracle DB instance classes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Oracle.Concepts.InstanceClasses.html).

**CLI option:**

`--db-instance-class`

**RDS API parameter:**

`DBInstanceClass`

All**DB instance identifier**

The name for your DB instance. Name your DB instances in the same way that you
name your on-premises servers. Your DB instance identifier can contain up
to 63 alphanumeric characters, and must be unique for your account
in the AWS Region you chose.

**CLI option:**

`--db-instance-identifier`

**RDS API parameter:**

`DBInstanceIdentifier`

All**DB parameter group**

A parameter group for your DB instance. You can choose the default
parameter group, or you can create a custom parameter group.

If you are using the
BYOL model for RDS for Db2, before creating a DB instance, you must
first create a custom parameter group that contains your IBM Site ID
and IBM Customer ID. For more information, see [Bring your own license (BYOL) for Db2](db2-licensing.md#db2-licensing-options-byol).

For more information, see [Parameter groups for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html).

**CLI option:**

`--db-parameter-group-name`

**RDS API parameter:**

`DBParameterGroupName`

All**DB subnet group**

The DB subnet group you want to use for the DB cluster.

Select **Choose existing** to use an
existing DB subnet group. Then choose the required subnet group from the
**Existing DB subnet groups** dropdown
list.

Choose **Automatic setup**
to let RDS select a compatible DB subnet group. If none exist, RDS
creates a new subnet group for your cluster.

For more
information, see [Working with DB subnet groups](user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.Subnets).

**CLI option:**

`--db-subnet-group-name`

**RDS API parameter:**

`DBSubnetGroupName`

All**Dedicated Log Volume**

Use a dedicated log volume (DLV) to store database transaction
logs on a storage volume that's separate from the volume containing
the database tables.

For more information, see [Using a dedicated log volume (DLV)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.dlv.html).

**CLI option:**

`--dedicated-log-volume`

**RDS API parameter:**

`DedicatedLogVolume`

All**Deletion protection**

**Enable deletion protection** to prevent your DB
instance from being deleted. If you create a production DB instance with
the AWS Management Console, deletion protection is enabled by default.

For more information, see [Deleting a DB instance](user-deleteinstance.md).

**CLI option:**

`--deletion-protection`

`--no-deletion-protection`

**RDS API parameter:**

`DeletionProtection`

All**Encryption**

**Enable Encryption** to enable encryption at
rest for this DB instance.

For more information, see [Encrypting Amazon RDS resources](overview-encryption.md).

**CLI option:**

`--storage-encrypted`

`--no-storage-encrypted`

**RDS API parameter:**

`StorageEncrypted`

All**Enhanced Monitoring**

**Enable enhanced monitoring** to enable
gathering metrics in real time for the operating system that your DB
instance runs on.

For more information, see [Monitoring OS metrics with Enhanced Monitoring](user-monitoring-os.md).

**CLI options:**

`--monitoring-interval`

`--monitoring-role-arn`

**RDS API parameters:**

`MonitoringInterval`

`MonitoringRoleArn`

All**Engine type**

Choose the database engine to be used for this DB instance.

**CLI option:**

`--engine`

**RDS API parameter:**

`Engine`

All**Initial database name**

The name for the database on your DB instance. If you don't
provide a name, Amazon RDS doesn't create a database on the DB
instance (except for Oracle and PostgreSQL).
The
name can't be a word reserved by the database engine, and has
other constraints depending on the DB engine.

Db2:

- It must contain 1–8 alphanumeric characters.

- It must start with a-z, A-Z, @, $, or #, and be followed
by a-z, A-Z, 0-9, \_, @, #, or $.

- It can't contain spaces.

- For more information, see [Additional considerations](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-db-instance-prereqs.html#db2-prereqs-additional-considerations).

MariaDB and MySQL:

- It must contain 1–64 alphanumeric
characters.

Oracle:

- It must contain 1–8 alphanumeric characters.

- It can't be `NULL`. The default value is
`ORCL`.

- It must begin with a letter.

PostgreSQL:

- It must contain 1–63 alphanumeric
characters.

- It must begin with a letter or an underscore. Subsequent
characters can be letters, underscores, or digits
(0-9).

- The initial database name is `postgres`.

**CLI option:**

`--db-name`

**RDS API parameter:**

`DBName`

All except SQL Server**License**

Valid values for the license model:

- **bring-your-own-license** or
**marketplace-license** for Db2.

- **general-public-license** for
MariaDB.

- **license-included** for Microsoft SQL
Server.

- **general-public-license** for
MySQL.

- **license-included** or
**bring-your-own-license** for
Oracle.

- **postgresql-license** for
PostgreSQL.

**CLI option:**

`--license-model`

**RDS API parameter:**

`LicenseModel`

All

**Log exports**

The types of database log files to publish to Amazon CloudWatch Logs.

For more information, see [Publishing database logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.Procedural.UploadtoCloudWatch.html).

**CLI option:**

`--enable-cloudwatch-logs-exports`

**RDS API parameter:**

`EnableCloudwatchLogsExports`

All

**Maintenance window**

The 30-minute window in which pending modifications to your DB
instance are applied. If the time period doesn't matter, choose
**No Preference**.

For more information, see [Amazon RDS maintenance window](user-upgradedbinstance-maintenance.md#Concepts.DBMaintenance).

**CLI option:**

`--preferred-maintenance-window`

**RDS API parameter:**

`PreferredMaintenanceWindow`

All**Manage master credentials in AWS Secrets Manager**

Select **Manage master credentials in AWS Secrets Manager**
to manage the master user password in a secret in Secrets Manager.

Optionally, choose a KMS key to use to protect the secret.
Choose from the KMS keys in your account, or enter the key from a
different account.

For more information, see [Password management with Amazon RDS and AWS Secrets Manager](rds-secrets-manager.md).

**CLI option:**

`--manage-master-user-password |
									--no-manage-master-user-password`

`--master-user-secret-kms-key-id`

**RDS API parameter:**

`ManageMasterUserPassword`

`MasterUserSecretKmsKeyId`

All**Master password**

The password for your master user account. The password has the
following number of printable ASCII characters (excluding
`/`, `"`, a space, and `@`)
depending on the DB engine:

- Db2: 8–255

- Oracle: 8–30

- MariaDB and MySQL: 8–41

- SQL Server and PostgreSQL: 8–128

**CLI option:**

`--master-user-password`

**RDS API parameter:**

`MasterUserPassword`

All**Master username**

The name that you use as the master username to log in to your
DB instance with all database privileges. Note the following naming
restrictions:

- The name can contain 1–16 alphanumeric characters and underscores.

- The first character must be a letter.

- The name can't be a word reserved by the database engine.

You can't change the master username after you create the
DB instance.

For Db2, we recommend that you use the same master username as
your self-managed Db2 instance name.

For more information on privileges granted to the master user, see
[Master user account privileges](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.MasterAccounts.html).

**CLI option:**

`--master-username`

**RDS API parameter:**

`MasterUsername`

All**Microsoft SQL Server Windows**
**Authentication**

**Enable Microsoft SQL Server Windows**
**authentication**, then **Browse**
**Directory** to choose the directory where you want to
allow authorized domain users to authenticate with this SQL Server
instance using Windows Authentication.

**CLI options:**

`--domain`

`--domain-iam-role-name`

**RDS API parameters:**

`Domain`

`DomainIAMRoleName`

SQL Server

**Multi-AZ deployment**

**Create a standby instance** to create a passive
secondary replica of your DB instance in another Availability Zone for
failover support. We recommend Multi-AZ for production workloads to
maintain high availability.

For development and testing, you can choose **Do not**
**create a standby instance**.

For more information, see [Configuring and managing a Multi-AZ deployment for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html).

**CLI option:**

`--multi-az`

`--no-multi-az`

**RDS API parameter:**

`MultiAZ`

All**National character set (NCHAR)**

The national character set for your DB instance, commonly called the
NCHAR character set. You can set the national character set to
either AL16UTF16 (default) or UTF-8. You can't change the
national character set after you create the DB instance.

The national character set is different from the DB character set.
Unlike the DB character set, the national character set specifies
the encoding only for NCHAR data types (NCHAR, NVARCHAR2, and NCLOB)
columns without affecting database metadata.

For more information, see [RDS for Oracle character sets](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html).

**CLI option:**

`--nchar-character-set-name`

**API parameter:**

`NcharCharacterSetName`

Oracle**Network type**

The IP addressing protocols supported by the DB instance.

**IPv4** (the default) to specify that resources
can communicate with the DB instance only over the Internet Protocol
version 4 (IPv4) addressing protocol.

**Dual-stack mode** to specify that resources can
communicate with the DB instance over IPv4, Internet Protocol version 6
(IPv6), or both. Use dual-stack mode if you have any resources that
must communicate with your DB instance over the IPv6 addressing protocol.
Also, make sure that you associate an IPv6 CIDR block with all
subnets in the DB subnet group that you specify.

For more information, see [Amazon RDS IP addressing](user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.IP_addressing).

**CLI option:**

`--network-type`

**RDS API parameter:**

`NetworkType`

All

**Option group**

An option group for your DB instance. You can choose the default option
group or you can create a custom option group.

For more information, see [Working with option groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html).

**CLI option:**

`--option-group-name`

**RDS API parameter:**

`OptionGroupName`

All

**Performance Insights**

**Enable Performance Insights** to monitor your
DB instance load so that you can analyze and troubleshoot your database
performance.

Choose a retention period to determine how much Performance
Insights data history to keep. The retention setting is **Default (7 days)**. To retain your performance
data for longer, specify 1–24 months. For more information about retention periods, see
[Pricing and data retention for Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.Overview.cost.html).

Choose a KMS key to use to protect the key used to encrypt this
database volume. Choose from the KMS keys in your account, or
enter the key from a different account.

For more information, see [Monitoring DB load with Performance Insights on Amazon RDS](user-perfinsights.md).

**CLI options:**

`--enable-performance-insights`

`--no-enable-performance-insights`

`--performance-insights-retention-period`

`--performance-insights-kms-key-id`

**RDS API parameters:**

`EnablePerformanceInsights`

`PerformanceInsightsRetentionPeriod`

`PerformanceInsightsKMSKeyId`

All except
Db2

**Provisioned IOPS**

The Provisioned IOPS (I/O operations per second) value for the DB
instance. This setting is available only if you choose one of the
following for **Storage type**:

- **General purpose SSD (gp3)**

- **Provisioned IOPS SSD (io1)**

- **Provisioned IOPS SSD (io2)**

For more information, see [Amazon RDS DB instance storage](chap-storage.md).

**CLI option:**

`--iops`

**RDS API parameter:**

`Iops`

All

**Public access**

**Yes** to give the DB instance a public IP address,
meaning that it's accessible outside the VPC. To be publicly
accessible, the DB instance also has to be in a public subnet in the
VPC.

**No** to make the DB instance accessible only from
inside the VPC.

For more information, see [Hiding a DB instance in a VPC from the internet](user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.Hiding).

To connect to a DB instance from outside of its VPC, the DB instance
must be publicly accessible. Also, access must be granted using the
inbound rules of the DB instance's security group. In addition, other
requirements must be met. For more information, see [Can't connect to Amazon RDS DB instance](chap-troubleshooting.md#CHAP_Troubleshooting.Connecting).

If your DB instance isn't publicly accessible, use an AWS
Site-to-Site VPN connection or an Direct Connect connection to access it
from a private network. For more information, see [Internetwork traffic privacy](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/inter-network-traffic-privacy.html).

**CLI option:**

`--publicly-accessible`

`--no-publicly-accessible`

**RDS API parameter:**

`PubliclyAccessible`

All**RDS Extended Support**

Select **Enable RDS Extended Support** to allow
supported major engine versions to continue running past the RDS end
of standard support date.

When you create a DB instance, Amazon RDS defaults to RDS Extended Support. To prevent the creation of
a new DB instance after the RDS end of standard support date and to
avoid charges for RDS Extended Support, disable this setting. Your existing DB
instances won't incur charges until the RDS Extended Support pricing start
date.

For more information, see [Amazon RDS Extended Support with Amazon RDS](extended-support.md).

**CLI option:**

`--engine-lifecycle-support`

**RDS API parameter:**

`EngineLifecycleSupport`

MySQL

PostgreSQL

**RDS Proxy**

Choose **Create an RDS Proxy** to create a proxy
for your DB instance. Amazon RDS automatically creates an IAM role and a
Secrets Manager secret for the proxy.

For more information, see [Amazon RDS Proxy](rds-proxy.md).

Not available when creating a DB instance.

MariaDB

MySQL

PostgreSQL

**Storage autoscaling**

**Enable storage autoscaling** to enable Amazon RDS to
automatically increase storage when needed to avoid having your DB
instance run out of storage space.

Use **Maximum storage threshold** to set the
upper limit for Amazon RDS to automatically increase storage for your DB
instance. The default is 1,000 GiB.

For more information, see [Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.Autoscaling.html).

**CLI option:**

`--max-allocated-storage`

**RDS API parameter:**

`MaxAllocatedStorage`

All

**Storage throughput**

The storage throughput value for the DB instance. This setting is
available only if you choose **General purpose SSD**
**(gp3)** for **Storage type**.

For more information, see [gp3 storage (recommended)](chap-storage.md#gp3-storage).

**CLI option:**

`--storage-throughput`

**RDS API parameter:**

`StorageThroughput`

All**Storage type**

The storage type for your DB instance.

If you choose **General Purpose SSD (gp3)**, you
can provision additional provisioned IOPS and storage throughput
under **Advanced settings**.

If you choose **Provisioned IOPS SSD (io1)** or **Provisioned IOPS SSD (io2)**, enter the
**Provisioned IOPS** value.

For more information, see [Amazon RDS storage types](chap-storage.md#Concepts.Storage).

**CLI option:**

`--storage-type`

**RDS API parameter:**

`StorageType`

All**Additional storage volumes**

You can add up to three additional storage volumes to your RDS for Oracle or RDS for SQL Server DB instance.
You can configure each additional volume to use gp3 or io2 storage
types. You can also specify different allocated storage, IOPS, and
throughput settings to optimize for your workload requirements.

Make sure you specify the volume names as `rdsdbdata2`,
`rdsdbdata3`, or `rdsdbdata4`.

The combined storage across the primary and additional volumes
cannot exceed 256 TiB.

**CLI option:**

`--additional-storage-volumes`

**RDS API parameter:**

`AdditionalStorageVolumes`

Oracle and SQL Server**Subnet group**

A DB subnet group to associate with this DB instance.

For more information, see [Working with DB subnet groups](user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.Subnets).

**CLI option:**

`--db-subnet-group-name`

**RDS API parameter:**

`DBSubnetGroupName`

All**Tenant database name**

The name of your initial PDB in the multi-tenant configuration of
the Oracle architecture. This setting is available only if you
choose **Multi-tenant configuration** for
**Architecture configuration**.

The tenant database name must differ from the name of your CDB,
which is named `RDSCDB`. You can't change the CDB
name.

**CLI option:**

`--db-name`

**RDS API parameter:**

`DBName`

Oracle

**Tenant database master username**

The name that you use as the master username to log in to your
tenant database (PDB) with all database privileges. This setting is
available only if you choose **Multi-tenant**
**configuration** for **Architecture**
**configuration**.

Note the following naming restrictions:

- The name can contain 1–16 alphanumeric characters and underscores.

- The first character must be a letter.

- The name can't be a word reserved by the database engine.

You can't do the following:

- Change the tenant master username after you create the
tenant database.

- Log in with the tenant master username to the CDB.

**CLI option:**

`--master-username`

**RDS API parameter:**

`MasterUsername`

Oracle

**Tenant database master password**

The password for the master user account of your tenant database
(PDB). This setting is available only if you choose
**Multi-tenant configuration** for
**Architecture configuration**.

The password has 8–30 printable ASCII characters, excluding
`/`, `"`, a space, and
`@`.

**CLI option:**

`--master-password`

**RDS API parameter:**

`MasterPassword`

Oracle

**Tenant database character set**

The character set of the initial tenant database. This setting is
available only if you choose **Multi-tenant**
**configuration** for **Architecture**
**configuration**. Only RDS for Oracle CDB instances are
supported.

The default value of **AL32UTF8** for the tenant
database character set is for the Unicode 5.0 UTF-8 Universal
character set. You can choose a tenant database character set that
is different from the character set of the CDB.

For more information, see [RDS for Oracle character sets](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html).

**CLI option:**

`--character-set-name`

**RDS API parameter:**

`CharacterSetName`

Oracle

**Tenant database national character set**

The national character set for your tenant database, commonly
called the `NCHAR` character set. This setting is
available only if you choose **Multi-tenant**
**configuration** for **Architecture**
**configuration**. Only RDS for Oracle CDB instances are
supported.

You can set the national character set to either
**AL16UTF16** (default) or
**UTF-8**. You can't change the national
character set after you create the tenant database.

The tenant database national character set is different from the
tenant database character set. The national character set specifies
the encoding only for columns that use the `NCHAR` data
type ( `NCHAR`, `NVARCHAR2`, and
`NCLOB`) and doesn't affect database metadata.

For more information, see [RDS for Oracle character sets](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html).

**CLI option:**

`--nchar-character-set-name`

**API parameter:**

`NcharCharacterSetName`

Oracle

**Time zone**

The time zone for your DB instance. If you don't choose a time zone,
your DB instance uses the default time zone. You can't change the
time zone after the DB instance is created.

For more information, see [Local time zone for Amazon RDS for Db2 DB instances](db2-time-zone.md) and [Local time zone for Microsoft SQL Server DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.Concepts.General.TimeZone.html).

**CLI option:**

`--timezone`

**RDS API parameter:**

`Timezone`

Db2

SQL Server

RDS Custom for SQL Server

**Virtual Private Cloud (VPC)**

A VPC based on the Amazon VPC service to associate with this DB
instance.

For more information, see [Amazon VPC and Amazon RDS](user-vpc.md).

For the CLI and API, you specify the VPC security group
IDs.

All**VPC security group (firewall)**

The security group to associate with the DB instance.

For more information, see [Overview of VPC security groups](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.RDSSecurityGroups.html#Overview.RDSSecurityGroups.VPCSec).

**CLI option:**

`--vpc-security-group-ids`

**RDS API parameter:**

`VpcSecurityGroupIds`

All

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Creating a DB instance

Creating resources with AWS CloudFormation
