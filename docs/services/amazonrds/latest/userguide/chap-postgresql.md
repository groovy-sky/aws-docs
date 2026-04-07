# Amazon RDS for PostgreSQL

Amazon RDS supports DB instances running several versions of PostgreSQL. For a list of
available versions, see [Available PostgreSQL database versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.DBVersions.html).

You can create DB instances and DB snapshots, point-in-time restores and backups. DB
instances running PostgreSQL support Multi-AZ deployments, read replicas, Provisioned IOPS,
and can be created inside a virtual private cloud (VPC). You can also use Secure Socket
Layer (SSL) to connect to a DB instance running PostgreSQL.

Before creating a DB instance, make sure to complete the steps in [Setting up your Amazon RDS environment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SettingUp.html).

You can use any standard SQL client application to run commands for the instance from your
client computer. Such applications include pgAdmin, a popular Open Source administration and
development tool for PostgreSQL, or psql, a command line utility that is part of a
PostgreSQL installation. To deliver a managed service experience, Amazon RDS doesn't provide
host access to DB instances. Also, it restricts access to certain system procedures and
tables that require advanced privileges. Amazon RDS supports access to databases on a DB instance
using any standard SQL client application. Amazon RDS doesn't allow direct host access to a
DB instance by using Telnet or Secure Shell (SSH).

Amazon RDS for PostgreSQL is compliant with many industry standards. For example, you can use
Amazon RDS for PostgreSQL databases to build HIPAA-compliant applications and to store
healthcare-related information. This includes storage for protected health information (PHI)
under a completed Business Associate Agreement (BAA) with AWS. Amazon RDS for PostgreSQL also
meets Federal Risk and Authorization Management Program (FedRAMP) security requirements.
Amazon RDS for PostgreSQL has received a FedRAMP Joint Authorization Board (JAB) Provisional
Authority to Operate (P-ATO) at the FedRAMP HIGH Baseline within the AWS GovCloud (US) Regions.
For more information on supported compliance standards, see [AWS cloud compliance](https://aws.amazon.com/compliance).

To import PostgreSQL data into a DB instance, follow the information in the [Importing data into PostgreSQL on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html) section.

###### Important

If you encounter an issue with your RDS for PostgreSQL DB instance, your AWS support agent
might need more information about the health of your databases. The goal is to ensure
that AWS Support gets the required information as soon as possible.

You can use PG Collector to help gather valuable database information in a
consolidated HTML file. For more information on PG Collector, how to run it, and how to
download the HTML report, see [PG\
Collector](https://github.com/awslabs/pg-collector).

Upon successful completion, and unless otherwise noted, the script returns output in a
readable HTML format. The script is designed to exclude any data or security details
from the HTML that might compromise your business. It also makes no modifications to
your database or its environment. However, if you find any information in the HTML that
you are uncomfortable sharing, feel free to remove the problematic information before
uploading the HTML. When the HTML is acceptable, upload it using the attachments section
in the case details of your support case.

###### Topics

- [Common management tasks for Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.CommonTasks.html)

- [Working with the Database Preview environment](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/working-with-the-database-preview-environment.html)

- [Available PostgreSQL database versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.DBVersions.html)

- [Understanding the RDS for PostgreSQL incremental release process](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.ReleaseProcess.html)

- [Supported PostgreSQL extension versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.Extensions.html)

- [Working with PostgreSQL features supported by Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.html)

- [Connecting to a DB instance running the PostgreSQL database engine](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToPostgreSQLInstance.html)

- [Securing connections to RDS for PostgreSQL with SSL/TLS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.Security.html)

- [Using Kerberos authentication with Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/postgresql-kerberos.html)

- [Using a custom DNS server for outbound network access](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.CustomDNS.html)

- [Upgrades of the RDS for PostgreSQL DB engine](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.PostgreSQL.html)

- [Upgrading a PostgreSQL DB snapshot engine version](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBSnapshot.PostgreSQL.html)

- [Working with read replicas for Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.Replication.ReadReplicas.html)

- [Improving query performance for RDS for PostgreSQL with Amazon RDS Optimized Reads](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.optimizedreads.html)

- [Importing data into PostgreSQL on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html)

- [Exporting data from an RDS for PostgreSQL DB instance to Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/postgresql-s3-export.html)

- [Invoking an AWS Lambda function from an RDS for PostgreSQL DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL-Lambda.html)

- [Common DBA tasks for Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.html)

- [Tuning with wait events for RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Tuning.html)

- [Tuning RDS for PostgreSQL with Amazon DevOps Guru proactive insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Tuning_proactive_insights.html)

- [Using PostgreSQL extensions with Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Extensions.html)

- [Working with the supported foreign data wrappers for Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Extensions.foreign-data-wrappers.html)

- [Working with Trusted Language Extensions for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL_trusted_language_extension.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Oracle Database engine releases

Common management
tasks
