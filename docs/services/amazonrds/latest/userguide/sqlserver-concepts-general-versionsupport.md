# Microsoft SQL Server versions on Amazon RDS

You can specify any currently supported Microsoft SQL Server version when creating a new DB
instance. You can specify the Microsoft SQL Server major version (such as Microsoft SQL
Server 14.00), and any supported minor version for the specified major version. If no
version is specified, Amazon RDS defaults to a supported version, typically the most recent
version. If a major version is specified but a minor version is not, Amazon RDS defaults to a
recent release of the major version you have specified.

The following table shows the supported SQL Server versions for all editions and all
AWS Regions, except where noted.

###### Note

You can also use the [describe-db-engine-versions](../../../cli/latest/reference/rds/describe-db-engine-versions.md) AWS CLI command to see a list of supported
versions, as well as defaults for newly created DB instances. You can view the major
versions of your SQL Server databases by running the [describe-db-major-engine-versions](../../../cli/latest/reference/rds/describe-db-major-engine-versions.md) AWS CLI command or by using the [DescribeDBMajorEngineVersions](../../../../reference/amazonrds/latest/apireference/api-describedbmajorengineversions.md) RDS API operation.

Major versionMinor versionRDS API `EngineVersion` and CLI
`engine-version`SQL Server 2022

16.00.4236.2 (CU23)

16.00.4230.2 (CU22 GDR)

16.00.4225.2 (CU22)

16.00.4215.2 (CU21)

16.00.4210.1 (CU20 GDR)

16.00.4205.1 (CU20)

16.00.4195.2 (CU19)

16.00.4185.3 (CU18)

16.00.4175.1 (CU17)

16.00.4165.4 (CU16)

16.00.4150.1 (CU15)

16.00.4140.3 (CU14 GDR)

16.00.4135.4 (CU14)

16.00.4131.2 (CU13)

16.00.4125.3 (CU13)

16.00.4120.1 (CU12 GDR)

16.00.4115.5 (CU12)

16.00.4105.2 (CU11)

16.00.4095.4 (CU10)

16.00.4085.2 (CU9)

`16.00.4236.2.v1`

`16.00.4230.2.v1`

`16.00.4225.2.v1`

`16.00.4215.2.v1`

`16.00.4210.1.v1`

`16.00.4205.1.v1`

`16.00.4195.2.v1`

`16.00.4185.3.v1`

`16.00.4175.1.v1`

`16.00.4165.4.v1`

`16.00.4150.1.v1`

`16.00.4140.3.v1`

`16.00.4135.4.v1`

`16.00.4131.2.v1`

`16.00.4125.3.v1`

`16.00.4120.1.v1`

`16.00.4115.5.v1`

`16.00.4105.2.v1`

`16.00.4095.4.v1`

`16.00.4085.2.v1`

SQL Server 2019

15.00.4455.2 (CU32 GDR)

15.00.4445.1 (CU32 GDR)

15.00.4440.1 (CU32 GDR)

15.00.4435.7 (CU32)

15.00.4430.1 (CU32)

15.00.4420.2 (CU31)

15.00.4415.2 (CU30)

15.00.4410.1 (CU29 GDR)

15.00.4395.2 (CU28)

15.00.4390.2 (CU28)

15.00.4385.2 (CU28)

15.00.4382.1 (CU27)

15.00.4375.4 (CU27)

15.00.4365.2 (CU26)

15.00.4355.3 (CU25)

15.00.4345.5 (CU24)

15.00.4335.1 (CU23)

15.00.4322.2 (CU22)

15.00.4316.3 (CU21)

15.00.4312.2 (CU20)

15.00.4236.7 (CU16)

15.00.4198.2 (CU15)

15.00.4153.1 (CU12)

15.00.4073.23 (CU8)

15.00.4043.16 (CU5)

`15.00.4455.2.v1`

`15.00.4445.1.v1`

`15.00.4440.1.v1`

`15.00.4435.7.v1`

`15.00.4430.1.v1`

`15.00.4420.2.v1`

`15.00.4415.2.v1`

`15.00.4410.1.v1`

