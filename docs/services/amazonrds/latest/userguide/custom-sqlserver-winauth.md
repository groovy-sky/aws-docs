# Working with Microsoft Active Directory with RDS Custom for SQL Server

RDS Custom for SQL Server allows to join your instances to a Self-Managed Active Directory (AD) or AWS Managed Microsoft AD.
This is regardless of where your AD is hosted, like an On-premises data center,
Amazon EC2 or with any other cloud service providers.

For authentication of users and services, you can use NTLM or Kerberos
authentication on your RDS Custom for SQL Server DB instance without using intermediary domains and forest trusts.
When a user tries to authenticate on your RDS Custom for SQL Server DB instance with a self joined Active Directory,
requests for authentication are forwarded to a self-managed AD or AWS Managed Microsoft AD that you specify.

In the following sections, you can find information about working with
Self Managed Active Directory and AWS Managed Active Directory for RDS Custom for SQL Server.

###### Topics

- [Region and version availability](#custom-sqlserver-WinAuth.Regions)

- [Configure Self-Managed or On-premise AD](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.config-Self-Managed.html)

- [Configure Microsoft Active Directory using Directory Service](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.config-ADS.html)

- [Network configuration port rules](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.NWConfigPorts.html)

- [Network Validation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.NWValidation.html)

- [Setting up Windows Authentication for RDS Custom for SQL Server instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.settingUp.html)

- [Managing a DB instance in a Domain](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.ManagingDBI.html)

- [Understanding Domain membership](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.Understanding.html)

- [Troubleshooting Active Directory](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-sqlserver-WinAuth.Troubleshoot.html)

## Region and version availability

RDS Custom for SQL Server supports both Self Managed AD and AWS Managed Microsoft AD using NTLM or Kerberos in all Regions where RDS Custom for SQL Server is supported.
For more information, see [Supported Regions and DB engines for RDS Custom](concepts-rds-fea-regions-db-eng-feature-rdscustom.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Starting and stopping an RDS Custom for SQL Server DB instance

Configure Self-Managed or On-premise AD
