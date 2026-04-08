# Transitioning a Amazon RDS for SQL Server database from OFFLINE to ONLINE

You can transition your Microsoft SQL Server database on an Amazon RDS DB instance from `OFFLINE` to `ONLINE`.

SQL Server method

Amazon RDS method

ALTER DATABASE `db_name` SET ONLINE;

EXEC rdsadmin.dbo.rds\_set\_database\_online `db_name`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Restoring
license-terminated DB instances

Using CDC

All content copied from https://docs.aws.amazon.com/.