`15.00.4395.2.v1`

`15.00.4390.2.v1`

`15.00.4385.2.v1`

`15.00.4382.1.v1`

`15.00.4375.4.v1`

`15.00.4365.2.v1`

`15.00.4355.3.v1`

`15.00.4345.5.v1`

`15.00.4335.1.v1`

`15.00.4322.2.v1`

`15.00.4316.3.v1`

`15.00.4312.2.v1`

`15.00.4236.7.v1`

`15.00.4198.2.v1`

`15.00.4153.1.v1`

`15.00.4073.23.v1`

`15.00.4043.16.v1`

SQL Server 2017

14.00.3515.1 (CU31 GDR)

14.00.3505.1 (CU31 GDR)

14.00.3500.1.(CU31 GDR)

14.00.3495.9 (CU31 GDR)

14.00.3485.1 (CU31 GDR)

14.00.3480.1 (CU31)

14.00.3475.1 (CU31)

14.00.3471.2 (CU31)

14.00.3465.1 (CU31)

14.00.3460.9 (CU31)

14.00.3451.2 (CU30)

14.00.3421.10 (CU27)

14.00.3401.7 (CU25)

14.00.3381.3 (CU23)

14.00.3356.20 (CU22)

14.00.3294.2 (CU20)

14.00.3281.6 (CU19)

`14.00.3515.1.v1`

`14.00.3505.1.v1`

`14.00.3500.1.v1`

`14.00.3495.9.v1`

`14.00.3485.1.v1`

`14.00.3480.1.v1`

`14.00.3475.1.v1`

`14.00.3471.2.v1`

`14.00.3465.1.v1`

`14.00.3460.9.v1`

`14.00.3451.2.v1`

`14.00.3421.10.v1`

`14.00.3401.7.v1`

`14.00.3381.3.v1`

`14.00.3356.20.v1`

`14.00.3294.2.v1`

`14.00.3281.6.v1`

SQL Server 2016

13.00.6475.1 (GDR)

13.00.6470.1 (GDR)

13.00.6465.1 (GDR)

13.00.6460.7 (GDR)

13.00.6455.2 (GDR)

13.00.6450.1 (GDR)

13.00.6445.1 (GDR)

13.00.6441.1 (GDR)

13.00.6435.1 (GDR)

13.00.6430.49 (GDR)

13.00.6419.1 (SP3 + Hotfix)

13.00.6300.2 (SP3)

`14.00.6475.1.v1`

`14.00.6470.1.v1`

`13.00.6465.1.v1`

`13.00.6460.7.v1`

`13.00.6455.2.v1`

`13.00.6450.1.v1`

`13.00.6445.1.v1`

`13.00.6441.1.v1`

`13.00.6435.1.v1`

`13.00.6430.49.v1`

`13.00.6419.1.v1`

`13.00.6300.2.v1`

## Version management in Amazon RDS

Amazon RDS includes flexible version management that enables you to control when and how
your DB instance is patched or upgraded. This enables you to do the following for your
DB engine:

- Maintain compatibility with database engine patch versions.

- Test new patch versions to verify that they work with your application before
you deploy them in production.

- Plan and perform version upgrades to meet your service level agreements and
timing requirements.

### Microsoft SQL Server engine patching in Amazon RDS

