# Connecting to your Db2 DB instance

After Amazon RDS provisions your Amazon RDS for Db2 DB instance, you can use any standard SQL client
application to connect to the DB instance. Because Amazon RDS is a managed service, you can't
sign in as `SYSADM`, `SYSCTRL`, `SECADM`, or
`SYSMAINT`.

You can connect to a DB instance that is running the IBM Db2 database engine by using
IBM Db2 CLP, IBM CLPPlus, DBeaver, or IBM Db2 Data Management Console.

###### Note

Connecting to a Db2 database can fail if your RDS for Db2 DB instance doesn't have
sufficient memory. For more information, see [Database connection error](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-troubleshooting.html#db2-database-connection-error).

###### Topics

- [Finding the endpoint of your Amazon RDS for Db2 DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-finding-instance-endpoint.html)

- [Connecting to your Amazon RDS for Db2 DB instance with IBM Db2 CLP](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-connecting-with-clp-client.html)

- [Connecting to your Amazon RDS for Db2 DB instance with IBM CLPPlus](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-connecting-with-ibm-clpplus-client.html)

- [Connecting to your Amazon RDS for Db2 DB instance with DBeaver](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-connecting-with-dbeaver.html)

- [Connecting to your Amazon RDS for Db2 DB instance with IBM Db2 Data Management Console](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-connecting-with-ibm-data-management-console.html)

- [Considerations for security groups with Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-security-groups-considerations.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Multiple Db2 databases

Finding the endpoint
