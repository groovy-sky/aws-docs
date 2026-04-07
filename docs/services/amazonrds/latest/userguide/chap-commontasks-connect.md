# Connecting to an Amazon RDS DB instance

Before you can connect to a DB instance, you must create the DB instance. For
information, see [Creating an Amazon RDS DB instance](user-createdbinstance.md). After Amazon RDS provisions your DB instance, use any standard client application or utility
for your DB engine to connect to the DB instance. In the connection string, specify the DNS
address from the DB instance endpoint as the host parameter. Also, specify the port number
from the DB instance endpoint as the port parameter.

For more information about finding connection information for an Amazon RDS DB instance or scenarios for accessing a DB instance in a VPC, see the following topics.

- [Finding the connection information for an Amazon RDS DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_CommonTasks.Connect.EndpointAndPort.html)

- [Scenarios for accessing a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_CommonTasks.Connect.ScenariosForAccess.html)

## Connecting to DB instances with the AWS drivers

The AWS suite of drivers has been designed to provide support for faster switchover
and failover times, and authentication with AWS Secrets Manager, AWS Identity and Access Management (IAM), and Federated
Identity. The AWS drivers rely on monitoring DB instance status and being aware of the
instance topology to determine the new primary instance. This approach reduces
switchover and failover times to single-digit seconds, compared to tens of seconds for
open-source drivers.

The following table lists the features supported for each of the drivers. As new
service features are introduced, the goal of the AWS suite of drivers is to have
built-in support for these service features.

Feature[AWS JDBC Driver](https://github.com/awslabs/aws-advanced-jdbc-wrapper)[AWS Python Driver](https://github.com/awslabs/aws-advanced-python-wrapper)[AWS ODBC Driver for MySQL](https://github.com/aws/aws-mysql-odbc)Failover support[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheFailoverPlugin.md)[Yes](https://github.com/awslabs/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/using-plugins/UsingTheFailoverPlugin.md)[Yes](https://github.com/aws/aws-mysql-odbc/blob/main/docs/using-the-aws-driver/UsingTheAwsDriver.md)Enhanced failover monitoring[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheHostMonitoringPlugin.md)[Yes](https://github.com/awslabs/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/using-plugins/UsingTheHostMonitoringPlugin.md)[Yes](https://github.com/aws/aws-mysql-odbc/blob/main/docs/using-the-aws-driver/HostMonitoring.md)Read/write splitting[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheReadWriteSplittingPlugin.md)[Yes](https://github.com/awslabs/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/using-plugins/UsingTheReadWriteSplittingPlugin.md)NoDriver metadata connection[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheDriverMetadataConnectionPlugin.md)N/AN/ATelemetry[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/Telemetry.md)[Yes](https://github.com/aws/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/Telemetry.md)NoSecrets Manager[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheAwsSecretsManagerPlugin.md)[Yes](https://github.com/awslabs/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/using-plugins/UsingTheAwsSecretsManagerPlugin.md)[Yes](https://github.com/aws/aws-mysql-odbc/blob/main/docs/using-the-aws-driver/UsingTheAwsDriver.md)IAM authentication[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheIamAuthenticationPlugin.md)[Yes](https://github.com/awslabs/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/using-plugins/UsingTheIamAuthenticationPlugin.md)[Yes](https://github.com/aws/aws-mysql-odbc/blob/main/docs/using-the-aws-driver/UsingTheAwsDriver.md)Federated Identity (AD FS)[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheFederatedAuthPlugin.md)[Yes](https://github.com/awslabs/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/using-plugins/UsingTheFederatedAuthenticationPlugin.md)NoFederated Identity (Okta)[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/using-plugins/UsingTheFederatedAuthPlugin.md)NoNoMulti-AZ DB clusters[Yes](https://github.com/awslabs/aws-advanced-jdbc-wrapper/blob/main/docs/using-the-jdbc-driver/SupportForRDSMultiAzDBCluster.md)[Yes](https://github.com/aws/aws-advanced-python-wrapper/blob/main/docs/using-the-python-driver/SupportForRDSMultiAzDBCluster.md)No

For more information on the AWS drivers, see the corresponding language driver for
your [RDS for MariaDB](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MariaDB.Connecting.Drivers.html#MariaDB.Connecting.JDBCDriver), [RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/MySQL.Connecting.Drivers.html#MySQL.Connecting.JDBCDriver), or [RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Connecting.JDBCDriver.html) DB instance.

###### Note

The only features supported for RDS for MariaDB are authentication with AWS Secrets Manager,
AWS Identity and Access Management (IAM), and Federated Identity.

## Connecting to a DB instance that's running a specific DB engine

To learn how to connect to a DB instance that is running a specific DB engine, follow
the instructions for your DB engine:

- [RDS for Db2](user-connecttodb2dbinstance.md)

- [RDS for MariaDB](user-connecttomariadbinstance.md)

- [RDS for SQL Server](user-connecttomicrosoftsqlserverinstance.md)

- [RDS for MySQL](user-connecttoinstance.md)

- [RDS for Oracle](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToOracleInstance.html)

- [RDS for PostgreSQL](user-connecttopostgresqlinstance.md)

## Managing connections with RDS Proxy

You can also use Amazon RDS Proxy to manage connections to RDS for MariaDB, RDS for Microsoft SQL Server, RDS for MySQL, and RDS for PostgreSQL DB instances.
RDS Proxy allows applications to pool and share database connections to improve scalability. For more information, see
[Amazon RDS Proxy](rds-proxy.md).

## Database authentication options

Amazon RDS supports the following ways to authenticate database users:

- **Password authentication** – Your DB instance performs all administration of user accounts.
You create users and specify passwords with SQL statements. The SQL statements you can use depend on your DB engine.

- **AWS Identity and Access Management (IAM) database authentication** – You don't need to use a password when you connect to a DB
instance. Instead, you use an authentication token.

- **Kerberos authentication** – You use external authentication of database users using Kerberos and Microsoft
Active Directory. Kerberos is a network authentication protocol that uses tickets and
symmetric-key cryptography to eliminate the need to transmit passwords over the network.
Kerberos has been built into Active Directory and is designed to authenticate users to
network resources, such as databases.

IAM database authentication and Kerberos authentication are available only for
specific DB engines and versions.

For more information, see [Database authentication with Amazon RDS](database-authentication.md).

## Encrypted connections

You can use Secure Socket Layer (SSL) or Transport Layer Security (TLS) from your application to encrypt a
connection to a DB instance. Each DB engine has its own process for implementing SSL/TLS. For more information,
see [Using SSL/TLS to encrypt a connection to a DB instance or cluster](usingwithrds-ssl.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Creating resources with AWS CloudFormation

Finding the connection information