Amazon RDS periodically aggregates official Microsoft SQL Server database patches into
a DB instance engine version that's specific to Amazon RDS. For more information about
the Microsoft SQL Server patches in each engine version, see [Version and feature support on Amazon RDS](chap-sqlserver.md#SQLServer.Concepts.General.FeatureSupport).

Currently, you manually perform all engine upgrades on your DB instance. For more
information, see [Upgrades of the Microsoft SQL Server DB engine](user-upgradedbinstance-sqlserver.md).

### Deprecation schedule for major engine versions of Microsoft SQL Server on Amazon RDS

The following table displays the planned schedule of deprecations for major engine versions of Microsoft SQL Server.

DateInformationJuly 14, 2026

Microsoft will stop critical patch updates for SQL Server 2016. For more information, see [Microsoft SQL\
Server 2016](https://learn.microsoft.com/en-us/lifecycle/products/sql-server-2016) in the Microsoft documentation.

July 14, 2026

Amazon RDS plans to end support of Microsoft SQL Server 2016 on RDS for SQL Server. At that time, any remaining
instances will be scheduled to migrate to SQL Server 2019 (latest minor version available). For more
information, see [Announcement:\
Amazon RDS for SQL Server ending support for Microsoft SQL Server 2016](https://repost.aws/articles/ARGkeWligDSU-MQgBwUQj0nA/announcement-amazon-rds-for-sql-server-ending-support-for-microsoft-sql-server-2016).

To avoid an automatic upgrade from Microsoft SQL Server 2016, you can upgrade at a time that is
convenient to you. For more information, see [Upgrading a DB instance engine version](user-upgradedbinstance-upgrading.md).

January 15, 2026Amazon RDS is starting to disable the creation of new RDS for SQL Server DB instances using Microsoft SQL Server 2016.
For more information, see [Announcement:\
Amazon RDS for SQL Server ending support for Microsoft SQL Server 2016](https://repost.aws/articles/ARGkeWligDSU-MQgBwUQj0nA/announcement-amazon-rds-for-sql-server-ending-support-for-microsoft-sql-server-2016).July 9, 2024

Microsoft will stop critical patch updates for SQL Server 2014. For more information, see [Microsoft SQL\
Server 2014](https://learn.microsoft.com/en-us/lifecycle/products/sql-server-2014) in the Microsoft documentation.

June 1, 2024

Amazon RDS plans to end support of Microsoft SQL Server 2014 on RDS for SQL Server. At that time, any remaining
instances will be scheduled to migrate to SQL Server 2016 (latest minor version available). For more
information, see [Announcement:\
Amazon RDS for SQL Server ending support for SQL Server 2014 major versions](https://repost.aws/articles/AR-eyAH1PSSuevuZRUE9FV3A).

To avoid an automatic upgrade from Microsoft SQL Server 2014, you can upgrade at a time that is
convenient to you. For more information, see [Upgrading a DB instance engine version](user-upgradedbinstance-upgrading.md).

July 12, 2022

Microsoft will stop critical patch updates for SQL Server 2012. For more information, see [Microsoft SQL\
Server 2012](https://docs.microsoft.com/en-us/lifecycle/products/microsoft-sql-server-2012) in the Microsoft documentation.

June 1, 2022

Amazon RDS plans to end support of Microsoft SQL Server 2012 on RDS for SQL Server. At that time, any remaining
instances will be scheduled to migrate to SQL Server 2014 (latest minor version available). For more
information, see [Announcement:\
Amazon RDS for SQL Server ending support for SQL Server 2012 major versions](https://repost.aws/questions/QUFNiETqrMQ_WT_AXSxOYNOA).

To avoid an automatic upgrade from Microsoft SQL Server 2012, you can upgrade at a time that is
convenient to you. For more information, see [Upgrading a DB instance engine version](user-upgradedbinstance-upgrading.md).

September 1, 2021Amazon RDS is starting to disable the creation of new RDS for SQL Server DB instances using Microsoft SQL Server 2012.
For more information, see [Announcement:\
Amazon RDS for SQL Server ending support for SQL Server 2012 major versions](https://repost.aws/questions/QUFNiETqrMQ_WT_AXSxOYNOA).July 12, 2019

The Amazon RDS team deprecated support for Microsoft SQL Server 2008 R2 in June 2019. Remaining instances
of Microsoft SQL Server 2008 R2 are migrating to SQL Server 2012 (latest minor version available).

To avoid an automatic upgrade from Microsoft SQL Server 2008 R2, you can upgrade at a time that is
convenient to you. For more information, see [Upgrading a DB instance engine version](user-upgradedbinstance-upgrading.md).

April 25, 2019Before the end of April 2019, you will no longer be able to create new Amazon RDS for SQL Server database instances using Microsoft SQL
Server 2008R2.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Updating applications for new SSL/TLS certificates

Feature support

All content copied from https://docs.aws.amazon.com/.
