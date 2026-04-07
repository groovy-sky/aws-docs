# MariaDB on Amazon RDS SQL reference

Following, you can find descriptions of system stored procedures that are available for Amazon RDS
instances running the MariaDB DB engine.

You can use the system stored procedures that are available for MySQL DB
instances and MariaDB DB instances. These stored procedures are documented at
[RDS for MySQL stored procedure reference](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.SQLRef.html). MariaDB DB instances support all
of the stored procedures, except for `mysql.rds_start_replication_until` and
`mysql.rds_start_replication_until_gtid`.

Additionally, the following system stored procedures are supported only for Amazon RDS DB
instances running MariaDB:

- [mysql.rds\_replica\_status](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql_rds_replica_status.html)

- [mysql.rds\_set\_external\_master\_gtid](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql_rds_set_external_master_gtid.html)

- [mysql.rds\_kill\_query\_id](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql_rds_kill_query_id.html)

- [mysql.rds\_execute\_operation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql_rds_execute_operation.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Migrating data from a MySQL DB snapshot to a MariaDB DB instance

mysql.rds\_replica\_status
