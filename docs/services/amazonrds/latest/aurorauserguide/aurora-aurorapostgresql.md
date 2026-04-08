# Working with Amazon Aurora PostgreSQL

Amazon Aurora PostgreSQL is a fully managed, PostgreSQL–compatible, and ACID–compliant relational
database engine that combines the speed, reliability, and manageability of Amazon Aurora
with the simplicity and cost-effectiveness of open-source databases. Aurora PostgreSQL is a
drop-in replacement for PostgreSQL and makes it simple and cost-effective to set up,
operate, and scale your new and existing PostgreSQL deployments, thus freeing you to focus
on your business and applications. To learn more about Aurora in general, see
[What is Amazon Aurora?](chap-auroraoverview.md).

In addition to the benefits of Aurora, Aurora PostgreSQL offers a convenient migration pathway from
Amazon RDS into Aurora, with push-button migration tools that convert your existing RDS for PostgreSQL
applications to Aurora PostgreSQL. Routine database tasks such as provisioning, patching, backup, recovery,
failure detection, and repair are also easy to manage with Aurora PostgreSQL.

Aurora PostgreSQL can work with many industry standards. For example, you can use
Aurora PostgreSQL databases to build HIPAA-compliant applications and to store healthcare
related information, including protected health information (PHI), under a completed
Business Associate Agreement (BAA) with AWS.

Aurora PostgreSQL is FedRAMP HIGH eligible. For details about AWS and compliance efforts,
see [AWS services in scope by\
compliance program](https://aws.amazon.com/compliance/services-in-scope).

###### Important

If you encounter an issue with your Aurora PostgreSQL DB cluster, your AWS support agent might need more information about the health of your databases.
The goal is to ensure that AWS Support gets as much of the required information as possible in the shortest possible time.

You can use PG Collector to help gather valuable database information in a consolidated HTML file. For more information on PG Collector, how to run it, and how to download the HTML report, see
[PG Collector](https://github.com/awslabs/pg-collector).

Upon successful completion, and unless otherwise noted, the script returns output in a readable HTML format. The script is designed to exclude any data or security details from the HTML that might compromise your business.
It also makes no modifications to your database or its environment. However, if you find any information in the HTML that you are uncomfortable sharing, feel free to remove the problematic information before uploading the HTML.
When the HTML is acceptable, upload it using the attachments section in the case details of your support case.

###### Topics

- [Working with the database preview environment](working-with-the-apg-database-preview-environment.md)

- [Security with Amazon Aurora PostgreSQL](aurorapostgresql-security.md)

- [Updating applications to connect to Aurora PostgreSQL DB clusters using new SSL/TLS certificates](ssl-certificate-rotation-aurora-postgresql.md)

- [Using Kerberos authentication with Aurora PostgreSQL](postgresql-kerberos.md)

- [Migrating data to Amazon Aurora with PostgreSQL compatibility](aurorapostgresql-migrating.md)

- [Optimizing query performance in Aurora PostgreSQL](aurorapostgresql-optimizing-queries.md)

- [Working with unlogged tables in Aurora PostgreSQL](aurora-postgresql-unlogged-tables.md)

- [Working with PostgreSQL autovacuum on Amazon Aurora PostgreSQL](appendix-postgresql-commondbatasks-autovacuum.md)

- [Managing TOAST OID contention in Amazon Aurora PostgreSQL](appendix-postgresql-commondbatasks-toast-oid.md)

- [Using Babelfish for Aurora PostgreSQL](babelfish.md)

- [Performance and scaling for Amazon Aurora PostgreSQL](aurorapostgresql-managing.md)

- [Tuning with wait events for Aurora PostgreSQL](aurorapostgresql-tuning.md)

- [Tuning Aurora PostgreSQL with Amazon DevOps Guru proactive insights](postgresql-tuning-proactive-insights.md)

- [Best practices with Amazon Aurora PostgreSQL](aurorapostgresql-bestpractices.md)

- [Replication with Amazon Aurora PostgreSQL](aurorapostgresql-replication.md)

- [Local write forwarding in Aurora PostgreSQL](aurora-postgresql-write-forwarding.md)

- [Using Aurora PostgreSQL as a Knowledge Base for Amazon Bedrock](aurorapostgresql-vectordb.md)

- [Integrating Amazon Aurora PostgreSQL with other AWS services](aurorapostgresql-integrating.md)

- [Monitoring query execution plans and peak memory for Aurora PostgreSQL](aurorapostgresql-monitoring-query-plans.md)

- [Managing query execution plans for Aurora PostgreSQL](aurorapostgresql-optimize.md)

- [Working with extensions and foreign data wrappers](appendix-postgresql-commondbatasks.md)

- [Working with Trusted Language Extensions for PostgreSQL](postgresql-trusted-language-extension.md)

- [Amazon Aurora PostgreSQL reference](aurorapostgresql-reference.md)

- [Database engine updates for Amazon Aurora PostgreSQL](aurorapostgresql-updates.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Database engine updates and fixes for Amazon Aurora MySQL

The database preview environment

All content copied from https://docs.aws.amazon.com/.
