# Connecting to your DB instance using IAM authentication with the AWS drivers

The AWS suite of drivers has been designed to provide support for faster
switchover and failover times, and authentication with AWS Secrets Manager, AWS Identity and Access Management (IAM),
and Federated Identity. The AWS drivers rely on monitoring DB instance
status and being aware of the instance topology to determine the new writer.
This approach reduces switchover and failover times to single-digit seconds,
compared to tens of seconds for open-source drivers.

For more information on the AWS drivers, see the
corresponding language driver for your [RDS for MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MariaDB.Connecting.Drivers.html#MariaDB.Connecting.JDBCDriver), [RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Connecting.Drivers.html#MySQL.Connecting.JDBCDriver), or [RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Connecting.JDBCDriver.html) DB instance.

###### Note

The only features supported for RDS for MariaDB are authentication with AWS Secrets Manager,
AWS Identity and Access Management (IAM), and Federated Identity.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Connecting to your DB instance using IAM authentication

Connecting using IAM: AWS CLI and mysql client
