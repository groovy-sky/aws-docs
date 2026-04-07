# Support for SQL Server Analysis Services in Amazon RDS for SQL Server

Microsoft SQL Server Analysis Services (SSAS) is part of the Microsoft Business Intelligence (MSBI) suite. SSAS is an online
analytical processing (OLAP) and data mining tool that is installed within SQL Server. You use SSAS to analyze data to help make
business decisions. SSAS differs from the SQL Server relational database because SSAS is optimized for queries and calculations
common in a business intelligence environment.

You can turn on SSAS for existing or new DB instances. It's installed on the same DB
instance as your database engine. For more information on SSAS, see the Microsoft [Analysis services\
documentation](https://docs.microsoft.com/en-us/analysis-services).

Amazon RDS supports SSAS for SQL Server Standard and Enterprise Editions on the following versions:

- Tabular mode:

- SQL Server 2019, version 15.00.4043.16.v1 and higher

- SQL Server 2017, version 14.00.3223.3.v1 and higher

- SQL Server 2016, version 13.00.5426.0.v1 and higher

- Multidimensional mode:

- SQL Server 2019, version 15.00.4153.1.v1 and higher

- SQL Server 2017, version 14.00.3381.3.v1 and higher

- SQL Server 2016, version 13.00.5882.1.v1 and higher

###### Contents

- [Limitations](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.SSAS.html#SSAS.Limitations)

- [Turning on SSAS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Enabling.html)

  - [Creating an option group for SSAS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Enabling.html#SSAS.OptionGroup)

  - [Adding the SSAS option to the option group](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Enabling.html#SSAS.Add)

  - [Associating the option group with your DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Enabling.html#SSAS.Apply)

  - [Allowing inbound access to your VPC security group](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Enabling.html#SSAS.InboundRule)

  - [Enabling Amazon S3 integration](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Enabling.html#SSAS.EnableS3)
- [Deploying SSAS projects on Amazon RDS](ssas-deploy.md)

- [Monitoring the status of a deployment task](ssas-monitor.md)

- [Using SSAS on Amazon RDS](ssas-use.md)

  - [Setting up a Windows-authenticated user for SSAS](ssas-use.md#SSAS.Use.Auth)

  - [Adding a domain user as a database administrator](ssas-use.md#SSAS.Admin)

  - [Creating an SSAS proxy](ssas-use.md#SSAS.Use.Proxy)

  - [Scheduling SSAS database processing using SQL Server Agent](ssas-use.md#SSAS.Use.Schedule)

  - [Revoking SSAS access from the proxy](ssas-use.md#SSAS.Use.Revoke)
- [Backing up an SSAS database](ssas-backup.md)

- [Restoring an SSAS database](ssas-restore.md)

  - [Restoring a DB instance to a specified time](ssas-restore.md#SSAS.PITR)
- [Changing the SSAS mode](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.ChangeMode.html)

- [Turning off SSAS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Disable.html)

- [Troubleshooting SSAS issues](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.Trouble.html)

## Limitations

The following limitations apply to using SSAS on RDS for SQL Server:

- RDS for SQL Server supports running SSAS in Tabular or Multidimensional mode. For more information, see [Comparing tabular and \
multidimensional solutions](https://docs.microsoft.com/en-us/analysis-services/comparing-tabular-and-multidimensional-solutions-ssas) in the Microsoft documentation.

- You can only use one SSAS mode at a time. Before changing modes, make sure to
delete all of the SSAS databases.

For more information, see [Changing the SSAS mode](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SSAS.ChangeMode.html).

- Multi-AZ instances aren't supported.

- Instances must use self-managed Active Directory or AWS Directory Service for Microsoft Active Directory for SSAS authentication. For more information, see [Working with Active Directory with RDS for SQL Server](user-sqlserver-activedirectorywindowsauth.md).

- Users aren't given SSAS server administrator access, but they can be granted
database-level administrator access.

- The only supported port for accessing SSAS is 2383.

- You can't deploy projects directly. We provide an RDS stored procedure to do this. For more information, see [Deploying SSAS projects on Amazon RDS](ssas-deploy.md).

- Processing during deployment isn't supported.

- Using .xmla files for deployment isn't supported.

- SSAS project input files and database backup output files can only be in the
`D:\S3` folder on the DB instance.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Manually creating an IAM role for SQL Server Audit

Turning on SSAS
