# Best practices with Amazon Aurora PostgreSQL

Following, you can find several best practices for managing your Amazon Aurora PostgreSQL DB
cluster. Be sure to also review basic maintenance tasks. For more information, see [Performance and scaling for Amazon Aurora PostgreSQL](aurorapostgresql-managing.md).

###### Topics

- [Avoiding slow performance, automatic restart, and failover for Aurora PostgreSQL DB instances](#AuroraPostgreSQL.BestPractices.Avoiding)

- [Diagnosing table and index bloat](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.diag-table-ind-bloat.html)

- [Managing high object counts in Amazon Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/PostgreSQL.HighObjectCount.html)

- [Improved memory management in Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.BestPractices.memory.management.html)

- [Fast failover with Amazon Aurora PostgreSQL](aurorapostgresql-bestpractices-fastfailover.md)

- [Fast recovery after failover with cluster cache management for Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.cluster-cache-mgmt.html)

- [Managing Aurora PostgreSQL connection churn with pooling](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.BestPractices.connection_pooling.html)

- [Dead connection handling in PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Appendix.PostgreSQL.CommonDBATasks.DeadConnectionHandling.html)

- [Tuning memory parameters for Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.BestPractices.Tuning-memory-parameters.html)

- [Using Amazon CloudWatch metrics to analyze resource usage for Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL_AnayzeResourceUsage.html)

- [Using logical replication to perform a major version upgrade for Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.MajorVersionUpgrade.html)

- [Managing custom casts in Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/PostgreSQL.CustomCasts.html)

- [Best Practices for Parallel Queries in Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/PostgreSQL.ParallelQueries.html)

- [Troubleshooting storage issues in Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.BestPractices.TroubleshootingStorage.html)

## Avoiding slow performance, automatic restart, and failover for Aurora PostgreSQL DB instances

If you're running a heavy workload or workloads that spike beyond the allocated
resources of your DB instance, you can exhaust the resources on which you're running
your application and Aurora database. To get metrics on your database instance such as
CPU utilization, memory usage, and number of database connections used, you can refer to
the metrics provided by Amazon CloudWatch, Performance Insights, and Enhanced Monitoring. For more information on monitoring
your DB instance, see [Monitoring metrics in an Amazon Aurora cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/MonitoringAurora.html).

If your workload exhausts the resources you're using, your DB instance might slow
down, restart, or even fail over to another DB instance. To avoid this, monitor your
resource utilization, examine the workload running on your DB instance, and make
optimizations where necessary. If optimizations don't improve the instance metrics and
mitigate the resource exhaustion, consider scaling up your DB instance before you reach
its limits. For more information on available DB instance classes and their
specifications, see [Amazon AuroraDB instance classes](concepts-dbinstanceclass.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tuning Aurora PostgreSQL with Amazon DevOps Guru proactive insights

Diagnosing table and index bloat
