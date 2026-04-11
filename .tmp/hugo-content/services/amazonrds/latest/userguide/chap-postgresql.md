# Amazon RDS for PostgreSQL

Amazon RDS supports DB instances running several versions of PostgreSQL. For a list of
available versions, see [Available PostgreSQL database versions](postgresql-concepts-general-dbversions.md).

You can create DB instances and DB snapshots, point-in-time restores and backups. DB
instances running PostgreSQL support Multi-AZ deployments, read replicas, Provisioned IOPS,
and can be created inside a virtual private cloud (VPC). You can also use Secure Socket
Layer (SSL) to connect to a DB instance running PostgreSQL.

Before creating a DB instance, make sure to complete the steps in [Setting up your Amazon RDS environment](chap-settingup.md).

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

To import PostgreSQL data into a DB instance, follow the information in the [Importing data into PostgreSQL on Amazon RDS](postgresql-procedural-importing.md) section.

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

- [Common management tasks for Amazon RDS for PostgreSQL](chap-postgresql-commontasks.md)

- [Working with the Database Preview environment](working-with-the-database-preview-environment.md)

- [Available PostgreSQL database versions](postgresql-concepts-general-dbversions.md)

- [Understanding the RDS for PostgreSQL incremental release process](postgresql-concepts-general-releaseprocess.md)

- [Supported PostgreSQL extension versions](postgresql-concepts-general-featuresupport-extensions.md)

- [Working with PostgreSQL features supported by Amazon RDS for PostgreSQL](postgresql-concepts-general-featuresupport.md)

- [Connecting to a DB instance running the PostgreSQL database engine](user-connecttopostgresqlinstance.md)

- [Securing connections to RDS for PostgreSQL with SSL/TLS](postgresql-concepts-general-security.md)

- [Using Kerberos authentication with Amazon RDS for PostgreSQL](postgresql-kerberos.md)

- [Using a custom DNS server for outbound network access](appendix-postgresql-commondbatasks-customdns.md)

- [Upgrades of the RDS for PostgreSQL DB engine](user-upgradedbinstance-postgresql.md)

- [Upgrading a PostgreSQL DB snapshot engine version](user-upgradedbsnapshot-postgresql.md)

- [Working with read replicas for Amazon RDS for PostgreSQL](user-postgresql-replication-readreplicas.md)

- [Improving query performance for RDS for PostgreSQL with Amazon RDS Optimized Reads](user-postgresql-optimizedreads.md)

- [Importing data into PostgreSQL on Amazon RDS](postgresql-procedural-importing.md)

- [Exporting data from an RDS for PostgreSQL DB instance to Amazon S3](postgresql-s3-export.md)

- [Invoking an AWS Lambda function from an RDS for PostgreSQL DB instance](postgresql-lambda.md)

- [Common DBA tasks for Amazon RDS for PostgreSQL](appendix-postgresql-commondbatasks.md)

- [Tuning with wait events for RDS for PostgreSQL](postgresql-tuning.md)

- [Tuning RDS for PostgreSQL with Amazon DevOps Guru proactive insights](postgresql-tuning-proactive-insights.md)

- [Using PostgreSQL extensions with Amazon RDS for PostgreSQL](appendix-postgresql-commondbatasks-extensions.md)

- [Working with the supported foreign data wrappers for Amazon RDS for PostgreSQL](appendix-postgresql-commondbatasks-extensions-foreign-data-wrappers.md)

- [Working with Trusted Language Extensions for PostgreSQL](postgresql-trusted-language-extension.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Oracle Database engine releases

Common management
tasks

All content copied from https://docs.aws.amazon.com/.
