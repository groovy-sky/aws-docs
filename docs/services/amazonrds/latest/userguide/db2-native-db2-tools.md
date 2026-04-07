# Using native Db2 tools to migrate data from Db2 to Amazon RDS for Db2

You can use several native Db2 tools, utilities, and commands to move data directly from a
Db2 database to an Amazon RDS for Db2 database. To use these native Db2 tools, you must be able to
connect your client machine to an RDS for Db2 DB instance. For more information, see [Connecting a client machine to an Amazon RDS for Db2 DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-connecting-client-rds.html).

###### Note

Another way to move your data is to first save it to an Amazon S3 bucket, and then use the
`LOAD` command to transfer that data into a table in your RDS for Db2
database. This method provides the best performance when migrating a large amount of data
because of good network connectivity between RDS for Db2 and S3. For more information, see
[Migrating Db2 data through Amazon S3 to Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-migration-load-from-s3.html).

Tool nameUse caseLimitations

[db2look](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-native-db2-tools-db2look.html)

Copying metadata from a self-managed Db2 database to an RDS for Db2
database.

- You must modify the syntax for creating buffer pools, creating
tablespaces, and creating roles to match the syntax used by the
[RDS for Db2 stored procedures](db2-stored-procedures.md).

[IMPORT\
command](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-native-db2-tools-import.html)

Migrating small tables and tables with large objects (LOBs) from a client
machine to the RDS for Db2 DB instance.

- Slower than the `LOAD` utility because of
`INSERT` and `DELETE` logging
operations.

- Poor performance with limited network bandwidth.

[INGEST\
utility](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-native-db2-tools-ingest.html)

Continually streaming data from files and pipes _without_ large objects (LOBs) on the client machine to the
RDS for Db2 DB instance. Supports `INSERT` and `MERGE`
operations.

- Can't stream data files that contain LOBs. Use the
`IMPORT` command instead.

- Connectivity required between self-managed Db2 database and
RDS for Db2 database.

[INSERT\
command](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-native-db2-tools-insert.html)

Copying data in small tables from a self-managed Db2 database to an
RDS for Db2 database.

- Connectivity required between self-managed Db2 database and
RDS for Db2 database.

- Poor performance with limited network bandwidth.

[LOAD CLIENT\
command](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-native-db2-tools-load.html)

Migrating small tables _without_ large
objects (LOBs) from a client machine to the RDS for Db2 DB instance.

- Can't migrate data files that contain LOBs. Use the
`IMPORT` command instead.

- Poor performance with limited network bandwidth.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Migrating with AWS DMS

Connecting a client machine to RDS for Db2
