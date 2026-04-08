# MariaDB on Amazon RDS SQL reference

Following, you can find descriptions of system stored procedures that are available for Amazon RDS
instances running the MariaDB DB engine.

You can use the system stored procedures that are available for MySQL DB
instances and MariaDB DB instances. These stored procedures are documented at
[RDS for MySQL stored procedure reference](appendix-mysql-sqlref.md). MariaDB DB instances support all
of the stored procedures, except for `mysql.rds_start_replication_until` and
`mysql.rds_start_replication_until_gtid`.

Additionally, the following system stored procedures are supported only for Amazon RDS DB
instances running MariaDB:

- [mysql.rds\_replica\_status](mysql-rds-replica-status.md)

- [mysql.rds\_set\_external\_master\_gtid](mysql-rds-set-external-master-gtid.md)

- [mysql.rds\_kill\_query\_id](mysql-rds-kill-query-id.md)

- [mysql.rds\_execute\_operation](mysql-rds-execute-operation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Migrating data from a MySQL DB snapshot to a MariaDB DB instance

mysql.rds\_replica\_status

All content copied from https://docs.aws.amazon.com/.
