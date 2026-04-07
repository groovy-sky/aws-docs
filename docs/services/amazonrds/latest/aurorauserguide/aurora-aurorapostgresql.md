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

- [Working with the database preview environment](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/working-with-the-apg-database-preview-environment.html)

- [Security with Amazon Aurora PostgreSQL](aurorapostgresql-security.md)

- [Updating applications to connect to Aurora PostgreSQL DB clusters using new SSL/TLS certificates](ssl-certificate-rotation-aurora-postgresql.md)

- [Using Kerberos authentication with Aurora PostgreSQL](postgresql-kerberos.md)

- [Migrating data to Amazon Aurora with PostgreSQL compatibility](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html)

- [Optimizing query performance in Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.optimizing.queries.html)

- [Working with unlogged tables in Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-postgresql-unlogged-tables.html)

- [Working with PostgreSQL autovacuum on Amazon Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.Autovacuum.html)

- [Managing TOAST OID contention in Amazon Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.TOAST_OID.html)

- [Using Babelfish for Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/babelfish.html)

- [Performance and scaling for Amazon Aurora PostgreSQL](aurorapostgresql-managing.md)

- [Tuning with wait events for Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Tuning.html)

- [Tuning Aurora PostgreSQL with Amazon DevOps Guru proactive insights](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/PostgreSQL.Tuning_proactive_insights.html)

- [Best practices with Amazon Aurora PostgreSQL](aurorapostgresql-bestpractices.md)

- [Replication with Amazon Aurora PostgreSQL](aurorapostgresql-replication.md)

- [Local write forwarding in Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-postgresql-write-forwarding.html)

- [Using Aurora PostgreSQL as a Knowledge Base for Amazon Bedrock](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.VectorDB.html)

- [Integrating Amazon Aurora PostgreSQL with other AWS services](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Integrating.html)

- [Monitoring query execution plans and peak memory for Aurora PostgreSQL](aurorapostgresql-monitoring-query-plans.md)

- [Managing query execution plans for Aurora PostgreSQL](aurorapostgresql-optimize.md)

- [Working with extensions and foreign data wrappers](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.html)

- [Working with Trusted Language Extensions for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/PostgreSQL_trusted_language_extension.html)

- [Amazon Aurora PostgreSQL reference](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Reference.html)

- [Database engine updates for Amazon Aurora PostgreSQL](aurorapostgresql-updates.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Database engine updates and fixes for Amazon Aurora MySQL

The database preview environment
