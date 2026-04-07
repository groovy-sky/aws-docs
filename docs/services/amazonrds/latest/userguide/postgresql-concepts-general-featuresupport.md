# Working with PostgreSQL features supported by Amazon RDS for PostgreSQL

Amazon RDS for PostgreSQL supports many of the most common PostgreSQL features. For example,
PostgreSQL has an autovacuum feature that performs routine maintenance on the database.
The autovacuum feature is active by default. Although you can turn off this feature, we
highly recommend that you keep it on. Understanding this feature and what you can do to
make sure it works as it should is a basic task of any DBA. For more information about
the autovacuum, see [Working with PostgreSQL autovacuum on Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Autovacuum.html). To learn more about
other common DBA tasks, [Common DBA tasks for Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.html).

RDS for PostgreSQL also supports extensions that add important functionality to the DB
instance. For example, you can use the PostGIS extension to work with spatial data, or
use the pg\_cron extension to schedule maintenance from within the instance. For more
information about PostgreSQL extensions, see [Using PostgreSQL extensions with Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Extensions.html).

Foreign data wrappers are a specific type of extension designed to let your
RDS for PostgreSQL DB instance work with other commercial databases or data types. For more
information about foreign data wrappers supported by RDS for PostgreSQL, see [Working with the supported foreign data wrappers for Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.Extensions.foreign-data-wrappers.html).

Following, you can find information about some other features supported by
RDS for PostgreSQL.

###### Topics

- [Custom data types and enumerations with RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.AlterEnum.html)

- [Event triggers for RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.EventTriggers.html)

- [Huge pages for RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.HugePages.html)

- [Performing logical replication for Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.LogicalReplication.html)

- [Configuring IAM authentication for logical replication connections](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.IAMLogicalReplication.html)

- [RAM disk for the stats\_temp\_directory](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.RamDisk.html)

- [Tablespaces for RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.FeatureSupport.Tablespaces.html)

- [RDS for PostgreSQL collations for EBCDIC and other mainframe migrations](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Collations.mainframe.migration.html)

- [Managing logical slot synchronization for RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.PostgreSQL.CommonDBATasks.pglogical.slot.synchronization.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PostgreSQL extension versions

Custom data types and enumerations
